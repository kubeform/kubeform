package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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

type AwsS3BucketSpecLifecycleRuleNoncurrentVersionExpiration struct {
	Days int `json:"days"`
}

type AwsS3BucketSpecLifecycleRuleTransition struct {
	Date         string `json:"date"`
	Days         int    `json:"days"`
	StorageClass string `json:"storage_class"`
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
	Id                                 string                         `json:"id"`
	NoncurrentVersionExpiration        []AwsS3BucketSpecLifecycleRule `json:"noncurrent_version_expiration"`
	Transition                         []AwsS3BucketSpecLifecycleRule `json:"transition"`
	NoncurrentVersionTransition        []AwsS3BucketSpecLifecycleRule `json:"noncurrent_version_transition"`
	Prefix                             string                         `json:"prefix"`
	Tags                               map[string]string              `json:"tags"`
	Enabled                            bool                           `json:"enabled"`
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
	AccountId                string                                                    `json:"account_id"`
	Bucket                   string                                                    `json:"bucket"`
	StorageClass             string                                                    `json:"storage_class"`
	ReplicaKmsKeyId          string                                                    `json:"replica_kms_key_id"`
	AccessControlTranslation []AwsS3BucketSpecReplicationConfigurationRulesDestination `json:"access_control_translation"`
}

type AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	Enabled bool `json:"enabled"`
}

type AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria struct {
	SseKmsEncryptedObjects []AwsS3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria `json:"sse_kms_encrypted_objects"`
}

type AwsS3BucketSpecReplicationConfigurationRules struct {
	Status                  string                                         `json:"status"`
	Priority                int                                            `json:"priority"`
	Filter                  []AwsS3BucketSpecReplicationConfigurationRules `json:"filter"`
	Id                      string                                         `json:"id"`
	Destination             []AwsS3BucketSpecReplicationConfigurationRules `json:"destination"`
	SourceSelectionCriteria []AwsS3BucketSpecReplicationConfigurationRules `json:"source_selection_criteria"`
	Prefix                  string                                         `json:"prefix"`
}

type AwsS3BucketSpecReplicationConfiguration struct {
	Role  string                                    `json:"role"`
	Rules []AwsS3BucketSpecReplicationConfiguration `json:"rules"`
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

type AwsS3BucketSpecVersioning struct {
	Enabled   bool `json:"enabled"`
	MfaDelete bool `json:"mfa_delete"`
}

type AwsS3BucketSpecLogging struct {
	TargetBucket string `json:"target_bucket"`
	TargetPrefix string `json:"target_prefix"`
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
	Rule              []AwsS3BucketSpecObjectLockConfiguration `json:"rule"`
	ObjectLockEnabled string                                   `json:"object_lock_enabled"`
}

type AwsS3BucketSpecWebsite struct {
	IndexDocument         string `json:"index_document"`
	ErrorDocument         string `json:"error_document"`
	RedirectAllRequestsTo string `json:"redirect_all_requests_to"`
	RoutingRules          string `json:"routing_rules"`
}

type AwsS3BucketSpec struct {
	CorsRule                          []AwsS3BucketSpec `json:"cors_rule"`
	LifecycleRule                     []AwsS3BucketSpec `json:"lifecycle_rule"`
	ReplicationConfiguration          []AwsS3BucketSpec `json:"replication_configuration"`
	ServerSideEncryptionConfiguration []AwsS3BucketSpec `json:"server_side_encryption_configuration"`
	BucketDomainName                  string            `json:"bucket_domain_name"`
	BucketRegionalDomainName          string            `json:"bucket_regional_domain_name"`
	Acl                               string            `json:"acl"`
	HostedZoneId                      string            `json:"hosted_zone_id"`
	RequestPayer                      string            `json:"request_payer"`
	Tags                              map[string]string `json:"tags"`
	Bucket                            string            `json:"bucket"`
	BucketPrefix                      string            `json:"bucket_prefix"`
	WebsiteEndpoint                   string            `json:"website_endpoint"`
	Versioning                        []AwsS3BucketSpec `json:"versioning"`
	Logging                           []AwsS3BucketSpec `json:"logging"`
	ForceDestroy                      bool              `json:"force_destroy"`
	AccelerationStatus                string            `json:"acceleration_status"`
	Arn                               string            `json:"arn"`
	Region                            string            `json:"region"`
	WebsiteDomain                     string            `json:"website_domain"`
	ObjectLockConfiguration           []AwsS3BucketSpec `json:"object_lock_configuration"`
	Policy                            string            `json:"policy"`
	Website                           []AwsS3BucketSpec `json:"website"`
}



type AwsS3BucketStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketList is a list of AwsS3Buckets
type AwsS3BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3Bucket CRD objects
	Items []AwsS3Bucket `json:"items,omitempty"`
}