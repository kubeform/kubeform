package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type CloudhsmV2Hsm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudhsmV2HsmSpec   `json:"spec,omitempty"`
	Status            CloudhsmV2HsmStatus `json:"status,omitempty"`
}

type CloudhsmV2HsmSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AvailabilityZone string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	ClusterID        string `json:"clusterID" tf:"cluster_id"`
	// +optional
	HsmEniID string `json:"hsmEniID,omitempty" tf:"hsm_eni_id,omitempty"`
	// +optional
	HsmID string `json:"hsmID,omitempty" tf:"hsm_id,omitempty"`
	// +optional
	HsmState string `json:"hsmState,omitempty" tf:"hsm_state,omitempty"`
	// +optional
	IpAddress string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
}



type CloudhsmV2HsmStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *CloudhsmV2HsmSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudhsmV2HsmList is a list of CloudhsmV2Hsms
type CloudhsmV2HsmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudhsmV2Hsm CRD objects
	Items []CloudhsmV2Hsm `json:"items,omitempty"`
}