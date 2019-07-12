package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDataFactoryPipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataFactoryPipelineSpec   `json:"spec,omitempty"`
	Status            AzurermDataFactoryPipelineStatus `json:"status,omitempty"`
}

type AzurermDataFactoryPipelineSpec struct {
	Name              string            `json:"name"`
	DataFactoryName   string            `json:"data_factory_name"`
	ResourceGroupName string            `json:"resource_group_name"`
	Parameters        map[string]string `json:"parameters"`
	Variables         map[string]string `json:"variables"`
	Description       string            `json:"description"`
	Annotations       []string          `json:"annotations"`
}

type AzurermDataFactoryPipelineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDataFactoryPipelineList is a list of AzurermDataFactoryPipelines
type AzurermDataFactoryPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataFactoryPipeline CRD objects
	Items []AzurermDataFactoryPipeline `json:"items,omitempty"`
}
