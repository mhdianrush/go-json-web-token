package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mhdianrush/go-json-web-token/configs"
	"github.com/mhdianrush/go-json-web-token/helper"
	"github.com/mhdianrush/go-json-web-token/models"
	"gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	err := json.NewDecoder(r.Body).Decode(&register)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helper.Response(w, 400, "Password Not Match", nil)
		return
	}

	passwordHash, err := helper.HashPassword(register.Password)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	err = configs.DB.Create(&user).Error
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Register Successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User

	err = configs.DB.First(&user, "email = ?", login.Email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Wrong Email Or Password", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	// verify password match or not
	err = helper.VerifyPassword(user.Password, login.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Wrong Password", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	// create/generate token
	token, err := helper.CreateToken(&user)
	if err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "Login Successfully", token)
}
