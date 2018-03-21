.PHONY: create

tcho_DIR=$(HOME)/.config/tcho

DB_FILE=$(tcho_DIR)/db.sqlite

create:
	sqlite3 $(DB_FILE) < ./ddl/TAG.sql
	sqlite3 $(DB_FILE) < ./ddl/LOCATION.sql
	sqlite3 $(DB_FILE) < ./ddl/LINK.sql
	sqlite3 $(DB_FILE) < ./ddl/ATTACHMENT.sql
	sqlite3 $(DB_FILE) < ./ddl/NOTE.sql
	sqlite3 $(DB_FILE) < ./ddl/WORKING_KEY.sql
	sqlite3 $(DB_FILE) < ./ddl/CATEGORY.sql
	sqlite3 $(DB_FILE) < ./ddl/SERIES.sql
	sqlite3 $(DB_FILE) < ./ddl/SCHEDULE.sql
	sqlite3 $(DB_FILE) < ./ddl/TASK.sql
	sqlite3 $(DB_FILE) < ./ddl/WORKING_VALUE.sql
	sqlite3 $(DB_FILE) < ./ddl/KPT.sql
