package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
)

var env string
var kfile string

func main() {
	//read config
	flag.StringVar(&env, "env", "dev", "Environment to use, must have a corresponding [env].conf file")
	flag.StringVar(&env, "e", "dev", "Environment to use, must have a corresponding [env].conf file")
	flag.StringVar(&kfile, "file", "", "file to use")
	flag.StringVar(&kfile, "f", "", "file to use")
	flag.Parse()

	filename := fmt.Sprintf("%s.conf", env)
	config, err := ReadPropertiesFile(filename)
	if err != nil {
		log.Fatal("Error reading conf file: ", err)
		os.Exit(1)
	}

	t, err := template.ParseFiles(kfile)
	if err != nil {
		log.Fatal("Cannot find File: ", err)
		os.Exit(1)
	}
	f, err := os.Create("temp.yaml")
	if err != nil {
		log.Fatal("Error creating template file: ", err)
	}
	//I'm pretty sure we can make this a bit better by just bypassing file save/delete logic, and directly injecting to
	//kubectl
	t.Execute(f, config)
	f.Close()

	//exec kubectl command
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

	//remove temp file
	os.Remove("temp.yaml")
}
