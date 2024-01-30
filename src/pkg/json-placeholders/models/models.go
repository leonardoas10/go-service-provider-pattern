package models

type JsonPlaceHolder struct {
	UserId string `json:"userId"`
	Id string `json:"id"`
	Title string `json:"title"`
	Completed  bool `json:"completed"`
}
type UpdateJsonPlaceHolder struct {
	Id string `json:"id"`
	Title string `json:"title"`
}

type JsonPlaceHolderId struct {
	Id string 
}