package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type FilestoreInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FilestoreInstanceSpec   `json:"spec,omitempty"`
	Status            FilestoreInstanceStatus `json:"status,omitempty"`
}

type FilestoreInstanceSpecFileShares struct {
	CapacityGb int    `json:"capacityGb" tf:"capacity_gb"`
	Name       string `json:"name" tf:"name"`
}

type FilestoreInstanceSpecNetworks struct {
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`
	Modes       []string `json:"modes" tf:"modes"`
	Network     string   `json:"network" tf:"network"`
	// +optional
	ReservedIPRange string `json:"reservedIPRange,omitempty" tf:"reserved_ip_range,omitempty"`
}

type FilestoreInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreateTime string `json:"createTime,omitempty" tf:"create_time,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Etag string `json:"etag,omitempty" tf:"etag,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	FileShares []FilestoreInstanceSpecFileShares `json:"fileShares" tf:"file_shares"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	Name   string            `json:"name" tf:"name"`
	// +kubebuilder:validation:MinItems=1
	Networks []FilestoreInstanceSpecNetworks `json:"networks" tf:"networks"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	Tier    string `json:"tier" tf:"tier"`
	Zone    string `json:"zone" tf:"zone"`
}

type FilestoreInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FilestoreInstanceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FilestoreInstanceList is a list of FilestoreInstances
type FilestoreInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FilestoreInstance CRD objects
	Items []FilestoreInstance `json:"items,omitempty"`
}
