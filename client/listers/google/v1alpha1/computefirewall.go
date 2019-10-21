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

// ComputeFirewallLister helps list ComputeFirewalls.
type ComputeFirewallLister interface {
	// List lists all ComputeFirewalls in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeFirewall, err error)
	// ComputeFirewalls returns an object that can list and get ComputeFirewalls.
	ComputeFirewalls(namespace string) ComputeFirewallNamespaceLister
	ComputeFirewallListerExpansion
}

// computeFirewallLister implements the ComputeFirewallLister interface.
type computeFirewallLister struct {
	indexer cache.Indexer
}

// NewComputeFirewallLister returns a new ComputeFirewallLister.
func NewComputeFirewallLister(indexer cache.Indexer) ComputeFirewallLister {
	return &computeFirewallLister{indexer: indexer}
}

// List lists all ComputeFirewalls in the indexer.
func (s *computeFirewallLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeFirewall, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeFirewall))
	})
	return ret, err
}

// ComputeFirewalls returns an object that can list and get ComputeFirewalls.
func (s *computeFirewallLister) ComputeFirewalls(namespace string) ComputeFirewallNamespaceLister {
	return computeFirewallNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeFirewallNamespaceLister helps list and get ComputeFirewalls.
type ComputeFirewallNamespaceLister interface {
	// List lists all ComputeFirewalls in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeFirewall, err error)
	// Get retrieves the ComputeFirewall from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeFirewall, error)
	ComputeFirewallNamespaceListerExpansion
}

// computeFirewallNamespaceLister implements the ComputeFirewallNamespaceLister
// interface.
type computeFirewallNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeFirewalls in the indexer for a given namespace.
func (s computeFirewallNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeFirewall, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeFirewall))
	})
	return ret, err
}

// Get retrieves the ComputeFirewall from the indexer for a given namespace and name.
func (s computeFirewallNamespaceLister) Get(name string) (*v1alpha1.ComputeFirewall, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computefirewall"), name)
	}
	return obj.(*v1alpha1.ComputeFirewall), nil
}