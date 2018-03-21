.PHONY: sqlite doc

NO_MAKE_DIR= ./.git% ./data% ./doc%

PACKAGE     = bitbucket.org/kamiazya/tcho
PACKAGE_NAME     = tcho
BASE        = $(GOPATH)/src/$(PACKAGE)
# tcho_DIR  = $(HOME)/.config/tcho

SUB_DIRS = $(shell find . -type d)
SUB_PKGS = $(filter-out $(NO_MAKE_DIR), $(SUB_DIRS))

sqlite:
	sqlite3 ./.tcho.sqlite < ./infrastructure/database/sqlite/ddl/SCHEDULE.sql

doc:
	@SUB_PKGS="$(SUB_PKGS)";\
	for PKG in $$SUB_PKGS; do\
		mkdir ../tcho-doc/$(PACKAGE_NAME)/`echo $$PKG | cut -c 3-`;\
		godepgraph -s $(PACKAGE)/`echo $$PKG | cut -c 3-` | dot -Kdot -Grankdir=LR -Tpng -o ../tcho-doc/$(PACKAGE_NAME)/`echo $$PKG | cut -c 3-`/deps.png;\
	done

deps:
	 go run tools/deps/main.go | dot -Kdot -Grankdir=LR -Tpng -Glayout=dot  -o test.png; open test.png
