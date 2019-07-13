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

type AwsIotThingPrincipalAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIotThingPrincipalAttachmentSpec   `json:"spec,omitempty"`
	Status            AwsIotThingPrincipalAttachmentStatus `json:"status,omitempty"`
}

type AwsIotThingPrincipalAttachmentSpec struct {
	Principal string `json:"principal"`
	Thing     string `json:"thing"`
}



type AwsIotThingPrincipalAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIotThingPrincipalAttachmentList is a list of AwsIotThingPrincipalAttachments
type AwsIotThingPrincipalAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIotThingPrincipalAttachment CRD objects
	Items []AwsIotThingPrincipalAttachment `json:"items,omitempty"`
}