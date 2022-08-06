package schema

type Author struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type AuthorRequest struct {
	ID string `json:"id" binding:"required"`
}
