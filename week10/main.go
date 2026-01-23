package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
)

const saveFile = "public/memo.txt" // データファイルの保存先
const memoDir = "public/memos"

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)
	http.HandleFunc("/memo/save", memoSave)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func memo(w http.ResponseWriter, r *http.Request) {
	_ = os.MkdirAll(memoDir, 0755)

	message := ""
	if r.URL.Query().Get("saved") == "1" {
		message = "<p style='color:green;'>✔ 保存しました</p>"
	}

	// データファイルを開く
	text, err := os.ReadFile(saveFile)
	if err != nil {
		text = []byte("ここにメモを記入してください。")
	}

	files, _ := os.ReadDir(memoDir)
	list := "<ul>"
	for _, f := range files {
		name := html.EscapeString(f.Name())
		list += "<li>" + name + "</li>"
	}
	list += "</ul>"

	// HTMLのフォームを返す
	s := "<html>" +
		"<h1>メモ帳アプリ</h1>" +
		message +

		"<style>textarea { width:99%; height:200px; }</style>" +
		"<form method='get' action='/mwrite'>" +
		"<textarea name='text'>" + html.EscapeString(string(text)) + "</textarea><br>" +
		"<input type='submit' value='保存'>" +
		"</form>" +

		"<hr>" +

		"<h2>複数メモ一覧</h2>" +
		list +

		"<h3>新規メモ作成</h3>" +
		"<form method='post' action='/memo/save'>" +
		"<input type='text' name='name' placeholder='memo1.txt'><br>" +
		"<textarea name='text'></textarea><br>" +
		"<input type='submit' value='保存'>" +
		"</form>" +

		"</html>"

	_, err = w.Write([]byte(s))
	_ = err
}

func mwrite(w http.ResponseWriter, r *http.Request) {
	// 投稿されたフォームを解析
	_ = r.ParseForm()

	if len(r.Form["text"]) == 0 {
		_, _ = w.Write([]byte("フォームから投稿してください。"))
		return
	}

	text := r.Form["text"][0]
	// データファイルへ書き込む
	_ = os.WriteFile(saveFile, []byte(text), 0644)
	// ルートページへリダイレクトして戻る --- (*4)
	http.Redirect(w, r, "/memo?saved=1", 302)
}

func memoSave(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	name := r.FormValue("name")
	text := r.FormValue("text")

	if name == "" {
		_, _ = w.Write([]byte("メモ名を入力してください"))
		return
	}

	path := memoDir + "/" + name
	_ = os.WriteFile(path, []byte(text), 0644)

	http.Redirect(w, r, "/memo?saved=1", 302)
}
