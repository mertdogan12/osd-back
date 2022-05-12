package api

import (
	"net/http"

	"github.com/mertdogan12/osd-perm/pkg/helper"
	"github.com/mertdogan12/osd/pkg/user"
)

func SaveReplay(w http.ResponseWriter, r *http.Request) {
	perm, err := checkPermission(r, "replay.save")

	if err == user.AuthError {
		helper.ApiRespond(http.StatusUnauthorized, "Token is invalid", w)
		return
	}

	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	if !perm {
		helper.ApiRespond(http.StatusUnauthorized, "No permissions", w)
		return
	}

	helper.ApiRespond(http.StatusOK, ":)", w)
}
