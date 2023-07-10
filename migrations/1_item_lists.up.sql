CREATE TABLE IF NOT EXISTS item_lists (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE
);
