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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// LambdaEventSourceMappingLister helps list LambdaEventSourceMappings.
type LambdaEventSourceMappingLister interface {
	// List lists all LambdaEventSourceMappings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LambdaEventSourceMapping, err error)
	// LambdaEventSourceMappings returns an object that can list and get LambdaEventSourceMappings.
	LambdaEventSourceMappings(namespace string) LambdaEventSourceMappingNamespaceLister
	LambdaEventSourceMappingListerExpansion
}

// lambdaEventSourceMappingLister implements the LambdaEventSourceMappingLister interface.
type lambdaEventSourceMappingLister struct {
	indexer cache.Indexer
}

// NewLambdaEventSourceMappingLister returns a new LambdaEventSourceMappingLister.
func NewLambdaEventSourceMappingLister(indexer cache.Indexer) LambdaEventSourceMappingLister {
	return &lambdaEventSourceMappingLister{indexer: indexer}
}

// List lists all LambdaEventSourceMappings in the indexer.
func (s *lambdaEventSourceMappingLister) List(selector labels.Selector) (ret []*v1alpha1.LambdaEventSourceMapping, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LambdaEventSourceMapping))
	})
	return ret, err
}

// LambdaEventSourceMappings returns an object that can list and get LambdaEventSourceMappings.
func (s *lambdaEventSourceMappingLister) LambdaEventSourceMappings(namespace string) LambdaEventSourceMappingNamespaceLister {
	return lambdaEventSourceMappingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LambdaEventSourceMappingNamespaceLister helps list and get LambdaEventSourceMappings.
type LambdaEventSourceMappingNamespaceLister interface {
	// List lists all LambdaEventSourceMappings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LambdaEventSourceMapping, err error)
	// Get retrieves the LambdaEventSourceMapping from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LambdaEventSourceMapping, error)
	LambdaEventSourceMappingNamespaceListerExpansion
}

// lambdaEventSourceMappingNamespaceLister implements the LambdaEventSourceMappingNamespaceLister
// interface.
type lambdaEventSourceMappingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LambdaEventSourceMappings in the indexer for a given namespace.
func (s lambdaEventSourceMappingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LambdaEventSourceMapping, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LambdaEventSourceMapping))
	})
	return ret, err
}

// Get retrieves the LambdaEventSourceMapping from the indexer for a given namespace and name.
func (s lambdaEventSourceMappingNamespaceLister) Get(name string) (*v1alpha1.LambdaEventSourceMapping, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lambdaeventsourcemapping"), name)
	}
	return obj.(*v1alpha1.LambdaEventSourceMapping), nil
}