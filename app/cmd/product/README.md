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
