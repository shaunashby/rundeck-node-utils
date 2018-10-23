package formatters

import "fmt"
import "github.com/shaunashby/rundeck-node-utils/foreman"
import "bytes"

// NodeIniFormatter is a type to contain the data for a Foreman host entry
type NodeIniFormatter struct {
	HostData   []foreman.ForemanHost
	PuppetZone string
}

// Write out the data in host file format:
func (f *NodeIniFormatter) String() string {

	var nodeStringBuffer bytes.Buffer

	// Top-level XML tag is project:
	fmt.Fprintf(&nodeStringBuffer, "[%s]\n", f.PuppetZone)

	// Iterate over the hosts in the host array:
	for _, host := range f.HostData {
		fmt.Fprintf(&nodeStringBuffer, "%-30s ansible_ssh_host=\"%s\"\n", host.Name(), host.IpAddress())
	}
	return nodeStringBuffer.String()
}
