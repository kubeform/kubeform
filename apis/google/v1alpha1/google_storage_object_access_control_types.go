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

type GoogleStorageObjectAccessControl struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageObjectAccessControlSpec   `json:"spec,omitempty"`
	Status            GoogleStorageObjectAccessControlStatus `json:"status,omitempty"`
}

type GoogleStorageObjectAccessControlSpecProjectTeam struct {
	ProjectNumber string `json:"project_number"`
	Team          string `json:"team"`
}

type GoogleStorageObjectAccessControlSpec struct {
	Bucket      string                                 `json:"bucket"`
	Entity      string                                 `json:"entity"`
	Role        string                                 `json:"role"`
	Domain      string                                 `json:"domain"`
	EntityId    string                                 `json:"entity_id"`
	Object      string                                 `json:"object"`
	Email       string                                 `json:"email"`
	Generation  int                                    `json:"generation"`
	ProjectTeam []GoogleStorageObjectAccessControlSpec `json:"project_team"`
}



type GoogleStorageObjectAccessControlStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleStorageObjectAccessControlList is a list of GoogleStorageObjectAccessControls
type GoogleStorageObjectAccessControlList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageObjectAccessControl CRD objects
	Items []GoogleStorageObjectAccessControl `json:"items,omitempty"`
}