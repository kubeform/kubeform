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

type AzurermStreamAnalyticsOutputMssql struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsOutputMssqlSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsOutputMssqlStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsOutputMssqlSpec struct {
	ResourceGroupName      string `json:"resource_group_name"`
	Server                 string `json:"server"`
	Database               string `json:"database"`
	Table                  string `json:"table"`
	User                   string `json:"user"`
	Password               string `json:"password"`
	Name                   string `json:"name"`
	StreamAnalyticsJobName string `json:"stream_analytics_job_name"`
}



type AzurermStreamAnalyticsOutputMssqlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermStreamAnalyticsOutputMssqlList is a list of AzurermStreamAnalyticsOutputMssqls
type AzurermStreamAnalyticsOutputMssqlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsOutputMssql CRD objects
	Items []AzurermStreamAnalyticsOutputMssql `json:"items,omitempty"`
}