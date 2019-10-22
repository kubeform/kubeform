package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DataFactoryPipeline struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryPipelineSpec   `json:"spec,omitempty"`
	Status            DataFactoryPipelineStatus `json:"status,omitempty"`
}

type DataFactoryPipelineSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Annotations     []string `json:"annotations,omitempty" tf:"annotations,omitempty"`
	DataFactoryName string   `json:"dataFactoryName" tf:"data_factory_name"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Parameters        map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	ResourceGroupName string            `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Variables map[string]string `json:"variables,omitempty" tf:"variables,omitempty"`
}

type DataFactoryPipelineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataFactoryPipelineSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryPipelineList is a list of DataFactoryPipelines
type DataFactoryPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactoryPipeline CRD objects
	Items []DataFactoryPipeline `json:"items,omitempty"`
}
