package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"os"
	"os/exec"
)

var env string
var kfile string

func SingleFile(config AppConfigProperties) {
	t, err := template.ParseFiles(kfile)
	if err != nil {
		log.Fatal("Cannot find File: ", err)
		os.Exit(1)
	}
	f, err := os.Create("temp.yaml")
	if err != nil {
		log.Fatal("Error creating template file: ", err)
	}
	// I'm pretty sure we can make this a bit better by just bypassing file save/delete logic, and directly injecting to
	// kubectl
	t.Execute(f, config)
	f.Close()

	// exec kubectl command
	cmd := exec.Command("kubectl", "create", "-f", "temp.yaml")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	kubeErr := cmd.Run()
	if kubeErr != nil {
		os.Remove("temp.yaml")
		log.Println(fmt.Sprint(kubeErr) + " :- " + stderr.String())
		os.Exit(1)
	}

	log.Println(out.String())

	// remove temp file
	os.Remove("temp.yaml")
}

func dirStruct(config AppConfigProperties) {
	return
}

func main() {
	// read config
	flag.StringVar(&env, "env", "dev", "Environment to use, must have a corresponding [env].conf file")
	flag.StringVar(&env, "e", "dev", "Environment to use, must have a corresponding [env].conf file")
	flag.StringVar(&kfile, "file", "", "file to use")
	flag.StringVar(&kfile, "f", "", "file to use")
	flag.Parse()

	// filename := fmt.Sprintf("%s.conf", env)

	// viper logic
	viper.SetConfigName(env)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viperErr := viper.ReadInConfig()
	if viperErr != nil {
		log.Fatal(viperErr)
	}

	config := AppConfigProperties{}
	strConf := viper.AllSettings()
	for key, val := range strConf {
		strKey := fmt.Sprintf("%v", key)
		strVal := fmt.Sprintf("%v", val)
		config[strKey] = strVal
	}

	// check if file is a dir or not
	file, err := os.Stat(kfile)
	if err != nil {
		log.Fatal(err)
	}
	switch mode := file.Mode(); {
	case mode.IsDir():
		dirStruct(config)
	default:
		SingleFile(config)
	}
}
