package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Alb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbSpec   `json:"spec,omitempty"`
	Status            AlbStatus `json:"status,omitempty"`
}

type AlbSpecAccessLogs struct {
	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type AlbSpecSubnetMapping struct {
	// +optional
	AllocationID string `json:"allocationID,omitempty" tf:"allocation_id,omitempty"`
	SubnetID     string `json:"subnetID" tf:"subnet_id"`
}

type AlbSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogs []AlbSpecAccessLogs `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	ArnSuffix string `json:"arnSuffix,omitempty" tf:"arn_suffix,omitempty"`
	// +optional
	DnsName string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`
	// +optional
	EnableCrossZoneLoadBalancing bool `json:"enableCrossZoneLoadBalancing,omitempty" tf:"enable_cross_zone_load_balancing,omitempty"`
	// +optional
	EnableDeletionProtection bool `json:"enableDeletionProtection,omitempty" tf:"enable_deletion_protection,omitempty"`
	// +optional
	EnableHttp2 bool `json:"enableHttp2,omitempty" tf:"enable_http2,omitempty"`
	// +optional
	IdleTimeout int `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`
	// +optional
	Internal bool `json:"internal,omitempty" tf:"internal,omitempty"`
	// +optional
	IpAddressType string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`
	// +optional
	LoadBalancerType string `json:"loadBalancerType,omitempty" tf:"load_balancer_type,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubnetMapping []AlbSpecSubnetMapping `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets,omitempty" tf:"subnets,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
	// +optional
	ZoneID string `json:"zoneID,omitempty" tf:"zone_id,omitempty"`
}

type AlbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AlbSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AlbList is a list of Albs
type AlbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Alb CRD objects
	Items []Alb `json:"items,omitempty"`
}
