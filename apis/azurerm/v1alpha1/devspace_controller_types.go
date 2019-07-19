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

type DevspaceController struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevspaceControllerSpec   `json:"spec,omitempty"`
	Status            DevspaceControllerStatus `json:"status,omitempty"`
}

type DevspaceControllerSpecSku struct {
	Name string `json:"name" tf:"name"`
	Tier string `json:"tier" tf:"tier"`
}

type DevspaceControllerSpec struct {
	HostSuffix        string `json:"hostSuffix" tf:"host_suffix"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []DevspaceControllerSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// Sensitive Data. Provide secret name which contains one value only
	TargetContainerHostCredentialsBase64 core.LocalObjectReference `json:"targetContainerHostCredentialsBase64" tf:"target_container_host_credentials_base64"`
	TargetContainerHostResourceID        string                    `json:"targetContainerHostResourceID" tf:"target_container_host_resource_id"`
	ProviderRef                          core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DevspaceControllerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
