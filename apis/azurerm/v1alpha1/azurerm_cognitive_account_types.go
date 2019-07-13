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

type AzurermCognitiveAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCognitiveAccountSpec   `json:"spec,omitempty"`
	Status            AzurermCognitiveAccountStatus `json:"status,omitempty"`
}

type AzurermCognitiveAccountSpecSku struct {
	Name string `json:"name"`
	Tier string `json:"tier"`
}

type AzurermCognitiveAccountSpec struct {
	PrimaryAccessKey   string                        `json:"primary_access_key"`
	ResourceGroupName  string                        `json:"resource_group_name"`
	Endpoint           string                        `json:"endpoint"`
	Kind               string                        `json:"kind"`
	Sku                []AzurermCognitiveAccountSpec `json:"sku"`
	Tags               map[string]string             `json:"tags"`
	SecondaryAccessKey string                        `json:"secondary_access_key"`
	Name               string                        `json:"name"`
	Location           string                        `json:"location"`
}



type AzurermCognitiveAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermCognitiveAccountList is a list of AzurermCognitiveAccounts
type AzurermCognitiveAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCognitiveAccount CRD objects
	Items []AzurermCognitiveAccount `json:"items,omitempty"`
}