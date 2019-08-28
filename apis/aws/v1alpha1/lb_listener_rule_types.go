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

type LbListenerRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbListenerRuleSpec   `json:"spec,omitempty"`
	Status            LbListenerRuleStatus `json:"status,omitempty"`
}

type LbListenerRuleSpecActionAuthenticateCognito struct {
	// +optional
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`
	// +optional
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`
	// +optional
	Scope string `json:"scope,omitempty" tf:"scope,omitempty"`
	// +optional
	SessionCookieName string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`
	// +optional
	SessionTimeout   int    `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`
	UserPoolArn      string `json:"userPoolArn" tf:"user_pool_arn"`
	UserPoolClientID string `json:"userPoolClientID" tf:"user_pool_client_id"`
	UserPoolDomain   string `json:"userPoolDomain" tf:"user_pool_domain"`
}

type LbListenerRuleSpecActionAuthenticateOidc struct {
	// +optional
	AuthenticationRequestExtraParams map[string]string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`
	AuthorizationEndpoint            string            `json:"authorizationEndpoint" tf:"authorization_endpoint"`
	ClientID                         string            `json:"clientID" tf:"client_id"`
	ClientSecret                     string            `json:"-" sensitive:"true" tf:"client_secret"`
	Issuer                           string            `json:"issuer" tf:"issuer"`
	// +optional
	OnUnauthenticatedRequest string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`
	// +optional
	Scope string `json:"scope,omitempty" tf:"scope,omitempty"`
	// +optional
	SessionCookieName string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`
	// +optional
	SessionTimeout   int    `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`
	TokenEndpoint    string `json:"tokenEndpoint" tf:"token_endpoint"`
	UserInfoEndpoint string `json:"userInfoEndpoint" tf:"user_info_endpoint"`
}

type LbListenerRuleSpecActionFixedResponse struct {
	ContentType string `json:"contentType" tf:"content_type"`
	// +optional
	MessageBody string `json:"messageBody,omitempty" tf:"message_body,omitempty"`
	// +optional
	StatusCode string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type LbListenerRuleSpecActionRedirect struct {
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Port string `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	Query      string `json:"query,omitempty" tf:"query,omitempty"`
	StatusCode string `json:"statusCode" tf:"status_code"`
}

type LbListenerRuleSpecAction struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateCognito []LbListenerRuleSpecActionAuthenticateCognito `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateOidc []LbListenerRuleSpecActionAuthenticateOidc `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedResponse []LbListenerRuleSpecActionFixedResponse `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`
	// +optional
	Order int `json:"order,omitempty" tf:"order,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Redirect []LbListenerRuleSpecActionRedirect `json:"redirect,omitempty" tf:"redirect,omitempty"`
	// +optional
	TargetGroupArn string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`
	Type           string `json:"type" tf:"type"`
}

type LbListenerRuleSpecCondition struct {
	// +optional
	Field string `json:"field,omitempty" tf:"field,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Values []string `json:"values,omitempty" tf:"values,omitempty"`
}

type LbListenerRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	Action []LbListenerRuleSpecAction `json:"action" tf:"action"`
	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Condition   []LbListenerRuleSpecCondition `json:"condition" tf:"condition"`
	ListenerArn string                        `json:"listenerArn" tf:"listener_arn"`
	// +optional
	Priority int `json:"priority,omitempty" tf:"priority,omitempty"`
}



type LbListenerRuleStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *LbListenerRuleSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbListenerRuleList is a list of LbListenerRules
type LbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbListenerRule CRD objects
	Items []LbListenerRule `json:"items,omitempty"`
}