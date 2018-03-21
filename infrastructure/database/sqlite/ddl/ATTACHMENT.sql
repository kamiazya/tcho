-- --------------
-- Attachment
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_attachment` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` TEXT,
	`mime` TEXT,

	-- Description
	`description_type` INTEGER,
	`description_text` TEXT,

	`blob` BLOB,

	`created` DATETIME
);

CREATE INDEX `idx_attachment_name` ON `tbl_attachment` (`name`);
CREATE INDEX `idx_attachment_created` ON `tbl_attachment` (`created`);

CREATE TABLE IF NOT EXISTS `tbl_attachment_tag_relation` (
	`attachment_id` INTEGER,
	`tag_id` INTEGER,
	PRIMARY KEY (`attachment_id`, `tag_id`)
);
