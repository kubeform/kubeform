package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type ApiManagementLogger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementLoggerSpec   `json:"spec,omitempty"`
	Status            ApiManagementLoggerStatus `json:"status,omitempty"`
}

type ApiManagementLoggerSpecApplicationInsights struct {
	InstrumentationKey string `json:"-" sensitive:"true" tf:"instrumentation_key"`
}

type ApiManagementLoggerSpecEventhub struct {
	ConnectionString string `json:"-" sensitive:"true" tf:"connection_string"`
	Name             string `json:"name" tf:"name"`
}

type ApiManagementLoggerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ApplicationInsights []ApiManagementLoggerSpecApplicationInsights `json:"applicationInsights,omitempty" tf:"application_insights,omitempty"`
	// +optional
	Buffered bool `json:"buffered,omitempty" tf:"buffered,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Eventhub          []ApiManagementLoggerSpecEventhub `json:"eventhub,omitempty" tf:"eventhub,omitempty"`
	Name              string                            `json:"name" tf:"name"`
	ResourceGroupName string                            `json:"resourceGroupName" tf:"resource_group_name"`
}

type ApiManagementLoggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiManagementLoggerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementLoggerList is a list of ApiManagementLoggers
type ApiManagementLoggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementLogger CRD objects
	Items []ApiManagementLogger `json:"items,omitempty"`
}
