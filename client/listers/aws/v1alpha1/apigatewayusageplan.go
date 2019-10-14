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

// ApiGatewayUsagePlanLister helps list ApiGatewayUsagePlans.
type ApiGatewayUsagePlanLister interface {
	// List lists all ApiGatewayUsagePlans in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayUsagePlan, err error)
	// ApiGatewayUsagePlans returns an object that can list and get ApiGatewayUsagePlans.
	ApiGatewayUsagePlans(namespace string) ApiGatewayUsagePlanNamespaceLister
	ApiGatewayUsagePlanListerExpansion
}

// apiGatewayUsagePlanLister implements the ApiGatewayUsagePlanLister interface.
type apiGatewayUsagePlanLister struct {
	indexer cache.Indexer
}

// NewApiGatewayUsagePlanLister returns a new ApiGatewayUsagePlanLister.
func NewApiGatewayUsagePlanLister(indexer cache.Indexer) ApiGatewayUsagePlanLister {
	return &apiGatewayUsagePlanLister{indexer: indexer}
}

// List lists all ApiGatewayUsagePlans in the indexer.
func (s *apiGatewayUsagePlanLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayUsagePlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayUsagePlan))
	})
	return ret, err
}

// ApiGatewayUsagePlans returns an object that can list and get ApiGatewayUsagePlans.
func (s *apiGatewayUsagePlanLister) ApiGatewayUsagePlans(namespace string) ApiGatewayUsagePlanNamespaceLister {
	return apiGatewayUsagePlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiGatewayUsagePlanNamespaceLister helps list and get ApiGatewayUsagePlans.
type ApiGatewayUsagePlanNamespaceLister interface {
	// List lists all ApiGatewayUsagePlans in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayUsagePlan, err error)
	// Get retrieves the ApiGatewayUsagePlan from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiGatewayUsagePlan, error)
	ApiGatewayUsagePlanNamespaceListerExpansion
}

// apiGatewayUsagePlanNamespaceLister implements the ApiGatewayUsagePlanNamespaceLister
// interface.
type apiGatewayUsagePlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiGatewayUsagePlans in the indexer for a given namespace.
func (s apiGatewayUsagePlanNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiGatewayUsagePlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiGatewayUsagePlan))
	})
	return ret, err
}

// Get retrieves the ApiGatewayUsagePlan from the indexer for a given namespace and name.
func (s apiGatewayUsagePlanNamespaceLister) Get(name string) (*v1alpha1.ApiGatewayUsagePlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apigatewayusageplan"), name)
	}
	return obj.(*v1alpha1.ApiGatewayUsagePlan), nil
}
