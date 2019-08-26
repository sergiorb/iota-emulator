package main

import "fmt"
import "github.com/tkanos/gonfig"
import "github.com/sergiorb/iota-emulator/src/config"

func main() {

	configuration := config.Config{}
	
	err := gonfig.GetConf("./src/config.json", &configuration)
	
	if err != nil {
		panic(err)
	}
	
	fmt.Println("IoTA-emulator started!");
	fmt.Println(configuration);
}
