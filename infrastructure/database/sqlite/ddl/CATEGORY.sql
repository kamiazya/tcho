-- --------------
-- Category
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_category` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`type` INTEGER,
	`name` TEXT UNIQUE,
	`color` TEXT,
	-- Description
	`description_type` INTEGER,
	`description_text` TEXT,

	`created` DATETIME,
	`modified` DATETIME
);

CREATE INDEX `idx_category_name` ON `tbl_category` (`name`);

CREATE INDEX `idx_category_type` ON `tbl_category` (`type`);

CREATE TABLE IF NOT EXISTS `tbl_category_default_working_key` (
	`category_id` INTEGER,
	`working_key_id` INTEGER,
	PRIMARY KEY (`category_id`, `working_key_id`)
);
