ALTER TABLE posts DROP FOREIGN KEY fk_user_id_posts;

ALTER TABLE posts MODIFY user_id INT NOT NULL;