package main

import (
	// "bytes"
	"io"
	// "bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"text/tabwriter"

	args "github.com/akamensky/argparse"
	c "github.com/logrusorgru/aurora"
	"gopkg.in/yaml.v2"
)

// Config for YAML-file
type Config struct {
	Endpoints map[string]string `yaml:"endpoints,omitempty"`
}

var logger = log.New(os.Stdout, "healy - ", 0)

func main() {

	parser := args.NewParser("healy", "Healy is an easy-to-use and fast health check programm")

	// var myLogFile *os.File = parser.File("l", "log-file", os.O_RDWR, 0600, ...)
	var configFile = parser.File("c", "config", os.O_RDONLY, 0600, &args.Options{
		Required: false,
		Help:     "host to connect with (server mode)",
		Default:  "endpoints.yml",
	})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	conf := parseConfig(configFile)
	testEndpoints(conf)
}

func parseConfig(file *os.File) Config {
	fileInfo, _ := file.Stat()
	bytes := make([]byte, fileInfo.Size())

	_, err := io.ReadFull(file, bytes)
	file.Close()
	if err != nil {
		logger.Fatalf("Error parsing %s: %s", fileInfo.Name(), err)
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		logger.Fatalf("Error parsing %s: %s", fileInfo.Name(), err)
	}

	return config
}

func testEndpoints(conf Config) {

	if len(conf.Endpoints) < 1 {
		logger.Printf("no endpoints found, nothing to do...")
	}

	var wg sync.WaitGroup
	w := tabwriter.NewWriter(os.Stdout, 0, 100, 5, ' ', 0)
	defer w.Flush()

	var resultLog = log.New(w, "healy - ", 0)

	for name, url := range conf.Endpoints {

		wg.Add(1)

		go func(name, url string) {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				resultLog.Printf("%s\t%s\t%s", name, url, c.Bold(c.Red("Failed")))
			} else {
				resultLog.Printf("%s\t%s\t%s", name, url, c.Bold(c.Green("Ok")))
			}
		}(name, url)

	}

	wg.Wait()
}
