package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermSecurityCenterContact struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSecurityCenterContactSpec   `json:"spec,omitempty"`
	Status            AzurermSecurityCenterContactStatus `json:"status,omitempty"`
}

type AzurermSecurityCenterContactSpec struct {
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	AlertNotifications bool   `json:"alert_notifications"`
	AlertsToAdmins     bool   `json:"alerts_to_admins"`
}



type AzurermSecurityCenterContactStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermSecurityCenterContactList is a list of AzurermSecurityCenterContacts
type AzurermSecurityCenterContactList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSecurityCenterContact CRD objects
	Items []AzurermSecurityCenterContact `json:"items,omitempty"`
}