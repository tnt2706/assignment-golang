package model

import "time"

type Todo struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Text      string    `json:"text" bson:"text,omitempty"`
	Done      bool      `json:"done" bson:"done,omitempty"`
	User      string    `json:"user" bson:"user,omitempty"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func (t *Todo) Add() {
	t.UpdatedAt = time.Now()
	t.CreatedAt = time.Now()
}

func (t *Todo) Update() {
	t.UpdatedAt = time.Now()
}
