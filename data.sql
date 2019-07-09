/*
SQLyog Ultimate v11.21 (64 bit)
MySQL - 5.7.26-0ubuntu0.18.04.1 : Database - decision_data
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`decision_data` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `decision_data`;

/*Table structure for table `decisions` */

DROP TABLE IF EXISTS `decisions`;

CREATE TABLE `decisions` (
  `decision_id` int(5) NOT NULL,
  `decision` char(100) DEFAULT NULL,
  PRIMARY KEY (`decision_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `decisions` */

insert  into `decisions`(`decision_id`,`decision`) values (1,'Camilan'),(2,'Restoran'),(3,'Cafe'),(4,'Homestay'),(5,'Rumah Kost'),(6,'Laundry');

/*Table structure for table `pengetahuan` */

DROP TABLE IF EXISTS `pengetahuan`;

CREATE TABLE `pengetahuan` (
  `pengetahuan_id` int(5) NOT NULL,
  `pengetahuan_nama` char(140) DEFAULT NULL,
  PRIMARY KEY (`pengetahuan_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `pengetahuan` */

insert  into `pengetahuan`(`pengetahuan_id`,`pengetahuan_nama`) values (1,'Tinggi'),(2,'Rendah'),(3,'Panjang'),(4,'Pendek'),(5,'Lama'),(6,'Singkat');

/*Table structure for table `pengetahuan_value` */

DROP TABLE IF EXISTS `pengetahuan_value`;

CREATE TABLE `pengetahuan_value` (
  `pertanyaan_id` int(5) DEFAULT NULL,
  `pengetahuan_id` int(5) DEFAULT NULL,
  `decision_id` int(5) DEFAULT NULL,
  `is_yes` int(5) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `pengetahuan_value` */

insert  into `pengetahuan_value`(`pertanyaan_id`,`pengetahuan_id`,`decision_id`,`is_yes`) values (1,1,1,0),(1,2,1,1),(1,1,2,1),(1,2,2,0),(1,1,3,1),(1,2,3,0),(1,1,4,0),(1,2,4,1),(1,1,5,0),(1,2,5,1),(1,1,6,0),(1,2,6,1),(2,3,1,0),(2,4,1,1),(2,3,2,1),(2,4,2,0),(2,3,3,1),(2,4,3,0),(2,3,4,1),(2,4,4,0),(2,3,5,1),(2,4,5,0),(2,3,6,0),(2,4,6,1),(3,1,1,0),(3,2,1,1),(3,1,2,1),(3,2,2,0),(3,1,3,1),(3,2,3,0),(3,1,4,1),(3,2,4,0),(3,1,5,1),(3,2,5,0),(3,1,6,1),(3,2,6,0),(4,1,1,0),(4,2,1,1),(4,1,2,1),(4,2,2,0),(4,1,3,1),(4,2,3,0),(4,1,4,1),(4,2,4,0),(4,1,5,1),(4,2,5,0),(4,1,6,0),(4,2,6,1),(5,5,1,0),(5,6,1,1),(5,5,2,1),(5,6,2,0),(5,5,3,1),(5,6,3,0),(5,5,4,0),(5,6,4,1),(5,5,5,0),(5,6,5,1),(5,5,6,0),(5,6,6,1),(6,1,1,0),(6,2,1,1),(6,1,2,1),(6,2,2,0),(6,1,3,1),(6,2,3,0),(6,1,4,1),(6,2,4,0),(6,1,5,1),(6,2,5,0),(6,1,6,0),(6,2,6,1);

/*Table structure for table `pertanyaan` */

DROP TABLE IF EXISTS `pertanyaan`;

CREATE TABLE `pertanyaan` (
  `pertanyaan_id` int(5) NOT NULL,
  `pertanyaan_string` varchar(300) DEFAULT NULL,
  PRIMARY KEY (`pertanyaan_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `pertanyaan` */

insert  into `pertanyaan`(`pertanyaan_id`,`pertanyaan_string`) values (1,'Tingkat resiko kerugian'),(2,'Jangka waktu usaha'),(3,'Tingkat keuntungan'),(4,'Modal yang diperlukan'),(5,'Pengalaman anda berwirausaha'),(6,'Tingkat penggunaan teknologi informasi di daerah anda');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
