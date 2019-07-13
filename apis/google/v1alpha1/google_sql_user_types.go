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

type GoogleSqlUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSqlUserSpec   `json:"spec,omitempty"`
	Status            GoogleSqlUserStatus `json:"status,omitempty"`
}

type GoogleSqlUserSpec struct {
	Password string `json:"password"`
	Project  string `json:"project"`
	Host     string `json:"host"`
	Instance string `json:"instance"`
	Name     string `json:"name"`
}



type GoogleSqlUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSqlUserList is a list of GoogleSqlUsers
type GoogleSqlUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSqlUser CRD objects
	Items []GoogleSqlUser `json:"items,omitempty"`
}