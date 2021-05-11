CREATE TABLE IF NOT EXISTS `yh_order_master` (
    `id` INT unsigned NOT NULL AUTO_INCREMENT,
    `order_time` datetime NOT NULL,
    `seller_id` INT unsigned NOT NULL,
    `image_id` INT unsigned NOT NULL,
    `item_id` INT unsigned NOT NULL,
    `is_hidden_page` tinyint NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;