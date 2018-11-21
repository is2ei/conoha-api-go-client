package conoha

import (
	"encoding/json"
	"fmt"
)

// ComputeServer represents the server information.
type ComputeServer struct {
	OsDcfDiskConfig                string                             `json:"OS-DCF:diskConfig,omitempty"`
	OsExtAzAvailabilityZone        string                             `json:"OS-EXT-AZ:availability_zone,omitempty"`
	OsExtSrvAttrHost               string                             `json:"OS-EXT-SRV-ATTR:host,omitempty"`
	OsExtSrvAttrHypervisorHostname string                             `json:"OS-EXT-SRV-ATTR:hypervisor_hostname,omitempty"`
	OsExtSrvAttrInstanceName       string                             `json:"OS-EXT-SRV-ATTR:instance_name,omitempty"`
	OsExtStsPowerState             int                                `json:"OS-EXT-STS:power_state,omitempty"`
	OsExtStsTaskState              string                             `json:"OS-EXT-STS:task_state,omitempty"`
	OsExtStsVmState                string                             `json:"OS-EXT-STS:vm_state,omitempty"`
	OsSrvUsgLaunchedAt             string                             `json:"OS-SRV-USG:launched_at,omitempty"`
	OsSrvUsgTerminatedAt           string                             `json:"OS-SRV-USG:terminated_at,omitempty"`
	AccessIPv4                     string                             `json:"accessIPv4,omitempty"`
	AccessIPv6                     string                             `json:"accessIPv6,omitempty"`
	Addresses                      map[string][]*ComputeServerAddress `json:"addresses,omitempty"`
	ConfigDrive                    string                             `json:"config_drive,omitempty"`
	Created                        string                             `json:"created,omitempty"`
	Flavor                         *ComputeFlavor                     `json:"flavor,omitempty"`
	HostId                         string                             `json:"hostId,omitempty"`
	Id                             string                             `json:"id,omitempty"`
	Image                          *ComputeImage                      `json:"image,omitempty"`
	KeyName                        string                             `json:"key_name,omitempty"`
	Links                          []*ComputeServerLink               `json:"links,omitempty"`
	Metadata                       *ComputeServerMetadata             `json:"metadata,omitempty"`
	Name                           string                             `json:"name,omitempty"`
	Status                         string                             `json:"status,omitempty"`
	TenantId                       string                             `json:"tenant_id,omitempty"`
	Updated                        string                             `json:"updated"`
	UserId                         string                             `json:"user_id"`
	ImageRef                       string                             `json:"imageRef,omitempty"`
	FlavorRef                      string                             `json:"flavorRef,omitempty"`
	AdminPass                      string                             `json:"adminPass,omitempty"`
	SecurityGroups                 []*ComputeSecurityGroup            `json:"security_groups,omitempty"`
}

// ComputeServerAddress represents the server's address information.
type ComputeServerAddress struct {
	Addr string `json:"addr"`
}

// ComputeServerMetadata represents the metadata for the server.
type ComputeServerMetadata struct {
	BackupId        string `json:"backup_id"`
	BackupSet       string `json:"backup_set"`
	BackupStatus    string `json:"backup_status"`
	InstanceNameTag string `json:"instance_name_tag"`
	Properties      string `json:"properties"`
}

type getComputeServerResponseParam struct {
	Server ComputeServer `json:"server"`
}

type computeAddServerRequestParam struct {
	Server ComputeServer `json:"server"`
}

type computeAddServerResponseParam struct {
	Server ComputeServer `json:"server"`
}

type reboot struct {
	Type string `json:"type"`
}

type computeStartServerRequestParam struct {
	OsStart string `json:"os-start"`
}

type computeRebootServerRequestParam struct {
	Reboot reboot `json:"reboot"`
}

// ComputeServerLink represents the link for the server.
type ComputeServerLink struct {
}

// ComputeSecurityGroup represents the security group applied to the server.
type ComputeSecurityGroup struct {
	Name string `json:"name,omitempty"`
}

// ComputeServer fetches the server information.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-get_vms_detail_specified.html
func (c *Conoha) ComputeServer(serverId string) (*ComputeServer, *responseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/servers/" + serverId

	contents, meta, err := c.buildAndExecRequest("GET", u, nil)
	p := getComputeServerResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return &p.Server, meta, err
}

// AddComputeServer adds new server.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-create_vm.html
func (c *Conoha) AddComputeServer(
	imageRef,
	flavorRef,
	adminPass,
	keyName,
	securityGroupName string,
) (*ComputeServer, *responseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/servers"

	if securityGroupName == "" {
		securityGroupName = "default"
	}

	securityGroup := ComputeSecurityGroup{
		Name: securityGroupName,
	}

	server := ComputeServer{
		ImageRef:       imageRef,
		FlavorRef:      flavorRef,
		AdminPass:      adminPass,
		KeyName:        keyName,
		SecurityGroups: []*ComputeSecurityGroup{&securityGroup},
	}
	param := computeAddServerRequestParam{
		Server: server,
	}
	body, _ := json.Marshal(param)

	contents, meta, err := c.buildAndExecRequest("POST", u, body)
	p := computeAddServerResponseParam{}
	if err == nil {
		err = json.Unmarshal(contents, &p)
	}

	return &p.Server, meta, err
}

// DeleteComputeServer deletes the server.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-delete_vm.html
func (c *Conoha) DeleteComputeServer(serverId string) (*responseMeta, error) {
	u := c.ComputeServiceUrl + "/v2/" + c.TenantId + "/servers/" + serverId

	_, meta, err := c.buildAndExecRequest("DELETE", u, nil)

	return meta, err
}

// StartComputeServer starts the serer.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-power_on_vm.html
func (c *Conoha) StartComputeServer(serverId string) (*responseMeta, error) {
	u := fmt.Sprintf("%s/v2/%s/servers/%s/action", c.ComputeServiceUrl, c.TenantId, serverId)

	param := computeStartServerRequestParam{
		OsStart: "null",
	}
	body, _ := json.Marshal(param)

	_, meta, err := c.buildAndExecRequest("POST", u, body)

	return meta, err
}

// RebootComputeServer reboots the server.
//
// ConoHa API docs: https://www.conoha.jp/docs/compute-reboot_vm.html
func (c *Conoha) RebootComputeServer(serverId string, isSoft bool) (*responseMeta, error) {
	u := fmt.Sprintf("%s/v2/%s/servers/%s/action", c.ComputeServiceUrl, c.TenantId, serverId)

	var t string
	if isSoft {
		t = "SOFT"
	} else {
		t = "HARD"
	}

	r := reboot{
		Type: t,
	}

	p := computeRebootServerRequestParam{
		Reboot: r,
	}

	body, _ := json.Marshal(p)

	_, meta, err := c.buildAndExecRequest("POST", u, body)

	return meta, err
}
