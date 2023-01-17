package nmap

import (
	"encoding/xml"
	"io/ioutil"
)

type NmapRun struct {
	XMLName xml.Name `xml:"nmaprun"`
	Scanner string   `xml:"scanner,attr"`
	Args    string   `xml:"args,attr"`
	Hosts   []Host   `xml:"host"`
}

type Host struct {
	XMLName   xml.Name   `xml:"host"`
	Addresses []Address  `xml:"address"`
	Ports     []Port     `xml:"ports>port"`
	Hostnames []Hostname `xml:"hostnames>hostname"`
	Status    Status     `xml:"status"`
}

type Address struct {
	XMLName  xml.Name `xml:"address"`
	Addr     string   `xml:"addr,attr"`
	AddrType string   `xml:"addrtype,attr"`
}

type Port struct {
	XMLName  xml.Name `xml:"port"`
	Protocol string   `xml:"protocol,attr"`
	PortId   string   `xml:"portid,attr"`
	State    State    `xml:"state"`
	Service  Service  `xml:"service"`
}

type State struct {
	XMLName   xml.Name `xml:"state"`
	State     string   `xml:"state,attr"`
	Reason    string   `xml:"reason,attr"`
	ReasonTtl string   `xml:"reason_ttl,attr"`
}

type Service struct {
	XMLName xml.Name `xml:"service"`
	Name    string   `xml:"name,attr"`
	Product string   `xml:"product,attr"`
	Version string   `xml:"version,attr"`
}

type Hostname struct {
	XMLName xml.Name `xml:"hostname"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
}

type Status struct {
	XMLName xml.Name `xml:"status"`
	State   string   `xml:"state,attr"`
}

func Parse(data []byte) (NmapRun, error) {
	var n NmapRun
	err := xml.Unmarshal(data, &n)
	if err != nil {
		return NmapRun{}, err

	}
	return n, nil

}

func ParseFile(file string) (NmapRun, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return NmapRun{}, err

	}
	return Parse(data)

}
