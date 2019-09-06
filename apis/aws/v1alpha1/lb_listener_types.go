package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LbListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbListenerSpec   `json:"spec,omitempty"`
	Status            LbListenerStatus `json:"status,omitempty"`
}

type LbListenerSpecDefaultActionAuthenticateCognito struct {
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

type LbListenerSpecDefaultActionAuthenticateOidc struct {
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

type LbListenerSpecDefaultActionFixedResponse struct {
	ContentType string `json:"contentType" tf:"content_type"`
	// +optional
	MessageBody string `json:"messageBody,omitempty" tf:"message_body,omitempty"`
	// +optional
	StatusCode string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type LbListenerSpecDefaultActionRedirect struct {
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

type LbListenerSpecDefaultAction struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateCognito []LbListenerSpecDefaultActionAuthenticateCognito `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AuthenticateOidc []LbListenerSpecDefaultActionAuthenticateOidc `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedResponse []LbListenerSpecDefaultActionFixedResponse `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`
	// +optional
	Order int `json:"order,omitempty" tf:"order,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Redirect []LbListenerSpecDefaultActionRedirect `json:"redirect,omitempty" tf:"redirect,omitempty"`
	// +optional
	TargetGroupArn string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`
	Type           string `json:"type" tf:"type"`
}

type LbListenerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	CertificateArn  string                        `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`
	DefaultAction   []LbListenerSpecDefaultAction `json:"defaultAction" tf:"default_action"`
	LoadBalancerArn string                        `json:"loadBalancerArn" tf:"load_balancer_arn"`
	Port            int                           `json:"port" tf:"port"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	SslPolicy string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`
}

type LbListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbListenerSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbListenerList is a list of LbListeners
type LbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbListener CRD objects
	Items []LbListener `json:"items,omitempty"`
}
