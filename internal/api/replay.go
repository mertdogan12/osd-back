package api

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mertdogan12/osd-perm/pkg/helper"
	"github.com/mertdogan12/osd/pkg/user"
	parser "github.com/mertdogan12/osu-replay-converter/pkg/osu-replay-parser"
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

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	filepath := os.Getenv("BACK_SAVEDIR") + "/replay.osr"
	err = os.WriteFile(filepath, body, 0644)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	parsedObj, err := parser.ConvertToObject(filepath)

	helper.ApiRespond(http.StatusOK, "Replay from "+parsedObj.PlayerName+" saved.", w)
}
