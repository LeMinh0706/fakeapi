-- name: CreateUser :exec
INSERT INTO users(
  email, full_name, hashed_password
) VALUES (
  $1, $2, $3
);

-- name: GetUserByUsername :one
SELECT id, email, hashed_password FROM users WHERE email = $1;
