SET CHARSET UTF8;
DROP DATABASE IF EXISTS todo_db;
CREATE DATABASE todo_db;

USE todo_db;

DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks
(
  id           VARCHAR(64)  NOT NULL,
  title        varchar(256) NOT NULL,
  created_at   TIMESTAMP    NOT NULL,
  deadline     TIMESTAMP    NOT NULL,
  is_completed BOOLEAN      NOT NULL,
  description  varchar(256),
  CONSTRAINT PK_task_id PRIMARY KEY (id)
);
