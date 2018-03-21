-- --------------
-- Label
-- --------------
CREATE TABLE IF NOT EXISTS `tbl_label` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`model_type` INTEGER,
	`model_id` INTEGER,

	`value` TEXT,
	`maintainer` TEXT,

	`created` DATETIME
);

CREATE INDEX `idx_label_model` ON `tbl_label` (`model_type`, `model_id`);
CREATE INDEX `idx_label_maintainer` ON `tbl_label` (`maintainer`);
CREATE INDEX `idx_label_value` ON `tbl_label` (`value`);
