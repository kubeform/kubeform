package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDevspaceController struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDevspaceControllerSpec   `json:"spec,omitempty"`
	Status            AzurermDevspaceControllerStatus `json:"status,omitempty"`
}

type AzurermDevspaceControllerSpecSku struct {
	Name string `json:"name"`
	Tier string `json:"tier"`
}

type AzurermDevspaceControllerSpec struct {
	Name                                 string                          `json:"name"`
	Location                             string                          `json:"location"`
	Sku                                  []AzurermDevspaceControllerSpec `json:"sku"`
	HostSuffix                           string                          `json:"host_suffix"`
	TargetContainerHostResourceId        string                          `json:"target_container_host_resource_id"`
	TargetContainerHostCredentialsBase64 string                          `json:"target_container_host_credentials_base64"`
	Tags                                 map[string]string               `json:"tags"`
	ResourceGroupName                    string                          `json:"resource_group_name"`
	DataPlaneFqdn                        string                          `json:"data_plane_fqdn"`
}



type AzurermDevspaceControllerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDevspaceControllerList is a list of AzurermDevspaceControllers
type AzurermDevspaceControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDevspaceController CRD objects
	Items []AzurermDevspaceController `json:"items,omitempty"`
}