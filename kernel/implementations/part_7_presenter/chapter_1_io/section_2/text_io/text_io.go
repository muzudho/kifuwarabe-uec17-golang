package text_io

import (
	"fmt"

	logger "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/implementations/part_7_presenter/chapter_1_io/section_1/logger"
)

// TextIO - テキスト入出力
type TextIO struct {
	// ロガー
	log1 *logger.Logger
}

func NewTextIO(log1 *logger.Logger) *TextIO {
	var t = new(TextIO)
	t.log1 = log1
	return t
}

func (t *TextIO) SendCommand(command string) {
	fmt.Print(command)
	//t.log1.C.Info(command)
}

func (t *TextIO) ReceivedCommand(command string) {
	// FIXME: 大会の邪魔になるのでは？
	//text_io.SendCommand(fmt.Sprintf("# %s", command))	// 人間向けの出力

	t.log1.J.Infow("input", "command", command)
}
