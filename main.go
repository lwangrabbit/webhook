package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const DefaultPort = 5000

func main() {
	// 定义命令行参数
	listenAddr := flag.String("addr", fmt.Sprintf(":%v", DefaultPort), "address to listen on")
	// 解析命令行参数
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/" {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				return
			}
			defer r.Body.Close()

			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02 15:04:05.000")

			fmt.Println(formattedTime, "Received message:", string(body))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Message received"))
		} else {
			http.NotFound(w, r)
		}
	})

	fmt.Println("Server started, listening on port ", strings.Split(*listenAddr, ":")[1])
	http.ListenAndServe(*listenAddr, nil)
}
