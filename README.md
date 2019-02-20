# MariaDB_container
This is a quick (temporary) solution for containerising ENotary database. 

### Follow these commands to run the project 

1- Clone the repo first

```
git clone https://github.com/UsmanFazil/MariaDB_container.git
```

2- Start the Database container 
```
docker run --name mariadbEnotary -p 3306:3306 -e MYSQL_ROOT_PASSWORD=mypass -d mariadb/server:10.3
```

3- Now run the go file to create Database schema.
```
go run dbfirst.go
```