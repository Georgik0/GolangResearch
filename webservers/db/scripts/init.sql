CREATE TABLE my_user (
    id serial PRIMARY KEY,
    username varchar UNIQUE NOT NULL
);
