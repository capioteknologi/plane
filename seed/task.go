package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func createTask(ctx context.Context, user *Identity, workspaceSlug string, projectID string, req RequestCreateTask) (ResponseCreateTask, error) {

	var res ResponseCreateTask

	respCreateTask, _, err := DoRequest(ctx, "POST", fmt.Sprintf("/workspaces/%v/projects/%v/issues/", workspaceSlug, projectID), "application/json", req, user.Token)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(respCreateTask, &res)
	return res, err

}
