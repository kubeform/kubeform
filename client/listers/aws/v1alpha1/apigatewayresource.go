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

// ApiGatewayResourceLister helps list ApiGatewayResources.
type ApiGatewayResourceLister interface {
	// List lists all ApiGatewayResources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayResource, err error)
	// ApiGatewayResources returns an object that can list and get ApiGatewayResources.
	ApiGatewayResources(namespace string) ApiGatewayResourceNamespaceLister
	ApiGatewayResourceListerExpansion
}

// apiGatewayResourceLister implements the ApiGatewayResourceLister interface.
type apiGatewayResourceLister struct {
	indexer cache.Indexer
}

// NewApiGatewayResourceLister returns a new ApiGatewayResourceLister.
func NewApiGatewayResourceLister(indexer cache.Indexer) ApiGatewayResourceLister {
	return &apiGatewayResourceLister{indexer: indexer}
}

// List lists all ApiGatewayResources in the indexer.
func (s *apiGatewayResourceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayResource))
	})
	return ret, err
}

// ApiGatewayResources returns an object that can list and get ApiGatewayResources.
func (s *apiGatewayResourceLister) ApiGatewayResources(namespace string) ApiGatewayResourceNamespaceLister {
	return apiGatewayResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayResourceNamespaceLister helps list and get ApiGatewayResources.
type ApiGatewayResourceNamespaceLister interface {
	// List lists all ApiGatewayResources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayResource, err error)
	// Get retrieves the ApiGatewayResource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayResource, error)
	ApiGatewayResourceNamespaceListerExpansion
}

// apiGatewayResourceNamespaceLister implements the ApiGatewayResourceNamespaceLister
// interface.
type apiGatewayResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayResources in the indexer for a given namespace.
func (s apiGatewayResourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayResource))
	})
	return ret, err
}

// Get retrieves the ApiGatewayResource from the indexer for a given namespace and name.
func (s apiGatewayResourceNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewayresource"), name)
	}
	return obj.(*v1alpha1.ApiGatewayResource), nil
}
