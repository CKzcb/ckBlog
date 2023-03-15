/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2023/3/15 07:46
 */

package main

import (
	"flag"
	"fmt"
	"github.com/CKzcb/ckBlog/internal/routers"
	"net/http"
	"time"
)

var host string
var port int

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "please input listener host")
	flag.IntVar(&port, "port", 8080, "please input listener port")
	flag.Parse()
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", host, port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
