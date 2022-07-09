CREATE TABLE IF NOT EXISTS songs (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  title text NOT NULL,
  artist text NOT NULL,
  year integer NOT NULL,
  genres text[] NOT NULL
)
