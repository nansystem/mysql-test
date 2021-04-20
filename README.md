# MySQLのexplain学習

``` sh
$ docker-compose build
$ docker-compose up -d
$ docker-compose exec app go run main.go
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

## gen/xo

``` sh
$ gen --connstr="root:password@tcp(localhost:3306)/devdb" --out generated --sqltype mysql --database devdb --no-json --overwrite
```

``` sh
$ xo "mysql://root:password@localhost:3306/devdb" -o generated --verbose
error # ファイルはgenerateされるがerrorと表示される
```