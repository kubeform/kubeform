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

type AwsCloudwatchDashboard struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchDashboardSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchDashboardStatus `json:"status,omitempty"`
}

type AwsCloudwatchDashboardSpec struct {
	DashboardName string `json:"dashboard_name"`
	DashboardArn  string `json:"dashboard_arn"`
	DashboardBody string `json:"dashboard_body"`
}



type AwsCloudwatchDashboardStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudwatchDashboardList is a list of AwsCloudwatchDashboards
type AwsCloudwatchDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchDashboard CRD objects
	Items []AwsCloudwatchDashboard `json:"items,omitempty"`
}