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

type AwsAppmeshVirtualService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppmeshVirtualServiceSpec   `json:"spec,omitempty"`
	Status            AwsAppmeshVirtualServiceStatus `json:"status,omitempty"`
}

type AwsAppmeshVirtualServiceSpecSpecProviderVirtualNode struct {
	VirtualNodeName string `json:"virtual_node_name"`
}

type AwsAppmeshVirtualServiceSpecSpecProviderVirtualRouter struct {
	VirtualRouterName string `json:"virtual_router_name"`
}

type AwsAppmeshVirtualServiceSpecSpecProvider struct {
	VirtualNode   []AwsAppmeshVirtualServiceSpecSpecProvider `json:"virtual_node"`
	VirtualRouter []AwsAppmeshVirtualServiceSpecSpecProvider `json:"virtual_router"`
}

type AwsAppmeshVirtualServiceSpecSpec struct {
	Provider []AwsAppmeshVirtualServiceSpecSpec `json:"provider"`
}

type AwsAppmeshVirtualServiceSpec struct {
	MeshName        string                         `json:"mesh_name"`
	Spec            []AwsAppmeshVirtualServiceSpec `json:"spec"`
	Arn             string                         `json:"arn"`
	CreatedDate     string                         `json:"created_date"`
	LastUpdatedDate string                         `json:"last_updated_date"`
	Tags            map[string]string              `json:"tags"`
	Name            string                         `json:"name"`
}



type AwsAppmeshVirtualServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppmeshVirtualServiceList is a list of AwsAppmeshVirtualServices
type AwsAppmeshVirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppmeshVirtualService CRD objects
	Items []AwsAppmeshVirtualService `json:"items,omitempty"`
}