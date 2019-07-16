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

type AppmeshVirtualNode struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppmeshVirtualNodeSpec   `json:"spec,omitempty"`
	Status            AppmeshVirtualNodeStatus `json:"status,omitempty"`
}

type AppmeshVirtualNodeSpecSpecBackendVirtualService struct {
	VirtualServiceName string `json:"virtual_service_name"`
}

type AppmeshVirtualNodeSpecSpecBackend struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VirtualService *[]AppmeshVirtualNodeSpecSpecBackend `json:"virtual_service,omitempty"`
}

type AppmeshVirtualNodeSpecSpecListenerHealthCheck struct {
	HealthyThreshold int `json:"healthy_threshold"`
	IntervalMillis   int `json:"interval_millis"`
	// +optional
	Path               string `json:"path,omitempty"`
	Protocol           string `json:"protocol"`
	TimeoutMillis      int    `json:"timeout_millis"`
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
}

type AppmeshVirtualNodeSpecSpecListenerPortMapping struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type AppmeshVirtualNodeSpecSpecListener struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheck *[]AppmeshVirtualNodeSpecSpecListener `json:"health_check,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	PortMapping []AppmeshVirtualNodeSpecSpecListener `json:"port_mapping"`
}

type AppmeshVirtualNodeSpecSpecLoggingAccessLogFile struct {
	Path string `json:"path"`
}

type AppmeshVirtualNodeSpecSpecLoggingAccessLog struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	File *[]AppmeshVirtualNodeSpecSpecLoggingAccessLog `json:"file,omitempty"`
}

type AppmeshVirtualNodeSpecSpecLogging struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccessLog *[]AppmeshVirtualNodeSpecSpecLogging `json:"access_log,omitempty"`
}

type AppmeshVirtualNodeSpecSpecServiceDiscoveryAwsCloudMap struct {
	// +optional
	Attributes    map[string]string `json:"attributes,omitempty"`
	NamespaceName string            `json:"namespace_name"`
	ServiceName   string            `json:"service_name"`
}

type AppmeshVirtualNodeSpecSpecServiceDiscoveryDns struct {
	Hostname string `json:"hostname"`
}

type AppmeshVirtualNodeSpecSpecServiceDiscovery struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AwsCloudMap *[]AppmeshVirtualNodeSpecSpecServiceDiscovery `json:"aws_cloud_map,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Dns *[]AppmeshVirtualNodeSpecSpecServiceDiscovery `json:"dns,omitempty"`
}

type AppmeshVirtualNodeSpecSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=25
	// +kubebuilder:validation:UniqueItems=true
	Backend *[]AppmeshVirtualNodeSpecSpec `json:"backend,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Listener *[]AppmeshVirtualNodeSpecSpec `json:"listener,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Logging *[]AppmeshVirtualNodeSpecSpec `json:"logging,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ServiceDiscovery *[]AppmeshVirtualNodeSpecSpec `json:"service_discovery,omitempty"`
}

type AppmeshVirtualNodeSpec struct {
	MeshName string `json:"mesh_name"`
	Name     string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Spec []AppmeshVirtualNodeSpec `json:"spec"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type AppmeshVirtualNodeStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
