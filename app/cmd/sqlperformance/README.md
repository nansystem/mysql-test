# SQLパフォーマンス詳解
[SQLパフォーマンス詳解](https://sql-performance-explained.jp/)  

## 複合インデックスの順序
INDEXの2番目だけWHERE句に使ってもINDEXは使われない。
``` sql
mysql> EXPLAIN SELECT first_name, last_name
    -> FROM employees
    -> WHERE subsidiary_id = 20;
+----+-------------+-----------+------------+------+---------------+------+---------+------+--------+----------+-------------+
| id | select_type | table     | partitions | type | possible_keys | key  | key_len | ref  | rows   | filtered | Extra       |
+----+-------------+-----------+------------+------+---------------+------+---------+------+--------+----------+-------------+
|  1 | SIMPLE      | employees | NULL       | ALL  | NULL          | NULL | NULL    | NULL | 996538 |    10.00 | Using where |
+----+-------------+-----------+------------+------+---------------+------+---------+------+--------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```

``` sql
ALTER TABLE employees DROP PRIMARY KEY, ADD PRIMARY KEY (subsidiary_id, employee_id);
```

``` sql
mysql> EXPLAIN SELECT first_name, last_name FROM employees WHERE subsidiary_id = 20;
+----+-------------+-----------+------------+------+---------------+---------+---------+-------+------+----------+-------+
| id | select_type | table     | partitions | type | possible_keys | key     | key_len | ref   | rows | filtered | Extra |
+----+-------------+-----------+------------+------+---------------+---------+---------+-------+------+----------+-------+
|  1 | SIMPLE      | employees | NULL       | ref  | PRIMARY       | PRIMARY | 4       | const | 1000 |   100.00 | NULL  |
+----+-------------+-----------+------------+------+---------------+---------+---------+-------+------+----------+-------+
1 row in set, 1 warning (0.01 sec)

mysql> EXPLAIN ANALYZE SELECT first_name, last_name FROM employees WHERE subsidiary_id = 20;
+--------------------------------------------------------------------------------------------------------------------------------------+
| EXPLAIN
  |
+--------------------------------------------------------------------------------------------------------------------------------------+
| -> Index lookup on employees using PRIMARY (subsidiary_id=20)  (cost=102.22 rows=1000) (actual time=0.022..0.390 rows=1000 loops=1)
 |
+--------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```

## NO_INDEX
``` sql
mysql> EXPLAIN SELECT /*+ NO_INDEX(employees PRIMARY) */
    -> first_name, last_name FROM employees WHERE subsidiary_id = 20;
+----+-------------+-----------+------------+------+---------------+------+---------+------+--------+----------+-------------+
| id | select_type | table     | partitions | type | possible_keys | key  | key_len | ref  | rows   | filtered | Extra       |
+----+-------------+-----------+------------+------+---------------+------+---------+------+--------+----------+-------------+
|  1 | SIMPLE      | employees | NULL       | ALL  | NULL          | NULL | NULL    | NULL | 996750 |     0.10 | Using where |
+----+-------------+-----------+------------+------+---------------+------+---------+------+--------+----------+-------------+
1 row in set, 1 warning (0.00 sec)
```

## 関数を使った時のINDEX

``` sql
CREATE INDEX emp_name ON employees (last_name);
ERROR 1071 (42000): Specified key was too long; max key length is 3072 bytes
-- VARCHAR(1000)だとmaxこえてINDEXはれない。 4バイト文字対応するなら3072/4=768がmax
```

``` sql
mysql> EXPLAIN ANALYZE SELECT * FROM employees WHERE last_name ='横田' AND subsidiary_id = 30;
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| EXPLAIN
                                                                                                                     |
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| -> Filter: (employees.subsidiary_id = 30)  (cost=511.42 rows=197) (actual time=6.417..7.641 rows=2 loops=1)
    -> Index lookup on employees using emp_name (last_name='横田')  (cost=511.42 rows=1967) (actual time=0.096..7.508 rows=1967 loops=1)
   |
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.01 sec)
```

## INDEXの張られているカラムに関数を使うとINDEXは使われない
関数を使いたかったら`functional index`にするか、入力値に関数を使う。

``` sql
mysql> EXPLAIN ANALYZE SELECT * FROM employees WHERE UPPER(last_name) = '横田';
+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| EXPLAIN
                                                                                                             |
+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| -> Filter: (upper(employees.last_name) = '横田')  (cost=100609.60 rows=996561) (actual time=0.166..325.548 rows=1967 loops=1)
    -> Table scan on employees  (cost=100609.60 rows=996561) (actual time=0.034..255.746 rows=1000000 loops=1)
   |
+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.32 sec)
```

``` sql
CREATE INDEX emp_up_name ON employees ( (UPPER(last_name)) );
ERROR 3757 (HY000): Cannot create a functional index on an expression that returns a BLOB or TEXT. Please consider using CAST.
CREATE INDEX emp_up_name ON employees ( (CAST(UPPER(last_name) AS CHAR(255))) );
```

## 第2章 where句

### 範囲検索 p39-

### 部署を指定して誕生日の範囲で絞る

``` sql
SELECT first_name, last_name, date_of_birth, subsidiary_id
FROM employees
WHERE 
subsidiary_id = 298
AND date_of_birth >= '2010-04-01'
AND date_of_birth <= '2010-04-09';
```

``` sql
ALTER TABLE employees ADD INDEX idx_date_of_birth_sub (date_of_birth, subsidiary_id);
ALTER TABLE employees ADD INDEX idx_sub_date_of_birth (subsidiary_id, date_of_birth);
```

``` sql
mysql> ANALYZE EXPLAIN SELECT first_name, last_name, date_of_birth, subsidiary_id FROM employees WHERE  subsidiary_id = 298 AND date_
of_birth >= '2010-04-01' AND date_of_birth <= '2010-04-09';
ERROR 1064 (42000): You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'EXPLAIN SELECT first_name, last_name, date_of_birth, subsidiary_id FROM employee' at line 1
mysql> EXPLAIN ANALYZE SELECT first_name, last_name, date_of_birth, subsidiary_id FROM employees WHERE  subsidiary_id = 298 AND date_
of_birth >= '2010-04-01' AND date_of_birth <= '2010-04-09';
+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| EXPLAIN

                 |
+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| -> Index range scan on employees using idx_sub_date_of_birth, with index condition: ((employees.subsidiary_id = 298) and (employees.date_of_birth >= DATE'2010-04-01') and (employees.date_of_birth <= DATE'2010-04-09'))  (cost=2.06 rows=4) (actual time=0.062..0.067 rows=4 loops=1)
 |
+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```