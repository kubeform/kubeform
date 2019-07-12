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
	Bucket               string            `json:"bucket"`
	Acl                  string            `json:"acl"`
	ContentBase64        string            `json:"content_base64"`
	StorageClass         string            `json:"storage_class"`
	ServerSideEncryption string            `json:"server_side_encryption"`
	Key                  string            `json:"key"`
	ContentType          string            `json:"content_type"`
	Source               string            `json:"source"`
	ContentEncoding      string            `json:"content_encoding"`
	ContentLanguage      string            `json:"content_language"`
	Etag                 string            `json:"etag"`
	VersionId            string            `json:"version_id"`
	Tags                 map[string]string `json:"tags"`
	CacheControl         string            `json:"cache_control"`
	ContentDisposition   string            `json:"content_disposition"`
	Content              string            `json:"content"`
	KmsKeyId             string            `json:"kms_key_id"`
	WebsiteRedirect      string            `json:"website_redirect"`
}

type AwsS3BucketObjectStatus struct {
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
