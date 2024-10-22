CREATE TABLE airing_schedules (
    id SERIAL PRIMARY KEY,          -- Auto-incrementing primary key
    external_id INT NOT NULL,       -- Corresponds anilist id
    airing_at TIMESTAMP NOT NULL,         -- Corresponds to AiringAt
    episode INT NOT NULL,           -- Corresponds to Episode
    time_until_airing INT NOT NULL, -- Corresponds to TimeUntilAiring
    media_id INT NOT NULL,          -- Foreign key to the media table
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE -- Define foreign key constraint
);

CREATE TRIGGER set_updated_timestamp_airing
BEFORE UPDATE ON airing_schedules
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();