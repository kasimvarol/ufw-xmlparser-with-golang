package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const POLICY_FILE = "ufw.xml"
const USER_FILE = "deneme"

type ufw struct {
	XMLName xml.Name `xml:"ufw"`
	List    []rule   `xml:"rule"`
}

type rule struct {
	XMLName  xml.Name `xml:"rule"`
	Act      string   `xml:"act,attr"`
	IP       string   `xml:"ip,attr"`
	Port     string   `xml:"port,attr"`
	Protocol string   `xml:"protocol,attr"`
}

var ufwRules ufw

func pluginRun() {
	for _, item := range ufwRules.List {
		// Checking standards
		if item.Act == "" {
			fmt.Println("ERROR // One of rules action field is missing!")
		} else if item.IP == "" && item.Port == "" {
			fmt.Println("ERROR // Either Port or IP should be specified!")
		} else if item.Port != "" && strings.Contains(item.Port, ":") && (item.Protocol == "any" || item.Protocol == "") {
			fmt.Println("ERROR // Multiports require specific protocol!")
		} else {
			//After standard check, filling empty variables -if any-  by default.
			if item.IP == "" {
				item.IP = "0.0.0.0/0"
			}
			if item.Port == "" {
				item.Port = "any"
			}
			if item.Protocol == "" {
				item.Protocol = "any"
			}
			newrule := "### tuple ### " + " " + item.Act + " " + item.Protocol + " " + item.Port + " 0.0.0.0/0 any " + item.IP + " in"
			fmt.Println(newrule)
		}

	}

}

func main() {

	xmlFile, err := os.Open("/home/mkv/code/src/github.com/kasimvarol/xmlparse/ufw.xml")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &ufwRules)

	pluginRun()
}
