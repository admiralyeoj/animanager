CREATE TABLE scheduler (
    id SERIAL PRIMARY KEY,                     -- Unique identifier for each job
    job_name VARCHAR(255) NOT NULL,            -- Name of the job
    cron_expression VARCHAR(255) NOT NULL,     -- Cron expression for scheduling
    function_name VARCHAR(255) NOT NULL,       -- Identifier of the function to execute
    is_active BOOLEAN DEFAULT true,            -- Enable or disable the job
    last_run TIMESTAMP,                        -- Last execution time
    next_run TIMESTAMP,                        -- Optional: next scheduled run time
    params JSONB,                              -- Additional parameters in JSON format
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);