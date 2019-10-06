package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	ErrorCachingMinTtl int `json:"errorCachingMinTtl,omitempty" tf:"error_caching_min_ttl,omitempty"`
	ErrorCode          int `json:"errorCode" tf:"error_code"`
	// +optional
	ResponseCode int `json:"responseCode,omitempty" tf:"response_code,omitempty"`
	// +optional
	ResponsePagePath string `json:"responsePagePath,omitempty" tf:"response_page_path,omitempty"`
}

type CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies struct {
	Forward string `json:"forward" tf:"forward"`
	// +optional
	WhitelistedNames []string `json:"whitelistedNames,omitempty" tf:"whitelisted_names,omitempty"`
}

type CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues struct {
	// +kubebuilder:validation:MaxItems=1
	Cookies []CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies `json:"cookies" tf:"cookies"`
	// +optional
	Headers     []string `json:"headers,omitempty" tf:"headers,omitempty"`
	QueryString bool     `json:"queryString" tf:"query_string"`
	// +optional
	QueryStringCacheKeys []string `json:"queryStringCacheKeys,omitempty" tf:"query_string_cache_keys,omitempty"`
}

type CloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"eventType" tf:"event_type"`
	// +optional
	IncludeBody bool   `json:"includeBody,omitempty" tf:"include_body,omitempty"`
	LambdaArn   string `json:"lambdaArn" tf:"lambda_arn"`
}

type CloudfrontDistributionSpecDefaultCacheBehavior struct {
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	CachedMethods  []string `json:"cachedMethods" tf:"cached_methods"`
	// +optional
	Compress bool `json:"compress,omitempty" tf:"compress,omitempty"`
	// +optional
	DefaultTtl int `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`
	// +optional
	FieldLevelEncryptionID string `json:"fieldLevelEncryptionID,omitempty" tf:"field_level_encryption_id,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ForwardedValues []CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues `json:"forwardedValues" tf:"forwarded_values"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	LambdaFunctionAssociation []CloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation `json:"lambdaFunctionAssociation,omitempty" tf:"lambda_function_association,omitempty"`
	// +optional
	MaxTtl int `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`
	// +optional
	MinTtl int `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`
	// +optional
	SmoothStreaming bool   `json:"smoothStreaming,omitempty" tf:"smooth_streaming,omitempty"`
	TargetOriginID  string `json:"targetOriginID" tf:"target_origin_id"`
	// +optional
	TrustedSigners       []string `json:"trustedSigners,omitempty" tf:"trusted_signers,omitempty"`
	ViewerProtocolPolicy string   `json:"viewerProtocolPolicy" tf:"viewer_protocol_policy"`
}

type CloudfrontDistributionSpecLoggingConfig struct {
	Bucket string `json:"bucket" tf:"bucket"`
	// +optional
	IncludeCookies bool `json:"includeCookies,omitempty" tf:"include_cookies,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies struct {
	Forward string `json:"forward" tf:"forward"`
	// +optional
	WhitelistedNames []string `json:"whitelistedNames,omitempty" tf:"whitelisted_names,omitempty"`
}

type CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues struct {
	// +kubebuilder:validation:MaxItems=1
	Cookies []CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies `json:"cookies" tf:"cookies"`
	// +optional
	Headers     []string `json:"headers,omitempty" tf:"headers,omitempty"`
	QueryString bool     `json:"queryString" tf:"query_string"`
	// +optional
	QueryStringCacheKeys []string `json:"queryStringCacheKeys,omitempty" tf:"query_string_cache_keys,omitempty"`
}

type CloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation struct {
	EventType string `json:"eventType" tf:"event_type"`
	// +optional
	IncludeBody bool   `json:"includeBody,omitempty" tf:"include_body,omitempty"`
	LambdaArn   string `json:"lambdaArn" tf:"lambda_arn"`
}

type CloudfrontDistributionSpecOrderedCacheBehavior struct {
	AllowedMethods []string `json:"allowedMethods" tf:"allowed_methods"`
	CachedMethods  []string `json:"cachedMethods" tf:"cached_methods"`
	// +optional
	Compress bool `json:"compress,omitempty" tf:"compress,omitempty"`
	// +optional
	DefaultTtl int `json:"defaultTtl,omitempty" tf:"default_ttl,omitempty"`
	// +optional
	FieldLevelEncryptionID string `json:"fieldLevelEncryptionID,omitempty" tf:"field_level_encryption_id,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ForwardedValues []CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues `json:"forwardedValues" tf:"forwarded_values"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	LambdaFunctionAssociation []CloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation `json:"lambdaFunctionAssociation,omitempty" tf:"lambda_function_association,omitempty"`
	// +optional
	MaxTtl int `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`
	// +optional
	MinTtl      int    `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`
	PathPattern string `json:"pathPattern" tf:"path_pattern"`
	// +optional
	SmoothStreaming bool   `json:"smoothStreaming,omitempty" tf:"smooth_streaming,omitempty"`
	TargetOriginID  string `json:"targetOriginID" tf:"target_origin_id"`
	// +optional
	TrustedSigners       []string `json:"trustedSigners,omitempty" tf:"trusted_signers,omitempty"`
	ViewerProtocolPolicy string   `json:"viewerProtocolPolicy" tf:"viewer_protocol_policy"`
}

type CloudfrontDistributionSpecOriginCustomHeader struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type CloudfrontDistributionSpecOriginCustomOriginConfig struct {
	HttpPort  int `json:"httpPort" tf:"http_port"`
	HttpsPort int `json:"httpsPort" tf:"https_port"`
	// +optional
	OriginKeepaliveTimeout int    `json:"originKeepaliveTimeout,omitempty" tf:"origin_keepalive_timeout,omitempty"`
	OriginProtocolPolicy   string `json:"originProtocolPolicy" tf:"origin_protocol_policy"`
	// +optional
	OriginReadTimeout  int      `json:"originReadTimeout,omitempty" tf:"origin_read_timeout,omitempty"`
	OriginSSLProtocols []string `json:"originSSLProtocols" tf:"origin_ssl_protocols"`
}

type CloudfrontDistributionSpecOriginS3OriginConfig struct {
	OriginAccessIdentity string `json:"originAccessIdentity" tf:"origin_access_identity"`
}

type CloudfrontDistributionSpecOrigin struct {
	// +optional
	CustomHeader []CloudfrontDistributionSpecOriginCustomHeader `json:"customHeader,omitempty" tf:"custom_header,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomOriginConfig []CloudfrontDistributionSpecOriginCustomOriginConfig `json:"customOriginConfig,omitempty" tf:"custom_origin_config,omitempty"`
	DomainName         string                                               `json:"domainName" tf:"domain_name"`
	OriginID           string                                               `json:"originID" tf:"origin_id"`
	// +optional
	OriginPath string `json:"originPath,omitempty" tf:"origin_path,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	S3OriginConfig []CloudfrontDistributionSpecOriginS3OriginConfig `json:"s3OriginConfig,omitempty" tf:"s3_origin_config,omitempty"`
}

type CloudfrontDistributionSpecOriginGroupFailoverCriteria struct {
	StatusCodes []int64 `json:"statusCodes" tf:"status_codes"`
}

type CloudfrontDistributionSpecOriginGroupMember struct {
	OriginID string `json:"originID" tf:"origin_id"`
}

type CloudfrontDistributionSpecOriginGroup struct {
	// +kubebuilder:validation:MaxItems=1
	FailoverCriteria []CloudfrontDistributionSpecOriginGroupFailoverCriteria `json:"failoverCriteria" tf:"failover_criteria"`
	// +kubebuilder:validation:MinItems=2
	Member   []CloudfrontDistributionSpecOriginGroupMember `json:"member" tf:"member"`
	OriginID string                                        `json:"originID" tf:"origin_id"`
}

type CloudfrontDistributionSpecRestrictionsGeoRestriction struct {
	// +optional
	Locations       []string `json:"locations,omitempty" tf:"locations,omitempty"`
	RestrictionType string   `json:"restrictionType" tf:"restriction_type"`
}

type CloudfrontDistributionSpecRestrictions struct {
	// +kubebuilder:validation:MaxItems=1
	GeoRestriction []CloudfrontDistributionSpecRestrictionsGeoRestriction `json:"geoRestriction" tf:"geo_restriction"`
}

type CloudfrontDistributionSpecViewerCertificate struct {
	// +optional
	AcmCertificateArn string `json:"acmCertificateArn,omitempty" tf:"acm_certificate_arn,omitempty"`
	// +optional
	CloudfrontDefaultCertificate bool `json:"cloudfrontDefaultCertificate,omitempty" tf:"cloudfront_default_certificate,omitempty"`
	// +optional
	IamCertificateID string `json:"iamCertificateID,omitempty" tf:"iam_certificate_id,omitempty"`
	// +optional
	MinimumProtocolVersion string `json:"minimumProtocolVersion,omitempty" tf:"minimum_protocol_version,omitempty"`
	// +optional
	SslSupportMethod string `json:"sslSupportMethod,omitempty" tf:"ssl_support_method,omitempty"`
}

type CloudfrontDistributionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActiveTrustedSigners map[string]string `json:"activeTrustedSigners,omitempty" tf:"active_trusted_signers,omitempty"`
	// +optional
	Aliases []string `json:"aliases,omitempty" tf:"aliases,omitempty"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CallerReference string `json:"callerReference,omitempty" tf:"caller_reference,omitempty"`
	// +optional
	Comment string `json:"comment,omitempty" tf:"comment,omitempty"`
	// +optional
	CustomErrorResponse []CloudfrontDistributionSpecCustomErrorResponse `json:"customErrorResponse,omitempty" tf:"custom_error_response,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	DefaultCacheBehavior []CloudfrontDistributionSpecDefaultCacheBehavior `json:"defaultCacheBehavior" tf:"default_cache_behavior"`
	// +optional
	DefaultRootObject string `json:"defaultRootObject,omitempty" tf:"default_root_object,omitempty"`
	// +optional
	DomainName string `json:"domainName,omitempty" tf:"domain_name,omitempty"`
	Enabled    bool   `json:"enabled" tf:"enabled"`
	// +optional
	Etag string `json:"etag,omitempty" tf:"etag,omitempty"`
	// +optional
	HostedZoneID string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id,omitempty"`
	// +optional
	HttpVersion string `json:"httpVersion,omitempty" tf:"http_version,omitempty"`
	// +optional
	InProgressValidationBatches int `json:"inProgressValidationBatches,omitempty" tf:"in_progress_validation_batches,omitempty"`
	// +optional
	IsIpv6Enabled bool `json:"isIpv6Enabled,omitempty" tf:"is_ipv6_enabled,omitempty"`
	// +optional
	LastModifiedTime string `json:"lastModifiedTime,omitempty" tf:"last_modified_time,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LoggingConfig []CloudfrontDistributionSpecLoggingConfig `json:"loggingConfig,omitempty" tf:"logging_config,omitempty"`
	// +optional
	OrderedCacheBehavior []CloudfrontDistributionSpecOrderedCacheBehavior `json:"orderedCacheBehavior,omitempty" tf:"ordered_cache_behavior,omitempty"`
	Origin               []CloudfrontDistributionSpecOrigin               `json:"origin" tf:"origin"`
	// +optional
	OriginGroup []CloudfrontDistributionSpecOriginGroup `json:"originGroup,omitempty" tf:"origin_group,omitempty"`
	// +optional
	PriceClass string `json:"priceClass,omitempty" tf:"price_class,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Restrictions []CloudfrontDistributionSpecRestrictions `json:"restrictions" tf:"restrictions"`
	// +optional
	RetainOnDelete bool `json:"retainOnDelete,omitempty" tf:"retain_on_delete,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ViewerCertificate []CloudfrontDistributionSpecViewerCertificate `json:"viewerCertificate" tf:"viewer_certificate"`
	// +optional
	WaitForDeployment bool `json:"waitForDeployment,omitempty" tf:"wait_for_deployment,omitempty"`
	// +optional
	WebACLID string `json:"webACLID,omitempty" tf:"web_acl_id,omitempty"`
}

type CloudfrontDistributionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudfrontDistributionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
