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

type AzurermApiManagementLogger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementLoggerSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementLoggerStatus `json:"status,omitempty"`
}

type AzurermApiManagementLoggerSpecEventhub struct {
	Name             string `json:"name"`
	ConnectionString string `json:"connection_string"`
}

type AzurermApiManagementLoggerSpecApplicationInsights struct {
	InstrumentationKey string `json:"instrumentation_key"`
}

type AzurermApiManagementLoggerSpec struct {
	Name                string                           `json:"name"`
	ResourceGroupName   string                           `json:"resource_group_name"`
	ApiManagementName   string                           `json:"api_management_name"`
	Eventhub            []AzurermApiManagementLoggerSpec `json:"eventhub"`
	ApplicationInsights []AzurermApiManagementLoggerSpec `json:"application_insights"`
	Buffered            bool                             `json:"buffered"`
	Description         string                           `json:"description"`
}



type AzurermApiManagementLoggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementLoggerList is a list of AzurermApiManagementLoggers
type AzurermApiManagementLoggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementLogger CRD objects
	Items []AzurermApiManagementLogger `json:"items,omitempty"`
}