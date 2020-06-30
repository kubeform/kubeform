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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IapAppEngineServiceIamBindingLister helps list IapAppEngineServiceIamBindings.
type IapAppEngineServiceIamBindingLister interface {
	// List lists all IapAppEngineServiceIamBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IapAppEngineServiceIamBinding, err error)
	// IapAppEngineServiceIamBindings returns an object that can list and get IapAppEngineServiceIamBindings.
	IapAppEngineServiceIamBindings(namespace string) IapAppEngineServiceIamBindingNamespaceLister
	IapAppEngineServiceIamBindingListerExpansion
}

// iapAppEngineServiceIamBindingLister implements the IapAppEngineServiceIamBindingLister interface.
type iapAppEngineServiceIamBindingLister struct {
	indexer cache.Indexer
}

// NewIapAppEngineServiceIamBindingLister returns a new IapAppEngineServiceIamBindingLister.
func NewIapAppEngineServiceIamBindingLister(indexer cache.Indexer) IapAppEngineServiceIamBindingLister {
	return &iapAppEngineServiceIamBindingLister{indexer: indexer}
}

// List lists all IapAppEngineServiceIamBindings in the indexer.
func (s *iapAppEngineServiceIamBindingLister) List(selector labels.Selector) (ret []*v1alpha1.IapAppEngineServiceIamBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IapAppEngineServiceIamBinding))
	})
	return ret, err
}

// IapAppEngineServiceIamBindings returns an object that can list and get IapAppEngineServiceIamBindings.
func (s *iapAppEngineServiceIamBindingLister) IapAppEngineServiceIamBindings(namespace string) IapAppEngineServiceIamBindingNamespaceLister {
	return iapAppEngineServiceIamBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IapAppEngineServiceIamBindingNamespaceLister helps list and get IapAppEngineServiceIamBindings.
type IapAppEngineServiceIamBindingNamespaceLister interface {
	// List lists all IapAppEngineServiceIamBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IapAppEngineServiceIamBinding, err error)
	// Get retrieves the IapAppEngineServiceIamBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IapAppEngineServiceIamBinding, error)
	IapAppEngineServiceIamBindingNamespaceListerExpansion
}

// iapAppEngineServiceIamBindingNamespaceLister implements the IapAppEngineServiceIamBindingNamespaceLister
// interface.
type iapAppEngineServiceIamBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IapAppEngineServiceIamBindings in the indexer for a given namespace.
func (s iapAppEngineServiceIamBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IapAppEngineServiceIamBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IapAppEngineServiceIamBinding))
	})
	return ret, err
}

// Get retrieves the IapAppEngineServiceIamBinding from the indexer for a given namespace and name.
func (s iapAppEngineServiceIamBindingNamespaceLister) Get(name string) (*v1alpha1.IapAppEngineServiceIamBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iapappengineserviceiambinding"), name)
	}
	return obj.(*v1alpha1.IapAppEngineServiceIamBinding), nil
}
