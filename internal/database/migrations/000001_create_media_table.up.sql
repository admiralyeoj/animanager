CREATE TABLE media (
    id SERIAL PRIMARY KEY,                         -- Auto-incrementing primary key
    external_id BIGINT NOT NULL UNIQUE,            -- Corresponds to Anilist ID and must be unique
    site_url TEXT,                                 -- Media site URL
    type VARCHAR(100) NOT NULL,                    -- Media type
    format VARCHAR(100) NOT NULL,                  -- Format of the media
    duration INT,                                  -- Duration of the media in minutes
    episodes INT,                                  -- Number of episodes (if applicable)
    cover_img TEXT,                                -- Cover image URL
    banner_img TEXT,                               -- Banner image URL
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,-- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Automatically set to current time on update
);
