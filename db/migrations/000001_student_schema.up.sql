-- Create the schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS student_schema;

-- Create the students table within the student_schema schema (without the age column)
CREATE TABLE IF NOT EXISTS student_schema.students (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Insert sample data into the students table in the student_schema schema (without the age column)
INSERT INTO student_schema.students (first_name, last_name, email)
VALUES 
('John', 'Doe', 'john.doe@example.com'),
('Jane', 'Smith', 'jane.smith@example.com'),
('Michael', 'Johnson', 'michael.johnson@example.com'),
('Emily', 'Davis', 'emily.davis@example.com'),
('David', 'Martinez', 'david.martinez@example.com'),
('Sarah', 'Garcia', 'sarah.garcia@example.com'),
('James', 'Wilson', 'james.wilson@example.com'),
('Linda', 'Moore', 'linda.moore@example.com'),
('Robert', 'Taylor', 'robert.taylor@example.com'),
('Jennifer', 'Anderson', 'jennifer.anderson@example.com'),
('William', 'Thomas', 'william.thomas@example.com'),
('Mary', 'Jackson', 'mary.jackson@example.com'),
('Daniel', 'White', 'daniel.white@example.com'),
('Sophia', 'Harris', 'sophia.harris@example.com'),
('Joseph', 'Clark', 'joseph.clark@example.com'),
('Olivia', 'Lewis', 'olivia.lewis@example.com'),
('Charles', 'Young', 'charles.young@example.com'),
('Charlotte', 'Allen', 'charlotte.allen@example.com'),
('Benjamin', 'Scott', 'benjamin.scott@example.com'),
('Ava', 'King', 'ava.king@example.com');
