CREATE TABLE IF NOT EXISTS reviews (
    id UUID PRIMARY KEY,
    user_ip TEXT NOT NULL,
    content TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
