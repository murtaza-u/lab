# Securing MYSQL

## Lock user account after 3 failed login attempts

```sql
CREATE USER IF NOT EXISTS 'murtaza'@'%'
IDENTIFIED WITH caching_sha2_password BY 'password'
FAILED_LOGIN_ATTEMPTS 3 PASSWORD_LOCK_TIME 2;
```

## Require client validation

```sql
CREATE USER IF NOT EXISTS 'murtaza'@'%'
IDENTIFIED WITH caching_sha2_password BY 'password'
REQUIRE X509;
```

```sql
GRANT ALL ON mydb.* TO 'murtaza'@'%';
```

## Verify SSL connection

```sql
SELECT * FROM performance_schema.session_status
WHERE VARIABLE_NAME IN ('Ssl_version','Ssl_cipher');
```

## Disable remote root logins

```sql
DELETE FROM mysql.user
WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');
```

```sql
FLUSH PRIVILEGES;
```
