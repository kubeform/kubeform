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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IotPolicyLister helps list IotPolicies.
type IotPolicyLister interface {
	// List lists all IotPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IotPolicy, err error)
	// IotPolicies returns an object that can list and get IotPolicies.
	IotPolicies(namespace string) IotPolicyNamespaceLister
	IotPolicyListerExpansion
}

// iotPolicyLister implements the IotPolicyLister interface.
type iotPolicyLister struct {
	indexer cache.Indexer
}

// NewIotPolicyLister returns a new IotPolicyLister.
func NewIotPolicyLister(indexer cache.Indexer) IotPolicyLister {
	return &iotPolicyLister{indexer: indexer}
}

// List lists all IotPolicies in the indexer.
func (s *iotPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.IotPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotPolicy))
	})
	return ret, err
}

// IotPolicies returns an object that can list and get IotPolicies.
func (s *iotPolicyLister) IotPolicies(namespace string) IotPolicyNamespaceLister {
	return iotPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IotPolicyNamespaceLister helps list and get IotPolicies.
type IotPolicyNamespaceLister interface {
	// List lists all IotPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IotPolicy, err error)
	// Get retrieves the IotPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IotPolicy, error)
	IotPolicyNamespaceListerExpansion
}

// iotPolicyNamespaceLister implements the IotPolicyNamespaceLister
// interface.
type iotPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IotPolicies in the indexer for a given namespace.
func (s iotPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IotPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IotPolicy))
	})
	return ret, err
}

// Get retrieves the IotPolicy from the indexer for a given namespace and name.
func (s iotPolicyNamespaceLister) Get(name string) (*v1alpha1.IotPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iotpolicy"), name)
	}
	return obj.(*v1alpha1.IotPolicy), nil
}
