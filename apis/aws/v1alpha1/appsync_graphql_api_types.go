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

type AppsyncGraphqlApi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncGraphqlApiSpec   `json:"spec,omitempty"`
	Status            AppsyncGraphqlApiStatus `json:"status,omitempty"`
}

type AppsyncGraphqlApiSpecLogConfig struct {
	CloudwatchLogsRoleArn string `json:"cloudwatch_logs_role_arn"`
	FieldLogLevel         string `json:"field_log_level"`
}

type AppsyncGraphqlApiSpecOpenidConnectConfig struct {
	// +optional
	AuthTtl int `json:"auth_ttl,omitempty"`
	// +optional
	ClientId string `json:"client_id,omitempty"`
	// +optional
	IatTtl int    `json:"iat_ttl,omitempty"`
	Issuer string `json:"issuer"`
}

type AppsyncGraphqlApiSpecUserPoolConfig struct {
	// +optional
	AppIdClientRegex string `json:"app_id_client_regex,omitempty"`
	DefaultAction    string `json:"default_action"`
	UserPoolId       string `json:"user_pool_id"`
}

type AppsyncGraphqlApiSpec struct {
	AuthenticationType string `json:"authentication_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LogConfig *[]AppsyncGraphqlApiSpec `json:"log_config,omitempty"`
	Name      string                   `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OpenidConnectConfig *[]AppsyncGraphqlApiSpec `json:"openid_connect_config,omitempty"`
	// +optional
	Schema string `json:"schema,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	UserPoolConfig *[]AppsyncGraphqlApiSpec `json:"user_pool_config,omitempty"`
}

type AppsyncGraphqlApiStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppsyncGraphqlApiList is a list of AppsyncGraphqlApis
type AppsyncGraphqlApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppsyncGraphqlApi CRD objects
	Items []AppsyncGraphqlApi `json:"items,omitempty"`
}
