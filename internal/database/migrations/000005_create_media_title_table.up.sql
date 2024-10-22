CREATE TABLE media_title (
    id SERIAL PRIMARY KEY,                            -- Auto-incrementing primary key
    english VARCHAR(255),                  -- English Title
    media_id INT NOT NULL,                           -- Foreign key to the media table
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE -- Define foreign key constraint
);

CREATE TRIGGER set_updated_timestamp_media_title
BEFORE UPDATE ON media_title
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();