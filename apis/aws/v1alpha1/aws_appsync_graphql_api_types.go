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

type AwsAppsyncGraphqlApi struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppsyncGraphqlApiSpec   `json:"spec,omitempty"`
	Status            AwsAppsyncGraphqlApiStatus `json:"status,omitempty"`
}

type AwsAppsyncGraphqlApiSpecLogConfig struct {
	CloudwatchLogsRoleArn string `json:"cloudwatch_logs_role_arn"`
	FieldLogLevel         string `json:"field_log_level"`
}

type AwsAppsyncGraphqlApiSpecOpenidConnectConfig struct {
	AuthTtl  int    `json:"auth_ttl"`
	ClientId string `json:"client_id"`
	IatTtl   int    `json:"iat_ttl"`
	Issuer   string `json:"issuer"`
}

type AwsAppsyncGraphqlApiSpecUserPoolConfig struct {
	UserPoolId       string `json:"user_pool_id"`
	AppIdClientRegex string `json:"app_id_client_regex"`
	AwsRegion        string `json:"aws_region"`
	DefaultAction    string `json:"default_action"`
}

type AwsAppsyncGraphqlApiSpec struct {
	Schema              string                     `json:"schema"`
	Name                string                     `json:"name"`
	LogConfig           []AwsAppsyncGraphqlApiSpec `json:"log_config"`
	OpenidConnectConfig []AwsAppsyncGraphqlApiSpec `json:"openid_connect_config"`
	UserPoolConfig      []AwsAppsyncGraphqlApiSpec `json:"user_pool_config"`
	Tags                map[string]string          `json:"tags"`
	AuthenticationType  string                     `json:"authentication_type"`
	Uris                map[string]string          `json:"uris"`
	Arn                 string                     `json:"arn"`
}



type AwsAppsyncGraphqlApiStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppsyncGraphqlApiList is a list of AwsAppsyncGraphqlApis
type AwsAppsyncGraphqlApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppsyncGraphqlApi CRD objects
	Items []AwsAppsyncGraphqlApi `json:"items,omitempty"`
}