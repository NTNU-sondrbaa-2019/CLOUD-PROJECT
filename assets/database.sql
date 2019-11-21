SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";
CREATE DATABASE IF NOT EXISTS `gr8elo` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `gr8elo`;

CREATE TABLE `GROUP` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `league_id` bigint(20) UNSIGNED NOT NULL,
  `league_season_name` varchar(32) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `GROUP_USER` (
  `group_id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `LEAGUE` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `platform_elo_id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `groups_limit` int(10) UNSIGNED NOT NULL,
  `group_size_limit` int(10) UNSIGNED NOT NULL,
  `start` datetime NOT NULL,
  `end` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `LEAGUE_SEASON` (
  `league_id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(32) NOT NULL,
  `start` datetime NOT NULL,
  `end` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `PLATFORM` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `version` varchar(16) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `PLATFORM_AUTH` (
  `platform_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `PLATFORM_ELO` (
  `platform_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `PLATFORM_ENDPOINT` (
  `platform_id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `path` varchar(1000) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `PLATFORM_USER` (
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `platform_id` bigint(20) UNSIGNED NOT NULL,
  `access_token` varchar(500) NOT NULL,
  `verification_key` varchar(500) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `RESULT` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `group_id` bigint(20) UNSIGNED NOT NULL,
  `elo_before` int(11) NOT NULL,
  `elo_after` int(11) NOT NULL,
  `elo_difference` int(11) GENERATED ALWAYS AS ((`elo_after` - `elo_before`)) VIRTUAL,
  `outcome` enum('draw','loss','win') NOT NULL,
  `played` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `RESULT_PLATFORM_ELO` (
  `result_id` bigint(20) UNSIGNED NOT NULL,
  `platform_elo_id` bigint(20) UNSIGNED NOT NULL,
  `verification_key` varchar(500) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `USER` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(320) NOT NULL,
  `registered` datetime NOT NULL,
  `last_online` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


ALTER TABLE `GROUP`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `league_id` (`league_id`,`league_season_name`,`name`),
  ADD KEY `name` (`name`),
  ADD KEY `league_season_name` (`league_season_name`);

ALTER TABLE `GROUP_USER`
  ADD PRIMARY KEY (`group_id`,`user_id`),
  ADD KEY `user_id` (`user_id`);

ALTER TABLE `LEAGUE`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `platform_elo_id` (`platform_elo_id`,`name`);

ALTER TABLE `LEAGUE_SEASON`
  ADD PRIMARY KEY (`league_id`,`name`),
  ADD KEY `name` (`name`);

ALTER TABLE `PLATFORM`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`,`version`);

ALTER TABLE `PLATFORM_AUTH`
  ADD PRIMARY KEY (`platform_id`);

ALTER TABLE `PLATFORM_ELO`
  ADD PRIMARY KEY (`platform_id`);

ALTER TABLE `PLATFORM_ENDPOINT`
  ADD PRIMARY KEY (`platform_id`);

ALTER TABLE `PLATFORM_USER`
  ADD PRIMARY KEY (`user_id`,`platform_id`),
  ADD KEY `platform_id` (`platform_id`);

ALTER TABLE `RESULT`
  ADD PRIMARY KEY (`id`),
  ADD KEY `group_id` (`group_id`);

ALTER TABLE `RESULT_PLATFORM_ELO`
  ADD PRIMARY KEY (`result_id`),
  ADD KEY `platform_elo_id` (`platform_elo_id`);

ALTER TABLE `USER`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);


ALTER TABLE `GROUP`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE `LEAGUE`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE `PLATFORM`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE `RESULT`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
ALTER TABLE `USER`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

ALTER TABLE `GROUP`
  ADD CONSTRAINT `GROUP_ibfk_1` FOREIGN KEY (`league_id`) REFERENCES `LEAGUE_SEASON` (`league_id`),
  ADD CONSTRAINT `GROUP_ibfk_2` FOREIGN KEY (`league_season_name`) REFERENCES `LEAGUE_SEASON` (`name`);

ALTER TABLE `GROUP_USER`
  ADD CONSTRAINT `GROUP_USER_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `USER` (`id`),
  ADD CONSTRAINT `GROUP_USER_ibfk_2` FOREIGN KEY (`group_id`) REFERENCES `GROUP` (`id`);

ALTER TABLE `LEAGUE`
  ADD CONSTRAINT `LEAGUE_ibfk_1` FOREIGN KEY (`platform_elo_id`) REFERENCES `PLATFORM_ELO` (`platform_id`);

ALTER TABLE `LEAGUE_SEASON`
  ADD CONSTRAINT `LEAGUE_SEASON_ibfk_1` FOREIGN KEY (`league_id`) REFERENCES `LEAGUE` (`id`);

ALTER TABLE `PLATFORM_AUTH`
  ADD CONSTRAINT `PLATFORM_AUTH_ibfk_1` FOREIGN KEY (`platform_id`) REFERENCES `PLATFORM` (`id`);

ALTER TABLE `PLATFORM_ELO`
  ADD CONSTRAINT `PLATFORM_ELO_ibfk_1` FOREIGN KEY (`platform_id`) REFERENCES `PLATFORM` (`id`);

ALTER TABLE `PLATFORM_ENDPOINT`
  ADD CONSTRAINT `PLATFORM_ENDPOINT_ibfk_1` FOREIGN KEY (`platform_id`) REFERENCES `PLATFORM` (`id`);

ALTER TABLE `PLATFORM_USER`
  ADD CONSTRAINT `PLATFORM_USER_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `USER` (`id`),
  ADD CONSTRAINT `PLATFORM_USER_ibfk_2` FOREIGN KEY (`platform_id`) REFERENCES `PLATFORM` (`id`);

ALTER TABLE `RESULT`
  ADD CONSTRAINT `RESULT_ibfk_1` FOREIGN KEY (`group_id`) REFERENCES `GROUP` (`id`);

ALTER TABLE `RESULT_PLATFORM_ELO`
  ADD CONSTRAINT `RESULT_PLATFORM_ELO_ibfk_1` FOREIGN KEY (`platform_elo_id`) REFERENCES `PLATFORM_ELO` (`platform_id`),
  ADD CONSTRAINT `RESULT_PLATFORM_ELO_ibfk_2` FOREIGN KEY (`result_id`) REFERENCES `RESULT` (`id`);
