-- name: CreateUser :one

INSERT INTO users (id,createdAt,updatedAt,email)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1
)
RETURNING *;



-- name: FetchUserByEmail :one
SELECT * FROM users where email = $1;

