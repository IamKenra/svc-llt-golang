-- MySQL dump 10.13  Distrib 8.4.6, for macos15.4 (arm64)
--
-- Host: 127.0.0.1    Database: layanan_lansia
-- ------------------------------------------------------
-- Server version	8.4.6

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `layanan_lansia`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `layanan_lansia` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `layanan_lansia`;

--
-- Table structure for table `alamat`
--

DROP TABLE IF EXISTS `alamat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `alamat` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `alamat` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `rt` varchar(5) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `rw` varchar(5) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `kelurahan` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `kecamatan` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `kota` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `kodepos` varchar(10) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `flag_aktif` tinyint(1) NOT NULL DEFAULT '1',
  `user_input` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_update` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tgl_input` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `tgl_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_alamat_uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `alamat`
--

LOCK TABLES `alamat` WRITE;
/*!40000 ALTER TABLE `alamat` DISABLE KEYS */;
/*!40000 ALTER TABLE `alamat` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `auth`
--

DROP TABLE IF EXISTS `auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `auth` (
  `id` bigint NOT NULL,
  `username` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `tgl_input` datetime(3) DEFAULT NULL,
  `tgl_update` datetime(3) DEFAULT NULL,
  `user_input` longtext COLLATE utf8mb4_general_ci,
  `user_delete` longtext COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_auth_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth`
--

LOCK TABLES `auth` WRITE;
/*!40000 ALTER TABLE `auth` DISABLE KEYS */;
/*!40000 ALTER TABLE `auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `identitas`
--

DROP TABLE IF EXISTS `identitas`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `identitas` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `nik` varchar(32) COLLATE utf8mb4_general_ci NOT NULL,
  `nama_depan` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
  `nama_belakang` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tgl_lahir` date DEFAULT NULL,
  `agama` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status_perkawinan` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `pekerjaan` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `flag_aktif` tinyint(1) NOT NULL DEFAULT '1',
  `user_input` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_update` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tgl_input` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `tgl_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_biodata_uuid` (`uuid`),
  UNIQUE KEY `uq_biodata_nik` (`nik`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `identitas`
--

LOCK TABLES `identitas` WRITE;
/*!40000 ALTER TABLE `identitas` DISABLE KEYS */;
/*!40000 ALTER TABLE `identitas` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lansia`
--

DROP TABLE IF EXISTS `lansia`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lansia` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `id_identitas` bigint unsigned NOT NULL,
  `id_alamat` bigint unsigned NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `path_gambar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `path_qr` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `level` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `caregiver` varchar(150) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `flag_aktif` tinyint(1) NOT NULL DEFAULT '1',
  `flag_delete` tinyint(1) NOT NULL DEFAULT '0',
  `user_input` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_update` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tgl_input` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `tgl_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_lansia_uuid` (`uuid`),
  KEY `fk_lansia_identitas` (`id_identitas`),
  KEY `fk_lansia_alamat` (`id_alamat`),
  CONSTRAINT `fk_lansia_alamat` FOREIGN KEY (`id_alamat`) REFERENCES `alamat` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_lansia_identitas` FOREIGN KEY (`id_identitas`) REFERENCES `identitas` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lansia`
--

LOCK TABLES `lansia` WRITE;
/*!40000 ALTER TABLE `lansia` DISABLE KEYS */;
/*!40000 ALTER TABLE `lansia` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `riwayat_penyakit`
--

DROP TABLE IF EXISTS `riwayat_penyakit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `riwayat_penyakit` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `id_lansia` bigint unsigned NOT NULL,
  `diagnosa` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tgl_diagnosa` date DEFAULT NULL,
  `user_input` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `user_update` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `tgl_input` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `tgl_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_riwayat_uuid` (`uuid`),
  KEY `fk_riwayat_lansia` (`id_lansia`),
  CONSTRAINT `fk_riwayat_lansia` FOREIGN KEY (`id_lansia`) REFERENCES `lansia` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `riwayat_penyakit`
--

LOCK TABLES `riwayat_penyakit` WRITE;
/*!40000 ALTER TABLE `riwayat_penyakit` DISABLE KEYS */;
/*!40000 ALTER TABLE `riwayat_penyakit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` bigint NOT NULL,
  `uuid` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `nama` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
  `tgl_input` datetime(3) DEFAULT NULL,
  `tgl_update` datetime(3) DEFAULT NULL,
  `user_input` longtext COLLATE utf8mb4_general_ci,
  `user_update` longtext COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_uuid` (`uuid`),
  UNIQUE KEY `idx_user_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'layanan_lansia'
--

--
-- Dumping routines for database 'layanan_lansia'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-16 11:39:59
