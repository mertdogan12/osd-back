package api

import (
	"io/ioutil"
	"net/http"

	"github.com/mertdogan12/osd-perm/pkg/helper"
	"github.com/mertdogan12/osd/pkg/user"
	parser "github.com/mertdogan12/osu-replay-parser"
)

func SaveReplay(w http.ResponseWriter, r *http.Request) {
	// Checks for permissions
	id, err := checkPermission(r, "replay.save")

	if err == user.AuthError {
		helper.ApiRespond(http.StatusUnauthorized, "Token is invalid", w)
		return
	}

	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	if id == nil {
		helper.ApiRespond(http.StatusUnauthorized, "No permissions", w)
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	if len(body) == 0 {
		helper.ApiRespond(http.StatusBadGateway, "Replay file is not given", w)
		return
	}

	parsedObj, err := parser.Parse(body)

	// Saves the replay

	helper.ApiRespond(http.StatusOK, "Replay from "+parsedObj.PlayerName+" saved.", w)
}
