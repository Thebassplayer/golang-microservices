package github

type GitHubErrorResponse struct {
	Message          string         `json:"message"`
	Errors           []GitHubErrors `json:"errors"`
	DocumentationUrl string         `json:"documentation_url"`
	Status           int            `json:"status"`
}

type GitHubErrors struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
