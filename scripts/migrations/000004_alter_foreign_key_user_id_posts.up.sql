ALTER TABLE posts MODIFY user_id BIGINT NOT NULL;
ALTER TABLE posts ADD CONSTRAINT fk_user_id_posts FOREIGN KEY (user_id) REFERENCES users(id);