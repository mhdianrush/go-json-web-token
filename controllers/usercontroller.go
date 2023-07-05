package controllers

import (
	"net/http"

	"github.com/mhdianrush/go-json-web-token/helper"
	"github.com/mhdianrush/go-json-web-token/models"
)

func Me(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("userinfo").(*helper.MyCustomClaims)
	userResponse := &models.MyProfile{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	helper.Response(w, 200, "My Profile", userResponse)
}
