package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Week 04: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	http.HandleFunc("/info", info)

	fmt.Println("Week 04 課題")
	http.ListenAndServe(":8080", nil)
	// 以下に実装してください
}

func info(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(jst).Format("2006年01月02日 15:04:05")
	a := r.Header.Get("User-Agent")
	fmt.Fprintln(w, "今の時刻は", now, "で，利用しているブラウザは「", a, "」です。")
}
