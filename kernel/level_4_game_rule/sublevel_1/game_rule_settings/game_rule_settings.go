package kernel

import (
	// Level 2.1
	komi_float "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_2_conceptual/sublevel_1/komi_float"
	moves_num "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_2_conceptual/sublevel_1/moves_num"
)

// GameRuleSettings - 対局ルール設定
type GameRuleSettings struct {
	// コミ。 6.5 といった数字を入れるだけ。実行速度優先で 64bitに
	Komi komi_float.KomiFloat

	// 上限手数
	maxMovesNum moves_num.MovesNum
}

// NewGameRuleSettings - 新規作成
func NewGameRuleSettings(komi komi_float.KomiFloat, maxMovesNum moves_num.MovesNum) *GameRuleSettings {
	var gr = new(GameRuleSettings)

	gr.Komi = komi
	gr.maxMovesNum = maxMovesNum

	return gr
}

// GetKomi - コミ取得
func (gr *GameRuleSettings) GetKomi() komi_float.KomiFloat {
	return gr.Komi
}

// GetMaxPositionNumber - 上限手数
func (gr *GameRuleSettings) GetMaxPositionNumber() moves_num.MovesNum {
	return gr.maxMovesNum
}
