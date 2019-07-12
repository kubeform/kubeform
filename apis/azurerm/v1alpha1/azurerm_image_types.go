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

type AzurermImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermImageSpec   `json:"spec,omitempty"`
	Status            AzurermImageStatus `json:"status,omitempty"`
}

type AzurermImageSpecOsDisk struct {
	BlobUri       string `json:"blob_uri"`
	Caching       string `json:"caching"`
	SizeGb        int    `json:"size_gb"`
	OsType        string `json:"os_type"`
	OsState       string `json:"os_state"`
	ManagedDiskId string `json:"managed_disk_id"`
}

type AzurermImageSpecDataDisk struct {
	Lun           int    `json:"lun"`
	ManagedDiskId string `json:"managed_disk_id"`
	BlobUri       string `json:"blob_uri"`
	Caching       string `json:"caching"`
	SizeGb        int    `json:"size_gb"`
}

type AzurermImageSpec struct {
	Name                   string             `json:"name"`
	Location               string             `json:"location"`
	ResourceGroupName      string             `json:"resource_group_name"`
	ZoneResilient          bool               `json:"zone_resilient"`
	SourceVirtualMachineId string             `json:"source_virtual_machine_id"`
	OsDisk                 []AzurermImageSpec `json:"os_disk"`
	DataDisk               []AzurermImageSpec `json:"data_disk"`
	Tags                   map[string]string  `json:"tags"`
}

type AzurermImageStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermImageList is a list of AzurermImages
type AzurermImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermImage CRD objects
	Items []AzurermImage `json:"items,omitempty"`
}
