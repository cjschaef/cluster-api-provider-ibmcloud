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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	capiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	// ClusterFinalizer allows DockerClusterReconciler to clean up resources associated with DockerCluster before
	// removing it from the apiserver.
	ClusterFinalizer = "ibmvpccluster.infrastructure.cluster.x-k8s.io"
)

// IBMVPCClusterSpec defines the desired state of IBMVPCCluster.
type IBMVPCClusterSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// cisInstance defines the IBM Cloud CIS instance to create DNS Records for the cluster.
	CISInstance *CISInstance `json:"cisInstance,omitempty"`

	// controlPlaneEndpoint represents the endpoint used to communicate with the control plane.
	// +optional
	ControlPlaneEndpoint capiv1beta1.APIEndpoint `json:"controlPlaneEndpoint"`

	// controlPlaneLoadBalancer is optional configuration for customizing control plane behavior.
	// +optional
	ControlPlaneLoadBalancer []*VPCLoadBalancerSpec `json:"controlPlaneLoadBalancer,omitempty"`
	
	COSInstance *CosInstance `json:"cosInstance,omitempty"`

	// dnsServicesInstance defines the IBM Cloud DNS Services instance to create DNS Records for the cluster.
	DNSServicesInstance *DNSServicesInstance `json:"dnsServicesInstance,omitempty"`

	// networkSpec represents the VPC network to use for the cluster.
	NetworkSpec *VPCNetworkSpec `json:"networkSpec,omitempty"`

	// region defines the IBM Cloud Region the cluster resources will be deployed in.
	Region string `json:"region"`

	// resourceGroup defines the IBM Cloud resource group where the cluster resources should be created.
	ResourceGroup string `json:"resourceGroup"`
}

// VPCNetworkSpec defines the desired state of the network resources for the cluster.
type VPCNetworkSpec struct {
	ComputeSubnetsSpec []Subnet `json:"computeSubnetsSpec,omitempty"`

	ControlPlaneSubnetsSpec []Subnet `json:"controlPlaneSubentsSpec,omitempty"`

	ResourceGroup string `json:"resourceGroup,omitempty"`

	SecurityGroups []SecurityGroup `json:"securityGroups,omitempty"`

	VPC *VPCResourceReference `json:"vpc,omitempty"`
}

// VPCLoadBalancerSpec defines the desired state of an VPC load balancer.
type VPCLoadBalancerSpec struct {
	// name sets the name of the VPC load balancer.
	// +kubebuilder:validation:MaxLength:=63
	// +kubebuilder:validation:Pattern=`^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$`
	// +optional
	Name string `json:"name,omitempty"`

	// public indicates that load balancer is public or private
	// +kubebuilder:default=true
	// +optional
	Public bool `json:"public,omitempty"`

	// additionalListeners sets the additional listeners for the control plane load balancer. .
	// +listType=map
	// +listMapKey=port
	// +optional
	AdditionalListeners []AdditionalListenerSpec `json:"additionalListeners,omitempty"`
}

// AdditionalListenerSpec defines the desired state of an
// additional listener on an VPC load balancer.
type AdditionalListenerSpec struct {
	// port sets the port for the additional listener.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int64 `json:"port"`
}

// VPCLoadBalancerStatus defines the status VPC load balancer.
type VPCLoadBalancerStatus struct {
	// id of VPC load balancer.
	// +optional
	ID *string `json:"id,omitempty"`
	// state is the status of the load balancer.
	State VPCLoadBalancerState `json:"state,omitempty"`
	// hostname is the hostname of load balancer.
	// +optional
	Hostname *string `json:"hostname,omitempty"`
	// +kubebuilder:default=false
	// controllerCreated indicates whether the resource is created by the controller.
	ControllerCreated *bool `json:"controllerCreated,omitempty"`
}

// IBMVPCClusterStatus defines the observed state of IBMVPCCluster.
type IBMVPCClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	VPC VPC `json:"vpc,omitempty"`

	// ready is true when the provider resource is ready.
	// +optional
	Ready       bool        `json:"ready"`
	Subnet      Subnet      `json:"subnet,omitempty"`
	VPCEndpoint VPCEndpoint `json:"vpcEndpoint,omitempty"`

	// controlPlaneLoadBalancerState is the status of the load balancer.
	// +optional
	ControlPlaneLoadBalancerState VPCLoadBalancerState `json:"controlPlaneLoadBalancerState,omitempty"`

	// conditions defines current service state of the load balancer.
	// +optional
	Conditions capiv1beta1.Conditions `json:"conditions,omitempty"`
}

// VPC holds the VPC information.
type VPC struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=ibmvpcclusters,scope=Namespaced,categories=cluster-api
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Cluster",type="string",JSONPath=".metadata.labels.cluster\\.x-k8s\\.io/cluster-name",description="Cluster to which this IBMVPCCluster belongs"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.ready",description="Cluster infrastructure is ready for IBM VPC instances"

// IBMVPCCluster is the Schema for the ibmvpcclusters API.
type IBMVPCCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IBMVPCClusterSpec   `json:"spec,omitempty"`
	Status IBMVPCClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// IBMVPCClusterList contains a list of IBMVPCCluster.
type IBMVPCClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IBMVPCCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IBMVPCCluster{}, &IBMVPCClusterList{})
}

// GetConditions returns the observations of the operational state of the IBMVPCCluster resource.
func (r *IBMVPCCluster) GetConditions() capiv1beta1.Conditions {
	return r.Status.Conditions
}

// SetConditions sets the underlying service state of the IBMVPCCluster to the predescribed clusterv1.Conditions.
func (r *IBMVPCCluster) SetConditions(conditions capiv1beta1.Conditions) {
	r.Status.Conditions = conditions
}
