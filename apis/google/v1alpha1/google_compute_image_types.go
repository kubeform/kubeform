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

type GoogleComputeImage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeImageSpec   `json:"spec,omitempty"`
	Status            GoogleComputeImageStatus `json:"status,omitempty"`
}

type GoogleComputeImageSpecRawDisk struct {
	ContainerType string `json:"container_type"`
	Source        string `json:"source"`
	Sha1          string `json:"sha1"`
}

type GoogleComputeImageSpec struct {
	CreateTimeout    int                      `json:"create_timeout"`
	Labels           map[string]string        `json:"labels"`
	Licenses         []string                 `json:"licenses"`
	LabelFingerprint string                   `json:"label_fingerprint"`
	Name             string                   `json:"name"`
	Family           string                   `json:"family"`
	Project          string                   `json:"project"`
	SelfLink         string                   `json:"self_link"`
	Description      string                   `json:"description"`
	SourceDisk       string                   `json:"source_disk"`
	RawDisk          []GoogleComputeImageSpec `json:"raw_disk"`
}

type GoogleComputeImageStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeImageList is a list of GoogleComputeImages
type GoogleComputeImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeImage CRD objects
	Items []GoogleComputeImage `json:"items,omitempty"`
}
