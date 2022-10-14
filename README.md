# go-datagran

## Tecnologia Utilizada

- GO

## Instalacion

1. Clonar el repositorio en el folder del servidor web en uso o en el de su elección, este folder debe tener permisos para que php se pueda ejecutar por CLI y permisos de lectura y escritura para el archivo .env.

```
https://github.com/alejog1582//go-datagran.git
```
2. Crear base de datos:

-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: localhost    Database: datagran
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `datagran_modelo_1`
--

DROP TABLE IF EXISTS `datagran_modelo_1`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `datagran_modelo_1` (
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `tienda_id` varchar(255) NOT NULL,
  `codigo_producto` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `rating_prediction` varchar(255) NOT NULL,
  `rp_fuente` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `datagran_modelo_1`
--

LOCK TABLES `datagran_modelo_1` WRITE;
/*!40000 ALTER TABLE `datagran_modelo_1` DISABLE KEYS */;
/*!40000 ALTER TABLE `datagran_modelo_1` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-10-14 16:38:16


3. Crear el .env. con las respectivas credenciales de base de datos

## Descripcion del Proyecto

- El proyecto consume una api de Datagran las cuales trae los productos mas comprados por cada una de las tiendas asi como tambien los productos que mas consumen las tiendas cercabas a 500 metros
- Por medio de concurrencias se envian 5 pedidos a la vez y cada pedido trae de 100 registros
- Los registsros son guradados en base de datos
- Al finalizar el consumo de los registros se envia una notificacion a MATTERMOST que el proceso ha finalizado con el numero de registros totales traidos y guardados de la API


## Autor

Jose Alejandro Gonzalez Rondon alejog1582@gmail.com. 
