DROP TABLE IF EXISTS portfolio_hub.users
CREATE TABLE portfolio_hub.users (
  id VARCHAR(128) PRIMARY KEY NOT NULL,
  name VARCHAR(64) NOT NULL,
  photo_url VARCHAR(256) NOT NULL
);

DROP TABLE IF EXISTS portfolio_hub.portfolios
CREATE TABLE portfolio_hub.portfolios (
  id SERIAL PRIMARY KEY,
  CONSTRAINT uid
    FOREIGN KEY (id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  title VARCHAR(64) NOT NULL,
  portfolio_url VARCHAR(256) NOT NULL
);

DROP TABLE IF EXISTS portfolio_hub.likes
CREATE TABLE portfolio_hub.likes (
  id SERIAL PRIMARY KEY,
  CONSTRAINT uid
    FOREIGN KEY (id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT portfolio_id
    FOREIGN KEY (id)
    REFERENCES users (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
);