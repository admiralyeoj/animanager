CREATE TABLE external_links (
    id SERIAL PRIMARY KEY,                            -- Auto-incrementing primary key
    site TEXT NOT NULL,                               -- Site name
    url VARCHAR(100) NOT NULL,                                -- URL link
    type VARCHAR(100) NOT NULL,                             -- external link type
    language VARCHAR(50),                                    -- Language of the link (nullable)
    site_id INT,                                      -- Site ID (can be a foreign key if needed)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Automatically set to current time on insert
);

CREATE TRIGGER set_updated_timestamp_external_links
BEFORE UPDATE ON external_links
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
