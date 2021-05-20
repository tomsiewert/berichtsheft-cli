CREATE TABLE IF NOT EXISTS tasks
(
    taskID int auto_increment NOT NULL,
    taskDate DATE NOT NULL,
    taskTime BIGINT(8) NOT NULL,
    taskDescription VARCHAR,
    PRIMARY KEY (taskID)
);
