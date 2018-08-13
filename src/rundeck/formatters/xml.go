package formatters

import "fmt"
import "foreman"

type NodeXmlFormatter struct {
	HostData *foreman.ForemanHost
}

// Output node to XML format using a template (eventually):
func (f *NodeXmlFormatter) String() string {
	defType := "Node"
	return fmt.Sprintf("<node type=\"%s\" name=\"%s\" />", defType, f.HostData.Name())
}

// Node contents:
// <node>
// name -> hostname
// hostname -> IP address
// description -> comment
// osArch -> Arch
// tags: x,y,z
// username: rundeck
// attributes: puppetmaster,
// attributes: hostgroup
// </node>
