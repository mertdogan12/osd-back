package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	osd "github.com/mertdogan12/osd/pkg/user"
)

type UserObj struct {
	Id          int      `json:"id"`
	Permissions []string `json:"permissions"`
}

func checkPermission(r *http.Request, permission string) (bool, error) {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	if token[0] != "Bearer" {
		return false, errors.New("No token is given")
	}

	obj, err := ReqAuthGET[UserObj](token[1], os.Getenv("BACK_URL")+"users/me")
	if err != nil {
		// TODO error handling
		panic(err)
	}

    fmt.Println(obj.Permissions)

	return true, nil
}

func ReqAuthGET[T any](token string, uri string) (*T, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var out T

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, osd.AuthError
	}

	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
