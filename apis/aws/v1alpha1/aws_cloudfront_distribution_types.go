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

type AwsCloudfrontDistributionSpecLoggingConfig struct {
	Bucket         string `json:"bucket"`
	IncludeCookies bool   `json:"include_cookies"`
	Prefix         string `json:"prefix"`
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
	CachedMethods             []string                                     `json:"cached_methods"`
	Compress                  bool                                         `json:"compress"`
	DefaultTtl                int                                          `json:"default_ttl"`
	ForwardedValues           []AwsCloudfrontDistributionSpecCacheBehavior `json:"forwarded_values"`
	MaxTtl                    int                                          `json:"max_ttl"`
	TargetOriginId            string                                       `json:"target_origin_id"`
	PathPattern               string                                       `json:"path_pattern"`
	SmoothStreaming           bool                                         `json:"smooth_streaming"`
	TrustedSigners            []string                                     `json:"trusted_signers"`
	AllowedMethods            []string                                     `json:"allowed_methods"`
	FieldLevelEncryptionId    string                                       `json:"field_level_encryption_id"`
	MinTtl                    int                                          `json:"min_ttl"`
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecCacheBehavior `json:"lambda_function_association"`
	ViewerProtocolPolicy      string                                       `json:"viewer_protocol_policy"`
}

type AwsCloudfrontDistributionSpecCustomErrorResponse struct {
	ErrorCachingMinTtl int    `json:"error_caching_min_ttl"`
	ErrorCode          int    `json:"error_code"`
	ResponseCode       int    `json:"response_code"`
	ResponsePagePath   string `json:"response_page_path"`
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

type AwsCloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation struct {
	LambdaArn   string `json:"lambda_arn"`
	IncludeBody bool   `json:"include_body"`
	EventType   string `json:"event_type"`
}

type AwsCloudfrontDistributionSpecDefaultCacheBehavior struct {
	TrustedSigners            []string                                            `json:"trusted_signers"`
	AllowedMethods            []string                                            `json:"allowed_methods"`
	DefaultTtl                int                                                 `json:"default_ttl"`
	FieldLevelEncryptionId    string                                              `json:"field_level_encryption_id"`
	ForwardedValues           []AwsCloudfrontDistributionSpecDefaultCacheBehavior `json:"forwarded_values"`
	MaxTtl                    int                                                 `json:"max_ttl"`
	SmoothStreaming           bool                                                `json:"smooth_streaming"`
	TargetOriginId            string                                              `json:"target_origin_id"`
	CachedMethods             []string                                            `json:"cached_methods"`
	Compress                  bool                                                `json:"compress"`
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecDefaultCacheBehavior `json:"lambda_function_association"`
	MinTtl                    int                                                 `json:"min_ttl"`
	ViewerProtocolPolicy      string                                              `json:"viewer_protocol_policy"`
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

type AwsCloudfrontDistributionSpecOriginS3OriginConfig struct {
	OriginAccessIdentity string `json:"origin_access_identity"`
}

type AwsCloudfrontDistributionSpecOriginCustomOriginConfig struct {
	OriginProtocolPolicy   string   `json:"origin_protocol_policy"`
	OriginSslProtocols     []string `json:"origin_ssl_protocols"`
	HttpPort               int      `json:"http_port"`
	HttpsPort              int      `json:"https_port"`
	OriginKeepaliveTimeout int      `json:"origin_keepalive_timeout"`
	OriginReadTimeout      int      `json:"origin_read_timeout"`
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

type AwsCloudfrontDistributionSpecRestrictionsGeoRestriction struct {
	Locations       []string `json:"locations"`
	RestrictionType string   `json:"restriction_type"`
}

type AwsCloudfrontDistributionSpecRestrictions struct {
	GeoRestriction []AwsCloudfrontDistributionSpecRestrictions `json:"geo_restriction"`
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
	QueryStringCacheKeys []string                                                           `json:"query_string_cache_keys"`
	Cookies              []AwsCloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues `json:"cookies"`
	Headers              []string                                                           `json:"headers"`
	QueryString          bool                                                               `json:"query_string"`
}

type AwsCloudfrontDistributionSpecOrderedCacheBehavior struct {
	LambdaFunctionAssociation []AwsCloudfrontDistributionSpecOrderedCacheBehavior `json:"lambda_function_association"`
	SmoothStreaming           bool                                                `json:"smooth_streaming"`
	ViewerProtocolPolicy      string                                              `json:"viewer_protocol_policy"`
	CachedMethods             []string                                            `json:"cached_methods"`
	TargetOriginId            string                                              `json:"target_origin_id"`
	MaxTtl                    int                                                 `json:"max_ttl"`
	MinTtl                    int                                                 `json:"min_ttl"`
	TrustedSigners            []string                                            `json:"trusted_signers"`
	AllowedMethods            []string                                            `json:"allowed_methods"`
	Compress                  bool                                                `json:"compress"`
	DefaultTtl                int                                                 `json:"default_ttl"`
	FieldLevelEncryptionId    string                                              `json:"field_level_encryption_id"`
	ForwardedValues           []AwsCloudfrontDistributionSpecOrderedCacheBehavior `json:"forwarded_values"`
	PathPattern               string                                              `json:"path_pattern"`
}

type AwsCloudfrontDistributionSpec struct {
	Enabled                     bool                            `json:"enabled"`
	LoggingConfig               []AwsCloudfrontDistributionSpec `json:"logging_config"`
	IsIpv6Enabled               bool                            `json:"is_ipv6_enabled"`
	CacheBehavior               []AwsCloudfrontDistributionSpec `json:"cache_behavior"`
	PriceClass                  string                          `json:"price_class"`
	CallerReference             string                          `json:"caller_reference"`
	Arn                         string                          `json:"arn"`
	CustomErrorResponse         []AwsCloudfrontDistributionSpec `json:"custom_error_response"`
	DefaultCacheBehavior        []AwsCloudfrontDistributionSpec `json:"default_cache_behavior"`
	Status                      string                          `json:"status"`
	ActiveTrustedSigners        map[string]string               `json:"active_trusted_signers"`
	HostedZoneId                string                          `json:"hosted_zone_id"`
	Tags                        map[string]string               `json:"tags"`
	Comment                     string                          `json:"comment"`
	OriginGroup                 []AwsCloudfrontDistributionSpec `json:"origin_group"`
	Origin                      []AwsCloudfrontDistributionSpec `json:"origin"`
	WebAclId                    string                          `json:"web_acl_id"`
	ViewerCertificate           []AwsCloudfrontDistributionSpec `json:"viewer_certificate"`
	DomainName                  string                          `json:"domain_name"`
	RetainOnDelete              bool                            `json:"retain_on_delete"`
	Aliases                     []string                        `json:"aliases"`
	DefaultRootObject           string                          `json:"default_root_object"`
	Restrictions                []AwsCloudfrontDistributionSpec `json:"restrictions"`
	InProgressValidationBatches int                             `json:"in_progress_validation_batches"`
	OrderedCacheBehavior        []AwsCloudfrontDistributionSpec `json:"ordered_cache_behavior"`
	Etag                        string                          `json:"etag"`
	HttpVersion                 string                          `json:"http_version"`
	LastModifiedTime            string                          `json:"last_modified_time"`
	WaitForDeployment           bool                            `json:"wait_for_deployment"`
}

type AwsCloudfrontDistributionStatus struct {
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
