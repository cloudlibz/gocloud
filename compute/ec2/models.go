package ec2

import (
	"encoding/xml"
)

//runinstance to store all attribute to create EC2 instance

type RunInstances struct {
	ImageId               string
	MinCount              int
	MaxCount              int
	KeyName               string
	InstanceType          string
	SecurityGroups        []SecurityGroup
	KernelId              string
	RamdiskId             string
	UserData              []byte
	AvailZone             string
	PlacementGroupName    string
	Monitoring            bool
	SubnetId              string
	DisableAPITermination bool
	ShutdownBehavior      string
	PrivateIPAddress      string
	BlockDeviceMappings   []BlockDeviceMapping
	NetworkInterfaces     []RunNetworkInterface
}

//SecurityGroup struct

type SecurityGroup struct {
	Id   string `xml:"groupId"`
	Name string `xml:"groupName"`
}

// BlockDevice struct to attach device

type BlockDeviceMapping struct {
	DeviceName          string `xml:"deviceName"`
	VirtualName         string `xml:"virtualName"`
	SnapshotId          string `xml:"ebs>snapshotId"`
	VolumeType          string `xml:"ebs>volumeType"`
	VolumeSize          int64  `xml:"ebs>volumeSize"`
	DeleteOnTermination bool   `xml:"ebs>deleteOnTermination"`
	IOPS                int64  `xml:"ebs>iops"`
}

//NetworkInterface struct for Ec2

type RunNetworkInterface struct {
	Id                      string
	DeviceIndex             int
	SubnetId                string
	Description             string
	PrivateIPs              []PrivateIP
	SecurityGroupIds        []string
	DeleteOnTermination     bool
	SecondaryPrivateIPCount int
}

//Private ip to assign PrivateIP

type PrivateIP struct {
	Address   string `xml:"privateIpAddress"`
	DNSName   string `xml:"privateDnsName"`
	IsPrimary bool   `xml:"primary"`
}

// run instance response

type RunInstancesResp struct {
	RequestId      string          `xml:"requestId"`
	ReservationId  string          `xml:"reservationId"`
	OwnerId        string          `xml:"ownerId"`
	SecurityGroups []SecurityGroup `xml:"groupSet>item"`
	Instances      []Instance      `xml:"instancesSet>item"`
}

// this struct represents running instance

type Instance struct {
	InstanceId         string             `xml:"instanceId"`
	InstanceType       string             `xml:"instanceType"`
	ImageId            string             `xml:"imageId"`
	PrivateDNSName     string             `xml:"privateDnsName"`
	DNSName            string             `xml:"dnsName"`
	IPAddress          string             `xml:"ipAddress"`
	PrivateIPAddress   string             `xml:"privateIpAddress"`
	SubnetId           string             `xml:"subnetId"`
	VPCId              string             `xml:"vpcId"`
	SourceDestCheck    bool               `xml:"sourceDestCheck"`
	KeyName            string             `xml:"keyName"`
	AMILaunchIndex     int                `xml:"amiLaunchIndex"`
	Hypervisor         string             `xml:"hypervisor"`
	VirtType           string             `xml:"virtualizationType"`
	Monitoring         string             `xml:"monitoring>state"`
	AvailZone          string             `xml:"placement>availabilityZone"`
	PlacementGroupName string             `xml:"placement>groupName"`
	State              InstanceState      `xml:"instanceState"`
	Tags               []Tag              `xml:"tagSet>item"`
	SecurityGroups     []SecurityGroup    `xml:"groupSet>item"`
	NetworkInterfaces  []NetworkInterface `xml:"networkInterfaceSet>item"`
}

//This stuct represents instance state change

type InstanceStateChange struct {
	InstanceId    string        `xml:"instanceId"`
	CurrentState  InstanceState `xml:"currentState"`
	PreviousState InstanceState `xml:"previousState"`
}

type SimpleResp struct {
	XMLName   xml.Name
	RequestId string `xml:"requestId"`
}

//struct to TerminateInstance

type TerminateInstancesResp struct {
	RequestId    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

// InstanceState struct

type InstanceState struct {
	Code int    `xml:"code"`
	Name string `xml:"name"`
}

//reperent ruuing instance NetworkInterface

type NetworkInterface struct {
	Id               string                     `xml:"networkInterfaceId"`
	SubnetId         string                     `xml:"subnetId"`
	VPCId            string                     `xml:"vpcId"`
	AvailZone        string                     `xml:"availabilityZone"`
	Description      string                     `xml:"description"`
	OwnerId          string                     `xml:"ownerId"`
	RequesterId      string                     `xml:"requesterId"`
	RequesterManaged bool                       `xml:"requesterManaged"`
	Status           string                     `xml:"status"`
	MACAddress       string                     `xml:"macAddress"`
	PrivateIPAddress string                     `xml:"privateIpAddress"`
	PrivateDNSName   string                     `xml:"privateDnsName"`
	SourceDestCheck  bool                       `xml:"sourceDestCheck"`
	Groups           []SecurityGroup            `xml:"groupSet>item"`
	Attachment       NetworkInterfaceAttachment `xml:"attachment"`
	Tags             []Tag                      `xml:"tagSet>item"`
	PrivateIPs       []PrivateIP                `xml:"privateIpAddressesSet>item"`
}

type NetworkInterfaceAttachment struct {
	Id                  string `xml:"attachmentId"`
	InstanceId          string `xml:"instanceId"`
	InstanceOwnerId     string `xml:"instanceOwnerId"`
	DeviceIndex         int    `xml:"deviceIndex"`
	Status              string `xml:"status"`
	AttachTime          string `xml:"attachTime"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

// reperent tag assgin to instance
type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

//start instance response

type StartInstanceResp struct {
	RequestId    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

//stop instances response

type StopInstanceResp struct {
	RequestId    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}
