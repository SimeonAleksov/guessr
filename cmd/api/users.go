package main

import (
	"errors"
	"net/http"

	"guessr.net/internal/data"
	"guessr.net/internal/validator"
)


func (app * application) createUserHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
      Username string `json:"username"`
      Password string `json:"password"`
    }

    err := app.readJSON(w, r, &input)
      if err != nil {
          app.badRequestResponse(w, r, err)
          return
      }


      user := &data.User{
        Username: input.Username,
      }

      err = user.Password.Set(input.Password)
      if err != nil {
        app.serverErrorResponse(w, r, err)
        return
      }

      v := validator.New()

      if data.ValidateUser(v, user); !v.Valid() {
        app.failedValidationResponse(w,r, v.Errors)
        return
      }

      err = app.models.Users.Insert(user)
      if err != nil {
        switch {
          case errors.Is(err, data.ErrDuplicateUsername):
              v.AddError("username", "this username is taken")
              app.failedValidationResponse(w, r, v.Errors)
          default:
              app.serverErrorResponse(w, r, err)
        }
        return
      }

      err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
      if err != nil {
          app.serverErrorResponse(w, r, err)
      }
}
