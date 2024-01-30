package structs

type RequestId struct {
	Id string `json:"id" validate:"required,gte=1"`
}

type RequestUpdateJsonPlaceHolder struct {
	Id string `json:"id" validate:"required,gte=1"`
	Title string `json:"title" validate:"required"`
}