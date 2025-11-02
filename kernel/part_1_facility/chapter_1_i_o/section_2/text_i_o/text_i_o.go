package text_io

// Section 1.1.1
import logger "github.com/muzudho/kifuwarabe-uec17-golang/kernel/part_1_facility/chapter_1_i_o/section_1/logger"

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

func (t *TextIO) GoCommand(command string) {
	t.log1.C.Info(command)
}
