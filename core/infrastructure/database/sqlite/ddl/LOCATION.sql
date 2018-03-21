-- --------------
-- Location
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_location` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`name` TEXT,
	-- Description
	`description_type` INTEGER,
	`description_text` TEXT
);

CREATE INDEX `idx_location_name` ON `tbl_location` (`name`);
