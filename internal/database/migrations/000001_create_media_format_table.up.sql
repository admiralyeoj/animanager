-- Create the reusable function
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the table
CREATE TABLE media_format (
    id SERIAL PRIMARY KEY,                            -- Auto-incrementing primary key
    name TEXT NOT NULL,                               -- Name of the format
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Automatically set to current time on insert
);

CREATE TRIGGER set_updated_timestamp_media_format
BEFORE UPDATE ON media_format
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

INSERT INTO media_format (name) VALUES 
('TV'), 
('TV_SHORT'), 
('MOVIE'), 
('SPECIAL'), 
('OVA'), 
('ONA'), 
('MUSIC'), 
('MANGA'), 
('NOVEL'), 
('ONE_SHOT');