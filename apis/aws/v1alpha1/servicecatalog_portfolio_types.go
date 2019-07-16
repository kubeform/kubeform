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

type ServicecatalogPortfolio struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicecatalogPortfolioSpec   `json:"spec,omitempty"`
	Status            ServicecatalogPortfolioStatus `json:"status,omitempty"`
}

type ServicecatalogPortfolioSpec struct {
	Name string `json:"name"`
	// +optional
	ProviderName string `json:"provider_name,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type ServicecatalogPortfolioStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicecatalogPortfolioList is a list of ServicecatalogPortfolios
type ServicecatalogPortfolioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicecatalogPortfolio CRD objects
	Items []ServicecatalogPortfolio `json:"items,omitempty"`
}
