
CREATE TABLE IF NOT EXISTS Songs(
   id serial PRIMARY KEY,
   user_id INT REFERENCES users(id),
   song_name VARCHAR (50) NOT NULL,
   labels VARCHAR (50),
   description VARCHAR (50) NOT NULL,
   duration INT,
   color VARCHAR (50),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- TODO: onces file upload is created, add new column for s3 image