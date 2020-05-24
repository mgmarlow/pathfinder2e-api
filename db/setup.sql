DROP TABLE IF EXISTS beasts;

CREATE TABLE beasts (
  id serial primary key,
  name VARCHAR(200) UNIQUE NOT NULL,
  xp int NOT NULL
  -- sr int,
  -- initiative VARCHAR(50),
  -- senses VARCHAR(50),
  -- ac VARCHAR(50),
  -- hp VARCHAR(50),
  -- strength int NOT NULL,
  -- dexterity int NOT NULL,
  -- constitution int NOT NULL,
  -- intelligence int NOT NULL,
  -- wisdom int NOT NULL,
  -- charisma int NOT NULL,
  -- bab VARCHAR(50) NOT NULL,
  -- cmb VARCHAR(50) NOT NULL,
  -- cmd int NOT NULL
);
