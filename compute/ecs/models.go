package ecs

// CreateInstance to store all attributes to create Ali cloud ECS instance.
type CreateInstance struct {
	RegionID                string
	ZoneID                  string
	ImageID                 string
	InstanceType            string
	SecurityGroupID         string
	InstanceName            string
	Description             string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	HostName                string
	Password                string
	IoOptimized             string
	SystemDiskCategory      string
	SystemDiskSize          string
	SystemDiskName          string
	SystemDiskDescription   string
}

// StartInstance to store all attributes to start Ali cloud ECS instance.
type StartInstance struct {
	InstanceID string
}

// StopInstance to store all attributes to stop Ali cloud ECS instance.
type StopInstance struct {
	InstanceID string
}


// RebootInstance to store all attributes to reboot Ali cloud ECS instance.
type RebootInstance struct {
	InstanceID string
}


// DeleteInstance to store all attribute to delete Ali cloud ECS instance.
type DeleteInstance struct {
	InstanceID string
}
