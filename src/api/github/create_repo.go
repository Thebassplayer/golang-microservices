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
