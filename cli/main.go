package main

import (
	"os"

	"bitbucket.org/kamiazya/tcho-cli/module"
	"bitbucket.org/kamiazya/tcho-cli/module/local"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var m module.Module

	// TODO envや設定ファイルなどからモジュールの実装を切り替える
	m = local.New(app)

	defer m.Close()

	m.Build()

	m.Run(os.Args)
}
