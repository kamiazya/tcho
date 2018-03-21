-- --------------
-- Link
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_link` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`title` TEXT,
	`url` TEXT,
	-- Description
	`description_type` INTEGER,
	`description_text` TEXT,
	`created` DATETIME
);

CREATE INDEX `idx_link_created` ON `tbl_link` (`created`);
