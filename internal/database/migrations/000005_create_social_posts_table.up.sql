CREATE TABLE social_post (
    id SERIAL PRIMARY KEY,  -- Auto-incrementing primary key
    post_id text NOT NULL, -- Corresponds anilist id
    airing_schedule_id BIGINT NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,  -- Automatically set to current time on insert
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (airing_schedule_id) REFERENCES airing_schedule(id) ON DELETE CASCADE -- Define foreign key constraint
);