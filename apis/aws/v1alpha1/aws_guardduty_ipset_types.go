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

type AwsGuarddutyIpset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGuarddutyIpsetSpec   `json:"spec,omitempty"`
	Status            AwsGuarddutyIpsetStatus `json:"status,omitempty"`
}

type AwsGuarddutyIpsetSpec struct {
	Location   string `json:"location"`
	Activate   bool   `json:"activate"`
	DetectorId string `json:"detector_id"`
	Name       string `json:"name"`
	Format     string `json:"format"`
}

type AwsGuarddutyIpsetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGuarddutyIpsetList is a list of AwsGuarddutyIpsets
type AwsGuarddutyIpsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGuarddutyIpset CRD objects
	Items []AwsGuarddutyIpset `json:"items,omitempty"`
}
