package vo

import "time"

type User struct {
	ID        string                 `json:"id"`
	Username  string                 `json:"username"`
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	FullName  string                 `json:"full_name"`
	Email     string                 `json:"email"`
	Avatar    string                 `json:"avatar"`
	Raw       map[string]interface{} `json:"raw"` // Raw data
}

type UserInfo struct {
	Login                   string      `json:"login"`
	ID                      int         `json:"id"`
	NodeID                  string      `json:"node_id"`
	AvatarURL               string      `json:"avatar_url"`
	GravatarID              string      `json:"gravatar_id"`
	URL                     string      `json:"url"`
	HTMLURL                 string      `json:"html_url"`
	FollowersURL            string      `json:"followers_url"`
	FollowingURL            string      `json:"following_url"`
	GistsURL                string      `json:"gists_url"`
	StarredURL              string      `json:"starred_url"`
	SubscriptionsURL        string      `json:"subscriptions_url"`
	OrganizationsURL        string      `json:"organizations_url"`
	ReposURL                string      `json:"repos_url"`
	EventsURL               string      `json:"events_url"`
	ReceivedEventsURL       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    interface{} `json:"name"`
	Company                 string      `json:"company"`
	Blog                    string      `json:"blog"`
	Location                string      `json:"location"`
	Email                   string      `json:"email"`
	Hireable                bool        `json:"hireable"`
	Bio                     string      `json:"bio"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}
