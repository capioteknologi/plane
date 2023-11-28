package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	BaseURL    = ""
	Admin      *Identity
	Workspaces []Workspace
	uniqKey    = time.Now().Unix()
	Members    = []Identity{}
	// Projects   []Project
)

func main() {
	structs.DefaultTagName = "json"
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	BaseURL = os.Getenv("BASE_URL")

	// create workspaces
	respLogin, err := doLogin(ctx, RequestLogin{
		Email:    "captain@plane.so",
		Password: "password123",
		Medium:   "email",
	})
	if err != nil {
		panic(err)
	}

	Admin = &Identity{
		ID:        respLogin.User.ID,
		Username:  respLogin.User.Username,
		Email:     respLogin.User.Email,
		FirstName: respLogin.User.FirstName,
		LastName:  respLogin.User.LastName,
		Token:     respLogin.AccessToken,
	}

	createWorkspaces(ctx)

	for _, v := range Workspaces {
		createMembers(ctx, v.ID)
	}

	createProjects(ctx, Workspaces)

	// var wg sync.WaitGroup

	for _, v := range Workspaces {
		for _, w := range v.Projects {
			createTasks(ctx, v.Slug, w.ID)
		}
		// wg.Add(1)
		// 	go func(ctx context.Context, workspaceSlug string, projectID string) {
		// 		// defer wg.Done()
		// 		createTasks(ctx, workspaceSlug, projectID)
		// 	}(ctx, v.Slug, w.ID)
		// }
	}

	// wg.Wait()
}

func createWorkspaces(ctx context.Context) {

	for i := 0; i < 10; i++ {
		order := i + 1
		respCreateWorkspace, err := createWorkspace(ctx, Admin, RequestCreateWorkspace{
			Name:             fmt.Sprintf("PT Sukses %v %v", order, uniqKey),
			Slug:             fmt.Sprintf("pt-sukses-%v-%d", order, uniqKey),
			OrganizationSize: "11-50",
		})
		if err != nil {
			panic(err)
		}
		Workspaces = append(Workspaces, Workspace{
			ID:   respCreateWorkspace.ID,
			Slug: respCreateWorkspace.Slug,
		})
	}

}

func createMembers(ctx context.Context, workspaceID string) {

	users := []Identity{}
	// create members
	for i := 0; i < 15; i++ {

		order := len(Members) + 1

		Members = append(Members, Identity{
			ID:        uuid.NewString(),
			Username:  fmt.Sprintf("%v%d", os.Getenv("PREFIX_EMAIL"), order),
			Email:     fmt.Sprintf("%v%d@gmail.com", os.Getenv("PREFIX_EMAIL"), order),
			FirstName: "Test",
			LastName:  fmt.Sprintf("%d", order),
		})

		users = append(users, Identity{
			ID:        uuid.NewString(),
			Username:  fmt.Sprintf("%v%d", os.Getenv("PREFIX_EMAIL"), order),
			Email:     fmt.Sprintf("%v%d@gmail.com", os.Getenv("PREFIX_EMAIL"), order),
			FirstName: "Test",
			LastName:  fmt.Sprintf("%d", order),
		})
	}
	dbURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	fmt.Println("++++++++++++++++++++++++++++++++++++++ ", dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, v := range users {

		_, err = db.Exec(fmt.Sprintf(`
		INSERT INTO public.users
		("password", last_login, id, username, mobile_number, email, first_name, last_name, avatar, date_joined, created_at, updated_at, last_location, created_location, is_superuser, is_managed, is_password_expired, is_active, is_staff, is_email_verified, is_password_autoset, is_onboarded, "token", billing_address_country, billing_address, has_billing_address, user_timezone, last_active, last_login_time, last_logout_time, last_login_ip, last_logout_ip, last_login_medium, last_login_uagent, token_updated_at, last_workspace_id, my_issues_prop, "role", is_bot, theme, is_tour_completed, onboarding_step, cover_image, display_name)
		VALUES('pbkdf2_sha256$600000$gXSPc3SXnMUg6oS0sUfF4T$9lNoVk9z8L8/YYQKgJatflF/D1re5FjbLBml+IQWBNM=', NULL, '%v'::uuid, '%v', NULL, '%v', '%v', '%v', '', '2023-11-27 10:23:08.762', '2023-11-27 10:23:08.762', '2023-11-27 10:30:30.737', '', '', false, false, false, true, false, false, false, true, 'e22ec50d74b64170bbd4b4b04c927b4f2305fc046ce3432abd9e9af806d056d5', 'INDIA', NULL, false, 'Asia/Jakarta', '2023-11-27 10:26:33.909', '2023-11-27 10:26:33.909', NULL, '172.22.0.10', '', 'email', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36', '2023-11-27 10:30:30.737', NULL, NULL, 'Freelancer / Consultant', false, '{}'::jsonb, true, '{"workspace_join": true, "profile_complete": true, "workspace_create": true, "workspace_invite": true}'::jsonb, NULL, '%v');

		INSERT INTO public.workspace_members
		(created_at, updated_at, id, "role", created_by_id, member_id, updated_by_id, workspace_id, company_role, view_props, default_props, issue_props)
		VALUES('2023-11-27 11:23:36.100', '2023-11-27 11:23:36.100', '%v'::uuid, 15, '%v'::uuid, '%v'::uuid, NULL, '%v'::uuid, NULL, '{"filters": {"state": null, "labels": null, "priority": null, "assignees": null, "created_by": null, "start_date": null, "subscriber": null, "state_group": null, "target_date": null}, "display_filters": {"type": null, "layout": "list", "group_by": null, "order_by": "-created_at", "sub_issue": true, "show_empty_groups": true, "calendar_date_range": ""}, "display_properties": {"key": true, "link": true, "state": true, "labels": true, "assignee": true, "due_date": true, "estimate": true, "priority": true, "created_on": true, "start_date": true, "updated_on": true, "sub_issue_count": true, "attachment_count": true}}'::jsonb, '{"filters": {"state": null, "labels": null, "priority": null, "assignees": null, "created_by": null, "start_date": null, "subscriber": null, "state_group": null, "target_date": null}, "display_filters": {"type": null, "layout": "list", "group_by": null, "order_by": "-created_at", "sub_issue": true, "show_empty_groups": true, "calendar_date_range": ""}, "display_properties": {"key": true, "link": true, "state": true, "labels": true, "assignee": true, "due_date": true, "estimate": true, "priority": true, "created_on": true, "start_date": true, "updated_on": true, "sub_issue_count": true, "attachment_count": true}}'::jsonb, '{"created": true, "assigned": true, "all_issues": true, "subscribed": true}'::jsonb);
	`, v.ID, v.Username, v.Email, v.FirstName, v.LastName, v.FirstName, uuid.NewString(), Admin.ID, v.ID, workspaceID))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("SQL statements executed successfully")
	}
}

func createProjects(ctx context.Context, workspaces []Workspace) {

	projectOrder := 0
	for j, v := range workspaces {

		for i := 0; i < 5; i++ {
			order := i + 1
			projectOrder += 1
			respCreateProject, err := createProject(ctx, Admin, v.Slug, RequestCreateProject{
				Name:        fmt.Sprintf("Project %v %v", order, uniqKey),
				CoverImage:  "https://images.unsplash.com/photo-1531045535792-b515d59c3d1f?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=870&q=80",
				Description: fmt.Sprintf("Project %v %v", order, uniqKey),
				Identifier:  fmt.Sprintf("%v%v", projectOrder, time.Now().Unix()),
				Network:     2,
				ProjectLead: nil,
			})
			if err != nil {
				panic(err)
			}
			workspaces[j].Projects = append(workspaces[j].Projects, Project{
				ID: respCreateProject.ID,
			})
		}

	}

}

func createTasks(ctx context.Context, workspaceSlug string, projectID string) {

	for i := 0; i < 200; i++ {
		order := i + 1

		req := RequestCreateTask{
			Project: projectID,
			Name:    fmt.Sprintf("Test Task %v", order),
			Description: Description{
				Type: "doc",
				Content: []Content{
					{
						Type: "paragraph",
						Content: []Content2{
							{
								Type: "text",
								Text: fmt.Sprintf("Test %v", order),
							},
						},
					},
				},
			},

			DescriptionHTML: fmt.Sprintf("<p>Test %v</p>", order),
			EstimatePoint:   nil,
			State:           "",
			Parent:          nil,
			Priority:        "none",
			Assignees:       []any{},
			AssigneesList:   []string{Admin.ID},
			Labels:          []any{},
			LabelsList:      []any{},
			StartDate:       nil,
			TargetDate:      nil,
		}
		respCreateTask, err := createTask(ctx, Admin, workspaceSlug, projectID, req)
		if err != nil {
			panic(err)
		}
		fmt.Println(respCreateTask)
	}

}

func DoRequest(ctx context.Context, method string, endpoint string, contentType string, body interface{}, accessToken string) ([]byte, int, error) {
	resp, statusCode, err := doRequest(ctx, method, endpoint, contentType, body, accessToken)
	if err != nil {
		if strings.Contains(err.Error(), "You do not have permission to perform this action.") {
			// create workspaces
			respLogin, err := doLogin(ctx, RequestLogin{
				Email:    "captain@plane.so",
				Password: "password123",
				Medium:   "email",
			})
			if err != nil {
				panic(err)
			}
			Admin.Token = respLogin.AccessToken
			return doRequest(ctx, method, endpoint, contentType, body, Admin.Token)
		} else {
			return resp, statusCode, err
		}
	}

	return resp, statusCode, err
}

func doRequest(ctx context.Context, method string, endpoint string, contentType string, body interface{}, accessToken string) ([]byte, int, error) {

	url := fmt.Sprint(BaseURL, endpoint)

	mapBody := structs.Map(body)

	bodyByte, err := json.Marshal(mapBody)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "cannot marshal json, body: %v", mapBody)
	}
	fmt.Println(string(bodyByte))

	HTTPReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(bodyByte))
	if err != nil {
		return nil, 0, errors.Wrapf(err, "cannot make request, body: %v", mapBody)
	}
	HTTPReq.Header.Set("Content-Type", contentType)
	if accessToken != "" {
		HTTPReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}
	// HTTPReq.Header.Set("X-Request-ID", logger.GetRequestID(ctx))

	// otel.Inject(ctx, HTTPReq.Header)

	resp, err := http.DefaultClient.Do(HTTPReq)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "cannot do request, body: %v", mapBody)
	}
	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "cannot read response, body: %v", mapBody)
	}

	fmt.Println("++++++++++++++++++++++++++++++++++++", string(respByte))

	if resp.StatusCode > 300 {
		var respErr ResponseError
		err = json.Unmarshal(respByte, &respErr)
		if err != nil {
			return nil, 0, errors.Wrap(err, "cannot unmarshal error response")
		}

		err = errors.Wrapf(respErr, "error calling %s", endpoint)
		return respByte, resp.StatusCode, err
	}

	return respByte, resp.StatusCode, err
}
