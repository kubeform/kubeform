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

// RedshiftParameterGroupLister helps list RedshiftParameterGroups.
type RedshiftParameterGroupLister interface {
	// List lists all RedshiftParameterGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RedshiftParameterGroup, err error)
	// RedshiftParameterGroups returns an object that can list and get RedshiftParameterGroups.
	RedshiftParameterGroups(namespace string) RedshiftParameterGroupNamespaceLister
	RedshiftParameterGroupListerExpansion
}

// redshiftParameterGroupLister implements the RedshiftParameterGroupLister interface.
type redshiftParameterGroupLister struct {
	indexer cache.Indexer
}

// NewRedshiftParameterGroupLister returns a new RedshiftParameterGroupLister.
func NewRedshiftParameterGroupLister(indexer cache.Indexer) RedshiftParameterGroupLister {
	return &redshiftParameterGroupLister{indexer: indexer}
}

// List lists all RedshiftParameterGroups in the indexer.
func (s *redshiftParameterGroupLister) List(selector labels.Selector) (ret []*v1alpha1.RedshiftParameterGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedshiftParameterGroup))
	})
	return ret, err
}

// RedshiftParameterGroups returns an object that can list and get RedshiftParameterGroups.
func (s *redshiftParameterGroupLister) RedshiftParameterGroups(namespace string) RedshiftParameterGroupNamespaceLister {
	return redshiftParameterGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RedshiftParameterGroupNamespaceLister helps list and get RedshiftParameterGroups.
type RedshiftParameterGroupNamespaceLister interface {
	// List lists all RedshiftParameterGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RedshiftParameterGroup, err error)
	// Get retrieves the RedshiftParameterGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RedshiftParameterGroup, error)
	RedshiftParameterGroupNamespaceListerExpansion
}

// redshiftParameterGroupNamespaceLister implements the RedshiftParameterGroupNamespaceLister
// interface.
type redshiftParameterGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RedshiftParameterGroups in the indexer for a given namespace.
func (s redshiftParameterGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RedshiftParameterGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RedshiftParameterGroup))
	})
	return ret, err
}

// Get retrieves the RedshiftParameterGroup from the indexer for a given namespace and name.
func (s redshiftParameterGroupNamespaceLister) Get(name string) (*v1alpha1.RedshiftParameterGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("redshiftparametergroup"), name)
	}
	return obj.(*v1alpha1.RedshiftParameterGroup), nil
}
