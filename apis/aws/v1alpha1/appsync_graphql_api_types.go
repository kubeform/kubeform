package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AppsyncGraphqlAPI struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncGraphqlAPISpec   `json:"spec,omitempty"`
	Status            AppsyncGraphqlAPIStatus `json:"status,omitempty"`
}

type AppsyncGraphqlAPISpecLogConfig struct {
	CloudwatchLogsRoleArn string `json:"cloudwatchLogsRoleArn" tf:"cloudwatch_logs_role_arn"`
	FieldLogLevel         string `json:"fieldLogLevel" tf:"field_log_level"`
}

type AppsyncGraphqlAPISpecOpenidConnectConfig struct {
	// +optional
	AuthTtl int `json:"authTtl,omitempty" tf:"auth_ttl,omitempty"`
	// +optional
	ClientID string `json:"clientID,omitempty" tf:"client_id,omitempty"`
	// +optional
	IatTtl int    `json:"iatTtl,omitempty" tf:"iat_ttl,omitempty"`
	Issuer string `json:"issuer" tf:"issuer"`
}

type AppsyncGraphqlAPISpecUserPoolConfig struct {
	// +optional
	AppIDClientRegex string `json:"appIDClientRegex,omitempty" tf:"app_id_client_regex,omitempty"`
	// +optional
	AwsRegion     string `json:"awsRegion,omitempty" tf:"aws_region,omitempty"`
	DefaultAction string `json:"defaultAction" tf:"default_action"`
	UserPoolID    string `json:"userPoolID" tf:"user_pool_id"`
}

type AppsyncGraphqlAPISpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                string `json:"arn,omitempty" tf:"arn,omitempty"`
	AuthenticationType string `json:"authenticationType" tf:"authentication_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LogConfig []AppsyncGraphqlAPISpecLogConfig `json:"logConfig,omitempty" tf:"log_config,omitempty"`
	Name      string                           `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OpenidConnectConfig []AppsyncGraphqlAPISpecOpenidConnectConfig `json:"openidConnectConfig,omitempty" tf:"openid_connect_config,omitempty"`
	// +optional
	Schema string `json:"schema,omitempty" tf:"schema,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Uris map[string]string `json:"uris,omitempty" tf:"uris,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	UserPoolConfig []AppsyncGraphqlAPISpecUserPoolConfig `json:"userPoolConfig,omitempty" tf:"user_pool_config,omitempty"`
}



type AppsyncGraphqlAPIStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *AppsyncGraphqlAPISpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppsyncGraphqlAPIList is a list of AppsyncGraphqlAPIs
type AppsyncGraphqlAPIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppsyncGraphqlAPI CRD objects
	Items []AppsyncGraphqlAPI `json:"items,omitempty"`
}