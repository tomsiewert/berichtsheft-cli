CREATE TABLE IF NOT EXISTS tasks
(
    taskID int auto_increment NOT NULL,
    taskDate DATE NOT NULL,
    taskDescription VARCHAR,
    PRIMARY KEY (taskID)
);
