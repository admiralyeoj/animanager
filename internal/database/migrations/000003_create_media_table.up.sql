CREATE TABLE media (
    id SERIAL PRIMARY KEY,          -- Auto-incrementing primary key
    external_id INT NOT NULL,       -- Corresponds anilist id
    site_url TEXT,                  -- Media site URL
    type_id INT NOT NULL,
    format_id INT NOT NULL,
    duration INT,                   -- Duration of the media in minutes
    episodes INT,                   -- Number of episodes (if applicable)
    cover_image TEXT, 
    banner_image TEXT,             -- 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (type_id) REFERENCES media_type(id) ON DELETE CASCADE, -- Define foreign key constraint
    FOREIGN KEY (format_id) REFERENCES media_format(id) ON DELETE CASCADE -- Define foreign key constraint
);

CREATE TRIGGER set_updated_timestamp_media
BEFORE UPDATE ON media
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();