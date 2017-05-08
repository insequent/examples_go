package main

import (
	"flag"
	"fmt"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "", "Absolute path to the kubeconfig file (can be included in .conf)")
	flag.Parse()

	if len(*kubeconfig) == 0 {
		fmt.Println("Kubeconfig is empty")
	} else {
		fmt.Println("Kubeconfig is", *kubeconfig)
	}

}
