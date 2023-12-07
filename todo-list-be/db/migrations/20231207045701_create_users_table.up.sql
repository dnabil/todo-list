CREATE TABLE users
(
  id         INT          NOT NULL AUTO_INCREMENT,
  created_at timestamp    NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp    NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  email      VARCHAR(255) NOT NULL UNIQUE ,
  password   VARCHAR(255) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT uq_user_email UNIQUE (email)
);