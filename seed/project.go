package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func createProject(ctx context.Context, user *Identity, workspaceSLug string, req RequestCreateProject) (ResponseCreateProject, error) {

	var res ResponseCreateProject

	respCreateProject, _, err := DoRequest(ctx, "POST", fmt.Sprintf("/workspaces/%v/projects/", workspaceSLug), "application/json", req, user.Token)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(respCreateProject, &res)
	return res, err

}
