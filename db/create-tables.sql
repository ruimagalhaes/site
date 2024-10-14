DROP TABLE IF EXISTS article;

CREATE TABLE article (
  id     INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  title  TINYTEXT NOT NULL,
  body   TEXT NOT NULL
);

INSERT INTO article (title, body)
VALUES
  ('First Article', 'This is the content of the first article.'),
  ('Second Article', 'Content for the second article goes here.'),
  ('Third Article', 'And here is the content for the third article.');

DROP TABLE IF EXISTS user;

CREATE TABLE user (
  id        INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
  username  VARCHAR(255) NOT NULL,
  password  VARCHAR(255) NOT NULL
);