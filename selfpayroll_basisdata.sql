-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 15, 2023 at 05:47 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.1.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `selfpayroll_basisdata`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `id` bigint(20) NOT NULL,
  `admin_id` char(36) DEFAULT uuid(),
  `email` varchar(300) DEFAULT NULL,
  `name` varchar(300) DEFAULT NULL,
  `password` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`id`, `admin_id`, `email`, `name`, `password`) VALUES
(1, 'e3bb2838-9a80-11ee-b159-50ebf69182ee', 'admin1@dpr.co.id', 'Dionixius', '$2a$14$E0GZQZBElBgwTwDTGiwdte1KdkrMnlsg./aGNzgn3H0EIY5OyoKpC'),
(2, 'f6fe34bc-9a80-11ee-b159-50ebf69182ee', 'admin2@dpr.co.id', 'Riki Widyo Laksono', '$2a$14$6zdNx2AMAjz1mbtSdzE0puOW8CPjQg0s3V2hkYikSAZpZ5H6akbPW'),
(3, 'fcb17e60-9a80-11ee-b159-50ebf69182ee', 'admin3@dpr.co.id', 'Pandu Violana Mulya', '$2a$14$.7uA8fbPwRhvugs7NGUCSu40Wv6uFIaWovZlg57giGC/4PRk9W7vW');

-- --------------------------------------------------------

--
-- Table structure for table `company`
--

CREATE TABLE `company` (
  `company_id` bigint(20) NOT NULL,
  `name` varchar(300) DEFAULT NULL,
  `balance` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `company`
--

INSERT INTO `company` (`company_id`, `name`, `balance`, `created_at`, `updated_at`) VALUES
(1, 'PT. Database DPR (persero)', 328000000, '2023-12-14 20:06:13.964', '2023-12-14 20:28:30.154');

-- --------------------------------------------------------

--
-- Table structure for table `employee`
--

CREATE TABLE `employee` (
  `id` bigint(20) NOT NULL,
  `employee_id` char(36) DEFAULT uuid(),
  `name` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `phone` longtext DEFAULT NULL,
  `address` longtext DEFAULT NULL,
  `private_pin` longtext DEFAULT NULL,
  `position_id` bigint(20) DEFAULT NULL,
  `is_paid` tinyint(1) DEFAULT 0,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `employee`
--

INSERT INTO `employee` (`id`, `employee_id`, `name`, `email`, `phone`, `address`, `private_pin`, `position_id`, `is_paid`, `created_at`, `updated_at`) VALUES
(1, '7777e73a-9a83-11ee-b159-50ebf69182ee', 'Rizky Ari', 'rizkyari@dpr.co.id', '081100001111', 'Jl. Ambarawa 1, no.1', '001122', 2, 0, '2023-12-14 20:19:53.080', '2023-12-14 20:26:36.433'),
(2, '9e5c50f3-9a83-11ee-b159-50ebf69182ee', 'Galang Putra', 'galangputra@dpr.co.id', '081100001112', 'Jl. Ambarawa 1, no.2', '001123', 9, 0, '2023-12-14 20:20:58.330', '2023-12-14 20:27:16.733'),
(3, 'aa8a6609-9a83-11ee-b159-50ebf69182ee', 'Riko Laksono', 'rikolaksono@dpr.co.id', '081100001113', 'Jl. Ambarawa 1, no.3', '001124', 10, 0, '2023-12-14 20:21:18.765', '2023-12-14 20:26:04.126'),
(4, 'b9cee6bc-9a83-11ee-b159-50ebf69182ee', 'Fajar Febryanto', 'fajarfebryanto@dpr.co.id', '081100001114', 'Jl. Ambarawa 1, no.4', '001125', 1, 0, '2023-12-14 20:21:44.380', '2023-12-14 20:27:58.145'),
(5, 'c8acbde3-9a83-11ee-b159-50ebf69182ee', 'Kevin Ardiansah', 'kevinardiansah@dpr.co.id', '081100001115', 'Jl. Ambarawa 1, no.5', '001126', 4, 1, '2023-12-14 20:22:09.322', '2023-12-14 20:27:27.352'),
(6, 'dbd495bf-9a83-11ee-b159-50ebf69182ee', 'Satrio Wiryo', 'satriowiryo@dpr.co.id', '081100001116', 'Jl. Ambarawa 1, no.6', '001127', 12, 1, '2023-12-14 20:22:41.459', '2023-12-14 20:26:58.324'),
(7, 'ed0bb329-9a83-11ee-b159-50ebf69182ee', 'Wirda', 'wirda@dpr.co.id', '081100001117', 'Jl. Ambarawa 1, no.7', '001128', 3, 1, '2023-12-14 20:23:10.342', '2023-12-14 20:26:48.569'),
(8, 'f93ae90d-9a83-11ee-b159-50ebf69182ee', 'Kharisma Putri', 'kharismaputri@dpr.co.id', '081100001118', 'Jl. Ambarawa 1, no.8', '001129', 5, 0, '2023-12-14 20:23:30.784', '2023-12-14 20:28:18.162'),
(9, '0e48399e-9a84-11ee-b159-50ebf69182ee', 'Meidy Indhira', 'meidyindhira@dpr.co.id', '081100001119', 'Jl. Ambarawa 1, no.9', '001130', 14, 0, '2023-12-14 20:24:06.103', '2023-12-14 20:26:26.495'),
(10, '229bdbb1-9a84-11ee-b159-50ebf69182ee', 'Arlen Brilianisa', 'arlenbrilianisa@dpr.co.id', '081100001120', 'Jl. Ambarawa 1, no.10', '001131', 8, 1, '2023-12-14 20:24:40.206', '2023-12-14 20:28:30.143');

-- --------------------------------------------------------

--
-- Table structure for table `position`
--

CREATE TABLE `position` (
  `id` bigint(20) NOT NULL,
  `name_position` varchar(36) DEFAULT NULL,
  `salary` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `position`
--

INSERT INTO `position` (`id`, `name_position`, `salary`, `created_at`, `updated_at`) VALUES
(1, 'DevOps Engineer', 50000000, '2023-12-14 20:07:51.691', '2023-12-14 20:07:51.691'),
(2, 'Electrical Engineer', 25000000, '2023-12-14 20:07:59.525', '2023-12-14 20:07:59.525'),
(3, 'Machine Learning Engineer', 50000000, '2023-12-14 20:08:06.168', '2023-12-14 20:08:06.168'),
(4, 'Security', 3000000, '2023-12-14 20:08:22.398', '2023-12-14 20:08:22.398'),
(5, 'Staff HR', 4500000, '2023-12-14 20:08:31.665', '2023-12-14 20:08:31.665'),
(6, 'Janitor', 2000000, '2023-12-14 20:08:53.416', '2023-12-14 20:08:53.415'),
(7, 'Accounting Staff', 5000000, '2023-12-14 20:09:09.240', '2023-12-14 20:09:09.240'),
(8, 'Junior Data Analyst', 7500000, '2023-12-14 20:09:24.919', '2023-12-14 20:09:24.918'),
(9, 'Senior Data Analyst', 15000000, '2023-12-14 20:09:34.327', '2023-12-14 20:09:34.327'),
(10, 'IT Support', 10000000, '2023-12-14 20:10:18.230', '2023-12-14 20:10:18.229'),
(11, 'Front End Developer', 12500000, '2023-12-14 20:11:12.258', '2023-12-14 20:11:12.258'),
(12, 'Back End Developer', 25000000, '2023-12-14 20:11:29.831', '2023-12-14 20:11:29.830'),
(13, 'UI/UX Designer', 17500000, '2023-12-14 20:11:46.471', '2023-12-14 20:11:46.470'),
(14, 'Business Analyst', 17500000, '2023-12-14 20:12:05.137', '2023-12-14 20:12:05.136'),
(15, 'Marketing', 8000000, '2023-12-14 20:12:50.309', '2023-12-14 20:12:50.309');

-- --------------------------------------------------------

--
-- Table structure for table `transaction`
--

CREATE TABLE `transaction` (
  `id` bigint(20) NOT NULL,
  `type` varchar(36) DEFAULT NULL,
  `amount` bigint(20) DEFAULT NULL,
  `note` varchar(36) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transaction`
--

INSERT INTO `transaction` (`id`, `type`, `amount`, `note`, `created_at`, `updated_at`) VALUES
(1, 'debit', 250000000, 'topup saldo perusahaan', '2023-12-14 20:14:44.091', '2023-12-14 20:14:44.091'),
(2, 'kredit', 50000000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:26:04.137', '2023-12-14 20:26:04.137'),
(3, 'kredit', 15000000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:26:26.507', '2023-12-14 20:26:26.507'),
(4, 'kredit', 50000000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:26:36.443', '2023-12-14 20:26:36.443'),
(5, 'kredit', 5000000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:26:48.580', '2023-12-14 20:26:48.580'),
(6, 'kredit', 2000000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:26:58.370', '2023-12-14 20:26:58.370'),
(7, 'kredit', 25000000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:27:16.747', '2023-12-14 20:27:16.747'),
(8, 'kredit', 4500000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:27:27.364', '2023-12-14 20:27:27.364'),
(9, 'kredit', 3000000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:27:58.157', '2023-12-14 20:27:58.157'),
(10, 'kredit', 7500000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:28:18.167', '2023-12-14 20:28:18.167'),
(11, 'kredit', 10000000, 'pencairan gaji bulanan karyawan', '2023-12-14 20:28:30.156', '2023-12-14 20:28:30.156');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `company`
--
ALTER TABLE `company`
  ADD PRIMARY KEY (`company_id`);

--
-- Indexes for table `employee`
--
ALTER TABLE `employee`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_employee_position` (`position_id`);

--
-- Indexes for table `position`
--
ALTER TABLE `position`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction`
--
ALTER TABLE `transaction`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `company`
--
ALTER TABLE `company`
  MODIFY `company_id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `employee`
--
ALTER TABLE `employee`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `position`
--
ALTER TABLE `position`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `transaction`
--
ALTER TABLE `transaction`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `employee`
--
ALTER TABLE `employee`
  ADD CONSTRAINT `fk_employee_position` FOREIGN KEY (`position_id`) REFERENCES `position` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
