package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Week 03: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	http.HandleFunc("/webfortune", webfortune)

	fmt.Println("Week 03 課題")
	http.ListenAndServe(":8080", nil)
	// 以下に実装してください
}

func webfortune(w http.ResponseWriter, r *http.Request) {
	fortunes := []string{"大吉", "中吉", "吉", "凶"}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := fortunes[rnd.Intn(len(fortunes))]

	fmt.Fprintln(w, "今の運勢は「",result,"」です")
}
