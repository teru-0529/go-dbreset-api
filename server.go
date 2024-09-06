package main

import (
	_ "github.com/lib/pq"
	"github.com/teru-0529/go-dbreset-api/businesslogic"
	"github.com/teru-0529/go-dbreset-api/infra"
)

func main() {
	// PROCESS: envファイルのロード
	config := infra.LeadEnv()

	// PROCESS: データベース(Sqlboiler)の設定
	cleanUp := infra.InitDB(config)
	defer cleanUp()

	// PROCESS: Webサーバー(echo)の設定
	e := infra.InitServer(config.DebugMode)

	// PROCESS: routing
	e.POST("/schemas/:schema/action-truncate", businesslogic.Truncate)
	e.POST("/schemas/:schema/action-seq-reset", businesslogic.SequenceReset)

	// サーバースタート
	e.Logger.Fatal(e.Start(":8080"))
}
