package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sorajima/workhour/monthly"
)

func main() {
	rootURI := "/sorajima/api"
	http.HandleFunc(rootURI+"/v1/workhour/monthly", monthly.HTTPMethodHandler)
	log.Fatal(http.ListenAndServe(":8000", accesslogHandler(http.DefaultServeMux)))
}

func accesslogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// ファイルパスからファイル情報取得
		currentDirectory, _ := os.Getwd()
		filename := filepath.Join(currentDirectory, "access.log")
		fileinfo, staterr := os.Stat(filename)
		// ファイル ローテート
		if staterr == nil {
			// ファイルサイズを表示
			fmt.Println(fileinfo.Size())

			if fileinfo.Size() > 1048576 {
				os.Rename(filename, filename+"."+time.Now().Format("20060102150405"))
			}
		}

		logfile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("cannnot open access.log:" + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer logfile.Close()

		// io.MultiWriteで、標準出力とファイルの両方を束ねてlogの出力先に設定する
		log.SetOutput(io.MultiWriter(logfile, os.Stdout))

		bufbody := new(bytes.Buffer)
		bufbody.ReadFrom(req.Body)
		body := bufbody.String()
		log.Println("Remote Address : ", req.RemoteAddr)
		log.Println("Content-Type   : ", req.Header.Get("Content-Type"))
		log.Println("Http Method    : ", req.Method)
		log.Println("Request URL    : ", req.URL)
		log.Println("Query string   : ", req.URL.RawQuery)
		log.Println("Body           : ", body)

		handler.ServeHTTP(w, req)
	})
}
