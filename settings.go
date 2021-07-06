package awsummary

import (
	"fmt"
	logs "github.com/sirupsen/logrus"
	"log"
	"os"
	"os/user"
	"path"
)

//p2s return string value from pointer of string.
//if the pointer is nil, p2s returns empty string.
func p2s(ps *string) string {
	if ps == nil {
		return ""
	}
	return *ps
}

func Msg(service string) string {
	return fmt.Sprintf("Region specific %s data", service)
}

// Make sure the credentials exists
func CheckConfig() {
	config := path.Join(userHome(), ".aws/credentials")
	if _, err := os.Stat(config); os.IsNotExist(err) {
		log.Printf("No config file found at: %s", config)
		if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
			log.Println("No environment variables found: AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY")
			os.Exit(1)
		}
	}
}

// Make sure we can create log file
func CheckLogFile(logfile string) {
	if logfile != "" {
		if _, err := os.Stat(logfile); err != nil {
			c, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE, 0755)
			if err != nil {
				log.Fatal("Cannot create log file: ", logfile)
			}
			defer c.Close()
		}
		f, err := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			log.Fatal(err)
		}
		logs.SetOutput(f)
	}
}

// Get userhome on different OS
func userHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
