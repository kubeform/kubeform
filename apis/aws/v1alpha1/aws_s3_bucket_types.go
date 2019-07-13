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

type AwsS3Bucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketSpec   `json:"spec,omitempty"`
	Status            AwsS3BucketStatus `json:"status,omitempty"`
}

type AwsS3BucketSpecCorsRule struct {
	MaxAgeSeconds  int      `json:"max_age_seconds"`
	AllowedHeaders []string `json:"allowed_headers"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedOrigins []string `json:"allowed_origins"`
	ExposeHeaders  []string `json:"expose_headers"`
}

type AwsS3BucketSpecWebsite struct {
	IndexDocument         string `json:"index_document"`
	ErrorDocument         string `json:"error_document"`
	RedirectAllRequestsTo string `json:"redirect_all_requests_to"`
	RoutingRules          string `json:"routing_rules"`
}

type AwsS3BucketSpecVersioning struct {
	Enabled   bool `json:"enabled"`
	MfaDelete bool `json:"mfa_delete"`
}

type AwsS3BucketSpecLogging struct {
	TargetBucket string `json:"target_bucket"`
	TargetPrefix string `json:"target_prefix"`
}

type AwsS3BucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	Days int `json:"days"`
}

type AwsS3BucketSpecLifecycleRuleTransition struct {
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
	Date         string `json:"date"`
}

type AwsS3BucketSpecLifecycleRuleNoncurrentVersionTransition struct {
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
}

type AwsS3BucketSpecLifecycleRuleExpiration struct {
	Date                      string `json:"date"`
	Days                      int    `json:"days"`
	ExpiredObjectDeleteMarker bool   `json:"expired_object_delete_marker"`
}

type AwsS3BucketSpecLifecycleRule struct {
	NoncurrentVersionExpiration        []AwsS3BucketSpecLifecycleRule `json:"noncurrent_version_expiration"`
	Transition                         []AwsS3BucketSpecLifecycleRule `json:"transition"`
	NoncurrentVersionTransition        []AwsS3BucketSpecLifecycleRule `json:"noncurrent_version_transition"`
	Prefix                             string                         `json:"prefix"`
	Tags                               map[string]string              `json:"tags"`
	Enabled                            bool                           `json:"enabled"`
	Id                                 string                         `json:"id"`
	AbortIncompleteMultipartUploadDays int                            `json:"abort_incomplete_multipart_upload_days"`
	Expiration                         []AwsS3BucketSpecLifecycleRule `json:"expiration"`
}

type AwsS3BucketSpecReplicationConfigurationRulesFilter struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}

type AwsS3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation struct {
	Owner string `json:"owner"`
}

type AwsS3BucketSpecReplicationConfigurationRulesDestination struct {
	AccessControlTranslation []AwsS3BucketSpecReplicationConfigurationRulesDestination `json:"access_control_translation"`
	AccountId                string                                                    `json:"account_id"`
	Bucket                   string                                                    `json:"bucket"`
	StorageClass             string                                                    `json:"storage_class"`
	ReplicaKmsKeyId          string                                                    `json:"replica_kms_key_id"`
}

type AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	Enabled bool `json:"enabled"`
}

type AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria struct {
	SseKmsEncryptedObjects []AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria `json:"sse_kms_encrypted_objects"`
}

type AwsS3BucketSpecReplicationConfigurationRules struct {
	Prefix                  string                                         `json:"prefix"`
	Status                  string                                         `json:"status"`
	Priority                int                                            `json:"priority"`
	Filter                  []AwsS3BucketSpecReplicationConfigurationRules `json:"filter"`
	Id                      string                                         `json:"id"`
	Destination             []AwsS3BucketSpecReplicationConfigurationRules `json:"destination"`
	SourceSelectionCriteria []AwsS3BucketSpecReplicationConfigurationRules `json:"source_selection_criteria"`
}

type AwsS3BucketSpecReplicationConfiguration struct {
	Role  string                                    `json:"role"`
	Rules []AwsS3BucketSpecReplicationConfiguration `json:"rules"`
}

type AwsS3BucketSpecObjectLockConfigurationRuleDefaultRetention struct {
	Mode  string `json:"mode"`
	Days  int    `json:"days"`
	Years int    `json:"years"`
}

type AwsS3BucketSpecObjectLockConfigurationRule struct {
	DefaultRetention []AwsS3BucketSpecObjectLockConfigurationRule `json:"default_retention"`
}

type AwsS3BucketSpecObjectLockConfiguration struct {
	ObjectLockEnabled string                                   `json:"object_lock_enabled"`
	Rule              []AwsS3BucketSpecObjectLockConfiguration `json:"rule"`
}

type AwsS3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault struct {
	KmsMasterKeyId string `json:"kms_master_key_id"`
	SseAlgorithm   string `json:"sse_algorithm"`
}

type AwsS3BucketSpecServerSideEncryptionConfigurationRule struct {
	ApplyServerSideEncryptionByDefault []AwsS3BucketSpecServerSideEncryptionConfigurationRule `json:"apply_server_side_encryption_by_default"`
}

type AwsS3BucketSpecServerSideEncryptionConfiguration struct {
	Rule []AwsS3BucketSpecServerSideEncryptionConfiguration `json:"rule"`
}

type AwsS3BucketSpec struct {
	CorsRule                          []AwsS3BucketSpec `json:"cors_rule"`
	Website                           []AwsS3BucketSpec `json:"website"`
	BucketDomainName                  string            `json:"bucket_domain_name"`
	BucketRegionalDomainName          string            `json:"bucket_regional_domain_name"`
	HostedZoneId                      string            `json:"hosted_zone_id"`
	WebsiteEndpoint                   string            `json:"website_endpoint"`
	WebsiteDomain                     string            `json:"website_domain"`
	Versioning                        []AwsS3BucketSpec `json:"versioning"`
	Logging                           []AwsS3BucketSpec `json:"logging"`
	ForceDestroy                      bool              `json:"force_destroy"`
	Bucket                            string            `json:"bucket"`
	BucketPrefix                      string            `json:"bucket_prefix"`
	Acl                               string            `json:"acl"`
	LifecycleRule                     []AwsS3BucketSpec `json:"lifecycle_rule"`
	AccelerationStatus                string            `json:"acceleration_status"`
	RequestPayer                      string            `json:"request_payer"`
	ReplicationConfiguration          []AwsS3BucketSpec `json:"replication_configuration"`
	ObjectLockConfiguration           []AwsS3BucketSpec `json:"object_lock_configuration"`
	Tags                              map[string]string `json:"tags"`
	Arn                               string            `json:"arn"`
	Policy                            string            `json:"policy"`
	Region                            string            `json:"region"`
	ServerSideEncryptionConfiguration []AwsS3BucketSpec `json:"server_side_encryption_configuration"`
}



type AwsS3BucketStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsS3BucketList is a list of AwsS3Buckets
type AwsS3BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3Bucket CRD objects
	Items []AwsS3Bucket `json:"items,omitempty"`
}