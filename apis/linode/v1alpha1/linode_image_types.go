package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LinodeImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeImageSpec   `json:"spec,omitempty"`
	Status            LinodeImageStatus `json:"status,omitempty"`
}

type LinodeImageSpec struct {
	Deprecated  bool   `json:"deprecated"`
	Size        int    `json:"size"`
	Vendor      string `json:"vendor"`
	Description string `json:"description"`
	Created     string `json:"created"`
	CreatedBy   string `json:"created_by"`
	IsPublic    bool   `json:"is_public"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	DiskId      int    `json:"disk_id"`
	LinodeId    int    `json:"linode_id"`
	Expiry      string `json:"expiry"`
}



type LinodeImageStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinodeImageList is a list of LinodeImages
type LinodeImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeImage CRD objects
	Items []LinodeImage `json:"items,omitempty"`
}