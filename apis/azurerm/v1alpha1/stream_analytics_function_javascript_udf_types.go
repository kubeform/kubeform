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

type StreamAnalyticsFunctionJavascriptUdf struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsFunctionJavascriptUdfSpec   `json:"spec,omitempty"`
	Status            StreamAnalyticsFunctionJavascriptUdfStatus `json:"status,omitempty"`
}

type StreamAnalyticsFunctionJavascriptUdfSpecInput struct {
	Type string `json:"type"`
}

type StreamAnalyticsFunctionJavascriptUdfSpecOutput struct {
	Type string `json:"type"`
}

type StreamAnalyticsFunctionJavascriptUdfSpec struct {
	// +kubebuilder:validation:MinItems=1
	Input []StreamAnalyticsFunctionJavascriptUdfSpec `json:"input"`
	Name  string                                     `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	Output                 []StreamAnalyticsFunctionJavascriptUdfSpec `json:"output"`
	ResourceGroupName      string                                     `json:"resource_group_name"`
	Script                 string                                     `json:"script"`
	StreamAnalyticsJobName string                                     `json:"stream_analytics_job_name"`
}

type StreamAnalyticsFunctionJavascriptUdfStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
