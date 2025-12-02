package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Week 08: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	fmt.Println("Week 08 課題")

	http.HandleFunc("/ave", aveHandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
	// 以下に実装してください
}

func aveHandler(w http.ResponseWriter, r *http.Request) {
	tokens := strings.Split(r.FormValue("scores"), ",")
	sum := 0
	dist := make([]int, 11)
	for _, t := range tokens {
		t = strings.TrimSpace(t)
		if t == "" {
			continue
		}
		score, _ := strconv.Atoi(t)
		sum += score
		idx := score / 10
		if idx > 10 {
			idx = 10
		}
		dist[idx]++
	}
	avg := float64(sum) / float64(len(tokens))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "平均点: %.2f<br>", avg)
	fmt.Fprint(w, "10点ごとの分布<br>")
	for i, v := range dist {
		label := fmt.Sprintf("%d-%d", i*10, i*10+9)
		if i == 10 {
			label = "100"
		}
		fmt.Fprintf(w, "%s: %d<br>", label, v)
	}
}
