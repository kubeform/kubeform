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

// ServiceAccountKeyLister helps list ServiceAccountKeys.
type ServiceAccountKeyLister interface {
	// List lists all ServiceAccountKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceAccountKey, err error)
	// ServiceAccountKeys returns an object that can list and get ServiceAccountKeys.
	ServiceAccountKeys(namespace string) ServiceAccountKeyNamespaceLister
	ServiceAccountKeyListerExpansion
}

// serviceAccountKeyLister implements the ServiceAccountKeyLister interface.
type serviceAccountKeyLister struct {
	indexer cache.Indexer
}

// NewServiceAccountKeyLister returns a new ServiceAccountKeyLister.
func NewServiceAccountKeyLister(indexer cache.Indexer) ServiceAccountKeyLister {
	return &serviceAccountKeyLister{indexer: indexer}
}

// List lists all ServiceAccountKeys in the indexer.
func (s *serviceAccountKeyLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceAccountKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceAccountKey))
	})
	return ret, err
}

// ServiceAccountKeys returns an object that can list and get ServiceAccountKeys.
func (s *serviceAccountKeyLister) ServiceAccountKeys(namespace string) ServiceAccountKeyNamespaceLister {
	return serviceAccountKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceAccountKeyNamespaceLister helps list and get ServiceAccountKeys.
type ServiceAccountKeyNamespaceLister interface {
	// List lists all ServiceAccountKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceAccountKey, err error)
	// Get retrieves the ServiceAccountKey from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ServiceAccountKey, error)
	ServiceAccountKeyNamespaceListerExpansion
}

// serviceAccountKeyNamespaceLister implements the ServiceAccountKeyNamespaceLister
// interface.
type serviceAccountKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceAccountKeys in the indexer for a given namespace.
func (s serviceAccountKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceAccountKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceAccountKey))
	})
	return ret, err
}

// Get retrieves the ServiceAccountKey from the indexer for a given namespace and name.
func (s serviceAccountKeyNamespaceLister) Get(name string) (*v1alpha1.ServiceAccountKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serviceaccountkey"), name)
	}
	return obj.(*v1alpha1.ServiceAccountKey), nil
}
