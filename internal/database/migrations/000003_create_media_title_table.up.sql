CREATE TABLE media_title (
    id SERIAL PRIMARY KEY,                            -- Auto-incrementing primary key
    english VARCHAR(255),                             -- English Title
    media_id BIGINT NOT NULL UNIQUE,                  -- Foreign key to the media table, must be unique
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   -- Automatically set to current time on update
    FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE -- Define foreign key constraint
);
