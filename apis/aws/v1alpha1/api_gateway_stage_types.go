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

type ApiGatewayStage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayStageSpec   `json:"spec,omitempty"`
	Status            ApiGatewayStageStatus `json:"status,omitempty"`
}

type ApiGatewayStageSpecAccessLogSettings struct {
	DestinationArn string `json:"destination_arn"`
	Format         string `json:"format"`
}

type ApiGatewayStageSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLogSettings *[]ApiGatewayStageSpec `json:"access_log_settings,omitempty"`
	// +optional
	CacheClusterEnabled bool `json:"cache_cluster_enabled,omitempty"`
	// +optional
	CacheClusterSize string `json:"cache_cluster_size,omitempty"`
	// +optional
	ClientCertificateId string `json:"client_certificate_id,omitempty"`
	DeploymentId        string `json:"deployment_id"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	DocumentationVersion string `json:"documentation_version,omitempty"`
	RestApiId            string `json:"rest_api_id"`
	StageName            string `json:"stage_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	Variables map[string]string `json:"variables,omitempty"`
	// +optional
	XrayTracingEnabled bool `json:"xray_tracing_enabled,omitempty"`
}

type ApiGatewayStageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
