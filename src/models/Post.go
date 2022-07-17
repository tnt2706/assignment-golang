package models

type Post struct {
	Id string `json:"_id"`
	Title string `json."title"`
	Description string `json:"description"`
	Content string `json:"content"`
	Owner string `json:"owner"`
	isPus string `json."salt"`
	Hash string `json."hash"`
	Roles [string] `json."roles"`
	Status string `json."status"`
}