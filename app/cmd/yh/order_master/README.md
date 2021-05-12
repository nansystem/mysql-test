# ヤフー社内でやってるMySQLチューニングセミナー大公開 
[ヤフー社内でやってるMySQLチューニングセミナー大公開 ](https://www.slideshare.net/techblogyahoo/mysql-58540246)  

## type=ALL または type=indexでrowsが大きい
![image](https://image.slidesharecdn.com/mysqlsqltuningseminormynajpug201602-160222054211/95/mysql-36-638.jpg?cb=1456119774)

## 注文マスタ(order_master)から非表示でない(is_hidden_page=0)注文を、注文時間(order_time)の最新順に20件取得する

``` sql
# NG SQL
mysql> EXPLAIN SELECT order_time, seller_id, image_id, item_id  FROM yh_order_master WHERE is_hidden_page = '0' ORDER BY order_time D
ESC LIMIT 20;
+----+-------------+-----------------+------------+------+---------------+------+---------+------+--------+----------+-----------------------------+
| id | select_type | table           | partitions | type | possible_keys | key  | key_len | ref  | rows   | filtered | Extra
              |
+----+-------------+-----------------+------------+------+---------------+------+---------+------+--------+----------+-----------------------------+
|  1 | SIMPLE      | yh_order_master | NULL       | ALL  | NULL          | NULL | NULL    | NULL | 997875 |    10.00 | Using where; Using filesort |
+----+-------------+-----------------+------------+------+---------------+------+---------+------+--------+----------+-----------------------------+
1 row in set, 1 warning (0.00 sec)
```

typeがALL、Extraでfilesortが行われている。  

is_hidden_pageはほぼ0(テストデータは90% 0にした)
``` sql
mysql> select COUNT(is_hidden_page) FROM yh_order_master WHERE is_hidden_page = 0;
+-----------------------+
| COUNT(is_hidden_page) |
+-----------------------+
|                899657 |
+-----------------------+
```


``` sql
-- あえてカーディナリティが低いカラムにインデックスを張ってみる
ALTER TABLE yh_order_master ADD INDEX idx_is_hidden_page (is_hidden_page);
```

テーブルのstatus、indexのカーディナリティを確認。

``` sql
mysql> SHOW TABLE STATUS LIKE 'yh_order_master';
+-----------------+--------+---------+------------+--------+----------------+-------------+-----------------+--------------+-----------+----------------+---------------------+---------------------+------------+-------------+----------+----------------+---------+
| Name            | Engine | Version | Row_format | Rows   | Avg_row_length | Data_length | Max_data_length | Index_length | Data_free | Auto_increment | Create_time         | Update_time         | Check_time | Collation   | Checksum | Create_options | Comment |
+-----------------+--------+---------+------------+--------+----------------+-------------+-----------------+--------------+-----------+----------------+---------------------+---------------------+------------+-------------+----------+----------------+---------+
| yh_order_master | InnoDB |      10 | Dynamic    | 997875 |             43 |    43581440 |               0 |            0 |   5242880 |        1000001 | 2021-05-12 23:09:50 | 2021-05-11 23:36:07 | NULL       | utf8mb4_bin |     NULL |                |         |
+-----------------+--------+---------+------------+--------+----------------+-------------+-----------------+--------------+-----------+----------------+---------------------+---------------------+------------+-------------+----------+----------------+---------+
1 row in set (0.01 sec)

mysql> SHOW INDEX FROM yh_order_master;
+-----------------+------------+--------------------+--------------+----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| Table           | Non_unique | Key_name           | Seq_in_index | Column_name    | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment | Visible | Expression |
+-----------------+------------+--------------------+--------------+----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
| yh_order_master |          0 | PRIMARY            |            1 | id             | A         |      979875 |     NULL |   NULL |      | BTREE      |         |               | YES     | NULL       |
| yh_order_master |          1 | idx_is_hidden_page |            1 | is_hidden_page | A         |           1 |     NULL |   NULL |      | BTREE      |         |               | YES     | NULL       |
+-----------------+------------+--------------------+--------------+----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+---------+------------+
2 rows in set (0.01 sec)
```

``` sql
mysql> EXPLAIN SELECT order_time, seller_id, image_id, item_id  FROM yh_order_master WHERE is_hidden_page = '0' ORDER BY order_time DESC LIMIT 20;
+----+-------------+-----------------+------------+------+--------------------+--------------------+---------+-------+--------+----------+----------------+
| id | select_type | table           | partitions | type | possible_keys      | key                | key_len | ref   | rows   | filtered | Extra          |
+----+-------------+-----------------+------------+------+--------------------+--------------------+---------+-------+--------+----------+----------------+
|  1 | SIMPLE      | yh_order_master | NULL       | ref  | idx_is_hidden_page | idx_is_hidden_page | 1       | const | 498937 |   100.00 | Using filesort |
+----+-------------+-----------------+------------+------+--------------------+--------------------+---------+-------+--------+----------+----------------+
1 row in set, 1 warning (0.00 sec)
```

INDEX使われた!  
カーディナリティが低いからといっても計測したほうがよい。  

``` sql
-- あえて昇順のインデックスを張ってみる
ALTER TABLE yh_order_master ADD INDEX idex_order_time (order_time);
```

``` sql
mysql> EXPLAIN SELECT order_time, seller_id, image_id, item_id  FROM yh_order_master WHERE is_hidden_page = '0' ORDER BY order_time DESC LIMIT 20;
+----+-------------+-----------------+------------+-------+--------------------+-----------------+---------+------+------+----------+----------------------------------+
| id | select_type | table           | partitions | type  | possible_keys      | key             | key_len | ref  | rows | filtered | Extra                            |
+----+-------------+-----------------+------------+-------+--------------------+-----------------+---------+------+------+----------+----------------------------------+
|  1 | SIMPLE      | yh_order_master | NULL       | index | idx_is_hidden_page | idex_order_time | 5       | NULL |   40 |    50.00 | Using where; Backward index scan |
+----+-------------+-----------------+------------+-------+--------------------+-----------------+---------+------+------+----------+----------------------------------+
1 row in set, 1 warning (0.00 sec)

mysql> EXPLAIN ANALYZE SELECT order_time, seller_id, image_id, item_id  FROM yh_order_master WHERE is_hidden_page = '0' ORDER BY orde
r_time DESC LIMIT 20;
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| EXPLAIN

                                                                          |
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| -> Limit: 20 row(s)  (cost=0.08 rows=20) (actual time=0.250..0.258 rows=20 loops=1)
    -> Filter: (yh_order_master.is_hidden_page = 0)  (cost=0.08 rows=20) (actual time=0.249..0.255 rows=20 loops=1)
        -> Index scan on yh_order_master using idex_order_time (reverse)  (cost=0.08 rows=40) (actual time=0.247..0.251 rows=22 loops=1)
 |
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```

`Backward index scan` INDEXを張っておけば逆順でも使ってくれるぽい。  

> The index is sorted, with the lowest value first. To find the max value, a backward index scan would find the maximum value first :).
> https://stackoverflow.com/questions/5017327/index-scan-backward-vs-index-scan

MySQL5.7から`Backward index scan`採用されている。
> While it should be noted that the MySQL 5.7 optimizer is able to scan an ascending index backwards (to give descending order), it comes at a higher cost. As shown further down, we can see forward index scans are ~15% better than backward index scans.
> https://mysqlserverteam.com/mysql-8-0-labs-descending-indexes-in-mysql/

> Thanks for showing interest in the new feature. The ~15% cost benefit in forward scans can be attributed to the optimizations done in innodb to favor forward scans over backward scans.
> https://dba.stackexchange.com/questions/199551/why-scanning-an-index-backwards-is-slower


``` sql
-- 降順のインデックスを張って、昇順のインデックスとどちらが使われるか試してみる
ALTER TABLE yh_order_master ADD INDEX idex_order_time_desc (order_time DESC);
```

``` sql
mysql> EXPLAIN SELECT order_time, seller_id, image_id, item_id  FROM yh_order_master WHERE is_hidden_page = '0' ORDER BY order_time D
ESC LIMIT 20;
+----+-------------+-----------------+------------+-------+--------------------+----------------------+---------+------+------+----------+-------------+
| id | select_type | table           | partitions | type  | possible_keys      | key                  | key_len | ref  | rows | filtered | Extra       |
+----+-------------+-----------------+------------+-------+--------------------+----------------------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | yh_order_master | NULL       | index | idx_is_hidden_page | idex_order_time_desc | 5       | NULL |   80 |    50.00 | Using where |
+----+-------------+-----------------+------------+-------+--------------------+----------------------+---------+------+------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```

当然、降順のインデックスが使われる。  

``` sql
mysql> EXPLAIN ANALYZE SELECT order_time, seller_id, image_id, item_id  FROM yh_order_master WHERE is_hidden_page = '0' ORDER BY orde
r_time DESC LIMIT 20;
+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| EXPLAIN

                                                                     |
+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| -> Limit: 20 row(s)  (cost=0.16 rows=20) (actual time=0.177..0.185 rows=20 loops=1)
    -> Filter: (yh_order_master.is_hidden_page = 0)  (cost=0.16 rows=40) (actual time=0.176..0.182 rows=20 loops=1)
        -> Index scan on yh_order_master using idex_order_time_desc  (cost=0.16 rows=80) (actual time=0.174..0.178 rows=22 loops=1)
 |
+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```


``` sql
ALTER TABLE yh_order_master ADD INDEX idex_is_hidden_page_order_time_desc (is_hidden_page, order_time DESC);
```

``` sql
mysql> EXPLAIN SELECT order_time, seller_id, image_id, item_id  FROM yh_order_master WHERE is_hidden_page = '0' ORDER BY order_time DESC LIMIT 20;
+----+-------------+-----------------+------------+------+--------------------------------------------------------+-------------------------------------+---------+-------+--------+----------+-------+
| id | select_type | table           | partitions | type | possible_keys                                          | key
                   | key_len | ref   | rows   | filtered | Extra |
+----+-------------+-----------------+------------+------+--------------------------------------------------------+-------------------------------------+---------+-------+--------+----------+-------+
|  1 | SIMPLE      | yh_order_master | NULL       | ref  | idx_is_hidden_page,idex_is_hidden_page_order_time_desc | idex_is_hidden_page_order_time_desc | 1       | const | 498937 |   100.00 | NULL  |
+----+-------------+-----------------+------------+------+--------------------------------------------------------+-------------------------------------+---------+-------+--------+----------+-------+
1 row in set, 1 warning (0.00 sec)

mysql> EXPLAIN ANALYZE SELECT order_time, seller_id, image_id, item_id  FROM yh_order_master WHERE is_hidden_page = '0' ORDER BY orde
r_time DESC LIMIT 20;
+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| EXPLAIN
                                                                                                                                  |
+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| -> Limit: 20 row(s)  (cost=51888.70 rows=20) (actual time=0.144..0.147 rows=20 loops=1)
    -> Index lookup on yh_order_master using idex_is_hidden_page_order_time_desc (is_hidden_page=0)  (cost=51888.70 rows=498937) (actual time=0.143..0.145 rows=20 loops=1)
 |
+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```