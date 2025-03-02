SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Create Database
-- ----------------------------
CREATE DATABASE IF NOT EXISTS `main` DEFAULT CHARACTER SET utf8mb4;

-- ----------------------------
-- Use Database
-- ----------------------------

USE `main`;

SET NAMES utf8mb4;

-- Create the User table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    `role` VARCHAR(50),
    avatar LONGTEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY  `username_idx1` (`username`) USING BTREE
);

-- Create the Task table
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50),
    user_id INT,  -- Foreign key to User, but no constraint
    archived BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create the Comment table
CREATE TABLE IF NOT EXISTS comments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    content TEXT NOT NULL,
    task_id INT,  -- Foreign key to Task, but no constraint
    user_id INT,  -- Foreign key to User, but no constraint
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY  `task_id_comments_idx` (`task_id`) USING BTREE,
    KEY  `user_id_comments_idx` (`user_id`) USING BTREE
);

-- Create the ChangeLog table
CREATE TABLE IF NOT EXISTS changelogs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    task_id INT,
    old_value LONGTEXT,
    new_value LONGTEXT,
    user_id INT,
    `action` VARCHAR(255) NOT NULL,
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    KEY  `action_changelogs_idx` (`action`) USING BTREE,
    KEY  `user_id_changelogs_idx` (`user_id`) USING BTREE
);


-- Example data

-- Insert Users
INSERT INTO users (username, password, role) VALUES
    ('admin', '$2a$10$VL0rDMqDczc0mOXL1Uuybu9QAx.CYOPilt.WftWH5xWqPNspPaLfy', 'admin'),
    ('user1', '$2a$10$VL0rDMqDczc0mOXL1Uuybu9QAx.CYOPilt.WftWH5xWqPNspPaLfy', 'user'),
    ('user2', '$2a$10$VL0rDMqDczc0mOXL1Uuybu9QAx.CYOPilt.WftWH5xWqPNspPaLfy', 'user');

-- Insert Tasks
INSERT INTO tasks (`name`, description, status, user_id, archived) VALUES
    ('Task 1', 'This is the first task.', 'todo', 1, FALSE),
    ('Task 2', 'This task is in progress.', 'in_progress', 2, FALSE),
    ('Task 3', 'This is a completed task.', 'done', 2, TRUE),  -- Archived task
    ('Task 4', 'Another task assigned to admin.', 'todo', 1, FALSE);

-- Insert Comments
INSERT INTO comments (content, task_id, user_id) VALUES
    ('This is a comment on Task 1.', 1, 2),
    ('This is another comment on Task 1.', 1, 3),
    ('This is a comment on Task 2.', 2, 1),
    ('Task 3 comment by a user.', 3, 2),
    ('Another comment on Task 4.', 4, 1);


