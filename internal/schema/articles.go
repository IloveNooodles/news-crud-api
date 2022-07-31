package schema

type Articles struct {
	ID         string `json:"id" binding:"required"`
	Author_ID  string `json:"author_id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Body       string `json:"text" binding:"required"`
	Created_at string `json:"created_at" binding:"required"`
}

type GetArticlesRequest struct {
	Query  string
	Author string
}

type GetArticlesResponse struct {
	ListArticle []Articles
}