DROP TABLE IF EXISTS comments;

CREATE TABLE comments(
    id SERIAL PRIMARY KEY,
    postid BIGINT NOT NULL,
    content TEXT NOT NULL,
    author TEXT,
    pubtime BIGINT NOT NULL DEFAULT extract(epoch from now())
);