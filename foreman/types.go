package foreman

import "fmt"

// ForemanApiResponse is a type to store the bytes from a Foreman API response:
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
