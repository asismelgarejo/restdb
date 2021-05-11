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
admin INT,
lastlogin INT64,
active INT
);

INSERT INTO users (username, password, admin, active) VALUES ('admin', 'admin', 1, 1);
