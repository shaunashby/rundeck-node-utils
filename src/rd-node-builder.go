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

// ReadHostFile reads the JSON input stream and returns
// a ForemanHost instance containing the unmarshalled
// JSON data:
func ReadHostFile(filename string) *foreman.ForemanHost {

	// Open the host file and read in the JSON stream:
	hfileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading from file %s: %s\n", filename, err))
	}

	fhost := foreman.ForemanHost{}

	if err := json.Unmarshal(hfileBytes, &fhost); err != nil {
		panic(fmt.Sprintf("Error unmarshalling from file: %s\n", err))
	}

	return &fhost
}

func main() {
	var hostDataFilename = fmt.Sprintf("%s/host.json", os.Getenv("HOME"))
	fhost := ReadHostFile(hostDataFilename)
	// Print the host data:
	fmt.Printf("%v\n", fhost)
}
