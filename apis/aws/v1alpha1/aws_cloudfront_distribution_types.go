package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudfrontDistribution struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudfrontDistributionSpec   `json:"spec,omitempty"`
	Status            AwsCloudfrontDistributionStatus `json:"status,omitempty"`
}

type AwsCloudfrontDistributionSpecCacheBehaviorForwardedValuesCookies struct {
	WhitelistedNames []string `json:"whitelisted_names"`
	Forward          string   `json:"forward"`
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
	ForwardedValues           []AwsCloudfrontDistributionSpecCacheBehavior `json:"forwarded_values"`
	SmoothStreaming           bool                                         `json:"smooth_streaming"`
	ViewerProtocolPolicy      string                                       `json:"viewer_protocol_policy"`
	FieldLevelEncryptionId    string                                       `json:"field_level_encryption_id"`
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecCacheBehavior `json:"lambda_function_association"`
	Compress                  bool                                         `json:"compress"`
	TargetOriginId            string                                       `json:"target_origin_id"`
	TrustedSigners            []string                                     `json:"trusted_signers"`
	PathPattern               string                                       `json:"path_pattern"`
	AllowedMethods            []string                                     `json:"allowed_methods"`
	CachedMethods             []string                                     `json:"cached_methods"`
	DefaultTtl                int                                          `json:"default_ttl"`
	MaxTtl                    int                                          `json:"max_ttl"`
	MinTtl                    int                                          `json:"min_ttl"`
}

type AwsCloudfrontDistributionSpecCustomErrorResponse struct {
	ErrorCode          int    `json:"error_code"`
	ResponseCode       int    `json:"response_code"`
	ResponsePagePath   string `json:"response_page_path"`
	ErrorCachingMinTtl int    `json:"error_caching_min_ttl"`
}

type AwsCloudfrontDistributionSpecOriginCustomOriginConfig struct {
	OriginKeepaliveTimeout int      `json:"origin_keepalive_timeout"`
	OriginReadTimeout      int      `json:"origin_read_timeout"`
	OriginProtocolPolicy   string   `json:"origin_protocol_policy"`
	OriginSslProtocols     []string `json:"origin_ssl_protocols"`
	HttpPort               int      `json:"http_port"`
	HttpsPort              int      `json:"https_port"`
}

type AwsCloudfrontDistributionSpecOriginCustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsCloudfrontDistributionSpecOriginS3OriginConfig struct {
	OriginAccessIdentity string `json:"origin_access_identity"`
}

type AwsCloudfrontDistributionSpecOrigin struct {
	CustomOriginConfig []AwsCloudfrontDistributionSpecOrigin `json:"custom_origin_config"`
	DomainName         string                                `json:"domain_name"`
	CustomHeader       []AwsCloudfrontDistributionSpecOrigin `json:"custom_header"`
	OriginId           string                                `json:"origin_id"`
	OriginPath         string                                `json:"origin_path"`
	S3OriginConfig     []AwsCloudfrontDistributionSpecOrigin `json:"s3_origin_config"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation struct {
	LambdaArn   string `json:"lambda_arn"`
	IncludeBody bool   `json:"include_body"`
	EventType   string `json:"event_type"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues struct {
	Headers              []string                                                           `json:"headers"`
	QueryString          bool                                                               `json:"query_string"`
	QueryStringCacheKeys []string                                                           `json:"query_string_cache_keys"`
	Cookies              []AwsCloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues `json:"cookies"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehavior struct {
	DefaultTtl                int                                                 `json:"default_ttl"`
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecDefaultCacheBehavior `json:"lambda_function_association"`
	MaxTtl                    int                                                 `json:"max_ttl"`
	SmoothStreaming           bool                                                `json:"smooth_streaming"`
	TargetOriginId            string                                              `json:"target_origin_id"`
	TrustedSigners            []string                                            `json:"trusted_signers"`
	AllowedMethods            []string                                            `json:"allowed_methods"`
	Compress                  bool                                                `json:"compress"`
	ViewerProtocolPolicy      string                                              `json:"viewer_protocol_policy"`
	ForwardedValues           []AwsCloudfrontDistributionSpecDefaultCacheBehavior `json:"forwarded_values"`
	MinTtl                    int                                                 `json:"min_ttl"`
	CachedMethods             []string                                            `json:"cached_methods"`
	FieldLevelEncryptionId    string                                              `json:"field_level_encryption_id"`
}

type AwsCloudfrontDistributionSpecViewerCertificate struct {
	CloudfrontDefaultCertificate bool   `json:"cloudfront_default_certificate"`
	IamCertificateId             string `json:"iam_certificate_id"`
	MinimumProtocolVersion       string `json:"minimum_protocol_version"`
	SslSupportMethod             string `json:"ssl_support_method"`
	AcmCertificateArn            string `json:"acm_certificate_arn"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies struct {
	Forward          string   `json:"forward"`
	WhitelistedNames []string `json:"whitelisted_names"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues struct {
	QueryStringCacheKeys []string                                                           `json:"query_string_cache_keys"`
	Cookies              []AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues `json:"cookies"`
	Headers              []string                                                           `json:"headers"`
	QueryString          bool                                                               `json:"query_string"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation struct {
	EventType   string `json:"event_type"`
	LambdaArn   string `json:"lambda_arn"`
	IncludeBody bool   `json:"include_body"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehavior struct {
	ViewerProtocolPolicy      string                                              `json:"viewer_protocol_policy"`
	Compress                  bool                                                `json:"compress"`
	DefaultTtl                int                                                 `json:"default_ttl"`
	ForwardedValues           []AwsCloudfrontDistributionSpecOrderedCacheBehavior `json:"forwarded_values"`
	SmoothStreaming           bool                                                `json:"smooth_streaming"`
	FieldLevelEncryptionId    string                                              `json:"field_level_encryption_id"`
	MaxTtl                    int                                                 `json:"max_ttl"`
	PathPattern               string                                              `json:"path_pattern"`
	TargetOriginId            string                                              `json:"target_origin_id"`
	CachedMethods             []string                                            `json:"cached_methods"`
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecOrderedCacheBehavior `json:"lambda_function_association"`
	AllowedMethods            []string                                            `json:"allowed_methods"`
	MinTtl                    int                                                 `json:"min_ttl"`
	TrustedSigners            []string                                            `json:"trusted_signers"`
}

type AwsCloudfrontDistributionSpecLoggingConfig struct {
	IncludeCookies bool   `json:"include_cookies"`
	Prefix         string `json:"prefix"`
	Bucket         string `json:"bucket"`
}

type AwsCloudfrontDistributionSpecOriginGroupMember struct {
	OriginId string `json:"origin_id"`
}

type AwsCloudfrontDistributionSpecOriginGroupFailoverCriteria struct {
	StatusCodes []int64 `json:"status_codes"`
}

type AwsCloudfrontDistributionSpecOriginGroup struct {
	Member           []AwsCloudfrontDistributionSpecOriginGroup `json:"member"`
	OriginId         string                                     `json:"origin_id"`
	FailoverCriteria []AwsCloudfrontDistributionSpecOriginGroup `json:"failover_criteria"`
}

type AwsCloudfrontDistributionSpecRestrictionsGeoRestriction struct {
	Locations       []string `json:"locations"`
	RestrictionType string   `json:"restriction_type"`
}

type AwsCloudfrontDistributionSpecRestrictions struct {
	GeoRestriction []AwsCloudfrontDistributionSpecRestrictions `json:"geo_restriction"`
}

type AwsCloudfrontDistributionSpec struct {
	CacheBehavior               []AwsCloudfrontDistributionSpec `json:"cache_behavior"`
	WebAclId                    string                          `json:"web_acl_id"`
	CustomErrorResponse         []AwsCloudfrontDistributionSpec `json:"custom_error_response"`
	DefaultRootObject           string                          `json:"default_root_object"`
	Origin                      []AwsCloudfrontDistributionSpec `json:"origin"`
	Aliases                     []string                        `json:"aliases"`
	Comment                     string                          `json:"comment"`
	DefaultCacheBehavior        []AwsCloudfrontDistributionSpec `json:"default_cache_behavior"`
	ViewerCertificate           []AwsCloudfrontDistributionSpec `json:"viewer_certificate"`
	Arn                         string                          `json:"arn"`
	HostedZoneId                string                          `json:"hosted_zone_id"`
	IsIpv6Enabled               bool                            `json:"is_ipv6_enabled"`
	ActiveTrustedSigners        map[string]string               `json:"active_trusted_signers"`
	WaitForDeployment           bool                            `json:"wait_for_deployment"`
	OrderedCacheBehavior        []AwsCloudfrontDistributionSpec `json:"ordered_cache_behavior"`
	Enabled                     bool                            `json:"enabled"`
	HttpVersion                 string                          `json:"http_version"`
	LoggingConfig               []AwsCloudfrontDistributionSpec `json:"logging_config"`
	OriginGroup                 []AwsCloudfrontDistributionSpec `json:"origin_group"`
	InProgressValidationBatches int                             `json:"in_progress_validation_batches"`
	Etag                        string                          `json:"etag"`
	PriceClass                  string                          `json:"price_class"`
	DomainName                  string                          `json:"domain_name"`
	Tags                        map[string]string               `json:"tags"`
	Restrictions                []AwsCloudfrontDistributionSpec `json:"restrictions"`
	CallerReference             string                          `json:"caller_reference"`
	Status                      string                          `json:"status"`
	LastModifiedTime            string                          `json:"last_modified_time"`
	RetainOnDelete              bool                            `json:"retain_on_delete"`
}

type AwsCloudfrontDistributionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudfrontDistributionList is a list of AwsCloudfrontDistributions
type AwsCloudfrontDistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudfrontDistribution CRD objects
	Items []AwsCloudfrontDistribution `json:"items,omitempty"`
}
