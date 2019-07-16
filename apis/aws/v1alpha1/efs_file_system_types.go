package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type EfsFileSystem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EfsFileSystemSpec   `json:"spec,omitempty"`
	Status            EfsFileSystemStatus `json:"status,omitempty"`
}

type EfsFileSystemSpec struct {
	// +optional
	ProvisionedThroughputInMibps json.Number `json:"provisioned_throughput_in_mibps,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	ThroughputMode string `json:"throughput_mode,omitempty"`
}

type EfsFileSystemStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EfsFileSystemList is a list of EfsFileSystems
type EfsFileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EfsFileSystem CRD objects
	Items []EfsFileSystem `json:"items,omitempty"`
}
