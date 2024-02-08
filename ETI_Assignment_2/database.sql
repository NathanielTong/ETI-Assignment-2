CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'localhost';

CREATE DATABASE IF NOT EXISTS task_management_db

USE task_management_db;

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    priority INT,
    deadline DATE,
    status VARCHAR(20)
);