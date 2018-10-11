package main

import (
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	"strings"  // пакет для работы с  UTF-8 строками
)

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //анализ аргументов,
	fmt.Println(r.Form) // ввод информации о форме на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "  🍄  nothing more to do 💃💫   🚀") // отправляем данные на клиентскую сторону
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)  // установим роутер
	err := http.ListenAndServe(":2888", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
