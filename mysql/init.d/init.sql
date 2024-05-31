-- disable remote root logins
DELETE FROM mysql.user
WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');


-- create non-root user
CREATE USER IF NOT EXISTS 'murtaza'@'%'
IDENTIFIED WITH caching_sha2_password BY 'password'
REQUIRE X509;


GRANT ALL ON assemble.* TO 'murtaza'@'%';


FLUSH PRIVILEGES;
