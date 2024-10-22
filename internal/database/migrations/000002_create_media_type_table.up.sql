-- Create the table
CREATE TABLE media_type (
    id SERIAL PRIMARY KEY,                            -- Auto-incrementing primary key
    name TEXT NOT NULL,                               -- Name of the type
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Automatically set to current time on insert
);

CREATE TRIGGER set_updated_timestamp_media_type
BEFORE UPDATE ON media_type
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- Insert data into the table
INSERT INTO media_type (name) VALUES 
('Anime'), 
('Manga');
