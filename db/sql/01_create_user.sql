CREATE USER 'writer'@'localhost' IDENTIFIED BY 'writer_password';
CREATE USER 'reader'@'localhost' IDENTIFIED BY 'reader_password';
GRANT ALL PRIVILEGES ON customer_db.* TO 'writer'@'localhost';
GRANT USAGE ON *.* TO 'reader'@'localhost';
GRANT SELECT ON customer_db.* TO 'reader'@'localhost';
FLUSH PRIVILEGES;
USE customer_db