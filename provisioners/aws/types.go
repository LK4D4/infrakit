package aws

import (
	"fmt"
	"github.com/docker/libmachete/provisioners/api"
	"reflect"
)

// CreateInstanceRequest is the struct used to create new instances.
type CreateInstanceRequest struct {
	api.BaseMachineRequest   `yaml:",inline"`
	AssociatePublicIPAddress bool              `yaml:"associate_public_ip_address" json:"associate_public_ip_address"`
	AvailabilityZone         string            `yaml:"availability_zone" json:"availability_zone"`
	BlockDeviceName          string            `yaml:"block_device_name" json:"block_device_name"`
	DeleteOnTermination      bool              `yaml:"delete_on_termination" json:"delete_on_termination"`
	EbsOptimized             bool              `yaml:"ebs_optimized" json:"ebs_optimized"`
	IamInstanceProfile       string            `yaml:"iam_instance_profile" json:"iam_instance_profile"`
	ImageID                  string            `yaml:"image_id" json:"image_id"`
	InstanceType             string            `yaml:"instance_type" json:"instance_type"`
	KeyName                  string            `yaml:"key_name" json:"key_name"`
	Monitoring               bool              `yaml:"monitoring" json:"monitoring"`
	PublicIPAddress          string            `yaml:"public_ip_address" json:"public_ip_address"`
	PrivateIPAddress         string            `yaml:"private_ip_address" json:"private_ip_address"`
	PrivateIPOnly            bool              `yaml:"private_ip_only" json:"private_ip_only"`
	RootSize                 int64             `yaml:"root_size" json:"root_size"`
	SecurityGroupIds         []string          `yaml:"security_group_ids,flow" json:"security_group_ids"`
	SubnetID                 string            `yaml:"subnet_id" json:"subnet_id"`
	Tags                     map[string]string `yaml:"tags" json:"tags"`
	VolumeType               string            `yaml:"volume_type" json:"volume_type"`
	VpcID                    string            `yaml:"vpc_id" json:"vpc_id"`
}

// Validate checks the data and returns error if not valid
func (req CreateInstanceRequest) Validate() error {
	// TODO finish this.
	return nil
}

func checkCredential(cred api.Credential) (c *credential, err error) {
	is := false
	if c, is = cred.(*credential); !is {
		err = fmt.Errorf("credential type mismatch: %v", reflect.TypeOf(cred))
		return
	}
	return
}

func checkMachineRequest(req api.MachineRequest) (r *CreateInstanceRequest, err error) {
	is := false
	if r, is = req.(*CreateInstanceRequest); !is {
		err = fmt.Errorf("request type mismatch: %v", reflect.TypeOf(req))
		return
	}
	return
}
