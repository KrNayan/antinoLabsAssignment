package blog

import "time"

// BlogConfig - it contains the blog config
type BlogConfig struct {
	// BlogId contains the blog int
	BlogId int `json:"BlogId,omitempty"`
	//EmailId contains the email id
	EmailId string `json:"EmailId,omitempty"`
	// Blog contains the blog
	Blog string `json:"Blog,omitempty"`
	// PostedOn contains the timestamp of blog posting
	PostedOn *time.Time `json:"PostedOn,omitempty"`
}
