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

type AwsS3BucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketObjectSpec   `json:"spec,omitempty"`
	Status            AwsS3BucketObjectStatus `json:"status,omitempty"`
}

type AwsS3BucketObjectSpec struct {
	Source               string            `json:"source"`
	VersionId            string            `json:"version_id"`
	Tags                 map[string]string `json:"tags"`
	WebsiteRedirect      string            `json:"website_redirect"`
	Acl                  string            `json:"acl"`
	ContentLanguage      string            `json:"content_language"`
	Content              string            `json:"content"`
	ServerSideEncryption string            `json:"server_side_encryption"`
	Etag                 string            `json:"etag"`
	ContentType          string            `json:"content_type"`
	ContentBase64        string            `json:"content_base64"`
	CacheControl         string            `json:"cache_control"`
	ContentDisposition   string            `json:"content_disposition"`
	ContentEncoding      string            `json:"content_encoding"`
	StorageClass         string            `json:"storage_class"`
	KmsKeyId             string            `json:"kms_key_id"`
	Bucket               string            `json:"bucket"`
	Key                  string            `json:"key"`
}



type AwsS3BucketObjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsS3BucketObjectList is a list of AwsS3BucketObjects
type AwsS3BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3BucketObject CRD objects
	Items []AwsS3BucketObject `json:"items,omitempty"`
}