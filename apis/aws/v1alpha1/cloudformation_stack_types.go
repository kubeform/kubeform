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

type CloudformationStack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudformationStackSpec   `json:"spec,omitempty"`
	Status            CloudformationStackStatus `json:"status,omitempty"`
}

type CloudformationStackSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Capabilities []string `json:"capabilities,omitempty"`
	// +optional
	DisableRollback bool `json:"disable_rollback,omitempty"`
	// +optional
	IamRoleArn string `json:"iam_role_arn,omitempty"`
	Name       string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	NotificationArns []string `json:"notification_arns,omitempty"`
	// +optional
	OnFailure string `json:"on_failure,omitempty"`
	// +optional
	PolicyUrl string `json:"policy_url,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	TemplateUrl string `json:"template_url,omitempty"`
	// +optional
	TimeoutInMinutes int `json:"timeout_in_minutes,omitempty"`
}

type CloudformationStackStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudformationStackList is a list of CloudformationStacks
type CloudformationStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudformationStack CRD objects
	Items []CloudformationStack `json:"items,omitempty"`
}
