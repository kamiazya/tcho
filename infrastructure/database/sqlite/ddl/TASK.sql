-- --------------
-- Task
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_task` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`title` TEXT,
	`type` INTEGER,
	`status` INTEGER,

	`all_day` BOOLEAN,
	`deadline` DATETIME,

	`custom_color` TEXT,

	`location_id` INTEGER,

	`created` DATETIME,
	`modified` DATETIME
);

CREATE INDEX `idx_task_status` ON `tbl_task` (`type`, `status`);

CREATE INDEX `idx_task_location` ON `tbl_task` (`location_id`);

CREATE INDEX `idx_task_title` ON `tbl_task` (`title`);

CREATE TABLE IF NOT EXISTS `tbl_task_relation` (
	`parent_task_id` INTEGER,
	`child_task_id` INTEGER,
	PRIMARY KEY (`parent_task_id`, `child_task_id`)
);

CREATE TABLE IF NOT EXISTS `tbl_task_attachment_relation` (
	`task_id` INTEGER,
	`attachment_id` INTEGER,
	PRIMARY KEY (`task_id`, `attachment_id`)
);

CREATE TABLE IF NOT EXISTS `tbl_task_tag_relation` (
	`task_id` INTEGER,
	`tag_id` INTEGER,
	PRIMARY KEY (`task_id`, `tag_id`)
);
