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

type S3Bucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketSpec   `json:"spec,omitempty"`
	Status            S3BucketStatus `json:"status,omitempty"`
}

type S3BucketSpecCorsRule struct {
	// +optional
	AllowedHeaders []string `json:"allowed_headers,omitempty"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedOrigins []string `json:"allowed_origins"`
	// +optional
	ExposeHeaders []string `json:"expose_headers,omitempty"`
	// +optional
	MaxAgeSeconds int `json:"max_age_seconds,omitempty"`
}

type S3BucketSpecLifecycleRuleExpiration struct {
	// +optional
	Date string `json:"date,omitempty"`
	// +optional
	Days int `json:"days,omitempty"`
	// +optional
	ExpiredObjectDeleteMarker bool `json:"expired_object_delete_marker,omitempty"`
}

type S3BucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	// +optional
	Days int `json:"days,omitempty"`
}

type S3BucketSpecLifecycleRuleNoncurrentVersionTransition struct {
	// +optional
	Days         int    `json:"days,omitempty"`
	StorageClass string `json:"storage_class"`
}

type S3BucketSpecLifecycleRuleTransition struct {
	// +optional
	Date string `json:"date,omitempty"`
	// +optional
	Days         int    `json:"days,omitempty"`
	StorageClass string `json:"storage_class"`
}

type S3BucketSpecLifecycleRule struct {
	// +optional
	AbortIncompleteMultipartUploadDays int  `json:"abort_incomplete_multipart_upload_days,omitempty"`
	Enabled                            bool `json:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Expiration *[]S3BucketSpecLifecycleRule `json:"expiration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	NoncurrentVersionExpiration *[]S3BucketSpecLifecycleRule `json:"noncurrent_version_expiration,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NoncurrentVersionTransition *[]S3BucketSpecLifecycleRule `json:"noncurrent_version_transition,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Transition *[]S3BucketSpecLifecycleRule `json:"transition,omitempty"`
}

type S3BucketSpecLogging struct {
	TargetBucket string `json:"target_bucket"`
	// +optional
	TargetPrefix string `json:"target_prefix,omitempty"`
}

type S3BucketSpecObjectLockConfigurationRuleDefaultRetention struct {
	// +optional
	Days int    `json:"days,omitempty"`
	Mode string `json:"mode"`
	// +optional
	Years int `json:"years,omitempty"`
}

type S3BucketSpecObjectLockConfigurationRule struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	DefaultRetention []S3BucketSpecObjectLockConfigurationRule `json:"default_retention"`
}

type S3BucketSpecObjectLockConfiguration struct {
	ObjectLockEnabled string `json:"object_lock_enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Rule *[]S3BucketSpecObjectLockConfiguration `json:"rule,omitempty"`
}

type S3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation struct {
	Owner string `json:"owner"`
}

type S3BucketSpecReplicationConfigurationRulesDestination struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	AccessControlTranslation *[]S3BucketSpecReplicationConfigurationRulesDestination `json:"access_control_translation,omitempty"`
	// +optional
	AccountId string `json:"account_id,omitempty"`
	Bucket    string `json:"bucket"`
	// +optional
	ReplicaKmsKeyId string `json:"replica_kms_key_id,omitempty"`
	// +optional
	StorageClass string `json:"storage_class,omitempty"`
}

type S3BucketSpecReplicationConfigurationRulesFilter struct {
	// +optional
	Prefix string `json:"prefix,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	Enabled bool `json:"enabled"`
}

type S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SseKmsEncryptedObjects *[]S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria `json:"sse_kms_encrypted_objects,omitempty"`
}

type S3BucketSpecReplicationConfigurationRules struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Destination []S3BucketSpecReplicationConfigurationRules `json:"destination"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Filter *[]S3BucketSpecReplicationConfigurationRules `json:"filter,omitempty"`
	// +optional
	Id string `json:"id,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty"`
	// +optional
	Priority int `json:"priority,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SourceSelectionCriteria *[]S3BucketSpecReplicationConfigurationRules `json:"source_selection_criteria,omitempty"`
	Status                  string                                       `json:"status"`
}

type S3BucketSpecReplicationConfiguration struct {
	Role string `json:"role"`
	// +kubebuilder:validation:UniqueItems=true
	Rules []S3BucketSpecReplicationConfiguration `json:"rules"`
}

type S3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	// +optional
	KmsMasterKeyId string `json:"kms_master_key_id,omitempty"`
	SseAlgorithm   string `json:"sse_algorithm"`
}

type S3BucketSpecServerSideEncryptionConfigurationRule struct {
	// +kubebuilder:validation:MaxItems=1
	ApplyServerSideEncryptionByDefault []S3BucketSpecServerSideEncryptionConfigurationRule `json:"apply_server_side_encryption_by_default"`
}

type S3BucketSpecServerSideEncryptionConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	Rule []S3BucketSpecServerSideEncryptionConfiguration `json:"rule"`
}

type S3BucketSpecWebsite struct {
	// +optional
	ErrorDocument string `json:"error_document,omitempty"`
	// +optional
	IndexDocument string `json:"index_document,omitempty"`
	// +optional
	RedirectAllRequestsTo string `json:"redirect_all_requests_to,omitempty"`
	// +optional
	RoutingRules string `json:"routing_rules,omitempty"`
}

type S3BucketSpec struct {
	// +optional
	Acl string `json:"acl,omitempty"`
	// +optional
	BucketPrefix string `json:"bucket_prefix,omitempty"`
	// +optional
	CorsRule *[]S3BucketSpec `json:"cors_rule,omitempty"`
	// +optional
	ForceDestroy bool `json:"force_destroy,omitempty"`
	// +optional
	LifecycleRule *[]S3BucketSpec `json:"lifecycle_rule,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Logging *[]S3BucketSpec `json:"logging,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ObjectLockConfiguration *[]S3BucketSpec `json:"object_lock_configuration,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReplicationConfiguration *[]S3BucketSpec `json:"replication_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServerSideEncryptionConfiguration *[]S3BucketSpec `json:"server_side_encryption_configuration,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Website *[]S3BucketSpec `json:"website,omitempty"`
}

type S3BucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// S3BucketList is a list of S3Buckets
type S3BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of S3Bucket CRD objects
	Items []S3Bucket `json:"items,omitempty"`
}
