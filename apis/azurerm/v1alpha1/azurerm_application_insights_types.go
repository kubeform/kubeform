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

type AzurermApplicationInsights struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApplicationInsightsSpec   `json:"spec,omitempty"`
	Status            AzurermApplicationInsightsStatus `json:"status,omitempty"`
}

type AzurermApplicationInsightsSpec struct {
	ResourceGroupName  string            `json:"resource_group_name"`
	Location           string            `json:"location"`
	ApplicationType    string            `json:"application_type"`
	Tags               map[string]string `json:"tags"`
	AppId              string            `json:"app_id"`
	InstrumentationKey string            `json:"instrumentation_key"`
	Name               string            `json:"name"`
}



type AzurermApplicationInsightsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApplicationInsightsList is a list of AzurermApplicationInsightss
type AzurermApplicationInsightsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApplicationInsights CRD objects
	Items []AzurermApplicationInsights `json:"items,omitempty"`
}