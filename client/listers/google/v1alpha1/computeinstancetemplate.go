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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"
)

// ComputeInstanceTemplateLister helps list ComputeInstanceTemplates.
type ComputeInstanceTemplateLister interface {
	// List lists all ComputeInstanceTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeInstanceTemplate, err error)
	// ComputeInstanceTemplates returns an object that can list and get ComputeInstanceTemplates.
	ComputeInstanceTemplates(namespace string) ComputeInstanceTemplateNamespaceLister
	ComputeInstanceTemplateListerExpansion
}

// computeInstanceTemplateLister implements the ComputeInstanceTemplateLister interface.
type computeInstanceTemplateLister struct {
	indexer cache.Indexer
}

// NewComputeInstanceTemplateLister returns a new ComputeInstanceTemplateLister.
func NewComputeInstanceTemplateLister(indexer cache.Indexer) ComputeInstanceTemplateLister {
	return &computeInstanceTemplateLister{indexer: indexer}
}

// List lists all ComputeInstanceTemplates in the indexer.
func (s *computeInstanceTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeInstanceTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeInstanceTemplate))
	})
	return ret, err
}

// ComputeInstanceTemplates returns an object that can list and get ComputeInstanceTemplates.
func (s *computeInstanceTemplateLister) ComputeInstanceTemplates(namespace string) ComputeInstanceTemplateNamespaceLister {
	return computeInstanceTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComputeInstanceTemplateNamespaceLister helps list and get ComputeInstanceTemplates.
type ComputeInstanceTemplateNamespaceLister interface {
	// List lists all ComputeInstanceTemplates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ComputeInstanceTemplate, err error)
	// Get retrieves the ComputeInstanceTemplate from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ComputeInstanceTemplate, error)
	ComputeInstanceTemplateNamespaceListerExpansion
}

// computeInstanceTemplateNamespaceLister implements the ComputeInstanceTemplateNamespaceLister
// interface.
type computeInstanceTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComputeInstanceTemplates in the indexer for a given namespace.
func (s computeInstanceTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComputeInstanceTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComputeInstanceTemplate))
	})
	return ret, err
}

// Get retrieves the ComputeInstanceTemplate from the indexer for a given namespace and name.
func (s computeInstanceTemplateNamespaceLister) Get(name string) (*v1alpha1.ComputeInstanceTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("computeinstancetemplate"), name)
	}
	return obj.(*v1alpha1.ComputeInstanceTemplate), nil
}
