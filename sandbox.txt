package

import (
	"bytes"
	"fmt"
	"github.com/drand/tlock"
	tlockHttp "github.com/drand/tlock/networks/http"
	"log"
	"os"
	"time"
)

type CryptTask struct {
	FileName      string
	EncryptedName string
	Network       tlock.Network
}

func encrypt(t CryptTask) {
	//暗号化フェーズ
	fmt.Println("start encrypt")
	in, _ := os.Open(t.FileName)
	defer in.Close()

	duration := 10 * time.Second
	roundNumber := t.Network.Current(time.Now().Add(duration))

	var cipherData bytes.Buffer
	if err := tlock.New(t.Network).Encrypt(&cipherData, in, roundNumber); err != nil {
		log.Fatalf("ohh, lol: %v", err)
	}
	//書き出し
	os.WriteFile(t.EncryptedName, cipherData.Bytes(), 0644)
	fmt.Printf("Encryption complete, file export: %s (Round: %d)\n", t.EncryptedName, roundNumber)

}
func decrypt(t CryptTask) {
	//復号フェーズ
	encIn, _ := os.Open(t.EncryptedName)
	defer encIn.Close()

	var plainData bytes.Buffer
	if err := tlock.New(t.Network).Decrypt(&plainData, encIn); err != nil {
		fmt.Printf("Decryption failed... as expected: %v\n", err)
	} else {
		fmt.Printf("SUCCESSFUL \n source:%s\n", plainData.String())
	}
}

// ネットワークの初期化
func main() {

	// 今回は定数にしちゃってるけど本来はViperだからね！
	host := "https://api.drand.sh/"
	chainHashStr := "52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971"

	network, err := tlockHttp.NewNetwork(host, chainHashStr)
	if err != nil {
		log.Fatalf("ネットワーク作成失敗: %v", err)
	}

	task := CryptTask{
		FileName:      "test.txt",
		EncryptedName: "test.tle",
		Network:       network,
	}

	for {
		fmt.Println("\n--- MyTock Sandbox Menu ---")
		fmt.Println("1: ファイルを封印する(Encrypt)\n")
		fmt.Println("2: ファイルを解禁する(Decrypt)\n")
		fmt.Println("3: Exit\n")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			encrypt(task)
		case 2:
			decrypt(task)

		case 3:
			fmt.Println("bye")
			return
		default:
			fmt.Println("無効な選択")
		}

	}

}
