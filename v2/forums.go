package v2

type Forum struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	URL            string    `json:"url"`
	Description    *string   `json:"description"`
	Type           ForumType `json:"type"`
	ThreadCount    int       `json:"threadCount"`
	PostCount      int       `json:"postCount"`
	LastPostID     string    `json:"lastPostId"`
	LastPostUserID string    `json:"lastPostUserId"`
	LastPostDate   int       `json:"lastPostDate"`
	TouchDate      int       `json:"touchDate"`
}

type Thread struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	GameID            string `json:"gameId"`
	ForumID           string `json:"forumId"`
	UserID            string `json:"userId"`
	Replies           int    `json:"replies"`
	Created           int    `json:"created"`
	LastCommentID     string `json:"lastCommentId"`
	LastCommentUserID string `json:"lastCommentUserId"`
	LastCommentDate   int    `json:"lastCommentDate"`
	Sticky            bool   `json:"sticky"`
	Locked            bool   `json:"locked"`
}
