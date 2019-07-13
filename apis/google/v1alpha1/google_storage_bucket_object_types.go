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

type GoogleStorageBucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageBucketObjectSpec   `json:"spec,omitempty"`
	Status            GoogleStorageBucketObjectStatus `json:"status,omitempty"`
}

type GoogleStorageBucketObjectSpec struct {
	ContentType        string `json:"content_type"`
	PredefinedAcl      string `json:"predefined_acl"`
	Source             string `json:"source"`
	CacheControl       string `json:"cache_control"`
	Name               string `json:"name"`
	ContentDisposition string `json:"content_disposition"`
	Md5hash            string `json:"md5hash"`
	StorageClass       string `json:"storage_class"`
	Bucket             string `json:"bucket"`
	ContentLanguage    string `json:"content_language"`
	Content            string `json:"content"`
	ContentEncoding    string `json:"content_encoding"`
	DetectMd5hash      string `json:"detect_md5hash"`
	Crc32c             string `json:"crc32c"`
}



type GoogleStorageBucketObjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageBucketObjectList is a list of GoogleStorageBucketObjects
type GoogleStorageBucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageBucketObject CRD objects
	Items []GoogleStorageBucketObject `json:"items,omitempty"`
}