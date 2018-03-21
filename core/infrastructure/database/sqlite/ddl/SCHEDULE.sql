-- --------------
-- Schedule
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_schedule` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`series_id` INTEGER,
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

CREATE INDEX `idx_schedule_time` ON `tbl_schedule` (`start`, `end`);

CREATE INDEX `idx_schedule_task` ON `tbl_schedule` (`task_id`);

CREATE INDEX `idx_schedule_location` ON `tbl_schedule` (`location_id`);

CREATE TABLE IF NOT EXISTS `tbl_schedule_note_relation` (
	`schedule_id` INTEGER,
	`note_id` INTEGER,
	PRIMARY KEY (`schedule_id`, `note_id`)
);

CREATE TABLE IF NOT EXISTS `tbl_schedule_attachment_relation` (
	`schedule_id` INTEGER,
	`attachment_id` INTEGER,
	PRIMARY KEY (`schedule_id`, `attachment_id`)
);

CREATE TABLE IF NOT EXISTS `tbl_schedule_tag_relation` (
	`schedule_id` INTEGER,
	`tag_id` INTEGER,
	PRIMARY KEY (`schedule_id`, `tag_id`)
);
