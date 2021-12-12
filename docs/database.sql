-- Database: resources

CREATE TABLE `images` (
  `id` int NOT NULL AUTO_INCREMENT,
  `resource_id` int NOT NULL,
  `link` varchar(255) NOT NULL,
  `status` varchar(15) NOT NULL,
  `created_by` int NOT NULL,
  `created_at` bigint NOT NULL,
  `updated_by` int NOT NULL,
  `updated_at` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `resources` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `category` varchar(32) NOT NULL,
  `status` varchar(15) NOT NULL,
  `types` varchar(10) NOT NULL,
  `content` text,
  `file_link` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_by` int NOT NULL,
  `created_at` bigint NOT NULL,
  `updated_by` int NOT NULL,
  `updated_at` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `first_name` varchar(16) NOT NULL,
  `last_name` varchar(16) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(32) NOT NULL,
  `role` varchar(16) NOT NULL,
  `last_login` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;