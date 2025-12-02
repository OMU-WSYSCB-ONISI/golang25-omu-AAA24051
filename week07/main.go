package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func main() {
	// Week 07: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	fmt.Println("Week 07 課題")

	http.HandleFunc("/cal02", cal02Handler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
	// 以下に実装してください
}

func cal02Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))

	switch r.FormValue("cal02") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, x*y)
	case "/":
		if y == 0 {
			fmt.Fprintln(w, "0 では割れません")
			return
		}
		fmt.Fprintf(w, "%.2f", float64(x)/float64(y))
	}
}
