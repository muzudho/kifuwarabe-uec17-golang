package main

import (
	"fmt"

	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
	stone "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_3_physical/sublevel_1/stone"

	game_rule_settings "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_4_game_rule/sublevel_1/game_rule_settings"

	kernel_core "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_31_controller"

	logger "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/implementations/part_7_presenter/chapter_1_io/section_1/logger"

	moves_num "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_2_conceptual/sublevel_1/moves_num"

	"strings"

	komi_float "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/level_2_conceptual/sublevel_1/komi_float"
)

func LoopGTP(text_io1 i_text_io.ITextIO, log1 *logger.Logger, engineConfig *Config) {
	// [O12o__11o_4o0] 棋譜の初期化に利用
	var onUnknownTurn = func() stone.Stone {
		var errMsg = fmt.Sprintf("? unexpected play_first:%s", engineConfig.GetPlayFirst())
		text_io1.SendCommand(errMsg)
		log1.J.Infow("error", "play_first", engineConfig.GetPlayFirst())
		panic(errMsg)
	}

	// [O11o_3o0]
	var gameRuleSettings = game_rule_settings.NewGameRuleSettings(komi_float.KomiFloat(engineConfig.GetKomi()), moves_num.MovesNum(engineConfig.GetMaxPositionNumber()))
	var kernel1 = kernel_core.NewDirtyKernel(*gameRuleSettings, engineConfig.GetBoardSize(), engineConfig.GetBoardSize(),
		// [O12o__11o_4o0] 棋譜の初期化
		moves_num.MovesNum(engineConfig.GetMaxPositionNumber()),
		stone.GetStoneOrDefaultFromTurn(engineConfig.GetPlayFirst(), onUnknownTurn))
	// 設定ファイルの内容をカーネルへ反映
	kernel1.Position.Board.Init(engineConfig.GetBoardSize(), engineConfig.GetBoardSize())

	// [O11o_1o0] コンソール等からの文字列入力
	for virtualIo.ScannerScan() {
		var command = virtualIo.ScannerText()
		text_io1.ReceivedCommand(command)

		// [O11o_3o0]
		var isHandled = kernel1.ReadCommand(command, text_io1, log1)
		if isHandled {
			continue
		}

		// [O11o_1o0]
		var tokens = strings.Split(command, " ")
		switch tokens[0] {

		// この下にコマンドを挟んでいく
		// -------------------------

		case "quit": // [O11o_1o0]
			// os.Exit(0)
			return

		// この上にコマンドを挟んでいく
		// -------------------------

		default: // [O11o_1o0]
			text_io1.SendCommand(fmt.Sprintf("? unknown_command command:'%s'\n", tokens[0]))
			log1.J.Infow("? unknown_command", "command", tokens[0])
		}
	}
}
