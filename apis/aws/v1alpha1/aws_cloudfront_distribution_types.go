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

type AwsCloudfrontDistribution struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudfrontDistributionSpec   `json:"spec,omitempty"`
	Status            AwsCloudfrontDistributionStatus `json:"status,omitempty"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation struct {
	EventType   string `json:"event_type"`
	LambdaArn   string `json:"lambda_arn"`
	IncludeBody bool   `json:"include_body"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues struct {
	Cookies              []AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues `json:"cookies"`
	Headers              []string                                                           `json:"headers"`
	QueryString          bool                                                               `json:"query_string"`
	QueryStringCacheKeys []string                                                           `json:"query_string_cache_keys"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehavior struct {
	MaxTtl                    int                                                 `json:"max_ttl"`
	MinTtl                    int                                                 `json:"min_ttl"`
	TargetOriginId            string                                              `json:"target_origin_id"`
	Compress                  bool                                                `json:"compress"`
	FieldLevelEncryptionId    string                                              `json:"field_level_encryption_id"`
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecDefaultCacheBehavior `json:"lambda_function_association"`
	ForwardedValues           []AwsCloudfrontDistributionSpecDefaultCacheBehavior `json:"forwarded_values"`
	SmoothStreaming           bool                                                `json:"smooth_streaming"`
	TrustedSigners            []string                                            `json:"trusted_signers"`
	ViewerProtocolPolicy      string                                              `json:"viewer_protocol_policy"`
	AllowedMethods            []string                                            `json:"allowed_methods"`
	CachedMethods             []string                                            `json:"cached_methods"`
	DefaultTtl                int                                                 `json:"default_ttl"`
}

type AwsCloudfrontDistributionSpecCustomErrorResponse struct {
	ErrorCachingMinTtl int    `json:"error_caching_min_ttl"`
	ErrorCode          int    `json:"error_code"`
	ResponseCode       int    `json:"response_code"`
	ResponsePagePath   string `json:"response_page_path"`
}

type AwsCloudfrontDistributionSpecLoggingConfig struct {
	IncludeCookies bool   `json:"include_cookies"`
	Prefix         string `json:"prefix"`
	Bucket         string `json:"bucket"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation struct {
	EventType   string `json:"event_type"`
	LambdaArn   string `json:"lambda_arn"`
	IncludeBody bool   `json:"include_body"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues struct {
	Cookies              []AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues `json:"cookies"`
	Headers              []string                                                           `json:"headers"`
	QueryString          bool                                                               `json:"query_string"`
	QueryStringCacheKeys []string                                                           `json:"query_string_cache_keys"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehavior struct {
	CachedMethods             []string                                            `json:"cached_methods"`
	FieldLevelEncryptionId    string                                              `json:"field_level_encryption_id"`
	Compress                  bool                                                `json:"compress"`
	MaxTtl                    int                                                 `json:"max_ttl"`
	MinTtl                    int                                                 `json:"min_ttl"`
	PathPattern               string                                              `json:"path_pattern"`
	TargetOriginId            string                                              `json:"target_origin_id"`
	AllowedMethods            []string                                            `json:"allowed_methods"`
	DefaultTtl                int                                                 `json:"default_ttl"`
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecOrderedCacheBehavior `json:"lambda_function_association"`
	ForwardedValues           []AwsCloudfrontDistributionSpecOrderedCacheBehavior `json:"forwarded_values"`
	SmoothStreaming           bool                                                `json:"smooth_streaming"`
	TrustedSigners            []string                                            `json:"trusted_signers"`
	ViewerProtocolPolicy      string                                              `json:"viewer_protocol_policy"`
}

type AwsCloudfrontDistributionSpecRestrictionsGeoRestriction struct {
	Locations       []string `json:"locations"`
	RestrictionType string   `json:"restriction_type"`
}

type AwsCloudfrontDistributionSpecRestrictions struct {
	GeoRestriction []AwsCloudfrontDistributionSpecRestrictions `json:"geo_restriction"`
}

type AwsCloudfrontDistributionSpecOriginGroupFailoverCriteria struct {
	StatusCodes []int64 `json:"status_codes"`
}

type AwsCloudfrontDistributionSpecOriginGroupMember struct {
	OriginId string `json:"origin_id"`
}

type AwsCloudfrontDistributionSpecOriginGroup struct {
	FailoverCriteria []AwsCloudfrontDistributionSpecOriginGroup `json:"failover_criteria"`
	Member           []AwsCloudfrontDistributionSpecOriginGroup `json:"member"`
	OriginId         string                                     `json:"origin_id"`
}

type AwsCloudfrontDistributionSpecOriginS3OriginConfig struct {
	OriginAccessIdentity string `json:"origin_access_identity"`
}

type AwsCloudfrontDistributionSpecOriginCustomOriginConfig struct {
	HttpsPort              int      `json:"https_port"`
	OriginKeepaliveTimeout int      `json:"origin_keepalive_timeout"`
	OriginReadTimeout      int      `json:"origin_read_timeout"`
	OriginProtocolPolicy   string   `json:"origin_protocol_policy"`
	OriginSslProtocols     []string `json:"origin_ssl_protocols"`
	HttpPort               int      `json:"http_port"`
}

type AwsCloudfrontDistributionSpecOriginCustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsCloudfrontDistributionSpecOrigin struct {
	S3OriginConfig     []AwsCloudfrontDistributionSpecOrigin `json:"s3_origin_config"`
	CustomOriginConfig []AwsCloudfrontDistributionSpecOrigin `json:"custom_origin_config"`
	DomainName         string                                `json:"domain_name"`
	CustomHeader       []AwsCloudfrontDistributionSpecOrigin `json:"custom_header"`
	OriginId           string                                `json:"origin_id"`
	OriginPath         string                                `json:"origin_path"`
}

type AwsCloudfrontDistributionSpecViewerCertificate struct {
	AcmCertificateArn            string `json:"acm_certificate_arn"`
	CloudfrontDefaultCertificate bool   `json:"cloudfront_default_certificate"`
	IamCertificateId             string `json:"iam_certificate_id"`
	MinimumProtocolVersion       string `json:"minimum_protocol_version"`
	SslSupportMethod             string `json:"ssl_support_method"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorForwardedValues struct {
	Cookies              []AwsCloudfrontDistributionSpecCacheBehaviorForwardedValues `json:"cookies"`
	Headers              []string                                                    `json:"headers"`
	QueryString          bool                                                        `json:"query_string"`
	QueryStringCacheKeys []string                                                    `json:"query_string_cache_keys"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorLambdaFunctionAssociation struct {
	EventType   string `json:"event_type"`
	LambdaArn   string `json:"lambda_arn"`
	IncludeBody bool   `json:"include_body"`
}

type AwsCloudfrontDistributionSpecCacheBehavior struct {
	TrustedSigners            []string                                     `json:"trusted_signers"`
	CachedMethods             []string                                     `json:"cached_methods"`
	FieldLevelEncryptionId    string                                       `json:"field_level_encryption_id"`
	DefaultTtl                int                                          `json:"default_ttl"`
	ForwardedValues           []AwsCloudfrontDistributionSpecCacheBehavior `json:"forwarded_values"`
	MaxTtl                    int                                          `json:"max_ttl"`
	TargetOriginId            string                                       `json:"target_origin_id"`
	AllowedMethods            []string                                     `json:"allowed_methods"`
	Compress                  bool                                         `json:"compress"`
	PathPattern               string                                       `json:"path_pattern"`
	SmoothStreaming           bool                                         `json:"smooth_streaming"`
	ViewerProtocolPolicy      string                                       `json:"viewer_protocol_policy"`
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecCacheBehavior `json:"lambda_function_association"`
	MinTtl                    int                                          `json:"min_ttl"`
}

type AwsCloudfrontDistributionSpec struct {
	Arn                         string                          `json:"arn"`
	DefaultCacheBehavior        []AwsCloudfrontDistributionSpec `json:"default_cache_behavior"`
	Enabled                     bool                            `json:"enabled"`
	HttpVersion                 string                          `json:"http_version"`
	PriceClass                  string                          `json:"price_class"`
	CallerReference             string                          `json:"caller_reference"`
	HostedZoneId                string                          `json:"hosted_zone_id"`
	Tags                        map[string]string               `json:"tags"`
	CustomErrorResponse         []AwsCloudfrontDistributionSpec `json:"custom_error_response"`
	LoggingConfig               []AwsCloudfrontDistributionSpec `json:"logging_config"`
	Status                      string                          `json:"status"`
	ActiveTrustedSigners        map[string]string               `json:"active_trusted_signers"`
	Aliases                     []string                        `json:"aliases"`
	OrderedCacheBehavior        []AwsCloudfrontDistributionSpec `json:"ordered_cache_behavior"`
	Etag                        string                          `json:"etag"`
	RetainOnDelete              bool                            `json:"retain_on_delete"`
	IsIpv6Enabled               bool                            `json:"is_ipv6_enabled"`
	Restrictions                []AwsCloudfrontDistributionSpec `json:"restrictions"`
	DomainName                  string                          `json:"domain_name"`
	InProgressValidationBatches int                             `json:"in_progress_validation_batches"`
	OriginGroup                 []AwsCloudfrontDistributionSpec `json:"origin_group"`
	Origin                      []AwsCloudfrontDistributionSpec `json:"origin"`
	ViewerCertificate           []AwsCloudfrontDistributionSpec `json:"viewer_certificate"`
	WaitForDeployment           bool                            `json:"wait_for_deployment"`
	LastModifiedTime            string                          `json:"last_modified_time"`
	CacheBehavior               []AwsCloudfrontDistributionSpec `json:"cache_behavior"`
	Comment                     string                          `json:"comment"`
	DefaultRootObject           string                          `json:"default_root_object"`
	WebAclId                    string                          `json:"web_acl_id"`
}



type AwsCloudfrontDistributionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudfrontDistributionList is a list of AwsCloudfrontDistributions
type AwsCloudfrontDistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudfrontDistribution CRD objects
	Items []AwsCloudfrontDistribution `json:"items,omitempty"`
}