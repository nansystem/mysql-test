CREATE TABLE IF NOT EXISTS `yh_team` (
  `team_id` int unsigned NOT NULL AUTO_INCREMENT,
  `team_name` varchar(255) NOT NULL,
  PRIMARY KEY (`team_id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `yh_member` (
  `team_id` int unsigned NOT NULL,
  `member_name` varchar(255) NOT NULL,
  `skill` varchar(255) NOT NULL,
  PRIMARY KEY (`team_id`, `member_name`)
) ENGINE = InnoDB;
