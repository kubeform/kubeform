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
	DestinationArn string `json:"destination_arn"`
	Format         string `json:"format"`
}

type AwsApiGatewayStageSpec struct {
	DeploymentId         string                   `json:"deployment_id"`
	ExecutionArn         string                   `json:"execution_arn"`
	ClientCertificateId  string                   `json:"client_certificate_id"`
	StageName            string                   `json:"stage_name"`
	Variables            map[string]string        `json:"variables"`
	AccessLogSettings    []AwsApiGatewayStageSpec `json:"access_log_settings"`
	CacheClusterSize     string                   `json:"cache_cluster_size"`
	RestApiId            string                   `json:"rest_api_id"`
	Tags                 map[string]string        `json:"tags"`
	XrayTracingEnabled   bool                     `json:"xray_tracing_enabled"`
	CacheClusterEnabled  bool                     `json:"cache_cluster_enabled"`
	Description          string                   `json:"description"`
	DocumentationVersion string                   `json:"documentation_version"`
	InvokeUrl            string                   `json:"invoke_url"`
}

type AwsApiGatewayStageStatus struct {
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
