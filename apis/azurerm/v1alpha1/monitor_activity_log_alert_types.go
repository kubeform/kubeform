package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MonitorActivityLogAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActivityLogAlertSpec   `json:"spec,omitempty"`
	Status            MonitorActivityLogAlertStatus `json:"status,omitempty"`
}

type MonitorActivityLogAlertSpecAction struct {
	ActionGroupID string `json:"actionGroupID" tf:"action_group_id"`
	// +optional
	WebhookProperties map[string]string `json:"webhookProperties,omitempty" tf:"webhook_properties,omitempty"`
}

type MonitorActivityLogAlertSpecCriteria struct {
	// +optional
	Caller   string `json:"caller,omitempty" tf:"caller,omitempty"`
	Category string `json:"category" tf:"category"`
	// +optional
	Level string `json:"level,omitempty" tf:"level,omitempty"`
	// +optional
	OperationName string `json:"operationName,omitempty" tf:"operation_name,omitempty"`
	// +optional
	ResourceGroup string `json:"resourceGroup,omitempty" tf:"resource_group,omitempty"`
	// +optional
	ResourceID string `json:"resourceID,omitempty" tf:"resource_id,omitempty"`
	// +optional
	ResourceProvider string `json:"resourceProvider,omitempty" tf:"resource_provider,omitempty"`
	// +optional
	ResourceType string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	SubStatus string `json:"subStatus,omitempty" tf:"sub_status,omitempty"`
}

type MonitorActivityLogAlertSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Action []MonitorActivityLogAlertSpecAction `json:"action,omitempty" tf:"action,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Criteria []MonitorActivityLogAlertSpecCriteria `json:"criteria" tf:"criteria"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled           bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Scopes []string `json:"scopes" tf:"scopes"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActivityLogAlertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorActivityLogAlertSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorActivityLogAlertList is a list of MonitorActivityLogAlerts
type MonitorActivityLogAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorActivityLogAlert CRD objects
	Items []MonitorActivityLogAlert `json:"items,omitempty"`
}
