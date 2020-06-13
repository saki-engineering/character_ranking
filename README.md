# mysql
## sampledb database
### votes table
作成クエリは以下の通り。
```mysql
CREATE TABLE IF NOT EXISTS votes(
		chara       VARCHAR(20) NOT NULL,
		user        INT UNSIGNED NOT NULL,
		created_at  DATETIME NOT NULL,
		ip          VARCHAR(50),
		FOREIGN KEY (user) REFERENCES users (id)
	);
```
userキーは、usersテーブルのidの外部キーである。

### users table
作成クエリは以下の通り。
```mysql
CREATE TABLE IF NOT EXISTS users(
		id        INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
		age       INT NOT NULL,
		gender    INT NOT NULL,
		address   INT NOT NULL
	);
```

### charas table
作成クエリは以下の通り。
```mysql
CREATE TABLE IF NOT EXISTS charas(
		id          INT UNSIGNED NOT NULL,
		chara       VARCHAR(20) NOT NULL
	);
```

### adminusers
作成クエリは以下の通り。
```mysql
CREATE TABLE IF NOT EXISTS adminusers(
		userid           VARCHAR(50) NOT NULL PRIMARY KEY,
		hashedpassword   VARCHAR(500) NOT NULL,
		auth             INT NOT NULL
	);
```