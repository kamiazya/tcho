-- --------------
-- Kpt
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_kpt` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`type` INTEGER,
	`status` INTEGER,

	-- Content
	`content_type` INTEGER,
	`content_text` TEXT,

	`created` DATETIME,
	`modified` DATETIME
);

