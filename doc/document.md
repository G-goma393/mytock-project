# TL;DR
drandネットワークを利用したタイムロック暗号化ツールとその周辺機能をまとめたTUI/CLIツール。Go製 Linux/
drandネットワークを利用したタイムロック暗号化ツールとその周辺機能をまとめたTUI/CLIツール。Rust製も制作予定 windows/Linux


### encrypt
./mytock-project encrypt <file名> --<duration>

<file名>の暗号化、復号可能時間<duration>
drandネットワークの設定はviper

./mytock-project encrypt
これ単体で実行されると
`panic: runtime error: index out of range [0] with length 0`
っていうふうにパニックになっちゃう
ちゃんと返さなきゃ使い方を返してあげよう

それだけじゃなくて
パニックにならないための処理が必要不可欠だね

### ls
./mytock-project ls

mytockで作成・追加した全ての暗号化・復号可能ファイルの一覧表示

`-e`|--encrypted|暗号化したファイルの一覧表示
`-r`|--ready    |復号可能ファイルの一覧表示
`-d`|--decrypted|復号済みファイルを一覧表示
`-a`|--all      |復号済みファイルを含む一覧表示

### add
./mytock-project add <file名>

drandで暗号化されたファイルをmytockに追加する

### detach
./mytock-project detach <file名>

追跡対象から外す



drandネットワークの設定はviperが担当
手動で変更したけばviperを叩け

durationは毎回聞く形に
durationの設定は危険ゾーン
注意が必要

それから本命の追跡機能
どこにファイルがあろうが記憶する
→読み込み権限の問題が出てくる
実行するユーザによって....いや？パスの記憶なら問題はないか

動作としてはencryptする度にパスを保持する
復号化に際してファイル名だけで復号できるように

lsコマンドなるものを実装したい
TUIの管理画面のCLI版といったところ
どのハッシュで...か

info
fileName tag暗号化したファイル - 復号化できる大まかな日付時間
fileName tag復号化したファイル
fileName tag復号化できるファイル
[name│tag│time]

このlsコマンドでファイルの場所を見に行く
それで、あれば良いんだけど
無い場合や他の端末、もしくは別の手段で暗号化してきたか
それらを追加するコマンドが別途必要


chainHashとかの値って人それぞれじゃない？だとしたら別途初期設定をするための何かを入れたほうが良いね

encryptコマンドの実装
decryptコマンドの実装
lsコマンドの実装
addコマンドの実装
detachコマンドの実装
SQLite




