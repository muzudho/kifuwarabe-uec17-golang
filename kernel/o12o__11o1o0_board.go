package kernel

import (
	// Level 1
	color "github.com/muzudho/kifuwarabe-uec17/kernel/types/level1/color"
	point "github.com/muzudho/kifuwarabe-uec17/kernel/types/level1/point"

	// Level 2
	board_coordinate "github.com/muzudho/kifuwarabe-uec17/kernel/types/level2/board_coordinate"
	game_rule_settings "github.com/muzudho/kifuwarabe-uec17/kernel/types/level2/game_rule_settings"
	stone "github.com/muzudho/kifuwarabe-uec17/kernel/types/level2/stone"
)

// Board - 盤
type Board struct {
	// ゲームルール
	gameRuleSettings game_rule_settings.GameRuleSettings
	// 盤座標
	coordinate board_coordinate.BoardCoordinate

	// 交点
	//
	// * 英語で交点は node かも知れないが、表計算でよく使われる cell の方を使う
	cells []stone.Stone
}

// NewBoard - 新規作成
func NewBoard(gameRuleSettings game_rule_settings.GameRuleSettings, boardWidht int, boardHeight int) *Board {
	var b = new(Board)

	// 設定ファイルから読込むので動的設定
	b.gameRuleSettings = gameRuleSettings

	// 枠の分、２つ増える
	var memoryBoardWidth = boardWidht + 2
	var memoryBoardHeight = boardHeight + 2

	b.coordinate = board_coordinate.BoardCoordinate{
		MemoryWidth:  memoryBoardWidth,
		MemoryHeight: memoryBoardHeight,
		// ４方向（東、北、西、南）の番地への相対インデックス
		Cell4Directions: [4]point.Point{
			1,
			point.Point(-memoryBoardWidth),
			-1,
			point.Point(memoryBoardWidth),
		},
	}

	// 盤のサイズ指定と、盤面の初期化
	b.resize(boardWidht, boardHeight)

	return b
}

// GetGameRule - ゲームルール取得
func (b *Board) GetGameRule() *game_rule_settings.GameRuleSettings {
	return &b.gameRuleSettings
}

// SetGameRule - ゲームルール設定
func (b *Board) SetGameRule(gameRuleSettings *game_rule_settings.GameRuleSettings) {
	b.gameRuleSettings = *gameRuleSettings
}

// GetCoordinate - 盤座標取得
func (b *Board) GetCoordinate() *board_coordinate.BoardCoordinate {
	return &b.coordinate
}

// GetStoneAt - 指定座標の石を取得
func (b *Board) GetStoneAt(i point.Point) stone.Stone {
	return b.cells[i]
}

// SetStoneAt - 指定座標の石を設定
func (b *Board) SetStoneAt(i point.Point, s stone.Stone) {
	b.cells[i] = s
}

// GetColorAt - 指定座標の石の色を取得
func (b *Board) GetColorAt(i point.Point) color.Color {
	return b.cells[i].GetColor()
}

// IsEmpty - 指定の交点は空点か？
func (b *Board) IsSpaceAt(point point.Point) bool {
	return b.GetStoneAt(point) == stone.Stone_Space
}

// サイズ変更
func (b *Board) resize(width int, height int) {
	b.coordinate.MemoryWidth = width + board_coordinate.BothSidesWallThickness
	b.coordinate.MemoryHeight = height + board_coordinate.BothSidesWallThickness
	b.cells = make([]stone.Stone, b.coordinate.GetMemoryArea())

	// ４方向（東、北、西、南）の番地への相対インデックス
	b.coordinate.Cell4Directions = [4]point.Point{1, point.Point(-b.coordinate.GetMemoryWidth()), -1, point.Point(b.coordinate.GetMemoryWidth())}
}
