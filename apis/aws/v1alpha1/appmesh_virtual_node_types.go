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

type AppmeshVirtualNode struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualNodeSpec   `json:"spec,omitempty"`
	Status            AppmeshVirtualNodeStatus `json:"status,omitempty"`
}

type AppmeshVirtualNodeSpecSpecBackendVirtualService struct {
	VirtualServiceName string `json:"virtualServiceName" tf:"virtual_service_name"`
}

type AppmeshVirtualNodeSpecSpecBackend struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VirtualService []AppmeshVirtualNodeSpecSpecBackendVirtualService `json:"virtualService,omitempty" tf:"virtual_service,omitempty"`
}

type AppmeshVirtualNodeSpecSpecListenerHealthCheck struct {
	HealthyThreshold int `json:"healthyThreshold" tf:"healthy_threshold"`
	IntervalMillis   int `json:"intervalMillis" tf:"interval_millis"`
	// +optional
	Path string `json:"path,omitempty" tf:"path,omitempty"`
	// +optional
	Port               int    `json:"port,omitempty" tf:"port,omitempty"`
	Protocol           string `json:"protocol" tf:"protocol"`
	TimeoutMillis      int    `json:"timeoutMillis" tf:"timeout_millis"`
	UnhealthyThreshold int    `json:"unhealthyThreshold" tf:"unhealthy_threshold"`
}

type AppmeshVirtualNodeSpecSpecListenerPortMapping struct {
	Port     int    `json:"port" tf:"port"`
	Protocol string `json:"protocol" tf:"protocol"`
}

type AppmeshVirtualNodeSpecSpecListener struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheck []AppmeshVirtualNodeSpecSpecListenerHealthCheck `json:"healthCheck,omitempty" tf:"health_check,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	PortMapping []AppmeshVirtualNodeSpecSpecListenerPortMapping `json:"portMapping" tf:"port_mapping"`
}

type AppmeshVirtualNodeSpecSpecLoggingAccessLogFile struct {
	Path string `json:"path" tf:"path"`
}

type AppmeshVirtualNodeSpecSpecLoggingAccessLog struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	File []AppmeshVirtualNodeSpecSpecLoggingAccessLogFile `json:"file,omitempty" tf:"file,omitempty"`
}

type AppmeshVirtualNodeSpecSpecLogging struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLog []AppmeshVirtualNodeSpecSpecLoggingAccessLog `json:"accessLog,omitempty" tf:"access_log,omitempty"`
}

type AppmeshVirtualNodeSpecSpecServiceDiscoveryAwsCloudMap struct {
	// +optional
	Attributes    map[string]string `json:"attributes,omitempty" tf:"attributes,omitempty"`
	NamespaceName string            `json:"namespaceName" tf:"namespace_name"`
	ServiceName   string            `json:"serviceName" tf:"service_name"`
}

type AppmeshVirtualNodeSpecSpecServiceDiscoveryDns struct {
	Hostname string `json:"hostname" tf:"hostname"`
}

type AppmeshVirtualNodeSpecSpecServiceDiscovery struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AwsCloudMap []AppmeshVirtualNodeSpecSpecServiceDiscoveryAwsCloudMap `json:"awsCloudMap,omitempty" tf:"aws_cloud_map,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Dns []AppmeshVirtualNodeSpecSpecServiceDiscoveryDns `json:"dns,omitempty" tf:"dns,omitempty"`
}

type AppmeshVirtualNodeSpecSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=25
	// +kubebuilder:validation:UniqueItems=true
	Backend []AppmeshVirtualNodeSpecSpecBackend `json:"backend,omitempty" tf:"backend,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Listener []AppmeshVirtualNodeSpecSpecListener `json:"listener,omitempty" tf:"listener,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logging []AppmeshVirtualNodeSpecSpecLogging `json:"logging,omitempty" tf:"logging,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceDiscovery []AppmeshVirtualNodeSpecSpecServiceDiscovery `json:"serviceDiscovery,omitempty" tf:"service_discovery,omitempty"`
}

type AppmeshVirtualNodeSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn             string `json:"arn" tf:"arn"`
	CreatedDate     string `json:"createdDate" tf:"created_date"`
	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`
	MeshName        string `json:"meshName" tf:"mesh_name"`
	Name            string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Spec []AppmeshVirtualNodeSpecSpec `json:"spec" tf:"spec"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AppmeshVirtualNodeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppmeshVirtualNodeSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppmeshVirtualNodeList is a list of AppmeshVirtualNodes
type AppmeshVirtualNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppmeshVirtualNode CRD objects
	Items []AppmeshVirtualNode `json:"items,omitempty"`
}
