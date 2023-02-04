CREATE TABLE IF NOT EXISTS products (
  id serial PRIMARY KEY,
  name varchar,
  description varchar,
  price decimal,
  quantity integer
);

INSERT INTO
  products(name, description, price, quantity)
VALUES
  ('Shirt', 'Pretty', 29.50, 10),
  ('Laptop', 'fast', 1999.99, 2);