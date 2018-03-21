-- --------------
-- Tag
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_tag` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` TEXT UNIQUE,
	`color` TEXT,
	-- Description
	`description_type` INTEGER,
	`description_text` TEXT,
	`created` DATETIME
);

CREATE INDEX `idx_tag_name` ON `tbl_tag` (`name`);
