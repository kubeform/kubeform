package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	AllowedHeaders []string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	ExposeHeaders []string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`
	// +optional
	MaxAgeSeconds int `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type S3BucketSpecLifecycleRuleExpiration struct {
	// +optional
	Date string `json:"date,omitempty" tf:"date,omitempty"`
	// +optional
	Days int `json:"days,omitempty" tf:"days,omitempty"`
	// +optional
	ExpiredObjectDeleteMarker bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type S3BucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	// +optional
	Days int `json:"days,omitempty" tf:"days,omitempty"`
}

type S3BucketSpecLifecycleRuleNoncurrentVersionTransition struct {
	// +optional
	Days         int    `json:"days,omitempty" tf:"days,omitempty"`
	StorageClass string `json:"storageClass" tf:"storage_class"`
}

type S3BucketSpecLifecycleRuleTransition struct {
	// +optional
	Date string `json:"date,omitempty" tf:"date,omitempty"`
	// +optional
	Days         int    `json:"days,omitempty" tf:"days,omitempty"`
	StorageClass string `json:"storageClass" tf:"storage_class"`
}

type S3BucketSpecLifecycleRule struct {
	// +optional
	AbortIncompleteMultipartUploadDays int  `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`
	Enabled                            bool `json:"enabled" tf:"enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Expiration []S3BucketSpecLifecycleRuleExpiration `json:"expiration,omitempty" tf:"expiration,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	NoncurrentVersionExpiration []S3BucketSpecLifecycleRuleNoncurrentVersionExpiration `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NoncurrentVersionTransition []S3BucketSpecLifecycleRuleNoncurrentVersionTransition `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Transition []S3BucketSpecLifecycleRuleTransition `json:"transition,omitempty" tf:"transition,omitempty"`
}

type S3BucketSpecLogging struct {
	TargetBucket string `json:"targetBucket" tf:"target_bucket"`
	// +optional
	TargetPrefix string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type S3BucketSpecObjectLockConfigurationRuleDefaultRetention struct {
	// +optional
	Days int    `json:"days,omitempty" tf:"days,omitempty"`
	Mode string `json:"mode" tf:"mode"`
	// +optional
	Years int `json:"years,omitempty" tf:"years,omitempty"`
}

type S3BucketSpecObjectLockConfigurationRule struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	DefaultRetention []S3BucketSpecObjectLockConfigurationRuleDefaultRetention `json:"defaultRetention" tf:"default_retention"`
}

type S3BucketSpecObjectLockConfiguration struct {
	ObjectLockEnabled string `json:"objectLockEnabled" tf:"object_lock_enabled"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Rule []S3BucketSpecObjectLockConfigurationRule `json:"rule,omitempty" tf:"rule,omitempty"`
}

type S3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation struct {
	Owner string `json:"owner" tf:"owner"`
}

type S3BucketSpecReplicationConfigurationRulesDestination struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	AccessControlTranslation []S3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation `json:"accessControlTranslation,omitempty" tf:"access_control_translation,omitempty"`
	// +optional
	AccountID string `json:"accountID,omitempty" tf:"account_id,omitempty"`
	Bucket    string `json:"bucket" tf:"bucket"`
	// +optional
	ReplicaKmsKeyID string `json:"replicaKmsKeyID,omitempty" tf:"replica_kms_key_id,omitempty"`
	// +optional
	StorageClass string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`
}

type S3BucketSpecReplicationConfigurationRulesFilter struct {
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	Enabled bool `json:"enabled" tf:"enabled"`
}

type S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SseKmsEncryptedObjects []S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects `json:"sseKmsEncryptedObjects,omitempty" tf:"sse_kms_encrypted_objects,omitempty"`
}

type S3BucketSpecReplicationConfigurationRules struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Destination []S3BucketSpecReplicationConfigurationRulesDestination `json:"destination" tf:"destination"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Filter []S3BucketSpecReplicationConfigurationRulesFilter `json:"filter,omitempty" tf:"filter,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	// +optional
	Priority int `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SourceSelectionCriteria []S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria `json:"sourceSelectionCriteria,omitempty" tf:"source_selection_criteria,omitempty"`
	Status                  string                                                             `json:"status" tf:"status"`
}

type S3BucketSpecReplicationConfiguration struct {
	Role string `json:"role" tf:"role"`
	// +kubebuilder:validation:UniqueItems=true
	Rules []S3BucketSpecReplicationConfigurationRules `json:"rules" tf:"rules"`
}

type S3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	// +optional
	KmsMasterKeyID string `json:"kmsMasterKeyID,omitempty" tf:"kms_master_key_id,omitempty"`
	SseAlgorithm   string `json:"sseAlgorithm" tf:"sse_algorithm"`
}

type S3BucketSpecServerSideEncryptionConfigurationRule struct {
	// +kubebuilder:validation:MaxItems=1
	ApplyServerSideEncryptionByDefault []S3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault `json:"applyServerSideEncryptionByDefault" tf:"apply_server_side_encryption_by_default"`
}

type S3BucketSpecServerSideEncryptionConfiguration struct {
	// +kubebuilder:validation:MaxItems=1
	Rule []S3BucketSpecServerSideEncryptionConfigurationRule `json:"rule" tf:"rule"`
}

type S3BucketSpecVersioning struct {
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	MfaDelete bool `json:"mfaDelete,omitempty" tf:"mfa_delete,omitempty"`
}

type S3BucketSpecWebsite struct {
	// +optional
	ErrorDocument string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`
	// +optional
	IndexDocument string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`
	// +optional
	RedirectAllRequestsTo string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`
	// +optional
	RoutingRules string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type S3BucketSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	AccelerationStatus string `json:"accelerationStatus,omitempty" tf:"acceleration_status,omitempty"`
	// +optional
	Acl string `json:"acl,omitempty" tf:"acl,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Bucket string `json:"bucket,omitempty" tf:"bucket,omitempty"`
	// +optional
	BucketPrefix string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`
	// +optional
	CorsRule []S3BucketSpecCorsRule `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`
	// +optional
	ForceDestroy bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`
	// +optional
	HostedZoneID string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id,omitempty"`
	// +optional
	LifecycleRule []S3BucketSpecLifecycleRule `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Logging []S3BucketSpecLogging `json:"logging,omitempty" tf:"logging,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ObjectLockConfiguration []S3BucketSpecObjectLockConfiguration `json:"objectLockConfiguration,omitempty" tf:"object_lock_configuration,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ReplicationConfiguration []S3BucketSpecReplicationConfiguration `json:"replicationConfiguration,omitempty" tf:"replication_configuration,omitempty"`
	// +optional
	RequestPayer string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServerSideEncryptionConfiguration []S3BucketSpecServerSideEncryptionConfiguration `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Versioning []S3BucketSpecVersioning `json:"versioning,omitempty" tf:"versioning,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Website []S3BucketSpecWebsite `json:"website,omitempty" tf:"website,omitempty"`
	// +optional
	WebsiteDomain string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`
	// +optional
	WebsiteEndpoint string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type S3BucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
