package main

import (
	"context"
	"encoding/json"
)

func doLogin(ctx context.Context, req RequestLogin) (ResponseLogin, error) {

	var res ResponseLogin

	respLogin, _, err := doRequest(ctx, "POST", "/sign-in/", "application/json", req, "")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(respLogin, &res)
	return res, err

}
