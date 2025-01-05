-- Create a new UTF-8 `apptrak` database.
CREATE DATABASE apptrak CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Switch to using the `apptrak` database.
USE apptrak;

CREATE TABLE applications (
    id SERIAL PRIMARY KEY,                -- Unique ID for each application
    name VARCHAR(255) NOT NULL,            -- Name of the application (e.g., job or school)
    company_name VARCHAR(255) NOT NULL,    -- The company or school name
    position VARCHAR(255),                 -- The position applied for (for job applications)
    status VARCHAR(100) NOT NULL CHECK (status IN ('Applied', 'Interview', 'Offered', 'Rejected', 'N/A')),  -- Enforced status options
    application_date DATE NOT NULL,        -- Date when the application was submitted
    interview_date DATE,                  -- Date for the interview, if applicable
    notes TEXT,                            -- Additional notes related to the application
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the record was created
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- Timestamp when the record was last updated
);

-- Add an index on the created_at column.
CREATE INDEX idx_applications_created ON applications(created_at);


-- Insert dummy records into the `applications` table.
INSERT INTO applications (name, company_name, position, status, application_date, interview_date, notes) 
VALUES 
    ('Software Engineer Application', 'TechCorp', 'Software Engineer', 'Applied', '2025-01-01', NULL, 'Excited about their cutting-edge projects.'),
    ('Data Analyst Application', 'DataInsights', 'Data Analyst', 'Interview', '2024-12-20', '2025-01-05', 'Phone screen went well; waiting for next steps.'),
    ('Graduate Program Application', 'Top University', 'Computer Science Graduate Program', 'Rejected', '2024-11-15', NULL, 'Rejected due to limited spots; considering reapplying next year.');

