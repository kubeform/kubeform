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

type DevspaceController struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevspaceControllerSpec   `json:"spec,omitempty"`
	Status            DevspaceControllerStatus `json:"status,omitempty"`
}

type DevspaceControllerSpecSku struct {
	Name string `json:"name"`
	Tier string `json:"tier"`
}

type DevspaceControllerSpec struct {
	HostSuffix        string `json:"host_suffix"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku                                  []DevspaceControllerSpec `json:"sku"`
	TargetContainerHostCredentialsBase64 string                   `json:"target_container_host_credentials_base64"`
	TargetContainerHostResourceId        string                   `json:"target_container_host_resource_id"`
}

type DevspaceControllerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevspaceControllerList is a list of DevspaceControllers
type DevspaceControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DevspaceController CRD objects
	Items []DevspaceController `json:"items,omitempty"`
}
