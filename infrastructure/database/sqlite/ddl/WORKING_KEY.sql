-- --------------
-- WorkingKey
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_working_key` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`type` INTEGER,
	`name` TEXT UNIQUE,
	-- Description
	`description_type` INTEGER,
	`description_text` TEXT,
	`created` DATETIME,
	`modified` DATETIME
);

CREATE INDEX `idx_working_key_name` ON `tbl_working_key` (`name`);
