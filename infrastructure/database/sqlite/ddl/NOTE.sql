-- --------------
-- Note
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_note` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`title` TEXT,
	-- Body
	`body_type` INTEGER,
	`body_text` TEXT,
	`created` DATETIME,
	`modified` DATETIME
);

CREATE INDEX `idx_note_created` ON `tbl_note` (`created`);
CREATE INDEX `idx_note_modified` ON `tbl_note` (`modified`);

CREATE TABLE IF NOT EXISTS `tbl_note_link_relation` (
	`link_id` INTEGER,
	`tag_id` INTEGER,
	PRIMARY KEY (`link_id`, `tag_id`)
);

CREATE TABLE IF NOT EXISTS `tbl_note_tag_relation` (
	`note_id` INTEGER,
	`tag_id` INTEGER,
	PRIMARY KEY (`note_id`, `tag_id`)
);

CREATE TABLE IF NOT EXISTS `tbl_note_attachment_relation` (
	`attachment_id` INTEGER,
	`tag_id` INTEGER,
	PRIMARY KEY (`attachment_id`, `tag_id`)
);

