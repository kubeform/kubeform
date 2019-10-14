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

// ApiGatewayAPIKeyLister helps list ApiGatewayAPIKeys.
type ApiGatewayAPIKeyLister interface {
	// List lists all ApiGatewayAPIKeys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayAPIKey, err error)
	// ApiGatewayAPIKeys returns an object that can list and get ApiGatewayAPIKeys.
	ApiGatewayAPIKeys(namespace string) ApiGatewayAPIKeyNamespaceLister
	ApiGatewayAPIKeyListerExpansion
}

// apiGatewayAPIKeyLister implements the ApiGatewayAPIKeyLister interface.
type apiGatewayAPIKeyLister struct {
	indexer cache.Indexer
}

// NewApiGatewayAPIKeyLister returns a new ApiGatewayAPIKeyLister.
func NewApiGatewayAPIKeyLister(indexer cache.Indexer) ApiGatewayAPIKeyLister {
	return &apiGatewayAPIKeyLister{indexer: indexer}
}

// List lists all ApiGatewayAPIKeys in the indexer.
func (s *apiGatewayAPIKeyLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayAPIKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayAPIKey))
	})
	return ret, err
}

// ApiGatewayAPIKeys returns an object that can list and get ApiGatewayAPIKeys.
func (s *apiGatewayAPIKeyLister) ApiGatewayAPIKeys(namespace string) ApiGatewayAPIKeyNamespaceLister {
	return apiGatewayAPIKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayAPIKeyNamespaceLister helps list and get ApiGatewayAPIKeys.
type ApiGatewayAPIKeyNamespaceLister interface {
	// List lists all ApiGatewayAPIKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayAPIKey, err error)
	// Get retrieves the ApiGatewayAPIKey from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayAPIKey, error)
	ApiGatewayAPIKeyNamespaceListerExpansion
}

// apiGatewayAPIKeyNamespaceLister implements the ApiGatewayAPIKeyNamespaceLister
// interface.
type apiGatewayAPIKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayAPIKeys in the indexer for a given namespace.
func (s apiGatewayAPIKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayAPIKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayAPIKey))
	})
	return ret, err
}

// Get retrieves the ApiGatewayAPIKey from the indexer for a given namespace and name.
func (s apiGatewayAPIKeyNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayAPIKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewayapikey"), name)
	}
	return obj.(*v1alpha1.ApiGatewayAPIKey), nil
}
