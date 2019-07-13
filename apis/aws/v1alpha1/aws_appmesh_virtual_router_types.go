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

type AwsAppmeshVirtualRouter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppmeshVirtualRouterSpec   `json:"spec,omitempty"`
	Status            AwsAppmeshVirtualRouterStatus `json:"status,omitempty"`
}

type AwsAppmeshVirtualRouterSpecSpecListenerPortMapping struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
}

type AwsAppmeshVirtualRouterSpecSpecListener struct {
	PortMapping []AwsAppmeshVirtualRouterSpecSpecListener `json:"port_mapping"`
}

type AwsAppmeshVirtualRouterSpecSpec struct {
	ServiceNames []string                          `json:"service_names"`
	Listener     []AwsAppmeshVirtualRouterSpecSpec `json:"listener"`
}

type AwsAppmeshVirtualRouterSpec struct {
	Spec            []AwsAppmeshVirtualRouterSpec `json:"spec"`
	Arn             string                        `json:"arn"`
	CreatedDate     string                        `json:"created_date"`
	LastUpdatedDate string                        `json:"last_updated_date"`
	Tags            map[string]string             `json:"tags"`
	Name            string                        `json:"name"`
	MeshName        string                        `json:"mesh_name"`
}



type AwsAppmeshVirtualRouterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppmeshVirtualRouterList is a list of AwsAppmeshVirtualRouters
type AwsAppmeshVirtualRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppmeshVirtualRouter CRD objects
	Items []AwsAppmeshVirtualRouter `json:"items,omitempty"`
}