CREATE TABLE scheduler (
    id SERIAL PRIMARY KEY,                     -- Unique identifier for each job
    job_name VARCHAR(255) NOT NULL,            -- Name of the job
    cron_expression VARCHAR(255) NOT NULL,     -- Cron expression for scheduling
    function_name VARCHAR(255) NOT NULL,       -- Identifier of the function to execute
    is_active BOOLEAN DEFAULT true,            -- Enable or disable the job
    last_run TIMESTAMPTZ,                        -- Last execution time
    next_run TIMESTAMPTZ,                        -- Optional: next scheduled run time
    params JSONB,                              -- Additional parameters in JSON format
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

INSERT INTO scheduler (job_name, cron_expression, function_name, is_active, params, created_at, updated_at)
VALUES 
    ('Test Cron Job', '*/30 * * * * *', 'test', true, NULL, NOW(), NOW()),
    ('Import Scheduled Anime', '0 */2 * * * *', 'importScheduledAnime', true, NULL, NOW(), NOW())
    ('Announce New anime airing', '0 */1 * * * *', 'announceNewAnime', true, NULL, NOW(), NOW());