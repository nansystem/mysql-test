CREATE TABLE `products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `shop_id` int(10) unsigned NOT NULL,  -- 商品を掲載している店舗の ID
  `name` varchar(255) NOT NULL,         -- 商品名
  `price` int(10) unsigned NOT NULL,    -- 商品の価格
  `starts_at` datetime NOT NULL,        -- 商品の掲載開始日時
  `ends_at` datetime NOT NULL,          -- 商品の掲載終了日時
  PRIMARY KEY (`id`)
) ENGINE=InnoDB
