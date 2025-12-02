package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func main() {
	// Week 06: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))

	http.HandleFunc("/bmi", bmiHandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
	// 以下に実装してください
}

func bmiHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")

    weightStr := r.URL.Query().Get("weight")
    heightStr := r.URL.Query().Get("height")

    weight, err1 := strconv.ParseFloat(weightStr, 64)
    height, err2 := strconv.ParseFloat(heightStr, 64)

    if err1 != nil || err2 != nil || height == 0 {
        fmt.Fprintln(w, "正しい整数値を入力してください")
        return
    }

    h := height / 100.0
    bmi := weight / (h * h)

    fmt.Fprintln(w, "BMI =", bmi)
}
