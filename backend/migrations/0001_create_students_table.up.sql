CREATE TABLE students
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255) NOT NULL,
    group_id   INT          NOT NULL,
    birthday   DATE         NOT NULL,
    comments   TEXT
);