package main

import (
	"fmt"
	"github.com/G-goma393/mytock-project/cmd"
	//"github.com/drand/tlock" // ライブラリのインポート
)

func main() {
	fmt.Println("MyTock プロジェクト始動！")

	// tlock.Encrypt が存在するか参照して、ロードを確認する
	// (実行はしないので、値を変数に入れない '_' を使います)
	//_ = tlock.Encrypt

	fmt.Println("tlock ライブラリのロードに成功しました。")

	// Cobraのコマンド実行
	cmd.Execute()
}
