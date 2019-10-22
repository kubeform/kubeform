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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudformationStackSetLister helps list CloudformationStackSets.
type CloudformationStackSetLister interface {
	// List lists all CloudformationStackSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudformationStackSet, err error)
	// CloudformationStackSets returns an object that can list and get CloudformationStackSets.
	CloudformationStackSets(namespace string) CloudformationStackSetNamespaceLister
	CloudformationStackSetListerExpansion
}

// cloudformationStackSetLister implements the CloudformationStackSetLister interface.
type cloudformationStackSetLister struct {
	indexer cache.Indexer
}

// NewCloudformationStackSetLister returns a new CloudformationStackSetLister.
func NewCloudformationStackSetLister(indexer cache.Indexer) CloudformationStackSetLister {
	return &cloudformationStackSetLister{indexer: indexer}
}

// List lists all CloudformationStackSets in the indexer.
func (s *cloudformationStackSetLister) List(selector labels.Selector) (ret []*v1alpha1.CloudformationStackSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudformationStackSet))
	})
	return ret, err
}

// CloudformationStackSets returns an object that can list and get CloudformationStackSets.
func (s *cloudformationStackSetLister) CloudformationStackSets(namespace string) CloudformationStackSetNamespaceLister {
	return cloudformationStackSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudformationStackSetNamespaceLister helps list and get CloudformationStackSets.
type CloudformationStackSetNamespaceLister interface {
	// List lists all CloudformationStackSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudformationStackSet, err error)
	// Get retrieves the CloudformationStackSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudformationStackSet, error)
	CloudformationStackSetNamespaceListerExpansion
}

// cloudformationStackSetNamespaceLister implements the CloudformationStackSetNamespaceLister
// interface.
type cloudformationStackSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudformationStackSets in the indexer for a given namespace.
func (s cloudformationStackSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudformationStackSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudformationStackSet))
	})
	return ret, err
}

// Get retrieves the CloudformationStackSet from the indexer for a given namespace and name.
func (s cloudformationStackSetNamespaceLister) Get(name string) (*v1alpha1.CloudformationStackSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudformationstackset"), name)
	}
	return obj.(*v1alpha1.CloudformationStackSet), nil
}
