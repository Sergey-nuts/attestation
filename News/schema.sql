DROP TABLE IF EXISTS news;

CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT,
    pubtime BIGINT NOT NULL DEFAULT extract(epoch from now()),
    link TEXT NOT NULL UNIQUE
);