package structs

type RequestId struct {
	Id int `json:"id" validate:"required,gte=1"`
}

type RequestUpdateJsonPlaceHolder struct {
	Id int `json:"id" validate:"required,gte=1"`
	Title string `json:"title" validate:"required"`
}