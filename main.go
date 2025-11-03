// BOF [O9o0]

package main

import (
	"flag"
	"fmt"
	"os"

	dbg "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/debugger"

	// Implementation
	logger "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/implementations/part_7_presenter/chapter_1_io/section_1/logger"
	text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/implementations/part_7_presenter/chapter_1_io/section_2/text_io"

	// Interfaces
	i_text_io "github.com/muzudho/kifuwarabe-uec17-golang-from-uec14/kernel/interfaces/part_1_facility/chapter_1_io/section_1/i_text_io"
	// Section 1.1.2
	// Level 2.1
	// Level 3.1
	// Level 4.1
	// Level 31
)

// [O11o_1o0] グローバル変数として、バーチャルIOを１つ新規作成
// アプリケーションの中では 標準入出力は これを使うようにする
var virtualIo = dbg.NewVirtualIO()

func main() {
	// ========================================
	// flag - コマンドラインの解析
	// ========================================

	// 思考エンジン設定ファイル
	var (
		pEngineFilePath = flag.String("f", "engine.toml", "engine config file path")
		// [O11o__11o6o0] デバッグ用
		pIsDebug = flag.Bool("d", false, "for debug")
	)
	flag.Parse()
	// コマンドラインの第１引数で処理を振り分け
	var arg1 = flag.Arg(0)

	// ========================================
	// 思考エンジンの準備
	// ========================================

	// 思考エンジン設定ファイル
	var onError = func(err error) *Config {
		// ログファイルには出力できません。ログファイルはまだ読込んでいません

		// 強制終了
		panic(err)
	}
	var engineConfig = LoadEngineConfig(*pEngineFilePath, onError)

	// ========================================
	// 思考エンジンの準備　＞　ログ・ファイル
	// ========================================

	// ログファイル
	var plainTextLogFile, _ = os.OpenFile(engineConfig.GetPlainTextLog(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer plainTextLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// ログファイル
	var jsonLogFile, _ = os.OpenFile(engineConfig.GetJsonLog(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer jsonLogFile.Close() // ログファイル使用済み時にファイルを閉じる
	// カスタマイズしたロガーを使うなら
	var log1 = logger.NewSugaredLoggerForGame(plainTextLogFile, jsonLogFile)

	// [O11o__11o6o0] デバッグ用
	if *pIsDebug {
		virtualIo.ReplaceInputToFileLines("./debug.input.txt")
	}

	// ========================================
	// 思考エンジンの準備　＞　テキストＩＯ
	// ========================================

	var text_io1 i_text_io.ITextIO = text_io.NewTextIO(log1)

	// ========================================
	// コマンドラインの第一引数で処理を振り分ける
	// ========================================

	switch arg1 {
	// スモークテスト
	// コマンドライン例： `go run . hello`
	case "hello":
		fmt.Println("hello, world")

	// ロガーのテスト
	// コマンドライン例： `go run . welcome`
	case "welcome":
		text_io1.GoCommand(fmt.Sprintf("Welcome! name:'%s' weight:%.1f x:%d", "nihon taro", 92.6, 3))
		log1.J.Infow("Welcome!",
			"name", "nihon taro", "weight", 92.6, "x", 3)

		// この上に分岐を挟んでいく
		// ---------------------

	default:
		// fmt.Println("go run . {programName}")

		LoopGTP(text_io1, log1, engineConfig)
	}
}

// EOF [O9o0]
