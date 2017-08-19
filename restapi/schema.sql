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
<<<<<<< Updated upstream
use ewsdb
=======
USE ews;

>>>>>>> Stashed changes
DROP TABLE IF EXISTS `course`;

CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `participants` varchar(255) NOT NULL,
  `instructor_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `instructor_id` (`instructor_id`),
  CONSTRAINT `course_ibfk_1` FOREIGN KEY (`instructor_id`) REFERENCES `instructor` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;

INSERT INTO `course` (`id`, `name`, `date`, `participants`, `instructor_id`)
VALUES
	(1,'schwimmen','2017-08-07 16:42:03','alle',1),
	(2,'kochen','2017-08-07 16:42:03','',1),
	(3,'kochen','2017-08-07 16:42:03','',1),
	(4,'kochen','2017-08-07 16:42:03','',1);

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
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `instructor` WRITE;
/*!40000 ALTER TABLE `instructor` DISABLE KEYS */;

INSERT INTO `instructor` (`id`, `name`, `password`, `firstname`, `lastname`, `apikey`)
VALUES
	(1,'Klaus','klaus12345','Klaus','Kleber','0m3945c4987ap9mc78tr5c'),
	(2,'Peter','peter12345','Peter','Scheber','uc87943j75c98ev798scd'),
	(3,'Lennart','lennart1345','Lennart','Sweger','09384732840932c837c09er'),
	(4,'Adele','adele12345','Adele','Anika','c485094m385c098m4309v5m3'),
	(5,'Angelina','angelina12345','Angelina','Bertold','ipmci8509m3294oiewurcoinewr'),
	(6,'Anne','anne12345','Anne','Gros','u32n4932nu4cu324coewuoir'),
	(7,'Madlen','madlen12345','Madlen','Knecht','9u45c8329mcueoircsadvct5ezbr6ezb45'),
	(8,'Maja','maja12345','Maja','Blub','irc0ß8m23m098rc9es8urcoiesmrcoe'),
	(9,'Maxin','maxin12345','Maxin','Thor','urc3m4cpoicewpoircpowqimecwq'),
	(10,'Tia','tia12345','Tia','Rubble','omu32pppppesapoufcoineuru932842');

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
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `participant` WRITE;
/*!40000 ALTER TABLE `participant` DISABLE KEYS */;

INSERT INTO `participant` (`id`, `name`, `password`, `firstname`, `lastname`, `qrhash`)
VALUES
	(1,'Ulla','ulla123','Ulla','Urte','56vf4uv5u57u'),
	(2,'Ulrike','ulrike123','Ulrike','Grub','w324gtrcrc43324c245cx43q5'),
	(3,'Undine','undine123','Undine','Grub','mrcointcmtoirezufcp9qw4'),
	(4,'Ursula','ursula123','Ursila','Grub','ucr09m4309rcureciu'),
	(5,'Urte','urte123','Urte','Grub','uc3qufooimearutco'),
	(6,'Uschi','uschi123','Uschi','Grub','uemrc0943urcewoircuewoir'),
	(7,'Uta','uta123','Uta','Grub','u09mce0fewmfcoiwr'),
	(8,'Ute','ute123','Ute','Grub','mu2390u4c093wurcoieurc'),
	(9,'Ida','ida123','Ida','Grub','pm4394u09c0ru3209mc'),
	(10,'Ina','ina123','Ina','Grub','iu5coim43uoicr'),
	(11,'Isa','isa123','Isa','Grub','um532mu04c9u3rcuew'),
	(12,'Irene','irene123','Irene','Grub','c39m204c328u4c9oimruoic'),
	(13,'Isis','isis123','Isis','Grub','ficpomrwirc090mc28m9328rmuewoirc'),
	(14,'Ivana','ibana123','Ivana','Grub','ucromoimrrsac45ze5t'),
	(15,'Irem','irem123','Irem','Grub','wceoiomrowu093284n834cuewifomf'),
	(16,'Inka','inka123','Inka','Grub','rciuo23uc09u32cnourcor'),
	(17,'Insa','insa123','Insa','Grub','uc309ru432rcueruc982u3cm92'),
	(18,'Hanna','hanna123','Hanna','Grub','nu44444444444444cwurp23qmx3zro'),
	(19,'Hellen','hellen123','Hellen','Grub','aisoie7wtrc984wcmasxccsefeta'),
	(20,'Helga','helga123','Helga','Grub','asoicpmtccccwafc98ewp'),
	(99,'Hege','hege123','Hege','Grub','hewkjcslrhmfcshflicore587c09438r509m328c');

/*!40000 ALTER TABLE `participant` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
