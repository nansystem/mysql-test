# ad, cv
[[MySQL Workbench] VISUAL EXPLAIN でインデックスの挙動を確認する](https://engineering.linecorp.com/ja/blog/mysql-workbench-index-behavior-with-visual-explain/)

![image](https://user-images.githubusercontent.com/6994710/116802130-81973780-ab4b-11eb-9ab9-b6067030d3f6.png)

## インデックスなし

``` sql
SELECT * FROM cv WHERE status = 2 ORDER BY created_at LIMIT 100;
```

## WHERE の条件にインデックス

``` sql
ALTER TABLE cv ADD INDEX idx_status (status);
ALTER TABLE cv DROP INDEX idx_status;
```

## ORDER BYの条件にインデックス

``` sql
ALTER TABLE cv ADD INDEX idx_created_at (created_at);
ALTER TABLE cv DROP INDEX idx_created_at;
```

## GROUP BYの条件にインデックス

``` sql
SELECT ad_id, COUNT(*) FROM cv WHERE status = 2 GROUP BY ad_id;
```

``` sql
ALTER TABLE cv ADD INDEX idx_ad_id (ad_id);
ALTER TABLE cv DROP INDEX idx_ad_id;
```

## 複合インデックス

``` sql
SELECT * FROM cv WHERE status = 2 ORDER BY created_at LIMIT 100;
```

``` sql
ALTER TABLE cv ADD INDEX idx_status_created_at (status, created_at);
ALTER TABLE cv DROP INDEX idx_status_created_at;
```

## 複雑なSQL

一週間前からカウントして、コンバージョン数が多い順に ad テーブルのレコードを並べたい場合のクエリについて見ます。
細かい条件として cv の status は1、 ad の type は2のデータを抽出したいです。

``` sql
SELECT
    ad.id,
    COUNT(DISTINCT cv.user_id) as cv_count
FROM
    ad
    INNER JOIN cv
    ON cv.ad_id = ad.id
        AND cv.status = 1
        AND cv.created_at >= UNIX_TIMESTAMP(CURDATE() - INTERVAL 7 DAY)
WHERE
    ad.type = 2
    AND ad.end_at >= UNIX_TIMESTAMP(CURDATE() - INTERVAL 7 DAY)
GROUP BY
    ad.id
ORDER BY
    cv_count DESC
LIMIT 100
;
```

``` sql
-- EXPLAIN
# id, select_type, table, partitions, type, possible_keys, key, key_len, ref, rows, filtered, Extra
-- ALL -> Full Table Scan 
'1', 'SIMPLE', 'cv', NULL, 'ALL', NULL, NULL, NULL, NULL, '997899', '3.33', 'Using where; Using temporary; Using filesort'
'1', 'SIMPLE', 'ad', NULL, 'eq_ref', 'PRIMARY', 'PRIMARY', '4', 'devdb.cv.ad_id', '1', '5.00', 'Using where'

-- EXPLAIN ANALYZE
-> Limit: 100 row(s)  (actual time=470.995..471.002 rows=100 loops=1)
    -> Sort: cv_count DESC, limit input to 100 row(s) per chunk  (actual time=470.994..470.998 rows=100 loops=1)
        -> Stream results  (actual time=469.723..470.945 rows=386 loops=1)
            -> Group aggregate: count(distinct cv.user_id)  (actual time=469.720..470.892 rows=386 loops=1)
                -> Sort: ad.id  (actual time=469.712..469.993 rows=9660 loops=1)
                    -- buffer result ? 
                    -> Stream results  (cost=105378.90 rows=1663) (actual time=0.085..467.481 rows=9660 loops=1)
                        -- JOIN が nested loop で処理されている 
                        -> Nested loop inner join  (cost=105378.90 rows=1663) (actual time=0.084..466.207 rows=9660 loops=1)
                            -> Filter: ((cv.`status` = 1) and (cv.created_at >= <cache>(unix_timestamp((curdate() - interval 7 day)))))  (cost=93737.91 rows=33260) (actual time=0.037..223.265 rows=100049 loops=1)
                                -> Table scan on cv  (cost=93737.91 rows=997899) (actual time=0.027..179.085 rows=1000000 loops=1)
                            -> Filter: ((ad.`type` = 2) and (ad.end_at >= <cache>(unix_timestamp((curdate() - interval 7 day)))))  (cost=0.25 rows=0) (actual time=0.002..0.002 rows=0 loops=100049)
                                -- cv テーブルから読み取られた行に対して、対応する ad テーブルの行が PRIMARY KEY を使って読み取られている
                                -> Single-row index lookup on ad using PRIMARY (id=cv.ad_id)  (cost=0.25 rows=1) (actual time=0.002..0.002 rows=1 loops=100049)
```

### cvのWHEREのためのindex追加
``` sql
ALTER TABLE cv ADD INDEX idx_status_created_at (status, created_at);
```

``` sql
-- EXPLAIN
# id, select_type, table, partitions, type, possible_keys, key, key_len, ref, rows, filtered, Extra
-- range idx_status_created_atを使うようになった
'1', 'SIMPLE', 'cv', NULL, 'range', 'idx_status_created_at', 'idx_status_created_at', '5', NULL, '203596', '100.00', 'Using index condition; Using temporary; Using filesort'
'1', 'SIMPLE', 'ad', NULL, 'eq_ref', 'PRIMARY', 'PRIMARY', '4', 'devdb.cv.ad_id', '1', '5.00', 'Using where'

-- EXPLAIN ANALYZE
-> Limit: 100 row(s)  (actual time=499.779..499.787 rows=100 loops=1)
    -> Sort: cv_count DESC, limit input to 100 row(s) per chunk  (actual time=499.778..499.781 rows=100 loops=1)
        -> Stream results  (actual time=498.048..499.697 rows=386 loops=1)
            -> Group aggregate: count(distinct cv.user_id)  (actual time=498.045..499.629 rows=386 loops=1)
                -> Sort: ad.id  (actual time=498.034..498.428 rows=9660 loops=1)
                    -> Stream results  (cost=162877.06 rows=10180) (actual time=0.274..495.351 rows=9660 loops=1)
                        -> Nested loop inner join  (cost=162877.06 rows=10180) (actual time=0.273..493.708 rows=9660 loops=1)
                            -- idx_status_created_at見るようになった 
                            -> Index range scan on cv using idx_status_created_at, with index condition: ((cv.`status` = 1) and (cv.created_at >= <cache>(unix_timestamp((curdate() - interval 7 day)))))  (cost=91618.46 rows=203596) (actual time=0.256..220.947 rows=100049 loops=1)
                            -> Filter: ((ad.`type` = 2) and (ad.end_at >= <cache>(unix_timestamp((curdate() - interval 7 day)))))  (cost=0.25 rows=0) (actual time=0.003..0.003 rows=0 loops=100049)
                                -> Single-row index lookup on ad using PRIMARY (id=cv.ad_id)  (cost=0.25 rows=1) (actual time=0.002..0.002 rows=1 loops=100049)
```

### 結合条件のad_idも加えたindex追加
``` sql
ALTER TABLE cv ADD INDEX idx_ad_id_status_created_at (ad_id, status, created_at);
```

``` sql
-- EXPLAIN 
# id, select_type, table, partitions, type, possible_keys, key, key_len, ref, rows, filtered, Extra
-- 
'1', 'SIMPLE', 'ad', NULL, 'index', 'PRIMARY', 'PRIMARY', '4', NULL, '3930', '3.33', 'Using where; Using temporary; Using filesort'
-- idx_ad_id_status_created_atの方が使われた
'1', 'SIMPLE', 'cv', NULL, 'ref', 'idx_status_created_at,idx_ad_id_status_created_at', 'idx_ad_id_status_created_at', '5', 'devdb.ad.id,const', '25', '33.33', 'Using index condition'

-- EXPLAIN ANALYZE
-> Limit: 100 row(s)  (actual time=30.358..30.364 rows=100 loops=1)
    -> Sort: cv_count DESC, limit input to 100 row(s) per chunk  (actual time=30.357..30.360 rows=100 loops=1)
        -> Stream results  (actual time=0.324..30.195 rows=386 loops=1)
            -> Group aggregate: count(distinct cv.user_id)  (actual time=0.322..30.015 rows=386 loops=1)
                -> Nested loop inner join  (cost=2197.10 rows=3284) (actual time=0.200..28.738 rows=9660 loops=1)
                    -- adが駆動表
                    -> Filter: ((ad.`type` = 2) and (ad.end_at >= <cache>(unix_timestamp((curdate() - interval 7 day)))))  (cost=391.05 rows=131) (actual time=0.058..1.335 rows=386 loops=1)
                        -> Index scan on ad using PRIMARY  (cost=391.05 rows=3930) (actual time=0.050..1.069 rows=4000 loops=1)
                    -- idx_ad_id_status_created_atが使われた
                    -> Index lookup on cv using idx_ad_id_status_created_at (ad_id=ad.id, status=1), with index condition: (cv.created_at >= <cache>(unix_timestamp((curdate() - interval 7 day))))  (cost=6.27 rows=25) (actual time=0.063..0.070 rows=25 loops=386)
```

![image](https://user-images.githubusercontent.com/6994710/116247312-abe0a200-a7a5-11eb-8801-101bfbb134e6.png)


### adテーブルにWHERE句のためのindex追加

``` sql
ALTER TABLE ad ADD INDEX idx_ad_type_end_at (type, end_at);
ALTER TABLE ad ADD INDEX idx_ad_type_end_at_id (type, end_at, id);
```

``` sql
-- EXPLAIN
# id, select_type, table, partitions, type, possible_keys, key, key_len, ref, rows, filtered, Extra
-- idx_ad_type_end_atが使われた
'1', 'SIMPLE', 'ad', NULL, 'range', 'PRIMARY,idx_ad_type_end_at', 'idx_ad_type_end_at', '5', NULL, '386', '100.00', 'Using where; Using index; Using temporary; Using filesort'
'1', 'SIMPLE', 'cv', NULL, 'ref', 'idx_status_created_at,idx_ad_id_status_created_at', 'idx_ad_id_status_created_at', '5', 'devdb.ad.id,const', '25', '33.33', 'Using index condition'

-- EXPLAIN ANALYZE
-> Limit: 100 row(s)  (actual time=24.452..24.459 rows=100 loops=1)
    -> Sort: cv_count DESC, limit input to 100 row(s) per chunk  (actual time=24.451..24.455 rows=100 loops=1)
        -> Stream results  (actual time=0.470..24.338 rows=386 loops=1)
            -> Group aggregate: count(distinct cv.user_id)  (actual time=0.468..24.216 rows=386 loops=1)
                -> Nested loop inner join  (cost=5400.55 rows=9676) (actual time=0.370..23.127 rows=9660 loops=1)
                    -- ad.idソートしてる
                    -> Sort: ad.id  (cost=78.37 rows=386) (actual time=0.271..0.320 rows=386 loops=1)
                        -> Filter: ((ad.`type` = 2) and (ad.end_at >= <cache>(unix_timestamp((curdate() - interval 7 day)))))  (cost=78.37 rows=386) (actual time=0.046..0.186 rows=386 loops=1)
                            -- idx_ad_type_end_atが使われた
                            -> Index range scan on ad using idx_ad_type_end_at  (cost=78.37 rows=386) (actual time=0.040..0.144 rows=386 loops=1)
                    -- idx_ad_id_status_created_atのうち使われたのはad_id, status けどindex conditionでcreated_at使っているのでは?
                    -> Index lookup on cv using idx_ad_id_status_created_at (ad_id=ad.id, status=1), with index condition: (cv.created_at >= <cache>(unix_timestamp((curdate() - interval 7 day))))  (cost=6.27 rows=25) (actual time=0.052..0.058 rows=25 loops=386)
```
![image](https://user-images.githubusercontent.com/6994710/116248082-78524780-a7a6-11eb-96be-866d16bdb327.png)  
![image](https://user-images.githubusercontent.com/6994710/116251104-4e4e5480-a7a9-11eb-9aa3-73cfcb256bc2.png)  
![image](https://user-images.githubusercontent.com/6994710/116251161-602ff780-a7a9-11eb-871b-b14cdca92c1b.png)  


### used_key_partsに含まれていないcreated_atを消してみる
``` sql
ALTER TABLE cv DROP INDEX idx_ad_id_status_created_at;
ALTER TABLE cv ADD INDEX idx_ad_id_status (ad_id, status);
```

``` sql
-- EXPLAIN ANALYZE
-> Limit: 100 row(s)  (actual time=31.337..31.344 rows=100 loops=1)
    -> Sort: cv_count DESC, limit input to 100 row(s) per chunk  (actual time=31.337..31.340 rows=100 loops=1)
        -> Stream results  (actual time=0.564..31.131 rows=386 loops=1)
            -> Group aggregate: count(distinct cv.user_id)  (actual time=0.562..30.938 rows=386 loops=1)
                -> Nested loop inner join  (cost=3429.58 rows=3191) (actual time=0.462..29.454 rows=9660 loops=1)
                    -> Sort: ad.id  (cost=78.37 rows=386) (actual time=0.353..0.416 rows=386 loops=1)
                        -> Filter: ((ad.`type` = 2) and (ad.end_at >= <cache>(unix_timestamp((curdate() - interval 7 day)))))  (cost=78.37 rows=386) (actual time=0.078..0.259 rows=386 loops=1)
                            -> Index range scan on ad using idx_ad_type_end_at  (cost=78.37 rows=386) (actual time=0.073..0.212 rows=386 loops=1)
                    -- created_atを実テーブルから見ているため遅くなっているように見えるがcost=3429.58は低くなってる。used_key_partsを信用したほうがいいぽい。
                    -> Filter: (cv.created_at >= <cache>(unix_timestamp((curdate() - interval 7 day))))  (cost=6.20 rows=8) (actual time=0.065..0.074 rows=25 loops=386)
                        -> Index lookup on cv using idx_ad_id_status (ad_id=ad.id, status=1)  (cost=6.20 rows=25) (actual time=0.065..0.072 rows=25 loops=386)
```
