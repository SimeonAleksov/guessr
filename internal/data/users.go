package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"guessr.net/internal/validator"
)


type User struct {
    ID int64 `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Username string `json:"username"`
    Password password `json:"-"`
    Version int `json:"-"`
}


type password struct {
  plaintext *string
  hash []byte
}


func (p *password) Set(plaintextPassword string) error {
  hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)

  if err != nil {
    return err
  }

  p.plaintext = &plaintextPassword
  p.hash = hash

  return nil
}


func (p *password) Matches(plaintextPassword string) (bool, error) {
  err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))

  if err != nil {
      switch {
          case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
              return false, nil
          default:
              return false, nil
      }
  }

  return true, nil
}


func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}


func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Username != "", "username", "must be provided")
	v.Check(len(user.Username) <= 500, "username", "must not be more than 500 bytes long")

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	if user.Password.hash == nil {
        panic("missing password hash for user")
	}
}



var (
    ErrDuplicateUsername = errors.New("username has been taken")
)


type UserModel struct {
  DB *sql.DB
}


func (u UserModel) Insert(user *User) error {
    query := `
    INSERT INTO users (username, password_hash)
    VALUES ($1, $2)
    RETURNING id, created_at, version`
    args := []interface{}{user.Username, user.Password.hash}

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    err := u.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt, &user.Version)
    if err != nil {
      switch {
      case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:

          return ErrDuplicateUsername
      default:
          return err
      }
    }
    return nil
}
