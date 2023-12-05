package dto

type GitHubUser struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Bio       string `json:"bio,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
