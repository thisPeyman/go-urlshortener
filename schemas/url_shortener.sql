CREATE TABLE urls (
    short_url TEXT UNIQUE PRIMARY KEY,
    long_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);