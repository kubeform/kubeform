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

type AwsGuarddutyThreatintelset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGuarddutyThreatintelsetSpec   `json:"spec,omitempty"`
	Status            AwsGuarddutyThreatintelsetStatus `json:"status,omitempty"`
}

type AwsGuarddutyThreatintelsetSpec struct {
	DetectorId string `json:"detector_id"`
	Name       string `json:"name"`
	Format     string `json:"format"`
	Location   string `json:"location"`
	Activate   bool   `json:"activate"`
}

type AwsGuarddutyThreatintelsetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGuarddutyThreatintelsetList is a list of AwsGuarddutyThreatintelsets
type AwsGuarddutyThreatintelsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGuarddutyThreatintelset CRD objects
	Items []AwsGuarddutyThreatintelset `json:"items,omitempty"`
}
