package jsonplaceholders

type JsonPlaceHolder struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed  bool `json:"completed"`
}

type PostJsonPlaceHolder struct {
	Id int 
}