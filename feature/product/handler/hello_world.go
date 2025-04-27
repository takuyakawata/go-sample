package handler

import (
	"net/http"
)

// helloHandler は "helloworld" という文字列を返すシンプルなHTTPハンドラーです。
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// ステータスコード 200 OK を設定します。
	w.WriteHeader(http.StatusOK)
	// レスポンスボディに "helloworld" を書き込みます。
	_, err := w.Write([]byte("helloworld"))
	if err != nil {
		// 書き込み中にエラーが発生した場合の処理 (ログ出力など)
		// ここでは簡単のため省略しますが、実際にはエラーハンドリングを推奨します。
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
