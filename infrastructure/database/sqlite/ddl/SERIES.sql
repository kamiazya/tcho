
-- --------------
-- Series
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_series` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,

	`week0` BOOLEAN,
	`week1` BOOLEAN,
	`week2` BOOLEAN,
	`week3` BOOLEAN,
	`week4` BOOLEAN,
	`week5` BOOLEAN,
	`week6` BOOLEAN,

	`month0` BOOLEAN,
	`month1` BOOLEAN,
	`month2` BOOLEAN,
	`month3` BOOLEAN,
	`month4` BOOLEAN,
	`month5` BOOLEAN,

	`title` TEXT,
	-- Description
	`description_type` INTEGER,
	`description_text` TEXT,

	`all_day` BOOLEAN,
	`start` DATETIME,
	`end` DATETIME,

	`task_id` INTEGER,

	`custom_color` TEXT,

	`location_id` INTEGER,

	`created` DATETIME,
	`modified` DATETIME
);
