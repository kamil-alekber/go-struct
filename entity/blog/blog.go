package blog

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	Body    string `json:"body"`
	UserId  string `json:"user_id"`
	Created int    `json:"created"`
}
