package schema

type Articles struct {
	ID         string `json:"id"`
	Author_ID  string `json:"author_id"`
	Title      string `json:"title"`
	Body       string `json:"text"`
	Created_at string `json:"created_at"`
}
