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

type Image struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec,omitempty"`
	Status            ImageStatus `json:"status,omitempty"`
}

type ImageSpecDataDisk struct {
	// +optional
	Caching string `json:"caching,omitempty"`
	// +optional
	Lun int `json:"lun,omitempty"`
	// +optional
	ManagedDiskId string `json:"managed_disk_id,omitempty"`
}

type ImageSpecOsDisk struct {
	// +optional
	Caching string `json:"caching,omitempty"`
	// +optional
	OsState string `json:"os_state,omitempty"`
	// +optional
	OsType string `json:"os_type,omitempty"`
}

type ImageSpec struct {
	// +optional
	DataDisk *[]ImageSpec `json:"data_disk,omitempty"`
	Location string       `json:"location"`
	Name     string       `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OsDisk            *[]ImageSpec `json:"os_disk,omitempty"`
	ResourceGroupName string       `json:"resource_group_name"`
	// +optional
	SourceVirtualMachineId string `json:"source_virtual_machine_id,omitempty"`
	// +optional
	ZoneResilient bool `json:"zone_resilient,omitempty"`
}

type ImageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ImageList is a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Image CRD objects
	Items []Image `json:"items,omitempty"`
}
