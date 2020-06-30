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

// ShieldProtectionLister helps list ShieldProtections.
type ShieldProtectionLister interface {
	// List lists all ShieldProtections in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ShieldProtection, err error)
	// ShieldProtections returns an object that can list and get ShieldProtections.
	ShieldProtections(namespace string) ShieldProtectionNamespaceLister
	ShieldProtectionListerExpansion
}

// shieldProtectionLister implements the ShieldProtectionLister interface.
type shieldProtectionLister struct {
	indexer cache.Indexer
}

// NewShieldProtectionLister returns a new ShieldProtectionLister.
func NewShieldProtectionLister(indexer cache.Indexer) ShieldProtectionLister {
	return &shieldProtectionLister{indexer: indexer}
}

// List lists all ShieldProtections in the indexer.
func (s *shieldProtectionLister) List(selector labels.Selector) (ret []*v1alpha1.ShieldProtection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShieldProtection))
	})
	return ret, err
}

// ShieldProtections returns an object that can list and get ShieldProtections.
func (s *shieldProtectionLister) ShieldProtections(namespace string) ShieldProtectionNamespaceLister {
	return shieldProtectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ShieldProtectionNamespaceLister helps list and get ShieldProtections.
type ShieldProtectionNamespaceLister interface {
	// List lists all ShieldProtections in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ShieldProtection, err error)
	// Get retrieves the ShieldProtection from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ShieldProtection, error)
	ShieldProtectionNamespaceListerExpansion
}

// shieldProtectionNamespaceLister implements the ShieldProtectionNamespaceLister
// interface.
type shieldProtectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ShieldProtections in the indexer for a given namespace.
func (s shieldProtectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ShieldProtection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShieldProtection))
	})
	return ret, err
}

// Get retrieves the ShieldProtection from the indexer for a given namespace and name.
func (s shieldProtectionNamespaceLister) Get(name string) (*v1alpha1.ShieldProtection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("shieldprotection"), name)
	}
	return obj.(*v1alpha1.ShieldProtection), nil
}
