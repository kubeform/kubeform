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
	Source             string `json:"source"`
	Bucket             string `json:"bucket"`
	Name               string `json:"name"`
	ContentType        string `json:"content_type"`
	Crc32c             string `json:"crc32c"`
	ContentDisposition string `json:"content_disposition"`
	PredefinedAcl      string `json:"predefined_acl"`
	DetectMd5hash      string `json:"detect_md5hash"`
	Content            string `json:"content"`
	Md5hash            string `json:"md5hash"`
	StorageClass       string `json:"storage_class"`
	CacheControl       string `json:"cache_control"`
	ContentEncoding    string `json:"content_encoding"`
	ContentLanguage    string `json:"content_language"`
}

type GoogleStorageBucketObjectStatus struct {
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
