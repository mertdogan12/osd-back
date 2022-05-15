package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/mertdogan12/osd-perm/pkg/helper"
	"github.com/mertdogan12/osd/pkg/user"
	parser "github.com/mertdogan12/osu-replay-converter/pkg/osu-replay-parser"
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

	// Verifys the replay
	scoreID := r.URL.Query()["scoreID"][0]
	beatmapID := r.URL.Query()["beatmapID"][0]

	_, err = strconv.Atoi(scoreID)
	if err != nil {
		helper.ApiRespond(http.StatusInternalServerError, "Can't convert the scoreID to an int: "+scoreID, w)
		return
	}

	_, err = strconv.Atoi(beatmapID)
	if err != nil {
		helper.ApiRespond(http.StatusInternalServerError, "Can't convert the beatmapID to an int: "+beatmapID, w)
		return
	}

	// TODO verify the replay

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	// Saves the replay
	folder := filepath.Join(".", os.Getenv("BACK_SAVEDIR"), strconv.Itoa(*id), "Mert Dogan")
	err = os.MkdirAll(folder, os.ModePerm)

	// TODO save replay in <usernameID>/<player from replay>/<scoreid>;<beatmap id>;Replay online id (if exists).osr
	replaypath := fmt.Sprintf("%s/%d/%s/%s;%s;%s.osr", os.Getenv("BACK_SAVEDIR"), *id, "Mert Dogan", scoreID, beatmapID, "-1")
	err = os.WriteFile(replaypath, body, 0644)
	if err != nil {
		helper.ApiRespondErr(err, w)
		return
	}

	parsedObj, err := parser.ConvertToObject(replaypath)

	helper.ApiRespond(http.StatusOK, "Replay from "+parsedObj.PlayerName+" saved.", w)
}
