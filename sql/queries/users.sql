-- name: CreateUser :one

INSERT INTO users (id,createdAt,updatedAt,email,password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;



-- name: FetchUserByEmail :one
SELECT * FROM users where email = $1;

-- name: FetchUserWithoutPassword :one 
SELECT id,email,createdAt,updatedAt FROM users where email = $1;

