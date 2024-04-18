
CREATE TABLE IF NOT EXISTS Songs(
   id serial PRIMARY KEY,
   user_id INT REFERENCES users(id),
   name VARCHAR (50) UNIQUE NOT NULL,
   labels VARCHAR (50),
   description VARCHAR (50) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
