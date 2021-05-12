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

INSERT INTO yh_team(team_id, team_name) VALUES
  (1, 'Team1'),
  (2, 'Team2'),
  (3, 'Team3');

INSERT INTO yh_member(team_id, member_name, skill) VALUES
  (1, 'Taro', 'C'),
  (1, 'Jiro', 'Python'),
  (2, 'Hanako', 'Ruby'),
  (2, 'Saburo', 'PHP'),
  (3, 'Sirou', 'Perl');
