package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketObjectSpec   `json:"spec,omitempty"`
	Status            AwsS3BucketObjectStatus `json:"status,omitempty"`
}

type AwsS3BucketObjectSpec struct {
	Etag                 string            `json:"etag"`
	Key                  string            `json:"key"`
	CacheControl         string            `json:"cache_control"`
	ServerSideEncryption string            `json:"server_side_encryption"`
	KmsKeyId             string            `json:"kms_key_id"`
	VersionId            string            `json:"version_id"`
	Tags                 map[string]string `json:"tags"`
	Acl                  string            `json:"acl"`
	ContentDisposition   string            `json:"content_disposition"`
	ContentLanguage      string            `json:"content_language"`
	ContentBase64        string            `json:"content_base64"`
	StorageClass         string            `json:"storage_class"`
	WebsiteRedirect      string            `json:"website_redirect"`
	ContentEncoding      string            `json:"content_encoding"`
	ContentType          string            `json:"content_type"`
	Source               string            `json:"source"`
	Content              string            `json:"content"`
	Bucket               string            `json:"bucket"`
}

type AwsS3BucketObjectStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketObjectList is a list of AwsS3BucketObjects
type AwsS3BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3BucketObject CRD objects
	Items []AwsS3BucketObject `json:"items,omitempty"`
}
