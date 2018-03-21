.PHONY: run run-ja create clean deps doc

PACKAGE_NAME     = tcho-cli
PACKAGE     = bitbucket.org/kamiazya/$(PACKAGE_NAME)
VERSION = 0.0.0
Revision = $(shell git rev-parse HEAD)


NO_MAKE_DIR= ./.git% ./data% ./doc%

DB_FILE=./data/tcho.sqlite

SUB_DIRS = $(shell find . -type d)
SUB_PKGS = $(filter-out $(NO_MAKE_DIR), $(SUB_DIRS))


run:
	go run main.go app.go version.go

run-ja:
	go run -tags=ja main.go app.go version.go

create:
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/TAG.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/LOCATION.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/LINK.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/ATTACHMENT.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/NOTE.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/WORKING_KEY.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/CATEGORY.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/SERIES.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/SCHEDULE.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/TASK.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/WORKING_VALUE.sql
	sqlite3 $(DB_FILE) < $(MEMOMA_PACKAGE)/infrastructure/database/sqlite/ddl/KPT.sql

clean:
	@rm ./data/tcho.sqlite

deps:
	godepgraph -s bitbucket.org/kamiazya/tcho-cli | dot -Kdot -Grankdir=LR -Tpng -o deps.png

doc:
	@SUB_PKGS="$(SUB_PKGS)";\
	for PKG in $$SUB_PKGS; do\
		mkdir ../tcho-doc/$(PACKAGE_NAME)/`echo $$PKG | cut -c 3-`;\
		godepgraph -s $(PACKAGE)/`echo $$PKG | cut -c 3-` | dot -Kdot -Grankdir=LR -Tpng -o ../tcho-doc/$(PACKAGE_NAME)/`echo $$PKG | cut -c 3-`/deps.png;\
	done
