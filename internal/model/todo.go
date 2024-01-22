package model

type Todo struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	Text string `json:"text" bson:"text,omitempty"`
	Done bool   `json:"done" bson:"done,omitempty"`
	User string `json:"user" bson:"user,omitempty"`
}
