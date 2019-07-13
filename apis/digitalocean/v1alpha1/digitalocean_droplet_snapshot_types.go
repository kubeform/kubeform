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

type DigitaloceanDropletSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanDropletSnapshotSpec   `json:"spec,omitempty"`
	Status            DigitaloceanDropletSnapshotStatus `json:"status,omitempty"`
}

type DigitaloceanDropletSnapshotSpec struct {
	MinDiskSize int      `json:"min_disk_size"`
	Name        string   `json:"name"`
	DropletId   string   `json:"droplet_id"`
	Regions     []string `json:"regions"`
	Size        float64  `json:"size"`
	CreatedAt   string   `json:"created_at"`
}



type DigitaloceanDropletSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanDropletSnapshotList is a list of DigitaloceanDropletSnapshots
type DigitaloceanDropletSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanDropletSnapshot CRD objects
	Items []DigitaloceanDropletSnapshot `json:"items,omitempty"`
}