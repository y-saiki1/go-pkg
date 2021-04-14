-- CREATE DATABASE IF NOT EXISTS `aquarius`

--
-- データベース: `aquarius`
--
START TRANSACTION;
-- --------------------------------------------------------

-- ******************************
-- テーブルの構造 `users`
-- ******************************
CREATE TABLE `users` (
  `id` char(36) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `change_email` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `type` int UNSIGNED NOT NULL COMMENT '[1: 一般, 2: 弁護士, 3: cs]',
  `email_verified_at` timestamp NULL DEFAULT NULL,
  `email_verification_token` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;