CREATE TABLE tasks (
    id           NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    title        VARCHAR2(255) NOT NULL,
    description  VARCHAR2(1000),
    status       VARCHAR2(50) CHECK (status IN ('pending', 'completed')),
    due_date     DATE,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER update_task_timestamp
BEFORE UPDATE ON tasks
FOR EACH ROW
BEGIN
    :NEW.updated_at := CURRENT_TIMESTAMP;
END;
/




CREATE TABLE users (
    id NUMBER GENERATED BY DEFAULT ON NULL AS IDENTITY PRIMARY KEY,
    username VARCHAR2(50) NOT NULL,
    email VARCHAR2(100) NOT NULL UNIQUE,
    password VARCHAR2(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER users_before_update
BEFORE UPDATE ON users
FOR EACH ROW
BEGIN
    :new.updated_at := CURRENT_TIMESTAMP;
END;
/