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

type LinodeImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeImageSpec   `json:"spec,omitempty"`
	Status            LinodeImageStatus `json:"status,omitempty"`
}

type LinodeImageSpec struct {
	DiskId      int    `json:"disk_id"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	Deprecated  bool   `json:"deprecated"`
	IsPublic    bool   `json:"is_public"`
	Type        string `json:"type"`
	Expiry      string `json:"expiry"`
	Label       string `json:"label"`
	Vendor      string `json:"vendor"`
	Created     string `json:"created"`
	Size        int    `json:"size"`
	LinodeId    int    `json:"linode_id"`
}



type LinodeImageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LinodeImageList is a list of LinodeImages
type LinodeImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeImage CRD objects
	Items []LinodeImage `json:"items,omitempty"`
}