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

// GlueClassifierLister helps list GlueClassifiers.
type GlueClassifierLister interface {
	// List lists all GlueClassifiers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GlueClassifier, err error)
	// GlueClassifiers returns an object that can list and get GlueClassifiers.
	GlueClassifiers(namespace string) GlueClassifierNamespaceLister
	GlueClassifierListerExpansion
}

// glueClassifierLister implements the GlueClassifierLister interface.
type glueClassifierLister struct {
	indexer cache.Indexer
}

// NewGlueClassifierLister returns a new GlueClassifierLister.
func NewGlueClassifierLister(indexer cache.Indexer) GlueClassifierLister {
	return &glueClassifierLister{indexer: indexer}
}

// List lists all GlueClassifiers in the indexer.
func (s *glueClassifierLister) List(selector labels.Selector) (ret []*v1alpha1.GlueClassifier, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlueClassifier))
	})
	return ret, err
}

// GlueClassifiers returns an object that can list and get GlueClassifiers.
func (s *glueClassifierLister) GlueClassifiers(namespace string) GlueClassifierNamespaceLister {
	return glueClassifierNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlueClassifierNamespaceLister helps list and get GlueClassifiers.
type GlueClassifierNamespaceLister interface {
	// List lists all GlueClassifiers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GlueClassifier, err error)
	// Get retrieves the GlueClassifier from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GlueClassifier, error)
	GlueClassifierNamespaceListerExpansion
}

// glueClassifierNamespaceLister implements the GlueClassifierNamespaceLister
// interface.
type glueClassifierNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlueClassifiers in the indexer for a given namespace.
func (s glueClassifierNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GlueClassifier, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlueClassifier))
	})
	return ret, err
}

// Get retrieves the GlueClassifier from the indexer for a given namespace and name.
func (s glueClassifierNamespaceLister) Get(name string) (*v1alpha1.GlueClassifier, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("glueclassifier"), name)
	}
	return obj.(*v1alpha1.GlueClassifier), nil
}