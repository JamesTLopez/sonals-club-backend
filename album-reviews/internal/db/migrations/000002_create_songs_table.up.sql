
CREATE TABLE IF NOT EXISTS Songs(
   id serial PRIMARY KEY,
   user_id INT REFERENCES users(id),
   name VARCHAR (50) UNIQUE NOT NULL,
   description VARCHAR (50) NOT NULL
);

-- CREATE TABLE IF NOT EXISTS Songs(
--    id serial PRIMARY KEY,
--    name VARCHAR (50) UNIQUE NOT NULL,
-- );