import main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"github.com/drand/tlock/http"
	"github.com/drand/tlock"
)

//今回は定数にしちゃってるけど本来はViperだからね！
const (
	host := "https://api.drand.sh/"
	chainHash := "52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971"
	fileName := "test.txt"
	encryptedName := "test.tle"
)

func encrypt(){
	
	os.WriteFile(fileName, []byte("1234"), 0644)

	//暗号化フェーズ
	fmt.Pritln("start encrypt")
	in, _ := os.Open(fileName)
	defer in.Close()

	network, _ := http.NewNetwork(host, chainHash)
	duration := 10 * time.Second
	roundNumber := network.RoundNumber(time.Now().Add(duration))

	var cipherData bytes.Buffer
	if err := tlock.New(network).Encrypt(&cipherData, in, roundNumber); err != mil{
		log.Fatalf("ohh, lol: %v", err)
	}
	//書き出し
	os.WriteFile(encryptedName, cipherData.Bytes(), 0644)
	fmt.Println("Encryption complete, file export: %s (Round: %d)\n", encryptedName, roundNumber)
}
func decrypt(){}

//ネットワークの初期化
func main(){

	for{
		fmt.Println("\n--- MyTock Sandbox Menu ---")
		fmt.Println("1: ファイルを封印する(Encrypt)")
		fmt.Println("2: ファイルを解禁する(Decrypt)")
		fmt.Println("3: Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
			case 1:
				encrypt()
			case 2:
				decrypt()
			case 3:
				fmt.Println("bye")
				return
			default:
				fmt.Println("無効な選択")
		}
		
	}
	

	//復号フェーズ
	fmt.Println("try decrypted")
	time.Sleep(11*time.Second)

	
	encIn, _ := os.Open(encryptedName)
	defer encIn.Close()

	var plainData bytes.Buffer
	if err := tlock.New(network).Decrypt(&plainData, encIn); err != nil{
		fmt.Printf("Decryption failed... as expected: %v\n", err)
	}else{
		fmt.Printf("SUCCESSFUL \n source:%s\n", plainData.String())
	}
}
