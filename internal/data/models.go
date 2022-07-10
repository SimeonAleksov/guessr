package data

import (
  "database/sql"
  "errors"
)


var (
  ErrRecordNotFound = errors.New("resource not found")
)


type Models struct {
  Songs SongModel
  Users UserModel
}


func NewModels(db *sql.DB) Models {
    return Models{
      Songs: SongModel{DB: db},
      Users: UserModel{DB: db},
    }
}
