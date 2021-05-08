DROP DATABASE IF EXISTS restapi;
CREATE DATABASE restapi;
\c restapi;

/*
Users
*/
CREATE TABLE users (
id SERIAL PRIMARY KEY,
username VARCHAR NOT NULL,
password VARCHAR NOT NULL,
lastlogin INT,
active INT
);
