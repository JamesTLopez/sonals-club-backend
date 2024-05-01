
CREATE TABLE IF NOT EXISTS Samples(
   id serial PRIMARY KEY,
   user_id INT REFERENCES users(id),
   song_id INT REFERENCES songs(id),
   name VARCHAR (50) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- TODO: onces file upload is created, add new column for s3