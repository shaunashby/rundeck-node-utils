//____________________________________________________________________
// File: ansible-node-builder.go
//____________________________________________________________________
//
// Author: Shaun Ashby <shaun@ashby.ch>
// Created: 2018-08-20 10:58:18+0200
// Revision: $Id$
// Description: Read JSON file and create a node list compatible with Ansible
//
// Copyright (C) 2018 Shaun Ashby
//
//
//--------------------------------------------------------------------
package main

import "fmt"
import "encoding/json"
import "os"
import "io/ioutil"
import "github.com/shaunashby/rundeck-node-utils/foreman"
import "github.com/shaunashby/rundeck-node-utils/ansible/formatters"

// ReadHostFile reads the JSON input stream and returns
// a ForemanHost instance containing the unmarshalled
// JSON data:
func ReadHostFile(filename string) *foreman.ForemanApiResponse {

	// Open the host file and read in the JSON stream:
	hfileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading from file %s: %s\n", filename, err))
	}

	fresp := foreman.ForemanApiResponse{}

	if err := json.Unmarshal(hfileBytes, &fresp); err != nil {
		panic(fmt.Sprintf("Error unmarshalling from file: %s\n", err))
	}

	return &fresp
}

func main() {
	// Set the zone (office/infra/clients):
	puppetZone := "infra"
	var hostDataFilename = fmt.Sprintf("%s/%s.json", os.Getenv("HOME"), puppetZone)

	fresponse := ReadHostFile(hostDataFilename)

	// Choose a formatter:
	out := &formatters.NodeIniFormatter{
		HostData: fresponse.GetHosts(),
		PuppetZone: puppetZone,
	}
	// Print the formatted node data for Rundeck:
	fmt.Printf("%v\n",out)
}
