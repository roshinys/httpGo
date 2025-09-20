-- +goose Up
CREATE TABLE chirpy (
    id UUID PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    body TEXT NOT NULL, 
    userId UUID NOT NULL,
    FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE 
);

-- +goose Down
DROP TABLE chirpy;