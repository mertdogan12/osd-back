package api

import "net/http"

func SaveReplay(w http.ResponseWriter, r *http.Request) {
	checkPermission(r, "replay.save")
}
