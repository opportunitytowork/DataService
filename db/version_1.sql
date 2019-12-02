CREATE USER IF NOT EXISTS 'dataservice'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON otw . * TO 'dataservice'@'localhost';

-- Create Database
CREATE DATABASE IF NOT EXISTS otw CHARACTER SET utf8 COLLATE utf8_unicode_ci;

-- Create tables
