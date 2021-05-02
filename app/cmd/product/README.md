# products

[MySQL with InnoDB のインデックスの基礎知識とありがちな間違い](https://techlife.cookpad.com/entry/2017/04/18/092524)  

商品情報を管理するテーブルで、現在以降に掲載される商品の情報が日々追加されていく想定です。  

``` sql
CREATE TABLE `products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `shop_id` int(10) unsigned NOT NULL,  -- 商品を掲載している店舗の ID
  `name` varchar(255) NOT NULL,         -- 商品名
  `price` int(10) unsigned NOT NULL,    -- 商品の価格
  `starts_at` datetime NOT NULL,        -- 商品の掲載開始日時
  `ends_at` datetime NOT NULL,          -- 商品の掲載終了日時
  PRIMARY KEY (`id`)
) ENGINE=InnoDB
```

## 現在掲載されている商品を抽出する

``` sql
SELECT * FROM products WHERE starts_at <= NOW() AND ends_at >= NOW();
```

### インデックスの最初のカラムに範囲指定をする複合インデックス
``` sql
ALTER TABLE products ADD INDEX ix_ends_at_starts_at (ends_at, starts_at);
```

``` sql
-- EXPLAIN
# id, select_type, table, partitions, type, possible_keys, key, key_len, ref, rows, filtered, Extra
'1', 'SIMPLE', 'products', NULL, 'ALL', 'ix_ends_at_starts_at', NULL, NULL, NULL, '996820', '16.66', 'Using where'
```

INDEX使われないのでIndex Hintを使って実行

``` sql
SELECT * FROM products FORCE INDEX (ix_ends_at_starts_at) WHERE starts_at <= NOW() AND ends_at >= NOW();
```

```
-> Index range scan on products using ix_ends_at_starts_at, with index condition: ((products.starts_at <= <cache>(now())) and (products.ends_at >= <cache>(now())))  (cost=191055.77 rows=498410) (actual time=0.319..112.438 rows=27672 loops=1)
```

`used_key_parts`で`ends_at`しか使われない。  

![image](https://user-images.githubusercontent.com/6994710/116806326-2410e400-ab67-11eb-9567-c71eb7a9396f.png)

## 特定の店舗の現在掲載される商品を抽出する

``` sql
-- 件数の多いshop_idをさがす
SELECT COUNT(shop_id) cnt, shop_id FROM products GROUP BY shop_id ORDER BY cnt DESC;
```

``` sql
-- ダメな例
ALTER TABLE products ADD INDEX ix_ends_at_shop_id (ends_at, shop_id);
```

``` sql
-- ダメな例
ALTER TABLE products ADD INDEX ix_shop_id_starts_at (shop_id, starts_at);
```

``` sql
SELECT * FROM products WHERE shop_id = 767 AND starts_at <= NOW() AND ends_at >= NOW();
```

``` sql
-- よい例
ALTER TABLE products DROP INDEX ix_ends_at_shop_id,
  ADD INDEX ix_shop_id_ends_at (shop_id, ends_at);
```

``` sql
-- explain
mysql> EXPLAIN SELECT * FROM products WHERE shop_id = 767 AND starts_at <= NOW() AND ends_at >= NOW();
+----+-------------+----------+------------+-------+-----------------------------------------+--------------------+---------+------+------+----------+------------------------------------+
| id | select_type | table    | partitions | type  | possible_keys                           | key                | key_len | ref  | rows | filtered | Extra                              |
+----+-------------+----------+------------+-------+-----------------------------------------+--------------------+---------+------+------+----------+------------------------------------+
|  1 | SIMPLE      | products | NULL       | range | ix_ends_at_starts_at,ix_shop_id_ends_at | ix_shop_id_ends_at | 9       | NULL |  312 |    33.33 | Using index condition; Using where |
+----+-------------+----------+------------+-------+-----------------------------------------+--------------------+---------+------+------+----------+------------------------------------+
1 row in set, 1 warning (0.00 sec)

-- EXLAIN ANALYZE
mysql> EXPLAIN ANALYZE SELECT * FROM products WHERE shop_id = 767 AND starts_at <= NOW() AND ends_at >= NOW();
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| EXPLAIN

                                                                          |
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| -> Filter: (products.starts_at <= <cache>(now()))  (cost=119.86 rows=104) (actual time=0.539..1.631 rows=41 loops=1)
    -> Index range scan on products using ix_shop_id_ends_at, with index condition: ((products.shop_id = 767) and (products.ends_at >= <cache>(now())))  (cost=119.86 rows=312) (actual time=0.536..1.591 rows=312 loops=1)
 |
+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```

