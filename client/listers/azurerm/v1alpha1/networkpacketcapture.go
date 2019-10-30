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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkPacketCaptureLister helps list NetworkPacketCaptures.
type NetworkPacketCaptureLister interface {
	// List lists all NetworkPacketCaptures in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkPacketCapture, err error)
	// NetworkPacketCaptures returns an object that can list and get NetworkPacketCaptures.
	NetworkPacketCaptures(namespace string) NetworkPacketCaptureNamespaceLister
	NetworkPacketCaptureListerExpansion
}

// networkPacketCaptureLister implements the NetworkPacketCaptureLister interface.
type networkPacketCaptureLister struct {
	indexer cache.Indexer
}

// NewNetworkPacketCaptureLister returns a new NetworkPacketCaptureLister.
func NewNetworkPacketCaptureLister(indexer cache.Indexer) NetworkPacketCaptureLister {
	return &networkPacketCaptureLister{indexer: indexer}
}

// List lists all NetworkPacketCaptures in the indexer.
func (s *networkPacketCaptureLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkPacketCapture, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkPacketCapture))
	})
	return ret, err
}

// NetworkPacketCaptures returns an object that can list and get NetworkPacketCaptures.
func (s *networkPacketCaptureLister) NetworkPacketCaptures(namespace string) NetworkPacketCaptureNamespaceLister {
	return networkPacketCaptureNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkPacketCaptureNamespaceLister helps list and get NetworkPacketCaptures.
type NetworkPacketCaptureNamespaceLister interface {
	// List lists all NetworkPacketCaptures in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkPacketCapture, err error)
	// Get retrieves the NetworkPacketCapture from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkPacketCapture, error)
	NetworkPacketCaptureNamespaceListerExpansion
}

// networkPacketCaptureNamespaceLister implements the NetworkPacketCaptureNamespaceLister
// interface.
type networkPacketCaptureNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkPacketCaptures in the indexer for a given namespace.
func (s networkPacketCaptureNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkPacketCapture, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkPacketCapture))
	})
	return ret, err
}

// Get retrieves the NetworkPacketCapture from the indexer for a given namespace and name.
func (s networkPacketCaptureNamespaceLister) Get(name string) (*v1alpha1.NetworkPacketCapture, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkpacketcapture"), name)
	}
	return obj.(*v1alpha1.NetworkPacketCapture), nil
}
