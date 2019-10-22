package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ApiGatewayStage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayStageSpec   `json:"spec,omitempty"`
	Status            ApiGatewayStageStatus `json:"status,omitempty"`
}

type ApiGatewayStageSpecAccessLogSettings struct {
	DestinationArn string `json:"destinationArn" tf:"destination_arn"`
	Format         string `json:"format" tf:"format"`
}

type ApiGatewayStageSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogSettings []ApiGatewayStageSpecAccessLogSettings `json:"accessLogSettings,omitempty" tf:"access_log_settings,omitempty"`
	// +optional
	CacheClusterEnabled bool `json:"cacheClusterEnabled,omitempty" tf:"cache_cluster_enabled,omitempty"`
	// +optional
	CacheClusterSize string `json:"cacheClusterSize,omitempty" tf:"cache_cluster_size,omitempty"`
	// +optional
	ClientCertificateID string `json:"clientCertificateID,omitempty" tf:"client_certificate_id,omitempty"`
	DeploymentID        string `json:"deploymentID" tf:"deployment_id"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	DocumentationVersion string `json:"documentationVersion,omitempty" tf:"documentation_version,omitempty"`
	// +optional
	ExecutionArn string `json:"executionArn,omitempty" tf:"execution_arn,omitempty"`
	// +optional
	InvokeURL string `json:"invokeURL,omitempty" tf:"invoke_url,omitempty"`
	RestAPIID string `json:"restAPIID" tf:"rest_api_id"`
	StageName string `json:"stageName" tf:"stage_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Variables map[string]string `json:"variables,omitempty" tf:"variables,omitempty"`
	// +optional
	XrayTracingEnabled bool `json:"xrayTracingEnabled,omitempty" tf:"xray_tracing_enabled,omitempty"`
}

type ApiGatewayStageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayStageSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayStageList is a list of ApiGatewayStages
type ApiGatewayStageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayStage CRD objects
	Items []ApiGatewayStage `json:"items,omitempty"`
}
