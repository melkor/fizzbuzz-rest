package main

import (
	"fmt"
	"os"

	"github.com/melkor/fizzbuzz-rest/app"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	address       = pflag.String("address", ":8000", "The listening address of the web server")
	cacheAddress  = pflag.String("cache-address", "localhost:6379", "Address of the redis server used to store hits")
	cachePassword = pflag.String("cache-password", "", "Password of the redis server")
	debug         = pflag.BoolP("debug", "d", false, "Activate debug mode")
	help          = pflag.BoolP("help", "h", false, "Display usage message")
	verbose       = pflag.BoolP("verbose", "v", false, "Activate verbose mode")
)

// display help message and exit
func displayHelp() {
	fmt.Println(os.Args[0])
	fmt.Println("Web server that will expose a REST API to execute fizz-buzz.")
	pflag.PrintDefaults()
}

func main() {
	pflag.Parse()

	if *help {
		displayHelp()
		os.Exit(0)
	}

	if *debug {
		log.SetLevel(log.DebugLevel)
	} else if *verbose {
		log.SetLevel(log.InfoLevel)
	}

	app := app.New(*address, *cacheAddress, *cachePassword)
	app.Run()
}
