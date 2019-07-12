package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleStorageBucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageBucketObjectSpec   `json:"spec,omitempty"`
	Status            GoogleStorageBucketObjectStatus `json:"status,omitempty"`
}

type GoogleStorageBucketObjectSpec struct {
	Crc32c             string `json:"crc32c"`
	Source             string `json:"source"`
	DetectMd5hash      string `json:"detect_md5hash"`
	Bucket             string `json:"bucket"`
	CacheControl       string `json:"cache_control"`
	Content            string `json:"content"`
	PredefinedAcl      string `json:"predefined_acl"`
	Name               string `json:"name"`
	ContentLanguage    string `json:"content_language"`
	ContentType        string `json:"content_type"`
	StorageClass       string `json:"storage_class"`
	ContentDisposition string `json:"content_disposition"`
	ContentEncoding    string `json:"content_encoding"`
	Md5hash            string `json:"md5hash"`
}

type GoogleStorageBucketObjectStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleStorageBucketObjectList is a list of GoogleStorageBucketObjects
type GoogleStorageBucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageBucketObject CRD objects
	Items []GoogleStorageBucketObject `json:"items,omitempty"`
}
