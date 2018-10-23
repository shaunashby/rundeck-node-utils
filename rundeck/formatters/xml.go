package formatters

import "fmt"
import "github.com/shaunashby/rundeck-node-utils/foreman"
import "bytes"

// NodeXMLFormatter is a type to contain the data for a Foreman host entry
type NodeXMLFormatter struct {
	HostData   []foreman.ForemanHost
	PuppetZone string
}

// Output node to XML format using a template (eventually):
func (f *NodeXMLFormatter) String() string {
	defType := "Node"
	defUser := "rundeck"

	var nodeStringBuffer bytes.Buffer

	// Top-level XML tag is project:
	fmt.Fprintf(&nodeStringBuffer, "<project>\n")

	// Iterate over the hosts in the host array:
	for _, host := range f.HostData {
		fmt.Fprintf(&nodeStringBuffer, "     <node name=\"%s\" type=\"%s\"\n", host.Name(), defType)
		fmt.Fprintf(&nodeStringBuffer, "         hostname=\"%s\"\n", host.IpAddress())
		fmt.Fprintf(&nodeStringBuffer, "         description=\"%s\"\n", host.Comment())
		fmt.Fprintf(&nodeStringBuffer, "         osArch=\"%s\"\n", host.Arch())
		fmt.Fprintf(&nodeStringBuffer, "         puppetmaster=\"%s\"\n", f.PuppetZone)
		fmt.Fprintf(&nodeStringBuffer, "         hostgroup=\"%s\"\n", host.Hostgroup())
		fmt.Fprintf(&nodeStringBuffer, "         username=\"%s\" />\n", defUser)
		// TODO: add tags derived from Foreman metadata:
		// tags: x,y,z
	}
	fmt.Fprintf(&nodeStringBuffer, "</project>")

	return nodeStringBuffer.String()
}
