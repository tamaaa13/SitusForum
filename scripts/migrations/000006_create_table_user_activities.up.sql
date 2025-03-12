CREATE TABLE IF NOT EXISTS users_activities (
  id INT NOT NULL AUTO_INCREMENT,
  post_id INT NOT NULL,
  user_id BIGINT NOT NULL,
  is_liked BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by LONGTEXT NOT NULL,
  updated_by LONGTEXT NOT NULL,

  PRIMARY KEY (id),
  CONSTRAINT fk_post_id_users_activities FOREIGN KEY (post_id) REFERENCES posts(id),
  CONSTRAINT fk_user_id_users_activities FOREIGN KEY (user_id) REFERENCES users(id)
);