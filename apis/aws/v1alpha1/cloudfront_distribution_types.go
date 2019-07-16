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

type CloudfrontDistribution struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudfrontDistributionSpec   `json:"spec,omitempty"`
	Status            CloudfrontDistributionStatus `json:"status,omitempty"`
}

type CloudfrontDistributionSpecCustomErrorResponse struct {
	// +optional
	ErrorCachingMinTtl int `json:"error_caching_min_ttl,omitempty"`
	ErrorCode          int `json:"error_code"`
	// +optional
	ResponseCode int `json:"response_code,omitempty"`
	// +optional
	ResponsePagePath string `json:"response_page_path,omitempty"`
}

type CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies struct {
	Forward string `json:"forward"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WhitelistedNames []string `json:"whitelisted_names,omitempty"`
}

type CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues struct {
	// +kubebuilder:validation:MaxItems=1
	Cookies []CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues `json:"cookies"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Headers     []string `json:"headers,omitempty"`
	QueryString bool     `json:"query_string"`
	// +optional
	QueryStringCacheKeys []string `json:"query_string_cache_keys,omitempty"`
}

type CloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"event_type"`
	// +optional
	IncludeBody bool   `json:"include_body,omitempty"`
	LambdaArn   string `json:"lambda_arn"`
}

type CloudfrontDistributionSpecDefaultCacheBehavior struct {
	// +kubebuilder:validation:UniqueItems=true
	AllowedMethods []string `json:"allowed_methods"`
	// +kubebuilder:validation:UniqueItems=true
	CachedMethods []string `json:"cached_methods"`
	// +optional
	Compress bool `json:"compress,omitempty"`
	// +optional
	DefaultTtl int `json:"default_ttl,omitempty"`
	// +optional
	FieldLevelEncryptionId string `json:"field_level_encryption_id,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ForwardedValues []CloudfrontDistributionSpecDefaultCacheBehavior `json:"forwarded_values"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:UniqueItems=true
	LambdaFunctionAssociation *[]CloudfrontDistributionSpecDefaultCacheBehavior `json:"lambda_function_association,omitempty"`
	// +optional
	MaxTtl int `json:"max_ttl,omitempty"`
	// +optional
	MinTtl int `json:"min_ttl,omitempty"`
	// +optional
	SmoothStreaming bool   `json:"smooth_streaming,omitempty"`
	TargetOriginId  string `json:"target_origin_id"`
	// +optional
	TrustedSigners       []string `json:"trusted_signers,omitempty"`
	ViewerProtocolPolicy string   `json:"viewer_protocol_policy"`
}

type CloudfrontDistributionSpecLoggingConfig struct {
	Bucket string `json:"bucket"`
	// +optional
	IncludeCookies bool `json:"include_cookies,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty"`
}

type CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies struct {
	Forward string `json:"forward"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	WhitelistedNames []string `json:"whitelisted_names,omitempty"`
}

type CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues struct {
	// +kubebuilder:validation:MaxItems=1
	Cookies []CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues `json:"cookies"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Headers     []string `json:"headers,omitempty"`
	QueryString bool     `json:"query_string"`
	// +optional
	QueryStringCacheKeys []string `json:"query_string_cache_keys,omitempty"`
}

type CloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"event_type"`
	// +optional
	IncludeBody bool   `json:"include_body,omitempty"`
	LambdaArn   string `json:"lambda_arn"`
}

type CloudfrontDistributionSpecOrderedCacheBehavior struct {
	// +kubebuilder:validation:UniqueItems=true
	AllowedMethods []string `json:"allowed_methods"`
	// +kubebuilder:validation:UniqueItems=true
	CachedMethods []string `json:"cached_methods"`
	// +optional
	Compress bool `json:"compress,omitempty"`
	// +optional
	DefaultTtl int `json:"default_ttl,omitempty"`
	// +optional
	FieldLevelEncryptionId string `json:"field_level_encryption_id,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ForwardedValues []CloudfrontDistributionSpecOrderedCacheBehavior `json:"forwarded_values"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:UniqueItems=true
	LambdaFunctionAssociation *[]CloudfrontDistributionSpecOrderedCacheBehavior `json:"lambda_function_association,omitempty"`
	// +optional
	MaxTtl int `json:"max_ttl,omitempty"`
	// +optional
	MinTtl      int    `json:"min_ttl,omitempty"`
	PathPattern string `json:"path_pattern"`
	// +optional
	SmoothStreaming bool   `json:"smooth_streaming,omitempty"`
	TargetOriginId  string `json:"target_origin_id"`
	// +optional
	TrustedSigners       []string `json:"trusted_signers,omitempty"`
	ViewerProtocolPolicy string   `json:"viewer_protocol_policy"`
}

type CloudfrontDistributionSpecOriginCustomHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CloudfrontDistributionSpecOriginCustomOriginConfig struct {
	HttpPort  int `json:"http_port"`
	HttpsPort int `json:"https_port"`
	// +optional
	OriginKeepaliveTimeout int    `json:"origin_keepalive_timeout,omitempty"`
	OriginProtocolPolicy   string `json:"origin_protocol_policy"`
	// +optional
	OriginReadTimeout int `json:"origin_read_timeout,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	OriginSslProtocols []string `json:"origin_ssl_protocols"`
}

type CloudfrontDistributionSpecOriginS3OriginConfig struct {
	OriginAccessIdentity string `json:"origin_access_identity"`
}

type CloudfrontDistributionSpecOrigin struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CustomHeader *[]CloudfrontDistributionSpecOrigin `json:"custom_header,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomOriginConfig *[]CloudfrontDistributionSpecOrigin `json:"custom_origin_config,omitempty"`
	DomainName         string                              `json:"domain_name"`
	OriginId           string                              `json:"origin_id"`
	// +optional
	OriginPath string `json:"origin_path,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3OriginConfig *[]CloudfrontDistributionSpecOrigin `json:"s3_origin_config,omitempty"`
}

type CloudfrontDistributionSpecOriginGroupFailoverCriteria struct {
	// +kubebuilder:validation:UniqueItems=true
	StatusCodes []int64 `json:"status_codes"`
}

type CloudfrontDistributionSpecOriginGroupMember struct {
	OriginId string `json:"origin_id"`
}

type CloudfrontDistributionSpecOriginGroup struct {
	// +kubebuilder:validation:MaxItems=1
	FailoverCriteria []CloudfrontDistributionSpecOriginGroup `json:"failover_criteria"`
	// +kubebuilder:validation:MinItems=2
	Member   []CloudfrontDistributionSpecOriginGroup `json:"member"`
	OriginId string                                  `json:"origin_id"`
}

type CloudfrontDistributionSpecRestrictionsGeoRestriction struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Locations       []string `json:"locations,omitempty"`
	RestrictionType string   `json:"restriction_type"`
}

type CloudfrontDistributionSpecRestrictions struct {
	// +kubebuilder:validation:MaxItems=1
	GeoRestriction []CloudfrontDistributionSpecRestrictions `json:"geo_restriction"`
}

type CloudfrontDistributionSpecViewerCertificate struct {
	// +optional
	AcmCertificateArn string `json:"acm_certificate_arn,omitempty"`
	// +optional
	CloudfrontDefaultCertificate bool `json:"cloudfront_default_certificate,omitempty"`
	// +optional
	IamCertificateId string `json:"iam_certificate_id,omitempty"`
	// +optional
	MinimumProtocolVersion string `json:"minimum_protocol_version,omitempty"`
	// +optional
	SslSupportMethod string `json:"ssl_support_method,omitempty"`
}

type CloudfrontDistributionSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Aliases []string `json:"aliases,omitempty"`
	// +optional
	Comment string `json:"comment,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CustomErrorResponse *[]CloudfrontDistributionSpec `json:"custom_error_response,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	DefaultCacheBehavior []CloudfrontDistributionSpec `json:"default_cache_behavior"`
	// +optional
	DefaultRootObject string `json:"default_root_object,omitempty"`
	Enabled           bool   `json:"enabled"`
	// +optional
	HttpVersion string `json:"http_version,omitempty"`
	// +optional
	IsIpv6Enabled bool `json:"is_ipv6_enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig *[]CloudfrontDistributionSpec `json:"logging_config,omitempty"`
	// +optional
	OrderedCacheBehavior *[]CloudfrontDistributionSpec `json:"ordered_cache_behavior,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Origin []CloudfrontDistributionSpec `json:"origin"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OriginGroup *[]CloudfrontDistributionSpec `json:"origin_group,omitempty"`
	// +optional
	PriceClass string `json:"price_class,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Restrictions []CloudfrontDistributionSpec `json:"restrictions"`
	// +optional
	RetainOnDelete bool `json:"retain_on_delete,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ViewerCertificate []CloudfrontDistributionSpec `json:"viewer_certificate"`
	// +optional
	WaitForDeployment bool `json:"wait_for_deployment,omitempty"`
	// +optional
	WebAclId string `json:"web_acl_id,omitempty"`
}

type CloudfrontDistributionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudfrontDistributionList is a list of CloudfrontDistributions
type CloudfrontDistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudfrontDistribution CRD objects
	Items []CloudfrontDistribution `json:"items,omitempty"`
}
