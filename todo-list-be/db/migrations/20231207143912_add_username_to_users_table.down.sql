ALTER TABLE users 
    DROP INDEX uq_user_username,
    DROP COLUMN username;