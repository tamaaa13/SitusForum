CREATE TABLE IF NOT EXISTS comments (
  id int NOT NULL AUTO_INCREMENT,
  post_id int NOT NULL,
  user_id BIGINT NOT NULL,
  comments_content LONGTEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by LONGTEXT NOT NULL,
  updated_by LONGTEXT NOT NULL,

  PRIMARY KEY (id),
  CONSTRAINT fk_post_id_comments FOREIGN KEY (post_id) REFERENCES posts(id),
  CONSTRAINT fk_user_id_comments FOREIGN KEY (user_id) REFERENCES users(id)
);