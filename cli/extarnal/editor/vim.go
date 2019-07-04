package editor

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/kamiazya/tcho/core/domain/value/memo"
)

type vimEditor struct{}

func (e *vimEditor) edit(t memo.TextType, file *os.File) (err error) {
	options := []string{"-c"}
	if t == memo.Markdown {
		options = append(options, "setf markdown")
	} else {
		options = append(options, "setf text")
	}
	options = append(options, file.Name())

	cmd := exec.Command("vim", options...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	return nil
}

func (e *vimEditor) NewMemo(t memo.TextType) (file *os.File, err error) {

	file, err = ioutil.TempFile("", "tcho-cli")
	if err != nil {
		return nil, err
	}

	err = e.edit(t, file)

	return file, err
}

func (e *vimEditor) EditMemo(memo memo.Memo) (file *os.File, err error) {
	file, err = ioutil.TempFile("", "tcho-cli")
	_, err = file.WriteString(memo.Text)
	if err != nil {
		return nil, err
	}
	err = e.edit(memo.Type, file)
	return file, err
}
