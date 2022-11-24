# fibonacci Api
## 概要
指定したn番目のフィボナッチ数を返すAPI

クエリパラメーターでnは設定する

## 仕様

### エンドポイント
https://dry-depths-57184.herokuapp.com/fib

### メソッド
Get

### クエリパラメーター
n 正の数(0 < n)　1で先頭のフィボナッチ数を指定する

### リクエスト例
https://dry-depths-57184.herokuapp.com/fib?n=99

### レスポンス
200　成功した時

```json
{"result":218922995834555169026}
```

404　指定したパスがない時
```json
{"message":"no route","status":400}
```
404　クエリパラメーターが間違っているとき（マイナスの値、数値ではない）
```json
{"message":"please set positive number","status":400}
```

## ソースコードの構成
├── bin 

├── controller.go　//コントローラー

├── go.mod

├── go.sum

├── main.go　　　　//main

├── main_test.go　//テスト

└── router.go　   //ルーティング


