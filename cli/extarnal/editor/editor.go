package editor

import (
	"os"

	"bitbucket.org/kamiazya/tcho/core/domain/value/memo"
)

type TechoCliMemoEditor interface {
	NewMemo(t memo.TextType) (file *os.File, err error)
	EditMemo(memo memo.Memo) (file *os.File, err error)
}

var Default = &vimEditor{}
