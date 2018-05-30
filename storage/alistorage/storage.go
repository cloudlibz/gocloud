package alistorage

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"strconv"
)

// Createdisk create ECS-Disk accept map[string]interface{}
func (aliStorage *Alistorage) Createdisk(request interface{}) (resp interface{}, err error) {
	var options CreateDisk

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "RegionId":
			regionID, _ := value.(string)
			options.RegionID = regionID
		case "ZoneId":
			zoneID, _ := value.(string)
			options.ZoneID = zoneID
		case "DiskName":
			diskName, _ := value.(string)
			options.DiskName = diskName
		case "Description":
			description, _ := value.(string)
			options.Description = description
		case "Encrypted":
			switch value.(type) {
			case bool:
				options.Encrypted = value.(bool)
			case string:
				options.Encrypted = value.(string) == "true" || value.(string) == "True"
			}
		case "DiskCategory":
			diskCategory, _ := value.(string)
			options.DiskCategory = diskCategory
		case "Size":
			switch value.(type) {
			case int:
				options.Size = value.(int)
			case string:
				options.Size, _ = strconv.Atoi(value.(string))
			}
		case "SnapshotId":
			snapshotID, _ := value.(string)
			options.SnapshotID = snapshotID
		case "ClientToken":
			clientToken, _ := value.(string)
			options.ClientToken = clientToken
		}
	}

	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("CreateDisk", params, response)
	resp = response
	return resp, err
}

// Deletedisk delete ECS-Disk accept map[string]interface{}
func (aliStorage *Alistorage) Deletedisk(request interface{}) (resp interface{}, err error) {
	var options DeleteDisk

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "DiskId":
			diskID, _ := value.(string)
			options.DiskID = diskID
		}
	}
	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("DeleteDisk", params, response)
	resp = response
	return resp, err
}

// Createsnapshot create snapshot accept map[string]interface{}
func (aliStorage *Alistorage) Createsnapshot(request interface{}) (resp interface{}, err error) {
	var options CreateSnapshot

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "DiskId":
			diskID, _ := value.(string)
			options.DiskID = diskID
		case "SnapshotName":
			snapshotName, _ := value.(string)
			options.SnapshotName = snapshotName
		case "Description":
			description, _ := value.(string)
			options.Description = description
		case "ClientToken":
			clientToken, _ := value.(string)
			options.ClientToken = clientToken
		}
	}
	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("CreateSnapshot", params, response)
	resp = response
	return resp, err
}

// Deletesnapshot delete snapshot accept map[string]interface{}
func (aliStorage *Alistorage) Deletesnapshot(request interface{}) (resp interface{}, err error) {
	var options DeleteSnapshot

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "SnapshotId":
			snapshotID, _ := value.(string)
			options.SnapshotID = snapshotID
		}
	}
	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("DeleteSnapshot", params, response)
	resp = response
	return resp, err
}

// Attachdisk attach ECS-Disk to ECS, accept map[string]interface{}
func (aliStorage *Alistorage) Attachdisk(request interface{}) (resp interface{}, err error) {
	var options AttachDisk

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "InstanceId":
			instanceID, _ := value.(string)
			options.InstanceID = instanceID
		case "DiskId":
			diskID, _ := value.(string)
			options.DiskID = diskID
		case "DeleteWithInstance":
			switch value.(type) {
			case bool:
				options.DeleteWithInstance = value.(bool)
			case string:
				options.DeleteWithInstance = value.(string) == "true" || value.(string) == "True"
			}
		}
	}
	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("AttachDisk", params, response)
	resp = response
	return resp, err
}

// Detachdisk detach ECS-Disk from ECS, accept map[string]interface{}
func (aliStorage *Alistorage) Detachdisk(request interface{}) (resp interface{}, err error) {
	var options AttachDisk

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "InstanceId":
			instanceID, _ := value.(string)
			options.InstanceID = instanceID
		case "DiskId":
			diskID, _ := value.(string)
			options.DiskID = diskID
		}
	}
	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.ECSSignAndDoRequest("DetachDisk", params, response)
	resp = response
	return resp, err
}
