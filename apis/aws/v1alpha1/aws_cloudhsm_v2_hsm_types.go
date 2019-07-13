package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AwsCloudhsmV2Hsm struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudhsmV2HsmSpec   `json:"spec,omitempty"`
	Status            AwsCloudhsmV2HsmStatus `json:"status,omitempty"`
}

type AwsCloudhsmV2HsmSpec struct {
	AvailabilityZone string `json:"availability_zone"`
	IpAddress        string `json:"ip_address"`
	HsmId            string `json:"hsm_id"`
	HsmState         string `json:"hsm_state"`
	HsmEniId         string `json:"hsm_eni_id"`
	ClusterId        string `json:"cluster_id"`
	SubnetId         string `json:"subnet_id"`
}



type AwsCloudhsmV2HsmStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudhsmV2HsmList is a list of AwsCloudhsmV2Hsms
type AwsCloudhsmV2HsmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudhsmV2Hsm CRD objects
	Items []AwsCloudhsmV2Hsm `json:"items,omitempty"`
}