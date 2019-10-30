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

type ComputeHealthCheck struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeHealthCheckSpec   `json:"spec,omitempty"`
	Status            ComputeHealthCheckStatus `json:"status,omitempty"`
}

type ComputeHealthCheckSpecHttpHealthCheck struct {
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	RequestPath string `json:"requestPath,omitempty" tf:"request_path,omitempty"`
	// +optional
	Response string `json:"response,omitempty" tf:"response,omitempty"`
}

type ComputeHealthCheckSpecHttpsHealthCheck struct {
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	RequestPath string `json:"requestPath,omitempty" tf:"request_path,omitempty"`
	// +optional
	Response string `json:"response,omitempty" tf:"response,omitempty"`
}

type ComputeHealthCheckSpecSslHealthCheck struct {
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	Request string `json:"request,omitempty" tf:"request,omitempty"`
	// +optional
	Response string `json:"response,omitempty" tf:"response,omitempty"`
}

type ComputeHealthCheckSpecTcpHealthCheck struct {
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	Request string `json:"request,omitempty" tf:"request,omitempty"`
	// +optional
	Response string `json:"response,omitempty" tf:"response,omitempty"`
}

type ComputeHealthCheckSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CheckIntervalSec int `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	HealthyThreshold int `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpHealthCheck []ComputeHealthCheckSpecHttpHealthCheck `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HttpsHealthCheck []ComputeHealthCheckSpecHttpsHealthCheck `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`
	Name             string                                   `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SslHealthCheck []ComputeHealthCheckSpecSslHealthCheck `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	TcpHealthCheck []ComputeHealthCheckSpecTcpHealthCheck `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`
	// +optional
	TimeoutSec int `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	UnhealthyThreshold int `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type ComputeHealthCheckStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeHealthCheckSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeHealthCheckList is a list of ComputeHealthChecks
type ComputeHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeHealthCheck CRD objects
	Items []ComputeHealthCheck `json:"items,omitempty"`
}
