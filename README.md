# MySQLのexplain学習

``` sh
$ docker-compose build
$ docker-compose up -d
$ docker-compose exec app go run cmd/cvad/main.go
```

### 作り直し
``` sh
$ docker-compose down
$ docker-compose build
$ docker-compose up
```

### debug
``` sh
$ docker-compose up -d
$ docker-compose ps
$ docker-compose logs
```

### MYSQL_DATABASE変更時
volumeがすでにある場合、MYSQL_DATABASEで指定したデータベースを生成しない。  
そのためvolumeを削除してから再度起動する

``` sh
$ docker volume ls
$ docker volume rm [dbのvolume名]
$ docker-compose up -d
```

## xo

``` sh
$ xo "mysql://root:password@localhost:3306/devdb" -o generated --verbose
error # ファイルはgenerateされるがerrorと表示される
```

## file load
```
mysql> set persist local_infile=1;
mysql> select @@local_infile;
+----------------+
| @@local_infile |
+----------------+
|              1 |
+----------------+
1 row in set (0.01 sec)
```

``` sh
mysql --local-infile=1 -uroot -ppassword --protocol tcp
load data local
    infile '/home/nancy/git/mysql-test/app/clients.csv'
into table
    devdb.test_table
fields
    terminated by ',';
```

``` sql
CREATE TABLE IF NOT EXISTS `test_table` (
    `client_id` INT(10) NOT NULL,
    `client_name` VARCHAR(128) NOT NULL,
    `client_age` INT(10) NOT NULL
) ENGINE = InnoDB;
```