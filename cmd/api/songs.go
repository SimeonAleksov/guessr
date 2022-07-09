package main

import (
  "errors"
  "fmt"
  "net/http"
  "log"

  "guessr.net/internal/data"
  "guessr.net/internal/validator"
)

func (app *application) createSongHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Title string `json:"title"`
        Year int32 `json:"year"`
        Artist string `json:"artist"`
        Genres []string `json:"genres"`
    }
    err := app.readJSON(w, r, &input)
    if err != nil {
        app.badRequestResponse(w, r, err)
        return
    }
    song := &data.Song{
        Title: input.Title,
        Artist: input.Artist,
        Year: input.Year,
        Genres: input.Genres,
    }

    v := validator.New()
    if data.ValidateSong(v, song); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    err = app.models.Songs.Insert(song)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }

    headers := make(http.Header)
    headers.Set("Location", fmt.Sprintf("/v1/songs/%d", song.ID))
    err = app.writeJSON(w, http.StatusCreated, envelope{"song": song}, headers)

    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}

func (app *application) showSongHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
        http.NotFound(w, r)
        return
    }
    song, err := app.models.Songs.Get(id)
    if err != nil {
        switch {
            case errors.Is(err, data.ErrRecordNotFound):
              app.notFoundResponse(w, r)
            default:
              log.Println(err)
              app.serverErrorResponse(w, r, err)
            return
        }
    }
    err = app.writeJSON(w, http.StatusOK, envelope{"song": song}, nil)
    if err != nil {
        app.logger.Println(err)
        app.serverErrorResponse(w, r, err)
  }
}


func (app *application) updateSongHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam((r))

    if err != nil {
        http.NotFound(w, r)
        return
    }
    song, err := app.models.Songs.Get(id)
    if err != nil {
        switch {
            case errors.Is(err, data.ErrRecordNotFound):
              app.notFoundResponse(w, r)
            default:
              log.Println(err)
              app.serverErrorResponse(w, r, err)
            return
        }
    }

    var input struct {
        Title string `json:"title"`
        Artist string `json:"artist"`
        Year int32 `json:"Year"`
        Genres []string `json:"genres"`
    }
    song.Title = input.Title
    song.Artist = input.Artist
    song.Year = input.Year
    song.Genres = input.Genres

    v := validator.New()

    if data.ValidateSong(v, song); !v.Valid() {
        app.failedValidationResponse(w, r, v.Errors)
        return
    }

    err = app.models.Songs.Update(song)
    if err != nil {
        app.serverErrorResponse(w, r, err)
    }

    err = app.writeJSON(w, http.StatusOK, envelope{"song": song}, nil)

    if err != nil {
        app.serverErrorResponse(w, r, err)
        return
    }
}



func (app *application) deleteSongHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam((r))

    if err != nil {
        http.NotFound(w, r)
        return
    }
    err = app.models.Songs.Delete(id)
    if err != nil {
        switch {
            case errors.Is(err, data.ErrRecordNotFound):
              app.notFoundResponse(w, r)
            default:
              log.Println(err)
              app.serverErrorResponse(w, r, err)
            return
        }
    }
    err = app.writeJSON(w, http.StatusOK, envelope{"message": "song successfully deleted"}, nil)
}



func (app *application) listSongsHandler(w http.ResponseWriter, r *http.Request) {
    songs, err := app.models.Songs.GetAll()
    log.Println(songs)
    if err != nil {
        app.serverErrorResponse(w, r, err)
    }

    err = app.writeJSON(w, http.StatusOK, envelope{"songs": songs}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
    }
}
