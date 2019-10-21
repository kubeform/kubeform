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
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type StreamAnalyticsFunctionJavascriptUdf struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsFunctionJavascriptUdfSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsFunctionJavascriptUdfStatus `json:"status,omitempty"`
}

type StreamAnalyticsFunctionJavascriptUdfSpecInput struct {
	Type string `json:"type" tf:"type"`
}

type StreamAnalyticsFunctionJavascriptUdfSpecOutput struct {
	Type string `json:"type" tf:"type"`
}

type StreamAnalyticsFunctionJavascriptUdfSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MinItems=1
	Input []StreamAnalyticsFunctionJavascriptUdfSpecInput `json:"input" tf:"input"`
	Name  string                                          `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	Output                 []StreamAnalyticsFunctionJavascriptUdfSpecOutput `json:"output" tf:"output"`
	ResourceGroupName      string                                           `json:"resourceGroupName" tf:"resource_group_name"`
	Script                 string                                           `json:"script" tf:"script"`
	StreamAnalyticsJobName string                                           `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name"`
}

type StreamAnalyticsFunctionJavascriptUdfStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StreamAnalyticsFunctionJavascriptUdfSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StreamAnalyticsFunctionJavascriptUdfList is a list of StreamAnalyticsFunctionJavascriptUdfs
type StreamAnalyticsFunctionJavascriptUdfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StreamAnalyticsFunctionJavascriptUdf CRD objects
	Items []StreamAnalyticsFunctionJavascriptUdf `json:"items,omitempty"`
}