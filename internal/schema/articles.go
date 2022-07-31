package schema

type Articles struct {
	ID        string `json:"id" binding:"required"`
	Author_ID string `json:"author_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Body      string `json:"body" binding:"required"`
}

type articleResponse struct {
	Created_at string `json:"created_at" binding:"required"`
}

type authorName struct {
	Name string `json:"name" required:"true"`
}

type ArticlesAuthor struct {
	Articles
	articleResponse
	authorName
}
