package main

import "time"

// Config : 設定ファイル
type Config struct {
	Space        string `json:"space"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AccessToken : アクセストークン
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// Priorities : 優先度リスト
type Priorities []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Issue : 課題
type Issue struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"projectId"`
	IssueKey  string `json:"issueKey"`
	KeyID     int    `json:"keyId"`
	IssueType struct {
		ID           int    `json:"id"`
		ProjectID    int    `json:"projectId"`
		Name         string `json:"name"`
		Color        string `json:"color"`
		DisplayOrder int    `json:"displayOrder"`
	} `json:"issueType"`
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Resolutions interface{} `json:"resolutions"`
	Priority    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"priority"`
	Status struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	Assignee struct {
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		RoleType    int         `json:"roleType"`
		Lang        interface{} `json:"lang"`
		MailAddress string      `json:"mailAddress"`
	} `json:"assignee"`
	Category  []interface{} `json:"category"`
	Versions  []interface{} `json:"versions"`
	Milestone []struct {
		ID             int         `json:"id"`
		ProjectID      int         `json:"projectId"`
		Name           string      `json:"name"`
		Description    string      `json:"description"`
		StartDate      interface{} `json:"startDate"`
		ReleaseDueDate interface{} `json:"releaseDueDate"`
		Archived       bool        `json:"archived"`
		DisplayOrder   int         `json:"displayOrder"`
	} `json:"milestone"`
	StartDate      interface{} `json:"startDate"`
	DueDate        interface{} `json:"dueDate"`
	EstimatedHours interface{} `json:"estimatedHours"`
	ActualHours    interface{} `json:"actualHours"`
	ParentIssueID  interface{} `json:"parentIssueId"`
	CreatedUser    struct {
		ID          int    `json:"id"`
		UserID      string `json:"userId"`
		Name        string `json:"name"`
		RoleType    int    `json:"roleType"`
		Lang        string `json:"lang"`
		MailAddress string `json:"mailAddress"`
	} `json:"createdUser"`
	Created     time.Time `json:"created"`
	UpdatedUser struct {
		ID          int    `json:"id"`
		UserID      string `json:"userId"`
		Name        string `json:"name"`
		RoleType    int    `json:"roleType"`
		Lang        string `json:"lang"`
		MailAddress string `json:"mailAddress"`
	} `json:"updatedUser"`
	Updated      time.Time     `json:"updated"`
	CustomFields []interface{} `json:"customFields"`
	Attachments  []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Size int    `json:"size"`
	} `json:"attachments"`
	SharedFiles []interface{} `json:"sharedFiles"`
	Stars       []struct {
		ID        int         `json:"id"`
		Comment   interface{} `json:"comment"`
		URL       string      `json:"url"`
		Title     string      `json:"title"`
		Presenter struct {
			ID          int    `json:"id"`
			UserID      string `json:"userId"`
			Name        string `json:"name"`
			RoleType    int    `json:"roleType"`
			Lang        string `json:"lang"`
			MailAddress string `json:"mailAddress"`
		} `json:"presenter"`
		Created time.Time `json:"created"`
	} `json:"stars"`
}
