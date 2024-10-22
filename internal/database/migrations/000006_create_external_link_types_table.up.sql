CREATE TABLE external_link_types (
    id SERIAL PRIMARY KEY,                            -- Auto-incrementing primary key
    type TEXT NOT NULL,                               -- Site name
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP   -- Automatically set to current time on insert
);

CREATE TRIGGER set_updated_timestamp_external_link_types
BEFORE UPDATE ON external_link_types
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();


INSERT INTO external_link_types(type) VALUES ('Info'), ('Streaming'), ('Social');