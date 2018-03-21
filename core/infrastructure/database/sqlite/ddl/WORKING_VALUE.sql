-- --------------
-- WorkingValue
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_working_value` (
	`key_id` INTEGER,
	`task_id` INTEGER,
	`initial_value` INTEGER,
	`value` INTEGER,
	`created` DATETIME,
	`modified` DATETIME,
	PRIMARY KEY (`key_id`, `task_id`)
);
