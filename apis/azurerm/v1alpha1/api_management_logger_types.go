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

type ApiManagementLogger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementLoggerSpec   `json:"spec,omitempty"`
	Status            ApiManagementLoggerStatus `json:"status,omitempty"`
}

type ApiManagementLoggerSpecApplicationInsights struct {
	InstrumentationKey string `json:"instrumentation_key"`
}

type ApiManagementLoggerSpecEventhub struct {
	ConnectionString string `json:"connection_string"`
	Name             string `json:"name"`
}

type ApiManagementLoggerSpec struct {
	ApiManagementName string `json:"api_management_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ApplicationInsights *[]ApiManagementLoggerSpec `json:"application_insights,omitempty"`
	// +optional
	Buffered bool `json:"buffered,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Eventhub          *[]ApiManagementLoggerSpec `json:"eventhub,omitempty"`
	Name              string                     `json:"name"`
	ResourceGroupName string                     `json:"resource_group_name"`
}

type ApiManagementLoggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
