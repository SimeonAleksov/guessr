package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"guessr.net/internal/validator"
)


var AnonymousUser = &User{}

type (
  User struct {
    ID int64 `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Username string `json:"username"`
    Password password `json:"-"`
    Version int `json:"-"`
}

  UserModel struct {
    DB *sql.DB
  }

  password struct {
    plaintext *string
    hash []byte
  }
)


func (u *User) IsAnonymous() bool {
  return u == AnonymousUser
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


func (m UserModel) GetByUsername(username string) (*User, error) {
  query := `
  SELECT id, created_at, username, password_hash
  FROM users
  WHERE username = $1
  `

  var user User

  ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
  defer cancel()

  err := m.DB.QueryRowContext(ctx, query, username).Scan(
    &user.ID,
    &user.CreatedAt,
    &user.Username,
    &user.Password.hash,
  )
  if err != nil {
    switch {
    case errors.Is(err, sql.ErrNoRows):
      return nil, ErrRecordNotFound
    default:
      return nil, err
    }
  }
  return &user, nil
}

func (m UserModel) GetForToken(tokenScope, tokenPlaintext string) (*User, error) {
  // Calculate the SHA-256 hash of the plaintext token provided by the client.
  // Remember that this returns a byte *array* with length 32, not a slice.
  tokenHash := sha256.Sum256([]byte(tokenPlaintext))
  // Set up the SQL query.
  query := `
  SELECT users.id, users.created_at, users.username, users.password_hash, users.version
  FROM users
  INNER JOIN token
  ON users.id = token.user_id
  WHERE token.hash = $1
  AND token.scope = $2
  AND token.expiry > $3`
  // Create a slice containing the query arguments. Notice how we use the [:] operator
  // to get a slice containing the token hash, rather than passing in the array (which
  // is not supported by the pq driver), and that we pass the current time as the
  // value to check against the token expiry.
  args := []interface{}{tokenHash[:], tokenScope, time.Now()}
  var user User
  ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
  defer cancel()
  err := m.DB.QueryRowContext(ctx, query, args...).Scan(
    &user.ID,
    &user.CreatedAt,
    &user.Username,
    &user.Password.hash,
    &user.Version,
  )
  if err != nil {
    switch {
    case errors.Is(err, sql.ErrNoRows):
      return nil, ErrRecordNotFound
    default:
      return nil, err
    }
  }
  return &user, nil
}
