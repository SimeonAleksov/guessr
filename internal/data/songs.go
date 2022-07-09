package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lib/pq"
	"guessr.net/internal/validator"
)

type Song struct {
  ID int64 `json:"id"`
  CreatedAt time.Time `json:"created_at"`
  Title string `json:"title"`
  Artist string `json:"artist"`
  Year int32 `json:"year"`
  Genres []string `json:"genres"`
}

func ValidateSong(v *validator.Validator, song *Song) {
  v.Check(song.Title != "", "title", "must be provided")
  v.Check(len(song.Title) <= 500, "title", "must not be more than 500 bytes long")
  v.Check(song.Year != 0, "year", "must be provided")
  v.Check(song.Year >= 1888, "year", "must be greater than 1888")
  v.Check(song.Year <= int32(time.Now().Year()), "year", "must not be in the future")
  v.Check(song.Genres != nil, "genres", "must be provided")
  v.Check(len(song.Genres) >= 1, "genres", "must contain at least 1 genre")
  v.Check(len(song.Genres) <= 5, "genres", "must not contain more than 5 genres")
  v.Check(validator.Unique(song.Genres), "genres", "must not contain duplicate values")
}


type SongModel struct {
  DB *sql.DB
}


func (s SongModel) Insert(song *Song) error {
  query := `
    INSERT INTO songs (title, year, artist, genres)
    VALUES ($1, $2, $3, $4)
    RETURNING id, created_at`


    args := []interface{}{song.Title, song.Year, song.Artist, pq.Array(song.Genres)}
    return s.DB.QueryRow(query, args...).Scan(&song.ID, &song.CreatedAt)
}


func (s SongModel) Get(id int64) (*Song, error) {
  if id < 1 {
    return nil, ErrRecordNotFound
  }


  query := `
      SELECT id, created_at, title, artist, year, genres
      FROM songs
      WHERE id = $1`

    var song Song

    err := s.DB.QueryRow(query, id).Scan(
        &song.ID,
        &song.CreatedAt,
        &song.Title,
        &song.Artist,
        &song.Year,
        pq.Array(&song.Genres),
    )
    if err != nil {
        switch {
            case errors.Is(err, sql.ErrNoRows):
                return nil, ErrRecordNotFound
            default:
                return nil, err
        }
    }

    return &song, nil
}


func (s SongModel) Update(song *Song) error {
    query := `
        UPDATE songs
        SET title = $1, artist = $2, year = $3, genres = $4
        WHERE id = $5
        RETURNING id`
        args := []interface{}{
            song.Title,
            song.Artist,
            song.Year,
            pq.Array(song.Genres),
            song.ID,
        }
        return s.DB.QueryRow(query, args...).Scan(&song.ID)
}


func (s SongModel) Delete(id int64) error {
  query := `
      DELETE FROM songs where id = $1`

  result, err := s.DB.Exec(query, id)
  if err != nil {
    return err
  }

  rowsAffected, err := result.RowsAffected()
  if err != nil {
    return err
  }

  if rowsAffected == 0 {
    return ErrRecordNotFound
  }

  return nil
}


func (s SongModel) GetAll() ([]*Song, error) {
      query := `
        SELECT id, created_at, title, artist, year, genres
        FROM songs ORDER BY id`
      ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
      defer cancel()

      rows, err := s.DB.QueryContext(ctx, query)
      if err != nil {
        return nil, err
      }

      defer rows.Close()

      songs := []*Song{}

      for rows.Next() {
          var song Song
          err := rows.Scan(
              &song.ID,
              &song.CreatedAt,
              &song.Title,
              &song.Artist,
              &song.Year,
              pq.Array(&song.Genres),
          )
          if err != nil {
              return nil, err
          }
          songs = append(songs, &song)
      }

      if err = rows.Err(); err != nil {
          return nil, err
      }

      return songs, nil
}
