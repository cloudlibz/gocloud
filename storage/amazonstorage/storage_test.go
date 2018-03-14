package amazonstorage

import "testing"
import "github.com/cloudlibz/gocloud/auth"

func init() {
	awsAuth.LoadConfig()
}

func TestCreatedisk(t *testing.T) {
	var amazonstorage Amazonstorage

	createdisk := map[string]interface{}{
		"AvailZone":  "us-east-1a",
		"VolumeSize": 100,
		"Region":     "us-east-1",
	}
	_, err := amazonstorage.Createdisk(createdisk)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestDeletedisk(t *testing.T) {
	var amazonstorage Amazonstorage

	deletedisk := map[string]string{
		"VolumeId": "vol-0996a16ff8f032760",
		"Region":   "us-east-1",
	}
	_, err := amazonstorage.Deletedisk(deletedisk)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestAttachdisk(t *testing.T) {
	var amazonstorage Amazonstorage

	attachdisk := map[string]string{
		"VolumeId":   "vol-049426a70587418d7",
		"InstanceId": "i-0050d952f9b8d45d5",
		"Device":     "/dev/sdh",
		"Region":     "us-east-1",
	}

	_, err := amazonstorage.Attachdisk(attachdisk)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDetachdisk(t *testing.T) {
	var amazonstorage Amazonstorage

	detachdisk := map[string]string{
		"VolumeId":   "vol-049426a70587418d7",
		"InstanceId": "i-0050d952f9b8d45d5",
		"Device":     "/dev/sdh",
		"Force":      "true",
		"Region":     "us-east-1",
	}

	_, err := amazonstorage.Detachdisk(detachdisk)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreatesnapshot(t *testing.T) {
	var amazonstorage Amazonstorage

	createsnapshot := map[string]string{
		"VolumeId":    "vol-047d011f7536d2b7c",
		"Description": "",
		"Region":      "us-east-1",
	}
	_, err := amazonstorage.Createsnapshot(createsnapshot)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeletesnapshot(t *testing.T) {
	var amazonstorage Amazonstorage

	deletesnapshot := map[string]string{
		"SnapshotId": "snap-0f0839076356ce6cb",
		"Region":     "us-east-1",
	}

	_, err := amazonstorage.Deletesnapshot(deletesnapshot)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
