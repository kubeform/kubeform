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

type LightsailInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LightsailInstanceSpec   `json:"spec,omitempty"`
	Status            LightsailInstanceStatus `json:"status,omitempty"`
}

type LightsailInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn              string `json:"arn,omitempty" tf:"arn,omitempty"`
	AvailabilityZone string `json:"availabilityZone" tf:"availability_zone"`
	BlueprintID      string `json:"blueprintID" tf:"blueprint_id"`
	BundleID         string `json:"bundleID" tf:"bundle_id"`
	// +optional
	CpuCount int `json:"cpuCount,omitempty" tf:"cpu_count,omitempty"`
	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// +optional
	Ipv6Address string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`
	// +optional
	IsStaticIP bool `json:"isStaticIP,omitempty" tf:"is_static_ip,omitempty"`
	// +optional
	KeyPairName string `json:"keyPairName,omitempty" tf:"key_pair_name,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	PublicIPAddress string `json:"publicIPAddress,omitempty" tf:"public_ip_address,omitempty"`
	// +optional
	RamSize int `json:"ramSize,omitempty" tf:"ram_size,omitempty"`
	// +optional
	UserData string `json:"userData,omitempty" tf:"user_data,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type LightsailInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LightsailInstanceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LightsailInstanceList is a list of LightsailInstances
type LightsailInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LightsailInstance CRD objects
	Items []LightsailInstance `json:"items,omitempty"`
}
