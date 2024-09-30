package github

/*
{
    "name": "test-repo-from.api",
    "description": "This is test repo",
    "homepage": "https://github.com",
    "private": false,
    "is_template": true
}
*/

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
}
type CreateRepoResponse struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	Owner       RepoOwner   `json:"owner"`
	Permissions Permissions `json:"permissions"`
}

type RepoOwner struct {
	Id      int64  `json:"id"`
	Url     string `json:"url"`
	Login   string `json:"login"`
	HtmlUrl string `json:"html_url"`
}

type Permissions struct {
	Admin    bool `json:"admin"`
	Maintain bool `json:"maintain"`
	Push     bool `json:"push"`
	Triage   bool `json:"triage"`
	Pull     bool `json:"pull"`
}
