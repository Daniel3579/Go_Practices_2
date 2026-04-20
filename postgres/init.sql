CREATE TABLE auth (
  username VARCHAR(255) NOT NULL PRIMARY KEY,
  hash     VARCHAR(255)
);

CREATE TABLE task (
  id          SERIAL PRIMARY KEY,
  username    VARCHAR(50) NOT NULL REFERENCES auth(username) ON DELETE CASCADE,
  title       VARCHAR(50) NOT NULL,
  description VARCHAR(255),
  due_date    DATE,
  done        BOOLEAN DEFAULT FALSE
);