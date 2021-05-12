CREATE TABLE IF NOT EXISTS days
(
    daysID int auto_increment NOT NULL,
    daysDate DATE NOT NULL,
    PRIMARY KEY (daysID)
);

CREATE TABLE IF NOT EXISTS tasks
(
    taskID int auto_increment NOT NULL,
    dayID int,
    taskDescription VARCHAR,
    PRIMARY KEY (taskID),
    FOREIGN KEY (dayID) REFERENCES days(daysID)
);
