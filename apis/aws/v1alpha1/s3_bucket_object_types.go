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

type S3BucketObject struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketObjectSpec   `json:"spec,omitempty"`
	Status            S3BucketObjectStatus `json:"status,omitempty"`
}

type S3BucketObjectSpec struct {
	// +optional
	Acl    string `json:"acl,omitempty"`
	Bucket string `json:"bucket"`
	// +optional
	CacheControl string `json:"cache_control,omitempty"`
	// +optional
	Content string `json:"content,omitempty"`
	// +optional
	ContentBase64 string `json:"content_base64,omitempty"`
	// +optional
	ContentDisposition string `json:"content_disposition,omitempty"`
	// +optional
	ContentEncoding string `json:"content_encoding,omitempty"`
	// +optional
	ContentLanguage string `json:"content_language,omitempty"`
	Key             string `json:"key"`
	// +optional
	KmsKeyId string `json:"kms_key_id,omitempty"`
	// +optional
	Source string `json:"source,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	WebsiteRedirect string `json:"website_redirect,omitempty"`
}

type S3BucketObjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// S3BucketObjectList is a list of S3BucketObjects
type S3BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of S3BucketObject CRD objects
	Items []S3BucketObject `json:"items,omitempty"`
}
