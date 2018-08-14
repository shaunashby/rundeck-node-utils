package foreman

import "fmt"

// Type to hold the host information from parsed JSON
//
// {
//     "architecture_name": "i386",
//     "certname": "zmproxy1.dfinet.ch",
//     "comment": null,
//     "hostgroup_name": "zimbraserver",
//     "id": 49,
//     "ip": "195.70.10.142",
//     "mac": "7a:17:19:bd:dc:b4",
//     "name": "zmproxy1.dfinet.ch",
//     "operatingsystem_name": "CentOS 5.11",
//     "puppet_proxy_name": "inf-cf-01.dfinet.ch",
// }
//
type ForemanApiResponse struct {
	ForemanHosts []ForemanHost `json:"results"`
}

func (r *ForemanApiResponse) GetHosts() []ForemanHost {
	return r.ForemanHosts
}

type ForemanHost struct {
	ForemanHostArch            string `json:"architecture_name"`
	ForemanHostName            string `json:"name"`
	ForemanHostCertname        string `json:"certname"`
	ForemanHostComment         string `json:"comment"`
	ForemanHostMemberHostgroup string `json:"hostgroup_name"`
	ForemanHostId              int    `json:"id"`
	ForemanHostIpAddress       string `json:"ip"`
	ForemanHostMacAddress      string `json:"mac"`
	ForemanHostOs              string `json:"operatingsystem_name"`
	ForemanHostPuppetmaster    string `json:"puppet_proxy_name"`
}

// Functions for ForemanHost type:
func (fh *ForemanHost) Arch() string {
	return fh.ForemanHostArch
}

func (fh *ForemanHost) Name() string {
	return fh.ForemanHostName
}

func (fh *ForemanHost) Certname() string {
	return fh.ForemanHostCertname
}

func (fh *ForemanHost) Comment() string {
	return fh.ForemanHostComment
}

func (fh *ForemanHost) Hostgroup() string {
	return fh.ForemanHostMemberHostgroup
}

func (fh *ForemanHost) IpAddress() string {
	return fh.ForemanHostIpAddress
}

func (fh *ForemanHost) Puppetmaster() string {
	return fh.ForemanHostPuppetmaster
}

// Customize formatting for stdout:
func (fh *ForemanHost) String() string {
	return fmt.Sprintf("%s/%s/%s", fh.ForemanHostName, fh.ForemanHostIpAddress, fh.ForemanHostOs)
}
