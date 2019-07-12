package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppmeshVirtualRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppmeshVirtualRouterSpec   `json:"spec,omitempty"`
	Status            AwsAppmeshVirtualRouterStatus `json:"status,omitempty"`
}

type AwsAppmeshVirtualRouterSpecSpecListenerPortMapping struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type AwsAppmeshVirtualRouterSpecSpecListener struct {
	PortMapping []AwsAppmeshVirtualRouterSpecSpecListener `json:"port_mapping"`
}

type AwsAppmeshVirtualRouterSpecSpec struct {
	ServiceNames []string                          `json:"service_names"`
	Listener     []AwsAppmeshVirtualRouterSpecSpec `json:"listener"`
}

type AwsAppmeshVirtualRouterSpec struct {
	LastUpdatedDate string                        `json:"last_updated_date"`
	Name            string                        `json:"name"`
	MeshName        string                        `json:"mesh_name"`
	Spec            []AwsAppmeshVirtualRouterSpec `json:"spec"`
	Arn             string                        `json:"arn"`
	CreatedDate     string                        `json:"created_date"`
}



type AwsAppmeshVirtualRouterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppmeshVirtualRouterList is a list of AwsAppmeshVirtualRouters
type AwsAppmeshVirtualRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppmeshVirtualRouter CRD objects
	Items []AwsAppmeshVirtualRouter `json:"items,omitempty"`
}