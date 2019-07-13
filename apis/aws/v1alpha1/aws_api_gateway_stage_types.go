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

type AwsApiGatewayStage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayStageSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayStageStatus `json:"status,omitempty"`
}

type AwsApiGatewayStageSpecAccessLogSettings struct {
	Format         string `json:"format"`
	DestinationArn string `json:"destination_arn"`
}

type AwsApiGatewayStageSpec struct {
	ClientCertificateId  string                   `json:"client_certificate_id"`
	Description          string                   `json:"description"`
	Variables            map[string]string        `json:"variables"`
	AccessLogSettings    []AwsApiGatewayStageSpec `json:"access_log_settings"`
	DocumentationVersion string                   `json:"documentation_version"`
	InvokeUrl            string                   `json:"invoke_url"`
	CacheClusterSize     string                   `json:"cache_cluster_size"`
	ExecutionArn         string                   `json:"execution_arn"`
	RestApiId            string                   `json:"rest_api_id"`
	Tags                 map[string]string        `json:"tags"`
	DeploymentId         string                   `json:"deployment_id"`
	StageName            string                   `json:"stage_name"`
	XrayTracingEnabled   bool                     `json:"xray_tracing_enabled"`
	CacheClusterEnabled  bool                     `json:"cache_cluster_enabled"`
}



type AwsApiGatewayStageStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayStageList is a list of AwsApiGatewayStages
type AwsApiGatewayStageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayStage CRD objects
	Items []AwsApiGatewayStage `json:"items,omitempty"`
}