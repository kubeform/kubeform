package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type SesDomainMailFrom struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesDomainMailFromSpec   `json:"spec,omitempty"`
	Status            SesDomainMailFromStatus `json:"status,omitempty"`
}

type SesDomainMailFromSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BehaviorOnMxFailure string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`
	Domain              string `json:"domain" tf:"domain"`
	MailFromDomain      string `json:"mailFromDomain" tf:"mail_from_domain"`
}

type SesDomainMailFromStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SesDomainMailFromSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesDomainMailFromList is a list of SesDomainMailFroms
type SesDomainMailFromList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesDomainMailFrom CRD objects
	Items []SesDomainMailFrom `json:"items,omitempty"`
}
