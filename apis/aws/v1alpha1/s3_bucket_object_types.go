package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Acl    string `json:"acl,omitempty" tf:"acl,omitempty"`
	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	CacheControl string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`
	// +optional
	Content string `json:"content,omitempty" tf:"content,omitempty"`
	// +optional
	ContentBase64 string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`
	// +optional
	ContentDisposition string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`
	// +optional
	ContentEncoding string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`
	// +optional
	ContentLanguage string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`
	// +optional
	ContentType string `json:"contentType,omitempty" tf:"content_type,omitempty"`
	// +optional
	Etag string `json:"etag,omitempty" tf:"etag,omitempty"`
	Key  string `json:"key" tf:"key"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	ServerSideEncryption string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty"`
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	WebsiteRedirect string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type S3BucketObjectStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
