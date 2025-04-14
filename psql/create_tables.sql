CREATE TABLE Products(
    id serial PRIMARY KEY NOT NULL,
    sourceId INT NOT NULL UNIQUE, 
    title VARCHAR(100) NOT NULL,
    price  DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL
)