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

type AwsServicecatalogPortfolio struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsServicecatalogPortfolioSpec   `json:"spec,omitempty"`
	Status            AwsServicecatalogPortfolioStatus `json:"status,omitempty"`
}

type AwsServicecatalogPortfolioSpec struct {
	ProviderName string            `json:"provider_name"`
	Tags         map[string]string `json:"tags"`
	Arn          string            `json:"arn"`
	CreatedTime  string            `json:"created_time"`
	Name         string            `json:"name"`
	Description  string            `json:"description"`
}

type AwsServicecatalogPortfolioStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsServicecatalogPortfolioList is a list of AwsServicecatalogPortfolios
type AwsServicecatalogPortfolioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsServicecatalogPortfolio CRD objects
	Items []AwsServicecatalogPortfolio `json:"items,omitempty"`
}
