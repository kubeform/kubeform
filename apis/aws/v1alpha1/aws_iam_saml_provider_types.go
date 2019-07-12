package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamSamlProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamSamlProviderSpec   `json:"spec,omitempty"`
	Status            AwsIamSamlProviderStatus `json:"status,omitempty"`
}

type AwsIamSamlProviderSpec struct {
	Arn                  string `json:"arn"`
	ValidUntil           string `json:"valid_until"`
	Name                 string `json:"name"`
	SamlMetadataDocument string `json:"saml_metadata_document"`
}

type AwsIamSamlProviderStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamSamlProviderList is a list of AwsIamSamlProviders
type AwsIamSamlProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamSamlProvider CRD objects
	Items []AwsIamSamlProvider `json:"items,omitempty"`
}
