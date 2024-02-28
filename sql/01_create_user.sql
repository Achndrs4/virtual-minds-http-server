CREATE USER 'local'@'localhost' IDENTIFIED BY 'local_password';
GRANT ALL PRIVILEGES ON customer_db.* TO 'local'@'localhost';
FLUSH PRIVILEGES;
USE customer_db