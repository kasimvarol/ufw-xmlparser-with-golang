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
	act      string   `xml:"act,attr"`
	ip       string   `xml:"ip,attr"`
	port     string   `xml:"port,attr"`
	protocol string   `xml:"protocol,attr"`
}

var ufwRules ufw

func pluginRun() {

	for _, item := range ufwRules.List {
		newrule := map[string]string{"act": "allow", "protocol": "any", "port": "any", "ip": "0.0.0.0/0"}
		for k := range newrule {
			if len(item.k) != 0 {

			}
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
