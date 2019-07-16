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

type AppService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceSpec   `json:"spec,omitempty"`
	Status            AppServiceStatus `json:"status,omitempty"`
}

type AppServiceSpec struct {
	AppServicePlanId string `json:"app_service_plan_id"`
	// +optional
	ClientCertEnabled bool `json:"client_cert_enabled,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	HttpsOnly         bool   `json:"https_only,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type AppServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceList is a list of AppServices
type AppServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppService CRD objects
	Items []AppService `json:"items,omitempty"`
}
