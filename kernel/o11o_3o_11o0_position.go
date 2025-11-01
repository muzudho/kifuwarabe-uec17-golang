package kernel

import game_rule_settings "github.com/muzudho/kifuwarabe-uec17/kernel/types/level2/game_rule_settings"

type Position struct {
	// Board - 盤
	Board *Board

	// [O22o2o3o0]
	// CheckBoard - 呼吸点の探索時に使います
	CheckBoard *CheckBoard

	// CanNotPutOnMyEye - [O22o4o1o0] 自分の眼に石を置くことはできません
	CanNotPutOnMyEye bool
}

// NewDirtyKernel - カーネルの新規作成
// - 一部のメンバーは、初期化されていないので、別途初期化処理が要る
func NewDirtyPosition(gameRuleSettings game_rule_settings.GameRuleSettings, boardWidht int, boardHeight int) *Position {
	var p = new(Position)

	p.Board = NewBoard(gameRuleSettings, boardWidht, boardHeight)

	// [O22o2o3o0] チェックボード
	p.CheckBoard = NewDirtyCheckBoard()

	return p
}
