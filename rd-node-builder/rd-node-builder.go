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
import "github.com/shaunashby/rundeck-node-utils/foreman"
import "github.com/shaunashby/rundeck-node-utils/rundeck/formatters"

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

// MapPuppetZoneToMaster is a function to convert the Puppet zone into a server name
func MapPuppetZoneToMaster(puppetZone string) string {
	var puppetmaster string
	switch puppetZone {
	case "office":
		puppetmaster = "off-cf-01.dfinet.ch"
	case "infra":
		puppetmaster = "inf-cf-01.dfinet.ch"
	case "clients":
		puppetmaster = "inf-cf-02.dfinet.ch"
	}
	return puppetmaster
}

func main() {
	// Set the zone (office/infra/clients):
	puppetZone := "infra"
	var hostDataFilename = fmt.Sprintf("%s/%s.json", os.Getenv("HOME"), puppetZone)

	fresponse := ReadHostFile(hostDataFilename)

	// Choose a formatter:
	out := &formatters.NodeXMLFormatter{
		HostData:   fresponse.GetHosts(),
		PuppetZone: MapPuppetZoneToMaster(puppetZone),
	}
	// Print the formatted node data for Rundeck:
	fmt.Printf("%v\n", out)
}
