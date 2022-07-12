package models

import (
	"time"
)

type Song struct {
  ID int64 `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  Title string `json:"title"`
  Artist string `json:"artist"`
  Year int32 `json:"year"`
}


func (e *Song) Song() string {
	return "songs"
}
