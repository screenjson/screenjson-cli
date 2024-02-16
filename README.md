# ScreenJSON Command Line Tools

This small suite of command line programs written in Go are utilities to help processing between formats and make life easier.

## Example Workflow Cbains

### Ingest from Final Draft Pro to a database

```bash
# Import Final Draft Pro file
screenjson-import -in="screenplay.fdx" -out="screenplay.json"

# Check it's formatted correctly
screenjson-validate screenplay.json 

# Encrypt the key parts (location, action, dialogue etc)
screenjson-encrypt -in="screenplay.json" -out="encrypted.json" -password="12345" 

# Save it into Elasticsearch
screenjson-db -file="./screenplay.json" -engine="elasticsearch" -uri="http://username:password@localhost:9200/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")" 
```

### Export from a database back to a PDF

```bash
# Export from MongoDB/Elasticsearch
# ---> your logic, save to stored-encrypted.json

# Check it's formatted correctly
screenjson-validate stored-encrypted.json

# Decrypt the key parts (location, action, dialogue etc)
screenjson-decrypt -in="stored-encrypted.json" -out="decrypted.json" -password="12345" 

# Export to PDF and Fountain
screenjson-export -in "decrypted.json" -out="exported-for-print.pdf"
screenjson-export -in "decrypted.json" -out="exported-for-highland.fountain"
```

## screenjson-db

This utility saves the contents of a ScreenJSON file into most common databases.

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

##### Insert into Cassandra

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

#### Insert into Couchbase

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

#### Insert into CouchDB

This assumes:

- You have a CouchDB database named `screenplays`, containing a collection `imports` which has a field named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="couchdb" -uri="http://username:password@localhost:5984/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

#### Insert into AWS DynamoDB

This assumes:

- You have a DynamoDB database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="dynamodb" -uri="https://dynamodb.us-west-2.amazonaws.com" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

#### Insert into Elasticsearch

This assumes:

- You have an Elasticsearch database named `screenplays`, containing an index `imports` which has a field mapping named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="elasticsearch" -uri="http://username:password@localhost:9200/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

#### Insert into MongoDB

- You have a MongoDB database named `screenplays`, containing a collection `imports` which has a field named `screenjson`;
- You need to store a `custom_key` and `another_col` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="mongodb" -uri="mongodb://username:password@localhost:27017?authSource=admin" -database="screenplays" -table="imports" -field="screenjson" -additional="custom_key=12345,another_col=something_else"
```

#### Insert into MS SQL Server

This assumes:

- You have a SQL Server database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="mssql" -uri="sqlserver://username:password@localhost:1433;databaseName=screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

#### Insert into MySQL/MariaDB

This assumes:

- You have a MySQL database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="mysql" -uri="mysql://username:password@localhost:3306/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

#### Insert into Oracle

This assumes:

- You have an Oracle database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="oracle" -uri="jdbc:oracle:thin:username/password@//localhost:1521/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

#### Insert into PostgreSQL

This assumes:

- You have a PostgreSQL database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="postgresql" -uri="postgresql://username:password@localhost:5432/screenplays" -database="screenplays" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

#### Insert into Redis

This assumes:

- You have a Redis server up and running.
- You want to store a key  `screenjson`.
- You don't want to store any other keys.

```bash
screenjson-db -file="./screenplay.json" -engine="redis" -uri="redis://username:password@localhost:6379/0" -database="0" -field="screenjson"
```

#### Insert into SQLite

This assumes:

- You have a SQLite database named `screenplays`, containing a table `imports` which has a JSON column named `screenjson`;
- You need to store a `genre_id` with the record (optional);
- You want to store timestamps `created_at` and `updated_at` with the record (optional);

```bash
screenjson-db -file="./screenplay.json" -engine="sqlite" -uri="sqlite:///path/to/screenplays.db" -table="imports" -field="screenjson" -additional="genre_id=12345,created_at=$(date "+%Y-%m-%d %H:%M:%S"),updated_at=$(date "+%Y-%m-%d %H:%M:%S")"
```

## screenjson-decrypt

This utility decrypts the content values of scene elements which have been encrypted with a shared secret password because they are sensitive, for example: locations, characters, dialogue, action, and so on. 

To build the Docker image:

```bash
docker build -f docker/screenjson-decrypt/Dockerfile -t screenjson-decrypt:latest .
docker push screenjson/screenjson-decrypt:latest
```

Decrypt a file:

```bash
screenjson-decrypt -in="encrypted.json" -out="clear.json" -password="secret"
```

```bash
docker run screenjson/screenjson-decrypt -in="encrypted.json" -out="clear.json" -password="secret"
```

## screenjson-encrypt

This utility encrypts the content values of scene elements with a shared secret password because they are sensitive, for example: locations, characters, dialogue, action, and so on. The JSON structure is preserved.

To build the Docker image:

```bash
docker build -f docker/screenjson-encrypt/Dockerfile -t screenjson-encrypt:latest .
docker push screenjson/screenjson-encrypt:latest
```

Encrypt a file:

```bash
screenjson-encrypt -in="screenplay.json" -out="encrypted.json" -password="secret"
```

```bash
docker run screenjson/screenjson-encrypt -in="screenplay.json" -out="encrypted.json" -password="secret"
```

## screenjson-export

This utility converts a ScreenJSON data file into different screenwriting formats, for example: Final Draft Pro, Fade In Pro, Fountain markdown etc.

To build the Docker image:

```bash
docker build -f docker/screenjson-export/Dockerfile -t screenjson-export:latest .
docker push screenjson/screenjson-export:latest
```

#### Convert/export ScreenJSON to PDF

```bash
screenjson-export -in="screenplay.json" -out="screenplay.pdf" -password="secret"
```

```bash
docker run screenjson/screenjson-export -in="screenplay.json" -out="screenplay.pdf" -password="secret"
```

#### Convert/export ScreenJSON to Final Draft Pro

```bash
screenjson-export -in="screenplay.json" -out="screenplay.fdx"
```

```bash
docker run screenjson/screenjson-export -in="screenplay.json" -out="screenplay.fdx"
```

#### Convert/export ScreenJSON to FadeIn Pro

```bash
screenjson-export -in="screenplay.json" -out="screenplay.fadein"
```

```bash
docker run screenjson/screenjson-export -in="screenplay.json" -out="screenplay.fadein"
```

#### Convert/export ScreenJSON to Fountain

```bash
screenjson-export -in="screenplay.json" -out="screenplay.fountain"
```

```bash
docker run screenjson/screenjson-export -in="screenplay.json" -out="screenplay.fountain"
```

#### Convert/export ScreenJSON to Celtx

```bash
screenjson-export -in="screenplay.json" -out="screenplay.celtx"
```

```bash
docker run screenjson/screenjson-export -in="screenplay.json" -out="screenplay.celtx"
```


## screenjson-import

This utility converts major screenwriting formats, for example: Final Draft Pro, Fade In Pro, Fountain markdown etc, into ScreenJSON.

To build the Docker image:

```bash
docker build -f docker/screenjson-import/Dockerfile -t screenjson-import:latest .
docker push screenjson/screenjson-import:latest
```

#### Convert/import PDF to ScreenJSON

```bash
screenjson-import -in="screenplay.pdf" -out="screenplay.json"
```

```bash
docker run screenjson/screenjson-import -in="screenplay.pdf" -out="screenplay.json"
```

#### Convert/import Final Draft Pro to ScreenJSON

```bash
screenjson-import -in="screenplay.fdx" -out="screenplay.json"
```

```bash
docker run screenjson/screenjson-import -in="screenplay.fdx" -out="screenplay.json"
```

#### Convert/import FadeIn Pro to ScreenJSON

```bash
screenjson-import -in="screenplay.fadein" -out="screenplay.json"
```

```bash
docker run screenjson/screenjson-import -in="screenplay.fadein" -out="screenplay.json"
```

#### Convert/import Fountain to ScreenJSON

```bash
screenjson-import -in="screenplay.fountain" -out="screenplay.json"
```

```bash
docker run screenjson/screenjson-import -in="screenplay.fountain" -out="screenplay.json"
```

#### Convert/import Celtx to ScreenJSON

```bash
screenjson-import -in="screenplay.celtx" -out="screenplay.json"
```

```bash
docker run screenjson/screenjson-import -in="screenplay.celtx" -out="screenplay.json"
```

## screenjson-validate

This utility checks a JSON file to determine whether it is formatted correctly as ScreenJSON, using the official *JSON Schema* (https://json-schema.org/) specification.

To build the Docker image:

```bash
docker build -f docker/screenjson-validate/Dockerfile -t screenjson-validate:latest .
docker push screenjson/screenjson-validate:latest
```

#### Validate a JSON file

```bash
screenjson-validate my_screenplay.json
```

```bash
docker run screenjson/screenjson-validate my_screenplay.json
```
