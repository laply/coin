package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/laply/coin/explorer"
	"github.com/laply/coin/rest"
)

func usage () {
	fmt.Printf("Welcome to coin\n\n")
	fmt.Printf("Can use the following flags:\n\n")
	fmt.Printf("-port : 	Set the PORT of the server\n")
	fmt.Printf("-mode : 	Choose between 'html' and 'rest'  \n\n")
	runtime.Goexit()
	// os.Exit(0)
}

func Start() {
	if len(os.Args) == 1{
		usage()
	}


	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode{
		case "rest":
			rest.Start(*port)
		case "html":
			explorer.Start(*port)
		default:
			usage()

	}

}
