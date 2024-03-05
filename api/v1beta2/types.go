/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta2

import (
	"github.com/IBM/vpc-go-sdk/vpcv1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PowerVSInstanceState describes the state of an IBM Power VS instance.
type PowerVSInstanceState string

var (
	// PowerVSInstanceStateACTIVE is the string representing an instance in a ACTIVE state.
	PowerVSInstanceStateACTIVE = PowerVSInstanceState("ACTIVE")

	// PowerVSInstanceStateBUILD is the string representing an instance in a BUILD state.
	PowerVSInstanceStateBUILD = PowerVSInstanceState("BUILD")

	// PowerVSInstanceStateSHUTOFF is the string representing an instance in a SHUTOFF state.
	PowerVSInstanceStateSHUTOFF = PowerVSInstanceState("SHUTOFF")

	// PowerVSInstanceStateREBOOT is the string representing an instance in a REBOOT state.
	PowerVSInstanceStateREBOOT = PowerVSInstanceState("REBOOT")

	// PowerVSInstanceStateERROR is the string representing an instance in a ERROR state.
	PowerVSInstanceStateERROR = PowerVSInstanceState("ERROR")
)

// PowerVSImageState describes the state of an IBM Power VS image.
type PowerVSImageState string

var (
	// PowerVSImageStateACTIVE is the string representing an image in a active state.
	PowerVSImageStateACTIVE = PowerVSImageState("active")

	// PowerVSImageStateQue is the string representing an image in a queued state.
	PowerVSImageStateQue = PowerVSImageState("queued")

	// PowerVSImageStateFailed is the string representing an image in a failed state.
	PowerVSImageStateFailed = PowerVSImageState("failed")

	// PowerVSImageStateImporting is the string representing an image in a failed state.
	PowerVSImageStateImporting = PowerVSImageState("importing")
)

// VPCLoadBalancerState describes the state of the load balancer.
type VPCLoadBalancerState string

var (
	// VPCLoadBalancerStateActive is the string representing the load balancer in a active state.
	VPCLoadBalancerStateActive = VPCLoadBalancerState("active")

	// VPCLoadBalancerStateCreatePending is the string representing the load balancer in a queued state.
	VPCLoadBalancerStateCreatePending = VPCLoadBalancerState("create_pending")

	// VPCLoadBalancerStateDeletePending is the string representing the load balancer in a failed state.
	VPCLoadBalancerStateDeletePending = VPCLoadBalancerState("delete_pending")
)

// DeletePolicy defines the policy used to identify images to be preserved.
type DeletePolicy string

var (
	// DeletePolicyRetain is the string representing an image to be retained.
	DeletePolicyRetain = DeletePolicy("retain")
)

type SecurityGroupRuleAction string

var (
	SecurityGroupRuleActionAllow = vpcv1.NetworkACLRuleActionAllowConst
	SecurityGroupRuleActionDeny = vpcv1.NetworkACLRuleActionDenyConst
)

type SecurityGroupRuleDirection string

var (
	SecurityGroupRuleDirectionInbound = vpcv1.NetworkACLRuleDirectionInboundConst

	SecurityGroupRuleDirectionOutbound = vpcv1.NetworkACLRuleDirectionOutboundConst
)

type SecurityGroupRuleProtocol string

var (
	SecurityGroupRuleProtocolAll = vpcv1.NetworkACLRuleProtocolAllConst
	SecurityGroupRuleProtocolICMP = vpcv1.NetworkACLRuleProtocolIcmpConst
	SecurityGroupRuleProtocolTCP = vpcv1.NetworkACLRuleProtocolTCPConst
	SecurityGroupRuleProtocolUDP = vpcv1.NetworkACLRuleProtocolUDPConst
)

type SecurityGroupRuleRemoteType string

var (
	SecurityGroupRuleRemoteTypeAny = SecurityGroupRuleRemoteType("any")
	SecurityGroupRuleRemoteTypeCIDR = SecurityGroupRuleRemoteType("cidr")
	SecurityGroupRuleRemoteTypeIP = SecurityGroupRuleRemoteType("ip")
	SecurityGroupRuleRemoteTypeSG = SecurityGroupRuleRemoteType("sg")
)

type CISInstance struct {
	Domain string `json:"domain"`
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// CosInstance represents IBM Cloud COS instance.
type CosInstance struct {
	// PresignedURLDuration defines the duration for which presigned URLs are valid.
	//
	// This is used to generate presigned URLs for S3 Bucket objects, which are used by
	// control-plane and worker nodes to fetch bootstrap data.
	//
	// When enabled, the IAM instance profiles specified are not used.
	// +optional
	PresignedURLDuration *metav1.Duration `json:"presignedURLDuration,omitempty"`

	// Name defines name of IBM cloud COS instance to be created.
	// +kubebuilder:validation:MinLength:=3
	// +kubebuilder:validation:MaxLength:=63
	// +kubebuilder:validation:Pattern=`^[a-z0-9][a-z0-9.-]{1,61}[a-z0-9]$`
	Name string `json:"name,omitempty"`

	// bucketName is IBM cloud COS bucket name
	BucketName string `json:"bucketName,omitempty"`

	// bucketRegion is IBM cloud COS bucket region
	BucketRegion string `json:"bucketRegion,omitempty"`
}

type DNSServicesInstance struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Zone string `json:"zone"`
	ZoneLabel string `json:"zoneLabel,omitempty"`
}

// NetworkInterface holds the network interface information like subnet id.
type NetworkInterface struct {
	// Subnet ID of the network interface.
	Subnet string `json:"subnet,omitempty"`
}

type PortRange struct {
	MaximumPort int `json:"maximumPort,omitempty"`
	MinimumPort int `json:"minimumPort,omitempty"`
}

type SecurityGroup struct {
	Name string
	ResourceGroup string
	Rules []*SecurityGroupRule
	Tags []string
	VPC *VPCResourceReference
}

type SecurityGroupRule struct {
	Action SecurityGroupRuleAction
	Destination *SecurityGroupRuleRemoteSpec
	Direction SecurityGroupRuleDirection
	SecurityGroupID string
	Source *SecurityGroupRuleRemoteSpec
}

type SecurityGroupRuleRemote struct {
	CIDRSubnetName string `json:"cidrsubnetname,omitempty"`
	RemoteType SecurityGroupRuleRemoteType `json:"remoteType"`
	SecurityGroupName string `json:"securityGroupName,omitempty"`
}

type SecurityGroupRuleRemoteSpec struct {
	ICMPCode string `json:"icmpCode,omitempty"`
	ICMPType string `json:"icmpType,omitempty"`
	PortRange *PortRange `json:"portRange,omitempty"`
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	Remotes []SecurityGroupRuleRemote `json:"remotes"`
}

// Subnet describes a subnet.
type Subnet struct {
	Ipv4CidrBlock *string `json:"cidr,omitempty"`
	Name          *string `json:"name,omitempty"`
	ID            *string `json:"id,omitempty"`
	Zone          *string `json:"zone,omitempty"`
}

// VPCEndpoint describes a VPCEndpoint.
type VPCEndpoint struct {
	Address *string `json:"address"`
	// +optional
	// Deprecated: This field has no function and is going to be removed in the next release.
	FIPID *string `json:"floatingIPID,omitempty"`
	// +optional
	LBID *string `json:"loadBalancerIPID,omitempty"`
}

// VPCResourceReference is a reference to a specific VPC resource by ID or Name
// Only one of ID or Name may be specified. Specifying more than one will result in
// a validation error.
type VPCResourceReference struct {
	// ID of resource
	// +kubebuilder:validation:MinLength=1
	// +optional
	ID *string `json:"id,omitempty"`

	// Name of resource
	// +kubebuilder:validation:MinLength=1
	// +optional
	Name *string `json:"name,omitempty"`

	// IBM Cloud VPC region
	Region *string `json:"region,omitempty"`
}
