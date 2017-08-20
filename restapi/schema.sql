# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: localhost (MySQL 5.7.18)
# Datenbank: testtt
# Erstellt am: 2017-08-15 15:16:21 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Export von Tabelle course
# ------------------------------------------------------------
USE ews;

USE ewsdb;
DROP TABLE IF EXISTS `course`;

CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `instructor_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `instructor_id` (`instructor_id`),
  CONSTRAINT `course_ibfk_1` FOREIGN KEY (`instructor_id`) REFERENCES `instructor` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;

INSERT INTO `course` (`id`, `name`, `date`, `instructor_id`)
VALUES
	(1,'schwimmen','2017-08-07 16:42:03',1),
	(2,'kochen','2017-08-07 16:42:03',1),
	(3,'kochen','2017-08-07 16:42:03',1),
	(4,'kochen','2017-08-07 16:42:03',1);

/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;


# Export von Tabelle instructor
# ------------------------------------------------------------

DROP TABLE IF EXISTS `instructor`;

CREATE TABLE `instructor` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `apikey` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `apikey` (`apikey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `instructor` WRITE;
/*!40000 ALTER TABLE `instructor` DISABLE KEYS */;

INSERT INTO `instructor` (`id`, `name`, `password`, `firstname`, `lastname`, `apikey`)
VALUES
(1,'Klaus','klaus12345','Klaus','Kleber','M7hfjphd3abDUwxwxt8r4gO5q'),
(2,'Peter','peter12345','Peter','Scheber','rtGI6Cfvhdr1F05yVIHF0wOAA'),
(3,'Lennart','lennart1345','Lennart','Sweger','IefdrM5kMmpbSONLYX7Vr6ksz'),
(4,'Adele','adele12345','Adele','Anika','cJezmSKnGao6wv06Rh6beNXdl'),
(5,'Angelina','angelina12345','Angelina','Bertold','0D5BIR7uI9cjVC2aHjjgVDelj'),
(6,'Anne','anne12345','Anne','Gros','0NKNdUKxXgGAVj9hv6ZdgA9Oe'),
(7,'Madlen','madlen12345','Madlen','Knecht','HG3pZddhWqNUO8AyV2BEvI3SM'),
(8,'Maja','maja12345','Maja','Blub','rVWAK3k0Eh2XzdcJjV4HZeEoy'),
(9,'Maxin','maxin12345','Maxin','Thor','1jfOPrwiSeeVUypCTS9psPDzk'),
(10,'Tia','tia12345','Tia','Rubble','mclJxwLzIsnN5O61PeSbU5j1b');

/*!40000 ALTER TABLE `instructor` ENABLE KEYS */;
UNLOCK TABLES;


# Export von Tabelle participant
# ------------------------------------------------------------

DROP TABLE IF EXISTS `participant`;

CREATE TABLE `participant` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `qrhash` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  UNIQUE KEY `qrhash` (`qrhash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `participant` WRITE;
/*!40000 ALTER TABLE `participant` DISABLE KEYS */;

INSERT INTO `participant` (`id`, `name`, `password`, `firstname`, `lastname`, `qrhash`)
VALUES
(1,'Ulla','ulla123','Ulla','Urte','SskVvFqGlDYy2vBSxF8wgYyS0'),
(2,'Ulrike','ulrike123','Ulrike','Grub','JNygMLPyAR9czjttaK9CVRIkT'),
(3,'Undine','undine123','Undine','Grub','L35td3HR7JsJPNiOY73MAM45T'),
(4,'Ursula','ursula123','Ursila','Grub','asENjnjr07c4TWItbnUlsgvSb'),
(5,'Urte','urte123','Urte','Grub','61XbC6LQXRBHDuD2gNdrVm3je'),
(6,'Uschi','uschi123','Uschi','Grub','vrcpZJTMb4Tc7d7thEpMEVVIg'),
(7,'Uta','uta123','Uta','Grub','vY6Bxxbb2hUzuiTCRJZy9kRkH'),
(8,'Ute','ute123','Ute','Grub','5o9Sype4b57aSlLmAFNuz5Y4K'),
(9,'Ida','ida123','Ida','Grub','cCizrGhMLVrXqik4ZfqH9th9L'),
(10,'Ina','ina123','Ina','Grub','p88minmGz7RJYlAgTvNw9qtNf'),
(11,'Isa','isa123','Isa','Grub','I2KW6SVWEWVpMaTTTIhbKwccP'),
(12,'Irene','irene123','Irene','Grub','0kDd2FF8XGqZFK0guWosIlVog'),
(13,'Isis','isis123','Isis','Grub','SR4g0MkagU6L7HkYDL1kiQMmW'),
(14,'Ivana','ibana123','Ivana','Grub','kL2n07EK0B6liVLwfxgDpRGNF'),
(15,'Irem','irem123','Irem','Grub','VsMA37DQ1H3vtSXBUCmK2YRAr'),
(16,'Inka','inka123','Inka','Grub','3Oziv0SaC0oNrQl3VSnTaZfFn'),
(17,'Insa','insa123','Insa','Grub','pmg9ZIvKxo3res0XJQcqEY8lL'),
(18,'Hanna','hanna123','Hanna','Grub','uHiUAwEnIgsq7Yw4u2XvPgSa2'),
(19,'Hellen','hellen123','Hellen','Grub','TITzRMrGQC0BHtlwGbO6HsnOw');

/*!40000 ALTER TABLE `participant` ENABLE KEYS */;
UNLOCK TABLES;

DROP TABLE IF EXISTS `courseparticipant`;

CREATE TABLE `courseparticipant` (
  `courseid` int(11) NOT NULL,
  `participantid` int(11) NOT NULL,
  KEY `courseid` (`courseid`),
  KEY `participantid` (`participantid`),
  CONSTRAINT `courseid` FOREIGN KEY (`courseid`) REFERENCES `course` (`id`),
  CONSTRAINT `participantid` FOREIGN KEY (`participantid`) REFERENCES `participant` (`id`),
  CONSTRAINT pk_courseparticipant PRIMARY KEY (`courseid`,`participantid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `courseparticipant` WRITE, `participant` WRITE;

/*!40000 ALTER TABLE `participant` DISABLE KEYS */;

INSERT INTO `courseparticipant` (`courseid`, `participantid`)
VALUES
	(1,1),
  (1,2),
  (1,3),
  (2,4),
  (2,5),
  (2,6),
  (2,7),
  (2,8),
  (3,9),
  (3,10),
  (3,11),
  (3,12),
  (3,15),
  (4,16),
  (4,17),
  (4,18);

  /*!40000 ALTER TABLE `participant` ENABLE KEYS */;
UNLOCK TABLES;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
