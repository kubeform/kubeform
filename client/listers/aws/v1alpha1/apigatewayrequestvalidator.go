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

// ApiGatewayRequestValidatorLister helps list ApiGatewayRequestValidators.
type ApiGatewayRequestValidatorLister interface {
	// List lists all ApiGatewayRequestValidators in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayRequestValidator, err error)
	// ApiGatewayRequestValidators returns an object that can list and get ApiGatewayRequestValidators.
	ApiGatewayRequestValidators(namespace string) ApiGatewayRequestValidatorNamespaceLister
	ApiGatewayRequestValidatorListerExpansion
}

// apiGatewayRequestValidatorLister implements the ApiGatewayRequestValidatorLister interface.
type apiGatewayRequestValidatorLister struct {
	indexer cache.Indexer
}

// NewApiGatewayRequestValidatorLister returns a new ApiGatewayRequestValidatorLister.
func NewApiGatewayRequestValidatorLister(indexer cache.Indexer) ApiGatewayRequestValidatorLister {
	return &apiGatewayRequestValidatorLister{indexer: indexer}
}

// List lists all ApiGatewayRequestValidators in the indexer.
func (s *apiGatewayRequestValidatorLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayRequestValidator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayRequestValidator))
	})
	return ret, err
}

// ApiGatewayRequestValidators returns an object that can list and get ApiGatewayRequestValidators.
func (s *apiGatewayRequestValidatorLister) ApiGatewayRequestValidators(namespace string) ApiGatewayRequestValidatorNamespaceLister {
	return apiGatewayRequestValidatorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayRequestValidatorNamespaceLister helps list and get ApiGatewayRequestValidators.
type ApiGatewayRequestValidatorNamespaceLister interface {
	// List lists all ApiGatewayRequestValidators in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayRequestValidator, err error)
	// Get retrieves the ApiGatewayRequestValidator from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayRequestValidator, error)
	ApiGatewayRequestValidatorNamespaceListerExpansion
}

// apiGatewayRequestValidatorNamespaceLister implements the ApiGatewayRequestValidatorNamespaceLister
// interface.
type apiGatewayRequestValidatorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayRequestValidators in the indexer for a given namespace.
func (s apiGatewayRequestValidatorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayRequestValidator, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayRequestValidator))
	})
	return ret, err
}

// Get retrieves the ApiGatewayRequestValidator from the indexer for a given namespace and name.
func (s apiGatewayRequestValidatorNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayRequestValidator, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewayrequestvalidator"), name)
	}
	return obj.(*v1alpha1.ApiGatewayRequestValidator), nil
}