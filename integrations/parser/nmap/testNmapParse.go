package nmap

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	nmap, err := ParseFile("sample.xml")
	if err != nil {
		t.Errorf("Error parsing file: %v", err)

	}

	if nmap.Scanner != "nmap" {
		t.Errorf("Expected scanner to be 'nmap', got %s", nmap.Scanner)

	}

	if len(nmap.Hosts) != 1 {
		t.Errorf("Expected 1 host, got %d", len(nmap.Hosts))

	}

	host := nmap.Hosts[0]
	if host.Status.State != "up" {
		t.Errorf("Expected host to be up, got %s", host.Status.State)

	}

	if len(host.Ports) != 22 {
		t.Errorf("Expected 22 ports, got %d", len(host.Ports))

	}

	port := host.Ports[0]
	if port.Protocol != "tcp" {
		t.Errorf("Expected protocol to be 'tcp', got %s", port.Protocol)

	}

	if port.PortId != "22" {
		t.Errorf("Expected port ID to be 22, got %s", port.PortId)

	}

}
