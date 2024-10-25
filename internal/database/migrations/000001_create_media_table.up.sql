CREATE TABLE media (
    id SERIAL PRIMARY KEY,          -- Auto-incrementing primary key
    external_id INT NOT NULL,       -- Corresponds anilist id
    site_url TEXT,                  -- Media site URL
    type VARCHAR(100) NOT NULL,
    format VARCHAR(100) NOT NULL,
    duration INT,                   -- Duration of the media in minutes
    episodes INT,                   -- Number of episodes (if applicable)
    cover_img TEXT, 
    banner_img TEXT,             -- 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
