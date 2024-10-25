CREATE TABLE external_link (
    id SERIAL PRIMARY KEY,  -- Auto-incrementing primary key
    external_id BIGINT NOT NULL UNIQUE,       -- Corresponds anilist id
    site_id BIGINT, -- Site ID 
    name TEXT NOT NULL, -- Site name
    url text NOT NULL, -- URL link
    type VARCHAR(100) NOT NULL, -- external link type
    language VARCHAR(50), -- Language of the link (nullable)
    media_id BIGINT NOT NULL,                            
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    FOREIGN KEY (media_id) REFERENCES media(id) ON DELETE CASCADE -- Define foreign key constraint
);