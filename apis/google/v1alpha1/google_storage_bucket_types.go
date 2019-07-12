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

type GoogleStorageBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageBucketSpec   `json:"spec,omitempty"`
	Status            GoogleStorageBucketStatus `json:"status,omitempty"`
}

type GoogleStorageBucketSpecLifecycleRuleAction struct {
	Type         string `json:"type"`
	StorageClass string `json:"storage_class"`
}

type GoogleStorageBucketSpecLifecycleRuleCondition struct {
	CreatedBefore       string   `json:"created_before"`
	IsLive              bool     `json:"is_live"`
	MatchesStorageClass []string `json:"matches_storage_class"`
	NumNewerVersions    int      `json:"num_newer_versions"`
	Age                 int      `json:"age"`
}

type GoogleStorageBucketSpecLifecycleRule struct {
	Action    []GoogleStorageBucketSpecLifecycleRule `json:"action"`
	Condition []GoogleStorageBucketSpecLifecycleRule `json:"condition"`
}

type GoogleStorageBucketSpecWebsite struct {
	NotFoundPage   string `json:"not_found_page"`
	MainPageSuffix string `json:"main_page_suffix"`
}

type GoogleStorageBucketSpecEncryption struct {
	DefaultKmsKeyName string `json:"default_kms_key_name"`
}

type GoogleStorageBucketSpecLogging struct {
	LogBucket       string `json:"log_bucket"`
	LogObjectPrefix string `json:"log_object_prefix"`
}

type GoogleStorageBucketSpecVersioning struct {
	Enabled bool `json:"enabled"`
}

type GoogleStorageBucketSpecCors struct {
	Method         []string `json:"method"`
	ResponseHeader []string `json:"response_header"`
	MaxAgeSeconds  int      `json:"max_age_seconds"`
	Origin         []string `json:"origin"`
}

type GoogleStorageBucketSpec struct {
	Name          string                    `json:"name"`
	Location      string                    `json:"location"`
	Project       string                    `json:"project"`
	LifecycleRule []GoogleStorageBucketSpec `json:"lifecycle_rule"`
	Website       []GoogleStorageBucketSpec `json:"website"`
	Encryption    []GoogleStorageBucketSpec `json:"encryption"`
	ForceDestroy  bool                      `json:"force_destroy"`
	Labels        map[string]string         `json:"labels"`
	PredefinedAcl string                    `json:"predefined_acl"`
	Url           string                    `json:"url"`
	StorageClass  string                    `json:"storage_class"`
	Logging       []GoogleStorageBucketSpec `json:"logging"`
	SelfLink      string                    `json:"self_link"`
	Versioning    []GoogleStorageBucketSpec `json:"versioning"`
	Cors          []GoogleStorageBucketSpec `json:"cors"`
}

type GoogleStorageBucketStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageBucketList is a list of GoogleStorageBuckets
type GoogleStorageBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageBucket CRD objects
	Items []GoogleStorageBucket `json:"items,omitempty"`
}
