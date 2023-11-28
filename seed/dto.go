package main

import "time"

type ResponseError struct {
	Detail string `json:"detail"`
}

func (e ResponseError) Error() string {
	return e.Detail
}

type Workspace struct {
	ID       string    `json:"id"`
	Slug     string    `json:"slug"`
	Projects []Project `json:"projects"`
}

type Project struct {
	ID string `json:"id"`
}

type Identity struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Medium   string `json:"medium"`
}

type ResponseLogin struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type Theme struct {
}
type OnboardingStep struct {
	WorkspaceJoin   bool `json:"workspace_join"`
	ProfileComplete bool `json:"profile_complete"`
	WorkspaceCreate bool `json:"workspace_create"`
	WorkspaceInvite bool `json:"workspace_invite"`
}
type User struct {
	ID                    string         `json:"id"`
	LastLogin             any            `json:"last_login"`
	Username              string         `json:"username"`
	MobileNumber          any            `json:"mobile_number"`
	Email                 string         `json:"email"`
	FirstName             string         `json:"first_name"`
	LastName              string         `json:"last_name"`
	Avatar                string         `json:"avatar"`
	CoverImage            any            `json:"cover_image"`
	DateJoined            time.Time      `json:"date_joined"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	LastLocation          string         `json:"last_location"`
	CreatedLocation       string         `json:"created_location"`
	IsSuperuser           bool           `json:"is_superuser"`
	IsManaged             bool           `json:"is_managed"`
	IsPasswordExpired     bool           `json:"is_password_expired"`
	IsActive              bool           `json:"is_active"`
	IsStaff               bool           `json:"is_staff"`
	IsEmailVerified       bool           `json:"is_email_verified"`
	IsPasswordAutoset     bool           `json:"is_password_autoset"`
	IsOnboarded           bool           `json:"is_onboarded"`
	Token                 string         `json:"token"`
	BillingAddressCountry string         `json:"billing_address_country"`
	BillingAddress        any            `json:"billing_address"`
	HasBillingAddress     bool           `json:"has_billing_address"`
	UserTimezone          string         `json:"user_timezone"`
	LastActive            time.Time      `json:"last_active"`
	LastLoginTime         time.Time      `json:"last_login_time"`
	LastLogoutTime        time.Time      `json:"last_logout_time"`
	LastLoginIP           string         `json:"last_login_ip"`
	LastLogoutIP          string         `json:"last_logout_ip"`
	LastLoginMedium       string         `json:"last_login_medium"`
	LastLoginUagent       string         `json:"last_login_uagent"`
	TokenUpdatedAt        time.Time      `json:"token_updated_at"`
	LastWorkspaceID       any            `json:"last_workspace_id"`
	MyIssuesProp          any            `json:"my_issues_prop"`
	Role                  string         `json:"role"`
	IsBot                 bool           `json:"is_bot"`
	Theme                 Theme          `json:"theme"`
	DisplayName           string         `json:"display_name"`
	IsTourCompleted       bool           `json:"is_tour_completed"`
	OnboardingStep        OnboardingStep `json:"onboarding_step"`
	Groups                []any          `json:"groups"`
	UserPermissions       []any          `json:"user_permissions"`
}

type RequestCreateWorkspace struct {
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	OrganizationSize string `json:"organization_size"`
}

type ResponseCreateWorkspace struct {
	ID    string `json:"id"`
	Owner struct {
		ID          string `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Avatar      string `json:"avatar"`
		IsBot       bool   `json:"is_bot"`
		DisplayName string `json:"display_name"`
	} `json:"owner"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Name             string    `json:"name"`
	Logo             any       `json:"logo"`
	Slug             string    `json:"slug"`
	OrganizationSize string    `json:"organization_size"`
	CreatedBy        string    `json:"created_by"`
	UpdatedBy        string    `json:"updated_by"`
}

type RequestCreateProject struct {
	CoverImage  string `json:"cover_image"`
	Description string `json:"description"`
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	Network     int    `json:"network"`
	ProjectLead any    `json:"project_lead"`
	Emoji       string `json:"emoji"`
}

type ResponseCreateProject struct {
	ID              string          `json:"id"`
	WorkspaceDetail WorkspaceDetail `json:"workspace_detail"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	DescriptionText any             `json:"description_text"`
	DescriptionHTML any             `json:"description_html"`
	Network         int             `json:"network"`
	Identifier      string          `json:"identifier"`
	Emoji           string          `json:"emoji"`
	IconProp        any             `json:"icon_prop"`
	ModuleView      bool            `json:"module_view"`
	CycleView       bool            `json:"cycle_view"`
	IssueViewsView  bool            `json:"issue_views_view"`
	PageView        bool            `json:"page_view"`
	InboxView       bool            `json:"inbox_view"`
	CoverImage      string          `json:"cover_image"`
	ArchiveIn       int             `json:"archive_in"`
	CloseIn         int             `json:"close_in"`
	CreatedBy       string          `json:"created_by"`
	UpdatedBy       string          `json:"updated_by"`
	Workspace       string          `json:"workspace"`
	DefaultAssignee any             `json:"default_assignee"`
	ProjectLead     any             `json:"project_lead"`
	Estimate        any             `json:"estimate"`
	DefaultState    any             `json:"default_state"`
	SortOrder       float64         `json:"sort_order"`
	MemberRole      int             `json:"member_role"`
	IsMember        bool            `json:"is_member"`
}
type WorkspaceDetail struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	ID   string `json:"id"`
}

type Content2 struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type Content struct {
	Type    string     `json:"type"`
	Content []Content2 `json:"content"`
}
type Description struct {
	Type    string    `json:"type"`
	Content []Content `json:"content"`
}

type RequestCreateTask struct {
	Project         string      `json:"project"`
	Name            string      `json:"name"`
	Description     Description `json:"description"`
	DescriptionHTML string      `json:"description_html"`
	EstimatePoint   any         `json:"estimate_point"`
	State           string      `json:"state"`
	Parent          any         `json:"parent"`
	Priority        string      `json:"priority"`
	Assignees       []any       `json:"assignees"`
	AssigneesList   []string    `json:"assignees_list"`
	Labels          []any       `json:"labels"`
	LabelsList      []any       `json:"labels_list"`
	StartDate       any         `json:"start_date"`
	TargetDate      any         `json:"target_date"`
}

type ResponseCreateTask struct {
	ID string `json:"id"`
}
