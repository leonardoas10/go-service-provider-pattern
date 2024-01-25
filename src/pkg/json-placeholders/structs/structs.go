package structs

type RequestId struct {
	Id int `json:"id" validate:"required,gte=1"`
}