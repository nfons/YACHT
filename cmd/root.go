package cmd

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"html/template"
	"log"
	"os"
	"os/exec"
)

type AppConfigProperties map[string]string

var env string
var kfile string
var config AppConfigProperties
var rootCmd = &cobra.Command{
	Use:   "yacht",
	Short: "yacht is a simple template based parser for k8s",
	Long:  `yacht is a simple template based parser for k8s Check out http://github.com/nfons/yacht`,
}

func SingleFile(subcmd string) {
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
	cmd := exec.Command("kubectl", subcmd, "-f", "temp.yaml")
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

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Use Kubectl Apply instead",
	Run: func(cmd *cobra.Command, args []string) {
		SingleFile("apply")
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Use Kubectl Create with Templating",
	Run: func(cmd *cobra.Command, args []string) {
		SingleFile("create")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&env, "env", "e", "dev", "Env file you wish to apply, it wills search for a env file from where you call")
	rootCmd.PersistentFlags().StringVarP(&kfile, "file", "f", "", "Kubernetes yaml file you want to use")
	rootCmd.AddCommand(applyCmd)
	rootCmd.AddCommand(createCmd)
}

func initConfig() {
	config = AppConfigProperties{}
	// Don't forget to read config either from cfgFile or from home directory!
	if env != "" {
		// Use config file from the flag.
		viper.SetConfigFile(env)
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Error("Can't read config:", err)
		os.Exit(1)
	} else {
		strConf := viper.AllSettings()
		for key, val := range strConf {
			strKey := fmt.Sprintf("%v", key)
			strVal := fmt.Sprintf("%v", val)
			config[strKey] = strVal
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
