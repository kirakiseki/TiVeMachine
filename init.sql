-- database
DROP DATABASE IF EXISTS `tivemachine`;

CREATE DATABASE `tivemachine` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

USE `tivemachine`;

-- tables
CREATE TABLE `user` (
  `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `sex` int(1) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- data
INSERT INTO `user` (`username`,`password`) VALUES ('ishirai','y1YB9mYC9+/gbANs79+cZgUGVeLTT0EB4w5h3xwYAgA=');