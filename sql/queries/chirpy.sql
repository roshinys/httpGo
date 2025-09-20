-- name: CreateChirpy :one

INSERT INTO chirpy (id,createdAt,updatedAt,body,userId)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;


-- name: FetchChirps :many 
SELECT * FROM chirpy ORDER BY createdAt;

