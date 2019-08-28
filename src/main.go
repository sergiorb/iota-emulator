package main

import "fmt"
import "os"
import "os/signal"
import "syscall"
import "github.com/tkanos/gonfig"
import "github.com/sergiorb/iota-emulator/src/config"
import "github.com/sergiorb/iota-emulator/src/scheduler"
import "github.com/sergiorb/iota-emulator/src/jobs"
import "github.com/sergiorb/iota-emulator/src/generator"
import JsonAgent "github.com/sergiorb/iota-emulator/src/agents/json"

func main() {

	configuration := config.Config{}

	configPath := os.Getenv("IOTA_EMULATOR_CONFIG")

	if configPath == "" {
		configPath = "./config.json"
	}
	
	err := gonfig.GetConf(configPath, &configuration)
	
	if err != nil {
		panic(err)
	}

	var jsonAgent = &JsonAgent.JsonAgent{
		Config: configuration.Agents.Json,
	}

	var store = scheduler.BuildStore();

	for _, dev := range configuration.Devices {
		for _, attr := range dev.Attributes {

			store.AddJob(attr.Cron, &jobs.AttributeJob{
				DeviceId: dev.Id,
				Key: attr.Key,
				Cron: attr.Cron,
				Generator: generator.BuildGenerator(attr.ValueGenerator),
				Agent: jsonAgent,
			});
		}
	}

	store.Start();

	fmt.Println("IoTA-emulator started!");

	// Thanks to: https://www.reddit.com/r/golang/comments/4hktbe/read_user_input_until_he_press_ctrlc/
	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	msg := make(chan string, 1)

	go func() {
		// Receive input in a loop
		for {
			var s string
			fmt.Scan(&s)
			// Send what we read over the channel
			msg <- s
		}
	}()

	loop:
		for {
			select {
			case <-sigs:
				store.Stop();
				fmt.Println("Got shutdown, exiting")
				// Break out of the outer for statement and end the program
				break loop
			case s := <-msg:
				fmt.Println("Echoing: ", s)
			}
		}
	
}
