CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    created_at DATE DEFAULT CURRENT_DATE,
    is_completed BOOLEAN DEFAULT FALSE,
    item_list BIGINT REFERENCES item_lists(id)
);
