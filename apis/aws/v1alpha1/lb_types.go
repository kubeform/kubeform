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

type Lb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSpec   `json:"spec,omitempty"`
	Status            LbStatus `json:"status,omitempty"`
}

type LbSpecAccessLogs struct {
	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type LbSpecSubnetMapping struct {
	// +optional
	AllocationID string `json:"allocationID,omitempty" tf:"allocation_id,omitempty"`
	SubnetID     string `json:"subnetID" tf:"subnet_id"`
}

type LbSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogs []LbSpecAccessLogs `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`
	Arn        string             `json:"arn" tf:"arn"`
	ArnSuffix  string             `json:"arnSuffix" tf:"arn_suffix"`
	DnsName    string             `json:"dnsName" tf:"dns_name"`
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
	SubnetMapping []LbSpecSubnetMapping `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Subnets []string `json:"subnets,omitempty" tf:"subnets,omitempty"`
	// +optional
	Tags   map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VpcID  string            `json:"vpcID" tf:"vpc_id"`
	ZoneID string            `json:"zoneID" tf:"zone_id"`
}

type LbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbList is a list of Lbs
type LbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Lb CRD objects
	Items []Lb `json:"items,omitempty"`
}
