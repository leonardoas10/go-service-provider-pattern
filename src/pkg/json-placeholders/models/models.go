package models

type JsonPlaceHolder struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed  bool `json:"completed"`
}

type UpdateJsonPlaceHolder struct {
	Id int `json:"id"`
	Title string `json:"title"`
}

type JsonPlaceHolderId struct {
	Id int 
}