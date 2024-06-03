CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   spotify_id VARCHAR(255),
   display_name VARCHAR (50) UNIQUE NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);