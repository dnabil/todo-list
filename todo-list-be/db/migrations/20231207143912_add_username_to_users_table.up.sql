ALTER TABLE users 
    ADD COLUMN username VARCHAR(36) NOT NULL,
    ADD CONSTRAINT uq_user_username UNIQUE (username);