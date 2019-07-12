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

type AwsDbParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbParameterGroupSpec   `json:"spec,omitempty"`
	Status            AwsDbParameterGroupStatus `json:"status,omitempty"`
}

type AwsDbParameterGroupSpecParameter struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	ApplyMethod string `json:"apply_method"`
}

type AwsDbParameterGroupSpec struct {
	Parameter   []AwsDbParameterGroupSpec `json:"parameter"`
	Tags        map[string]string         `json:"tags"`
	Arn         string                    `json:"arn"`
	Name        string                    `json:"name"`
	NamePrefix  string                    `json:"name_prefix"`
	Family      string                    `json:"family"`
	Description string                    `json:"description"`
}

type AwsDbParameterGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbParameterGroupList is a list of AwsDbParameterGroups
type AwsDbParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbParameterGroup CRD objects
	Items []AwsDbParameterGroup `json:"items,omitempty"`
}
