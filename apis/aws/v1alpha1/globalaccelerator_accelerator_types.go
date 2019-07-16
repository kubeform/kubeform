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

type GlobalacceleratorAccelerator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalacceleratorAcceleratorSpec   `json:"spec,omitempty"`
	Status            GlobalacceleratorAcceleratorStatus `json:"status,omitempty"`
}

type GlobalacceleratorAcceleratorSpecAttributes struct {
	// +optional
	FlowLogsEnabled bool `json:"flow_logs_enabled,omitempty"`
	// +optional
	FlowLogsS3Bucket string `json:"flow_logs_s3_bucket,omitempty"`
	// +optional
	FlowLogsS3Prefix string `json:"flow_logs_s3_prefix,omitempty"`
}

type GlobalacceleratorAcceleratorSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Attributes *[]GlobalacceleratorAcceleratorSpec `json:"attributes,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	IpAddressType string `json:"ip_address_type,omitempty"`
	Name          string `json:"name"`
}

type GlobalacceleratorAcceleratorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlobalacceleratorAcceleratorList is a list of GlobalacceleratorAccelerators
type GlobalacceleratorAcceleratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlobalacceleratorAccelerator CRD objects
	Items []GlobalacceleratorAccelerator `json:"items,omitempty"`
}
