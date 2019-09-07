/*
Copyright 2019 The Kubeform Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// ComputeVPNTunnelLister helps list ComputeVPNTunnels.
type ComputeVPNTunnelLister interface {
	// List lists all ComputeVPNTunnels in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeVPNTunnel, err error)
	// ComputeVPNTunnels returns an object that can list and get ComputeVPNTunnels.
	ComputeVPNTunnels(namespace string) ComputeVPNTunnelNamespaceLister
	ComputeVPNTunnelListerExpansion
}

// computeVPNTunnelLister implements the ComputeVPNTunnelLister interface.
type computeVPNTunnelLister struct {
	indexer cache.Indexer
}

// NewComputeVPNTunnelLister returns a new ComputeVPNTunnelLister.
func NewComputeVPNTunnelLister(indexer cache.Indexer) ComputeVPNTunnelLister {
	return &computeVPNTunnelLister{indexer: indexer}
}

// List lists all ComputeVPNTunnels in the indexer.
func (s *computeVPNTunnelLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeVPNTunnel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeVPNTunnel))
	})
	return ret, err
}

// ComputeVPNTunnels returns an object that can list and get ComputeVPNTunnels.
func (s *computeVPNTunnelLister) ComputeVPNTunnels(namespace string) ComputeVPNTunnelNamespaceLister {
	return computeVPNTunnelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeVPNTunnelNamespaceLister helps list and get ComputeVPNTunnels.
type ComputeVPNTunnelNamespaceLister interface {
	// List lists all ComputeVPNTunnels in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeVPNTunnel, err error)
	// Get retrieves the ComputeVPNTunnel from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeVPNTunnel, error)
	ComputeVPNTunnelNamespaceListerExpansion
}

// computeVPNTunnelNamespaceLister implements the ComputeVPNTunnelNamespaceLister
// interface.
type computeVPNTunnelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeVPNTunnels in the indexer for a given namespace.
func (s computeVPNTunnelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeVPNTunnel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeVPNTunnel))
	})
	return ret, err
}

// Get retrieves the ComputeVPNTunnel from the indexer for a given namespace and name.
func (s computeVPNTunnelNamespaceLister) Get(name string) (*v1alpha1.ComputeVPNTunnel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computevpntunnel"), name)
	}
	return obj.(*v1alpha1.ComputeVPNTunnel), nil
}