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

type AzurermAzureadApplication struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAzureadApplicationSpec   `json:"spec,omitempty"`
	Status            AzurermAzureadApplicationStatus `json:"status,omitempty"`
}

type AzurermAzureadApplicationSpec struct {
	AvailableToOtherTenants bool     `json:"available_to_other_tenants"`
	Oauth2AllowImplicitFlow bool     `json:"oauth2_allow_implicit_flow"`
	ApplicationId           string   `json:"application_id"`
	Name                    string   `json:"name"`
	Homepage                string   `json:"homepage"`
	IdentifierUris          []string `json:"identifier_uris"`
	ReplyUrls               []string `json:"reply_urls"`
}



type AzurermAzureadApplicationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAzureadApplicationList is a list of AzurermAzureadApplications
type AzurermAzureadApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAzureadApplication CRD objects
	Items []AzurermAzureadApplication `json:"items,omitempty"`
}