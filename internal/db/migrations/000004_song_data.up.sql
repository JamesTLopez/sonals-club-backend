
CREATE TABLE IF NOT EXISTS Song_Data(
   id serial PRIMARY KEY,
   song_id INT REFERENCES songs(id),
   tuning VARCHAR(50),
   key VARCHAR(50),
   capo VARCHAR(50),
   difficulty VARCHAR(50),
   chords VARCHAR(50),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

