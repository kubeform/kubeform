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

type AlbListenerRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AlbListenerRuleSpec   `json:"spec,omitempty"`
	Status            AlbListenerRuleStatus `json:"status,omitempty"`
}

type AlbListenerRuleSpecActionAuthenticateCognito struct {
	// +optional
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params,omitempty"`
	UserPoolArn                      string            `json:"user_pool_arn"`
	UserPoolClientId                 string            `json:"user_pool_client_id"`
	UserPoolDomain                   string            `json:"user_pool_domain"`
}

type AlbListenerRuleSpecActionAuthenticateOidc struct {
	// +optional
	AuthenticationRequestExtraParams map[string]string `json:"authentication_request_extra_params,omitempty"`
	AuthorizationEndpoint            string            `json:"authorization_endpoint"`
	ClientId                         string            `json:"client_id"`
	ClientSecret                     string            `json:"client_secret"`
	Issuer                           string            `json:"issuer"`
	TokenEndpoint                    string            `json:"token_endpoint"`
	UserInfoEndpoint                 string            `json:"user_info_endpoint"`
}

type AlbListenerRuleSpecActionFixedResponse struct {
	ContentType string `json:"content_type"`
	// +optional
	MessageBody string `json:"message_body,omitempty"`
}

type AlbListenerRuleSpecActionRedirect struct {
	// +optional
	Host string `json:"host,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
	// +optional
	Port string `json:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty"`
	// +optional
	Query      string `json:"query,omitempty"`
	StatusCode string `json:"status_code"`
}

type AlbListenerRuleSpecAction struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateCognito *[]AlbListenerRuleSpecAction `json:"authenticate_cognito,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateOidc *[]AlbListenerRuleSpecAction `json:"authenticate_oidc,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedResponse *[]AlbListenerRuleSpecAction `json:"fixed_response,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Redirect *[]AlbListenerRuleSpecAction `json:"redirect,omitempty"`
	// +optional
	TargetGroupArn string `json:"target_group_arn,omitempty"`
	Type           string `json:"type"`
}

type AlbListenerRuleSpecCondition struct {
	// +optional
	Field string `json:"field,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Values []string `json:"values,omitempty"`
}

type AlbListenerRuleSpec struct {
	Action []AlbListenerRuleSpec `json:"action"`
	// +kubebuilder:validation:UniqueItems=true
	Condition   []AlbListenerRuleSpec `json:"condition"`
	ListenerArn string                `json:"listener_arn"`
}

type AlbListenerRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AlbListenerRuleList is a list of AlbListenerRules
type AlbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AlbListenerRule CRD objects
	Items []AlbListenerRule `json:"items,omitempty"`
}
