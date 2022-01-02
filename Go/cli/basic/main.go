package main

import (
	"flag"
	"fmt"
	"os"

)

func main(){
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[1]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	switch len(os.Args){
	case 2:
		fmt.Println(os.Args[1:])
	case 3:
		flagParse(os.Args)
	default:
		fmt.Println("No Args")
	}

}
func flagParse(args []string ){
	pPtr := flag.String("p", "p", "port forwarding")
	flag.Parse()
	fmt.Println("Port PTR: ", *pPtr)
	args = flag.Args()
	fmt.Println("args: ", args)

}
