CREATE TABLE external_links (
    id SERIAL PRIMARY KEY,                            -- Auto-incrementing primary key
    site TEXT NOT NULL,                               -- Site name
    url TEXT NOT NULL,                                -- URL link
    type_id INT NOT NULL,                             -- Foreign key to external_link_types
    language TEXT,                                    -- Language of the link (nullable)
    site_id INT,                                      -- Site ID (can be a foreign key if needed)
    color TEXT,                                       -- Color associated with the link (nullable)
    icon TEXT,                                        -- Icon associated with the link (nullable)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    FOREIGN KEY (type_id) REFERENCES external_link_types(id) ON DELETE CASCADE  -- Foreign key constraint
);

CREATE TRIGGER set_updated_timestamp_external_links
BEFORE UPDATE ON external_links
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
