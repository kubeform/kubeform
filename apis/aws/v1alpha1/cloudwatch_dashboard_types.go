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

type CloudwatchDashboard struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchDashboardSpec   `json:"spec,omitempty"`
	Status            CloudwatchDashboardStatus `json:"status,omitempty"`
}

type CloudwatchDashboardSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DashboardArn  string `json:"dashboardArn,omitempty" tf:"dashboard_arn,omitempty"`
	DashboardBody string `json:"dashboardBody" tf:"dashboard_body"`
	DashboardName string `json:"dashboardName" tf:"dashboard_name"`
}

type CloudwatchDashboardStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CloudwatchDashboardSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchDashboardList is a list of CloudwatchDashboards
type CloudwatchDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchDashboard CRD objects
	Items []CloudwatchDashboard `json:"items,omitempty"`
}
