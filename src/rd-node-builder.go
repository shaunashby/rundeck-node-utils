//____________________________________________________________________
// File: rd-node-builder.go
//____________________________________________________________________
//
// Author: Shaun Ashby <shaun@ashby.ch>
// Created: 2018-08-13 14:58:18+0200
// Revision: $Id$
// Description: Read JSON file and create a node list compatible with Rundeck
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
import "foreman"
import "rundeck/formatters"

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
	var hostDataFilename = fmt.Sprintf("%s/hosts.json", os.Getenv("HOME"))
	fresponse := ReadHostFile(hostDataFilename)
	// Choose a formatter:
	out := &formatters.NodeXmlFormatter{
		HostData: fresponse.GetHosts(),
	}
	fmt.Printf("%v\n",out)
	// Print the host data:
	fmt.Printf("%v\n", fhost)
}
