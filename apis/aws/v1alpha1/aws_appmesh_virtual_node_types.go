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

type AwsAppmeshVirtualNode struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppmeshVirtualNodeSpec   `json:"spec,omitempty"`
	Status            AwsAppmeshVirtualNodeStatus `json:"status,omitempty"`
}

type AwsAppmeshVirtualNodeSpecSpecBackendVirtualService struct {
	VirtualServiceName string `json:"virtual_service_name"`
}

type AwsAppmeshVirtualNodeSpecSpecBackend struct {
	VirtualService []AwsAppmeshVirtualNodeSpecSpecBackend `json:"virtual_service"`
}

type AwsAppmeshVirtualNodeSpecSpecListenerHealthCheck struct {
	UnhealthyThreshold int    `json:"unhealthy_threshold"`
	HealthyThreshold   int    `json:"healthy_threshold"`
	IntervalMillis     int    `json:"interval_millis"`
	Path               string `json:"path"`
	Port               int    `json:"port"`
	Protocol           string `json:"protocol"`
	TimeoutMillis      int    `json:"timeout_millis"`
}

type AwsAppmeshVirtualNodeSpecSpecListenerPortMapping struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type AwsAppmeshVirtualNodeSpecSpecListener struct {
	HealthCheck []AwsAppmeshVirtualNodeSpecSpecListener `json:"health_check"`
	PortMapping []AwsAppmeshVirtualNodeSpecSpecListener `json:"port_mapping"`
}

type AwsAppmeshVirtualNodeSpecSpecLoggingAccessLogFile struct {
	Path string `json:"path"`
}

type AwsAppmeshVirtualNodeSpecSpecLoggingAccessLog struct {
	File []AwsAppmeshVirtualNodeSpecSpecLoggingAccessLog `json:"file"`
}

type AwsAppmeshVirtualNodeSpecSpecLogging struct {
	AccessLog []AwsAppmeshVirtualNodeSpecSpecLogging `json:"access_log"`
}

type AwsAppmeshVirtualNodeSpecSpecServiceDiscoveryAwsCloudMap struct {
	NamespaceName string            `json:"namespace_name"`
	ServiceName   string            `json:"service_name"`
	Attributes    map[string]string `json:"attributes"`
}

type AwsAppmeshVirtualNodeSpecSpecServiceDiscoveryDns struct {
	ServiceName string `json:"service_name"`
	Hostname    string `json:"hostname"`
}

type AwsAppmeshVirtualNodeSpecSpecServiceDiscovery struct {
	AwsCloudMap []AwsAppmeshVirtualNodeSpecSpecServiceDiscovery `json:"aws_cloud_map"`
	Dns         []AwsAppmeshVirtualNodeSpecSpecServiceDiscovery `json:"dns"`
}

type AwsAppmeshVirtualNodeSpecSpec struct {
	Backends         []string                        `json:"backends"`
	Backend          []AwsAppmeshVirtualNodeSpecSpec `json:"backend"`
	Listener         []AwsAppmeshVirtualNodeSpecSpec `json:"listener"`
	Logging          []AwsAppmeshVirtualNodeSpecSpec `json:"logging"`
	ServiceDiscovery []AwsAppmeshVirtualNodeSpecSpec `json:"service_discovery"`
}

type AwsAppmeshVirtualNodeSpec struct {
	CreatedDate     string                      `json:"created_date"`
	LastUpdatedDate string                      `json:"last_updated_date"`
	Tags            map[string]string           `json:"tags"`
	Name            string                      `json:"name"`
	MeshName        string                      `json:"mesh_name"`
	Spec            []AwsAppmeshVirtualNodeSpec `json:"spec"`
	Arn             string                      `json:"arn"`
}

type AwsAppmeshVirtualNodeStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppmeshVirtualNodeList is a list of AwsAppmeshVirtualNodes
type AwsAppmeshVirtualNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppmeshVirtualNode CRD objects
	Items []AwsAppmeshVirtualNode `json:"items,omitempty"`
}
