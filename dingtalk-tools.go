package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/icepy/go-dingtalk/src"
)

func main() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()
	if *configPath != "" {
		if _, err := os.Stat(*configPath); err != nil {
			fmt.Printf("Can't find config file `%s`\n", *configPath)
			os.Exit(1)
		} else {
			os.Setenv("RUNNER_CONFIG_PATH", *configPath)
		}
	}

	fmt.Printf("Current SDK VERSION=%s\n", dingtalk.VERSION)
	fmt.Printf("Current SDK OAPIURL=%s\n", dingtalk.OAPIURL)
	fmt.Printf("Current SDK TOPAPIURL=%s\n", dingtalk.TOPAPIURL)

}
