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

// EipLister helps list Eips.
type EipLister interface {
	// List lists all Eips in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Eip, err error)
	// Eips returns an object that can list and get Eips.
	Eips(namespace string) EipNamespaceLister
	EipListerExpansion
}

// eipLister implements the EipLister interface.
type eipLister struct {
	indexer cache.Indexer
}

// NewEipLister returns a new EipLister.
func NewEipLister(indexer cache.Indexer) EipLister {
	return &eipLister{indexer: indexer}
}

// List lists all Eips in the indexer.
func (s *eipLister) List(selector labels.Selector) (ret []*v1alpha1.Eip, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Eip))
	})
	return ret, err
}

// Eips returns an object that can list and get Eips.
func (s *eipLister) Eips(namespace string) EipNamespaceLister {
	return eipNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EipNamespaceLister helps list and get Eips.
type EipNamespaceLister interface {
	// List lists all Eips in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Eip, err error)
	// Get retrieves the Eip from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Eip, error)
	EipNamespaceListerExpansion
}

// eipNamespaceLister implements the EipNamespaceLister
// interface.
type eipNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Eips in the indexer for a given namespace.
func (s eipNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Eip, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Eip))
	})
	return ret, err
}

// Get retrieves the Eip from the indexer for a given namespace and name.
func (s eipNamespaceLister) Get(name string) (*v1alpha1.Eip, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eip"), name)
	}
	return obj.(*v1alpha1.Eip), nil
}
