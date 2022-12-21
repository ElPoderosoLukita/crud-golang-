CREATE DATABASE IF NOT EXISTS crud;

CREATE TABLE IF NOT EXISTS users(
    ID INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    age INT UNSIGNED NOT NULL
);
