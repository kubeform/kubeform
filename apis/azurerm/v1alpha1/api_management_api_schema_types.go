package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiManagementAPISchema struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementAPISchemaSpec   `json:"spec,omitempty"`
	Status            ApiManagementAPISchemaStatus `json:"status,omitempty"`
}

type ApiManagementAPISchemaSpec struct {
	ApiManagementName string                    `json:"apiManagementName" tf:"api_management_name"`
	ApiName           string                    `json:"apiName" tf:"api_name"`
	ContentType       string                    `json:"contentType" tf:"content_type"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	SchemaID          string                    `json:"schemaID" tf:"schema_id"`
	Value             string                    `json:"value" tf:"value"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ApiManagementAPISchemaStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementAPISchemaList is a list of ApiManagementAPISchemas
type ApiManagementAPISchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementAPISchema CRD objects
	Items []ApiManagementAPISchema `json:"items,omitempty"`
}
