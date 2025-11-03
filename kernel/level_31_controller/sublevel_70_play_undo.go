// BOF [O23o1o0]

package level_31_controller

import logger "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/implementations/part_7_presenter/chapter_1_io/section_1/logger"

// DoUndoPlay - 石を打ったのを戻す
//
// * `command` - Example: `undo`
// ........................----
// ........................0
func (kernel1 *Kernel) DoUndoPlay(command string, logg *logger.Logger) {
	kernel1.UndoPlay()
}

// UndoPlay - 石を打ったのを戻す
//
// Returns
// -------
// isOk : bool
//
//	石を置けたら真、置けなかったら偽
func (kernel1 *Kernel) UndoPlay() bool {

	// 初期局面から前には戻せない
	if kernel1.Record.GetMovesNum() < 1 {
		return false
	}

	// TODO 置いた石を盤上から消す
	// TODO アゲハマを取る前の盤上の場所を棋譜に記録しておく。連単位？
	// TODO アゲハマを盤上に戻す
	// TODO 棋譜から一手消す。カレントも減らす

	return false
}

// EOF [O23o1o0]
