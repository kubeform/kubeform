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

type StorageBucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketObjectSpec   `json:"spec,omitempty"`
	Status            StorageBucketObjectStatus `json:"status,omitempty"`
}

type StorageBucketObjectSpec struct {
	Bucket string `json:"bucket"`
	// +optional
	CacheControl string `json:"cache_control,omitempty"`
	// +optional
	Content string `json:"content,omitempty"`
	// +optional
	ContentDisposition string `json:"content_disposition,omitempty"`
	// +optional
	ContentEncoding string `json:"content_encoding,omitempty"`
	// +optional
	ContentLanguage string `json:"content_language,omitempty"`
	// +optional
	DetectMd5hash string `json:"detect_md5hash,omitempty"`
	Name          string `json:"name"`
	// +optional
	Source string `json:"source,omitempty"`
}

type StorageBucketObjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketObjectList is a list of StorageBucketObjects
type StorageBucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucketObject CRD objects
	Items []StorageBucketObject `json:"items,omitempty"`
}
