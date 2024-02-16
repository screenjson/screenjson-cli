# ScreenJSON Command Line Tools


## screenjson-db

To build the Docker image:

```bash
docker build -f docker/screenjson-db/Dockerfile -t screenjson-db:latest .
docker push screenjson/screenjson-db:latest
```

You can use the date command with the `--utc` or `-u` flag to get the current time in UTC, and use `+%Y-%m-%dT%H:%M:%SZ` to format the date according to the ISO 8601 standard.

```bash
date -u +%Y-%m-%dT%H:%M:%SZ
```

On Windows:

```
Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ" -AsUTC
$date = Get-Date -Format "yyyy-MM-ddTHH:mm:ssZ" -AsUTC
```

To get the current datetime in the `YYYY-MM-DD HH:mm:ss` format, which is commonly used in SQL databases, you can also use the date command on Unix-like systems or PowerShell/cmdlet on Windows. Use the date command with the `+%Y-%m-%d %H:%M:%S` format string:

```bash
date "+%Y-%m-%d %H:%M:%S"
```

On Windows:

```
Get-Date -Format "yyyy-MM-dd HH:mm:ss"
powershell -Command "Get-Date -Format 'yyyy-MM-dd HH:mm:ss'"
```

### Insert into Cassandra

This assumes:

- You have a Cassandra database named `screenplays`, containing a collection `imports` which has a field named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="cassandra" -uri="cassandra://username:password@localhost:9042/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

```bash
docker run screenjson/screenjson-db -file="./screenplay.json" -engine="cassandra" -uri="cassandra://username:password@localhost:9042/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into Couchbase

This assumes:

- You have a Couchbase database named `screenplays`, containing a collection `imports` which has a field named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="couchbase" -uri="couchbase://localhost" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

```bash
docker run screenjson/screenjson-db -file="./screenplay.json" -engine="couchbase" -uri="couchbase://localhost" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into CouchDB

This assumes:

- You have a CouchDB database named `screenplays`, containing a collection `imports` which has a field named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="couchdb" -uri="http://username:password@localhost:5984/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into AWS DynamoDB

This assumes:

- You have a DynamoDB database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="dynamodb" -uri="https://dynamodb.us-west-2.amazonaws.com" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into Elasticsearch

This assumes:

- You have an Elasticsearch database named `screenplays`, containing an index `imports` which has a field mapping named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="elasticsearch" -uri="http://username:password@localhost:9200/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into MongoDB

- You have a MongoDB database named `screenplays`, containing a collection `imports` which has a field named `screenjson`;
- You need to store a `custom_key` and `another_col` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="mongodb" -uri="mongodb://username:password@localhost:27017?authSource=admin" -database="screenplays" -table="imports" -field="screenjson" -additional="custom_key=12345,another_col=something_else"
```

### Insert into MS SQL Server

This assumes:

- You have a SQL Server database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="mssql" -uri="sqlserver://username:password@localhost:1433;databaseName=screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into MySQL/MariaDB

This assumes:

- You have a MySQL database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="mysql" -uri="mysql://username:password@localhost:3306/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into Oracle

This assumes:

- You have an Oracle database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="oracle" -uri="jdbc:oracle:thin:username/password@//localhost:1521/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into PostgreSQL

This assumes:

- You have a PostgreSQL database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="postgresql" -uri="postgresql://username:password@localhost:5432/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

### Insert into Redis

This assumes:

- You have a Redis server up and running.
- You want to store a key  `screenjson`.
- You don't want to store any other keys.

```bash
screenjson-db -file="./screenplay.json" -engine="redis" -uri="redis://username:password@localhost:6379/0" -database="0" -field="screenjson"
```

### Insert into SQLite

This assumes:

- You have a SQLite database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="sqlite" -uri="sqlite:///path/to/screenplays.db" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

## screenjson-decrypt

To build the Docker image:

```bash
docker build -f docker/screenjson-decrypt/Dockerfile -t screenjson-decrypt:latest .
docker push screenjson/screenjson-decrypt:latest
```

## screenjson-encrypt

To build the Docker image:

```bash
docker build -f docker/screenjson-encrypt/Dockerfile -t screenjson-encrypt:latest .
docker push screenjson/screenjson-encrypt:latest
```

## screenjson-export

To build the Docker image:

```bash
docker build -f docker/screenjson-export/Dockerfile -t screenjson-export:latest .
docker push screenjson/screenjson-export:latest
```

## screenjson-import

To build the Docker image:

```bash
docker build -f docker/screenjson-import/Dockerfile -t screenjson-import:latest .
docker push screenjson/screenjson-import:latest
```

## screenjson-validate

To build the Docker image:

```bash
docker build -f docker/screenjson-validate/Dockerfile -t screenjson-validate:latest .
docker push screenjson/screenjson-validate:latest
```

### Validate a JSON file

```bash
screenjson-validate my_screenplay.json
```

```bash
docker run screenjson/screenjson-validate my_screenplay.json
```
