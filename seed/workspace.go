package main

import (
	"context"
	"encoding/json"
)

func createWorkspace(ctx context.Context, user *Identity, req RequestCreateWorkspace) (ResponseCreateWorkspace, error) {

	var res ResponseCreateWorkspace

	respCreateWorkspace, _, err := DoRequest(ctx, "POST", "/workspaces/", "application/json", req, user.Token)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(respCreateWorkspace, &res)
	return res, err

}
