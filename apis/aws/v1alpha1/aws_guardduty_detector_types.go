package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGuarddutyDetector struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGuarddutyDetectorSpec   `json:"spec,omitempty"`
	Status            AwsGuarddutyDetectorStatus `json:"status,omitempty"`
}

type AwsGuarddutyDetectorSpec struct {
	Enable                     bool   `json:"enable"`
	AccountId                  string `json:"account_id"`
	FindingPublishingFrequency string `json:"finding_publishing_frequency"`
}

type AwsGuarddutyDetectorStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGuarddutyDetectorList is a list of AwsGuarddutyDetectors
type AwsGuarddutyDetectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGuarddutyDetector CRD objects
	Items []AwsGuarddutyDetector `json:"items,omitempty"`
}
