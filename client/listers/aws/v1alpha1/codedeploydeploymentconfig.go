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

// CodedeployDeploymentConfigLister helps list CodedeployDeploymentConfigs.
type CodedeployDeploymentConfigLister interface {
	// List lists all CodedeployDeploymentConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CodedeployDeploymentConfig, err error)
	// CodedeployDeploymentConfigs returns an object that can list and get CodedeployDeploymentConfigs.
	CodedeployDeploymentConfigs(namespace string) CodedeployDeploymentConfigNamespaceLister
	CodedeployDeploymentConfigListerExpansion
}

// codedeployDeploymentConfigLister implements the CodedeployDeploymentConfigLister interface.
type codedeployDeploymentConfigLister struct {
	indexer cache.Indexer
}

// NewCodedeployDeploymentConfigLister returns a new CodedeployDeploymentConfigLister.
func NewCodedeployDeploymentConfigLister(indexer cache.Indexer) CodedeployDeploymentConfigLister {
	return &codedeployDeploymentConfigLister{indexer: indexer}
}

// List lists all CodedeployDeploymentConfigs in the indexer.
func (s *codedeployDeploymentConfigLister) List(selector labels.Selector) (ret []*v1alpha1.CodedeployDeploymentConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodedeployDeploymentConfig))
	})
	return ret, err
}

// CodedeployDeploymentConfigs returns an object that can list and get CodedeployDeploymentConfigs.
func (s *codedeployDeploymentConfigLister) CodedeployDeploymentConfigs(namespace string) CodedeployDeploymentConfigNamespaceLister {
	return codedeployDeploymentConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CodedeployDeploymentConfigNamespaceLister helps list and get CodedeployDeploymentConfigs.
type CodedeployDeploymentConfigNamespaceLister interface {
	// List lists all CodedeployDeploymentConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CodedeployDeploymentConfig, err error)
	// Get retrieves the CodedeployDeploymentConfig from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CodedeployDeploymentConfig, error)
	CodedeployDeploymentConfigNamespaceListerExpansion
}

// codedeployDeploymentConfigNamespaceLister implements the CodedeployDeploymentConfigNamespaceLister
// interface.
type codedeployDeploymentConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CodedeployDeploymentConfigs in the indexer for a given namespace.
func (s codedeployDeploymentConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CodedeployDeploymentConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodedeployDeploymentConfig))
	})
	return ret, err
}

// Get retrieves the CodedeployDeploymentConfig from the indexer for a given namespace and name.
func (s codedeployDeploymentConfigNamespaceLister) Get(name string) (*v1alpha1.CodedeployDeploymentConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("codedeploydeploymentconfig"), name)
	}
	return obj.(*v1alpha1.CodedeployDeploymentConfig), nil
}
