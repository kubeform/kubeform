package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeImageSpec   `json:"spec,omitempty"`
	Status            GoogleComputeImageStatus `json:"status,omitempty"`
}

type GoogleComputeImageSpecRawDisk struct {
	Source        string `json:"source"`
	Sha1          string `json:"sha1"`
	ContainerType string `json:"container_type"`
}

type GoogleComputeImageSpec struct {
	Family           string                   `json:"family"`
	SourceDisk       string                   `json:"source_disk"`
	RawDisk          []GoogleComputeImageSpec `json:"raw_disk"`
	SelfLink         string                   `json:"self_link"`
	Labels           map[string]string        `json:"labels"`
	Licenses         []string                 `json:"licenses"`
	LabelFingerprint string                   `json:"label_fingerprint"`
	Description      string                   `json:"description"`
	Project          string                   `json:"project"`
	CreateTimeout    int                      `json:"create_timeout"`
	Name             string                   `json:"name"`
}



type GoogleComputeImageStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeImageList is a list of GoogleComputeImages
type GoogleComputeImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeImage CRD objects
	Items []GoogleComputeImage `json:"items,omitempty"`
}