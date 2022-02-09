package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/laply/coin/rest"
)

func usage () {
	fmt.Printf("Welcome to coin\n\n")
	fmt.Printf("Please use the following command:\n\n")
	fmt.Printf("explorer: start the HTML Explorer \n")
	fmt.Printf("rest	: start the REST API (recommended) \n\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	restFlags := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := restFlags.Int("port", 4210, "Set the port of the server")

	switch os.Args[1]{
		case "explorer", "Explorer":
			fmt.Println("Start Explorer")
		case "rest", "REST":
			restFlags.Parse(os.Args[2:])
		default:
			usage()			
	}

	if restFlags.Parsed() {
		rest.Start(*portFlag)
		// fmt.Println("Start REST API")
	}

}


