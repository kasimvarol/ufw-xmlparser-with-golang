package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//XML Parse Structure: <ufw> <rule ../> <rule ../> ... </ufw>
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

// Constants and global variables
const POLICY_FILE = "ufw.xml"
const USER_FILE = "deneme"

var ufwRules ufw
var rules []string

func pluginRun() {

	//Parsing XML file into global ufw variable
	xmlFile, err := os.Open(POLICY_FILE)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &ufwRules)

	//Creating rules as string and append to string slice
	for _, item := range ufwRules.List {
		var newrule string

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
			newrule = "### tuple ### " + " " + item.Act + " " + item.Protocol + " " + item.Port + " 0.0.0.0/0 any " + item.IP + " in"
			rules = append(rules, newrule)
		}

	}

	// For each rule read USER_FILE and write according to hierarchy.
	for _, rule := range rules {

		fi, err := os.Open(USER_FILE)
		if err != nil {
			log.Fatal(err)
		}
		fo, err := os.OpenFile(USER_FILE, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer fi.Close()
		defer fo.Close()

		scanner := bufio.NewScanner(fi)
		writer := bufio.NewWriter(fo)

		// Reading USER_FILE to decide which line to append/insert
		for scanner.Scan() {
			line := scanner.Text()
			if line == "### RULES ###" {
				line = line + "\n" + rule
			}
			writer.WriteString(line + "\n")
		}
		writer.Flush()

	}

}

func main() {
	pluginRun()
}
