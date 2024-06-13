
CREATE TABLE IF NOT EXISTS Samples(
   id serial PRIMARY KEY,
   user_id VARCHAR(255) REFERENCES users(spotify_id),
   song_id INT REFERENCES songs(id),
   sample_name VARCHAR (50) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- TODO: onces file upload is created, add new column for s3