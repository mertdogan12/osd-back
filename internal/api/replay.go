package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/mertdogan12/osd-back/internal/mongo"
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
	// TODO get id form parsedObj.PlayerName
	replaysId := uuid.New()

	// Saves the replay
	uploaderRes, playerRes, err := mongo.SaveReplay(*id, *id, replaysId)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	// Checks if the data actually got modified
	if uploaderRes.ModifiedCount == 0 {
		helper.ApiRespond(http.StatusBadGateway, fmt.Sprintf("Could not find user with the id: %d", *id), w)
		return
	}

	if playerRes.ModifiedCount == 0 {
		helper.ApiRespond(http.StatusBadGateway, fmt.Sprintf("Could not find user with the id: %d", *id), w)
		return
	}

	helper.ApiRespond(http.StatusOK, "Replay from "+parsedObj.PlayerName+" saved.", w)
}
