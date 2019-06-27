package main

import (
	"encoding/json"
	"fmt"
	"github.com/just1689/path-proxy/pprox"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func main() {
	if generate() {
		return
	}

	logrus.Println("Starting up...")
	b := getConfigByEnv()
	c := readConfigOrPanic(b)
	pprox.Proxy(c)
}

func generate() bool {
	if len(os.Args) == 2 && (os.Args[1] == "generate" || os.Args[1] == "g") {
		c := pprox.Config{
			Listen: ":8080",
			Routes: make([]pprox.Route, 1),
		}
		c.Routes[0] = pprox.Route{
			Target: "bbc.com",
			Prefix: "/bbc",
		}
		b, err := json.Marshal(c)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		return true
	}
	return false
}

func readConfigOrPanic(b []byte) *pprox.Config {
	logrus.Println("Unmarshal config")
	c := pprox.Config{}
	err := json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	return &c

}

func getConfigByEnv() []byte {
	f := os.Getenv("config")
	if f == "" {
		f = "config.json"
		logrus.Println("Defaulting to ", f)
	}
	logrus.Println("Loading from ", f)
	dat, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return dat

}
