/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type KeyVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultSpec   `json:"spec,omitempty"`
	Status            KeyVaultStatus `json:"status,omitempty"`
}

type KeyVaultSpecAccessPolicy struct {
	// +optional
	ApplicationID string `json:"applicationID,omitempty" tf:"application_id,omitempty"`
	// +optional
	CertificatePermissions []string `json:"certificatePermissions,omitempty" tf:"certificate_permissions,omitempty"`
	// +optional
	KeyPermissions []string `json:"keyPermissions,omitempty" tf:"key_permissions,omitempty"`
	ObjectID       string   `json:"objectID" tf:"object_id"`
	// +optional
	SecretPermissions []string `json:"secretPermissions,omitempty" tf:"secret_permissions,omitempty"`
	// +optional
	StoragePermissions []string `json:"storagePermissions,omitempty" tf:"storage_permissions,omitempty"`
	TenantID           string   `json:"tenantID" tf:"tenant_id"`
}

type KeyVaultSpecNetworkAcls struct {
	Bypass        string `json:"bypass" tf:"bypass"`
	DefaultAction string `json:"defaultAction" tf:"default_action"`
	// +optional
	IpRules []string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`
	// +optional
	VirtualNetworkSubnetIDS []string `json:"virtualNetworkSubnetIDS,omitempty" tf:"virtual_network_subnet_ids,omitempty"`
}

type KeyVaultSpecSku struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
}

type KeyVaultSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1024
	AccessPolicy []KeyVaultSpecAccessPolicy `json:"accessPolicy,omitempty" tf:"access_policy,omitempty"`
	// +optional
	EnabledForDeployment bool `json:"enabledForDeployment,omitempty" tf:"enabled_for_deployment,omitempty"`
	// +optional
	EnabledForDiskEncryption bool `json:"enabledForDiskEncryption,omitempty" tf:"enabled_for_disk_encryption,omitempty"`
	// +optional
	EnabledForTemplateDeployment bool   `json:"enabledForTemplateDeployment,omitempty" tf:"enabled_for_template_deployment,omitempty"`
	Location                     string `json:"location" tf:"location"`
	Name                         string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkAcls       []KeyVaultSpecNetworkAcls `json:"networkAcls,omitempty" tf:"network_acls,omitempty"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	Sku []KeyVaultSpecSku `json:"sku,omitempty" tf:"sku,omitempty"`
	// +optional
	SkuName string `json:"skuName,omitempty" tf:"sku_name,omitempty"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TenantID string            `json:"tenantID" tf:"tenant_id"`
	// +optional
	VaultURI string `json:"vaultURI,omitempty" tf:"vault_uri,omitempty"`
}

type KeyVaultStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KeyVaultSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KeyVaultList is a list of KeyVaults
type KeyVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyVault CRD objects
	Items []KeyVault `json:"items,omitempty"`
}
