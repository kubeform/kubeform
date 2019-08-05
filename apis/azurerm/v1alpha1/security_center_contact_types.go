package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SecurityCenterContact struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityCenterContactSpec   `json:"spec,omitempty"`
	Status            SecurityCenterContactStatus `json:"status,omitempty"`
}

type SecurityCenterContactSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AlertNotifications bool   `json:"alertNotifications" tf:"alert_notifications"`
	AlertsToAdmins     bool   `json:"alertsToAdmins" tf:"alerts_to_admins"`
	Email              string `json:"email" tf:"email"`
	Phone              string `json:"phone" tf:"phone"`
}

type SecurityCenterContactStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SecurityCenterContactSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SecurityCenterContactList is a list of SecurityCenterContacts
type SecurityCenterContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SecurityCenterContact CRD objects
	Items []SecurityCenterContact `json:"items,omitempty"`
}
