package main

import (
	"io"
	"log"
	"os"

	"github.com/alecthomas/template"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {

	// Set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Set paths to log files
	slogpath := os.Getenv("GOLOG")
	elogpath := os.Getenv("GOERR")
	ilogpath := os.Getenv("GOILOG")
	if slogpath == "" {
		slogpath = "/var/log/golang/golang-server-connection.log"
	}
	if elogpath == "" {
		elogpath = "/var/log/golang/golang-error.log"
	}
	if ilogpath == "" {
		ilogpath = "/var/log/golang/golang-server-internal.log"
	}

	// Initialize log files
	slog, _ := os.Create(slogpath)
	elog, _ := os.Create(elogpath)
	ilog, _ := os.Create(ilogpath)
	defer slog.Close()
	defer elog.Close()
	defer ilog.Close()

	// Set log outputs to respective log files
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(slog, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(elog, os.Stderr)
	log.SetOutput(io.MultiWriter(ilog, os.Stdout))

	// Initialize engine and files
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.BestCompression))
	html, err := template.New("").Delims("[[", "]]").ParseFiles("")
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(html)

	// Start server
	router.Run()
}
