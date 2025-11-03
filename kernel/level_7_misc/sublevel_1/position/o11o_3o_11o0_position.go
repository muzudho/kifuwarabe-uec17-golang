package position

import (
	// Entities
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/implementations/part_1_entities/chapter_2_rule_settings/section_1/game_rule_settings"

	// Level 4.1
	check_board "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_4_game_rule/sublevel_1/check_board"

	// Level 4.2
	board "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_4_game_rule/sublevel_2/board"
)

type Position struct {
	// Board - 盤
	Board *board.Board

	// [O22o2o3o0]
	// CheckBoard - 呼吸点の探索時に使います
	CheckBoard *check_board.CheckBoard

	// CanNotPutOnMyEye - [O22o4o1o0] 自分の眼に石を置くことはできません
	CanNotPutOnMyEye bool
}

// NewDirtyKernel - カーネルの新規作成
// - 一部のメンバーは、初期化されていないので、別途初期化処理が要る
func NewDirtyPosition(gameRuleSettings game_rule_settings.GameRuleSettings, boardWidht int, boardHeight int) *Position {
	var p = new(Position)

	p.Board = board.NewBoard(gameRuleSettings, boardWidht, boardHeight)

	// [O22o2o3o0] チェックボード
	p.CheckBoard = check_board.NewDirtyCheckBoard()

	return p
}
