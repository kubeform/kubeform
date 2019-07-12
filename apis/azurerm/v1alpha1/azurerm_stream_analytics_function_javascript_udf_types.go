package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStreamAnalyticsFunctionJavascriptUdf struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStreamAnalyticsFunctionJavascriptUdfSpec   `json:"spec,omitempty"`
	Status            AzurermStreamAnalyticsFunctionJavascriptUdfStatus `json:"status,omitempty"`
}

type AzurermStreamAnalyticsFunctionJavascriptUdfSpecInput struct {
	Type string `json:"type"`
}

type AzurermStreamAnalyticsFunctionJavascriptUdfSpecOutput struct {
	Type string `json:"type"`
}

type AzurermStreamAnalyticsFunctionJavascriptUdfSpec struct {
	StreamAnalyticsJobName string                                            `json:"stream_analytics_job_name"`
	ResourceGroupName      string                                            `json:"resource_group_name"`
	Input                  []AzurermStreamAnalyticsFunctionJavascriptUdfSpec `json:"input"`
	Output                 []AzurermStreamAnalyticsFunctionJavascriptUdfSpec `json:"output"`
	Script                 string                                            `json:"script"`
	Name                   string                                            `json:"name"`
}



type AzurermStreamAnalyticsFunctionJavascriptUdfStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStreamAnalyticsFunctionJavascriptUdfList is a list of AzurermStreamAnalyticsFunctionJavascriptUdfs
type AzurermStreamAnalyticsFunctionJavascriptUdfList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStreamAnalyticsFunctionJavascriptUdf CRD objects
	Items []AzurermStreamAnalyticsFunctionJavascriptUdf `json:"items,omitempty"`
}