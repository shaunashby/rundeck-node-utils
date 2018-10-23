//____________________________________________________________________
// File: ini.go
//____________________________________________________________________
//
// Author: Shaun Ashby <shaun@ashby.ch>
// Created: 2018-08-20 11:21:53+0200
// Revision: $Id$
// Description: Formatter compatible with Ansible hosts file
//
// Copyright (C) 2018 Shaun Ashby
//
//
//--------------------------------------------------------------------
package formatters

import "fmt"
import "github.com/shaunashby/rundeck-node-utils/foreman"
import "bytes"

type NodeIniFormatter struct {
	HostData []foreman.ForemanHost
	PuppetZone string
}

// Write out the data in host file format:
func (f *NodeIniFormatter) String() string {

	var nodeStringBuffer bytes.Buffer

	// Top-level XML tag is project:
	fmt.Fprintf(&nodeStringBuffer,"[%s]\n",f.PuppetZone)

	// Iterate over the hosts in the host array:
	for _, host := range f.HostData {
		fmt.Fprintf(&nodeStringBuffer,"%-30s ansible_ssh_host=\"%s\"\n", host.Name(), host.IpAddress())
	}
	return nodeStringBuffer.String()
}
