CREATE TABLE IF NOT EXISTS `yh_push_info` (
  `push_id` int unsigned NOT NULL AUTO_INCREMENT,
  `mod_date` datetime NOT NULL,
  `deleted_flg` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`push_id`)
) ENGINE = InnoDB;
