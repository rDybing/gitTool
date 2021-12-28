package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

type configT struct {
	Path    string
	Project []string
	Account string
	Token   string
}

const configFile = "./Config/config.json"

func main() {
	// load config
	var config configT
	if err := config.loadConfig(); err != nil {
		log.Panicln("Could not load config, exiting!")
	}
	// get push or pull
	if err := config.pull(); err != nil {
		log.Panicf("Could not pull project, exiting!\n%v\n", err)
	}
}

func (c configT) push() error {
	return nil
}

func (c configT) pull() error {
	for _, v := range c.Project {
		if _, err := exec.Command("cd", v).Output(); err != nil {
			return err
		}
		out, err := exec.Command("git", "pull").Output()
		go c.enterCreds()
		if err != nil {
			log.Print(err)
		}
		fmt.Println(string(out))
		if _, err := exec.Command("cd", "..").Output(); err != nil {
			return err
		}
	}
	return nil
}

func (c configT) enterCreds() {
	time.Sleep(1 * time.Second)
	fmt.Printf("%s\n", c.Account)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s\n", c.Token)
}

func (c *configT) loadConfig() error {
	f, err := ioutil.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("could not load config file: %s\n%v", configFile, err)
	}
	if err := json.Unmarshal(f, &c); err != nil {
		return fmt.Errorf("could not unmarshal config.json: %v", err)
	}
	return nil
}
