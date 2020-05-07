package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

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
		fmt.Println(item)
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
