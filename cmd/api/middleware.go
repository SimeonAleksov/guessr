package main

import (
	"errors"
	"net/http"
	"strings"

	"guessr.net/internal/data"
	"guessr.net/internal/validator"
)

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add the "Vary: Authorization" header to the response. This indicates to any caches
		// that the response may vary based on the value of the Authorization header in the request.
		w.Header().Set("Vary", "Authorization")

		// Retrieve the value of the Authorization header from teh request. This will return the
		// empty string "" if there is no such header found.
		authorizationHeader := r.Header.Get("Authorization")

		// If there is no Authorization header found, use the contextSetUser() helper to add
		// an AnonymousUser to the request context. Then we call the next handler in the chain
		// and return without executing any of the code below.
		if authorizationHeader == "" {
			r = app.contextSetUser(r, data.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		// Otherwise, we expect the value of the Authorization header to be in the format
		// "Bearer <token>". We try to split this into its constituent parts, and if the header
		// isn't in the expected format we return a 401 Unauthorized response using the
		// invalidAuthenticationTokenResponse helper.
		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		// Extract the actual authentication toekn from the header parts
		token := headerParts[1]

		// Validate the token to make sure it is in a sensible format.
		v := validator.New()

		// If the token isn't valid, use the invalidAuthenticationtokenResponse
		// helper to send a response, rather than the failedValidatedResponse helper.
		if data.ValidateTokenPlaintext(v, token); !v.Valid() {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		// Retrieve the details of the user associated with the authentication token.
		// call invalidAuthenticationTokenResponse if no matching record was found.
		user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.invalidAuthenticationTokenResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}

		// Call the contextSetUser healer to add the user information to the request context.
		r = app.contextSetUser(r, user)

		// Call next handler in chain
		next.ServeHTTP(w, r)
	})
}


func (app *application) requireAuthentication(next http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  user := app.contextGetUser(r)
  if user.IsAnonymous() {
    app.authenticationRequiredResponse(w, r)
    return
  }
  next.ServeHTTP(w, r)
  })
}


func (app *application) enableCORS(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    next.ServeHTTP(w, r)
  })
}
