CREATE TABLE airing_schedule (
    id SERIAL PRIMARY KEY,          -- Auto-incrementing primary key
    external_id BIGINT NOT NULL UNIQUE,       -- Corresponds anilist id
    airing_at BIGINT NOT NULL,         -- Corresponds to AiringAt
    episode INT NOT NULL,           -- Corresponds to Episode
    media_id BIGINT NOT NULL,          -- Foreign key to the media table
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE -- Define foreign key constraint
);
