package main

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
)


func (app *application) routes() http.Handler {
  router := httprouter.New()

  router.NotFound = http.HandlerFunc(app.notFoundResponse)


  router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

  router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

  router.HandlerFunc(http.MethodGet, "/v1/songs", app.requireAuthentication(app.listSongsHandler))
  router.HandlerFunc(http.MethodPost, "/v1/songs", app.createSongHandler)
  router.HandlerFunc(http.MethodGet, "/v1/songs/:id", app.showSongHandler)
  router.HandlerFunc(http.MethodPut, "/v1/songs/:id", app.updateSongHandler)
  router.HandlerFunc(http.MethodDelete, "/v1/songs/:id", app.deleteSongHandler)


  router.HandlerFunc(http.MethodPost, "/v1/users", app.createUserHandler)
  router.HandlerFunc(http.MethodPost, "/v1/token/", app.createAuthenticationTokenHandler)
  
  router.ServeFiles("/songs/*filepath", http.Dir("songs"))

  return app.enableCORS(app.authenticate(router))
}
