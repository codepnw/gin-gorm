package request

type UpdateTagsRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1,max=50" json:"name"`
}
