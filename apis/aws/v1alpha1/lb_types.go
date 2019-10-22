package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

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

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogs []LbSpecAccessLogs `json:"accessLogs,omitempty" tf:"access_logs,omitempty"`
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
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	SubnetMapping []LbSpecSubnetMapping `json:"subnetMapping,omitempty" tf:"subnet_mapping,omitempty"`
	// +optional
	Subnets []string `json:"subnets,omitempty" tf:"subnets,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
	// +optional
	ZoneID string `json:"zoneID,omitempty" tf:"zone_id,omitempty"`
}

type LbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
