package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SesDomainMailFrom struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesDomainMailFromSpec   `json:"spec,omitempty"`
	Status            SesDomainMailFromStatus `json:"status,omitempty"`
}

type SesDomainMailFromSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	BehaviorOnMxFailure string `json:"behaviorOnMxFailure,omitempty" tf:"behavior_on_mx_failure,omitempty"`
	Domain              string `json:"domain" tf:"domain"`
	MailFromDomain      string `json:"mailFromDomain" tf:"mail_from_domain"`
}

type SesDomainMailFromStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
