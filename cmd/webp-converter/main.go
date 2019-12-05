package main

import (
	"flag"
	"github.com/hzhyvinskyi/webp-converter/internal/app/http"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	listenAddr string
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	flag.StringVar(&listenAddr, "listen-addr", os.Getenv("LISTEN_ADDR"), "TCP address to listen on")
	flag.Parse()

	http.Serve(listenAddr)
}
