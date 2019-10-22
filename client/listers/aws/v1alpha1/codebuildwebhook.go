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

// CodebuildWebhookLister helps list CodebuildWebhooks.
type CodebuildWebhookLister interface {
	// List lists all CodebuildWebhooks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CodebuildWebhook, err error)
	// CodebuildWebhooks returns an object that can list and get CodebuildWebhooks.
	CodebuildWebhooks(namespace string) CodebuildWebhookNamespaceLister
	CodebuildWebhookListerExpansion
}

// codebuildWebhookLister implements the CodebuildWebhookLister interface.
type codebuildWebhookLister struct {
	indexer cache.Indexer
}

// NewCodebuildWebhookLister returns a new CodebuildWebhookLister.
func NewCodebuildWebhookLister(indexer cache.Indexer) CodebuildWebhookLister {
	return &codebuildWebhookLister{indexer: indexer}
}

// List lists all CodebuildWebhooks in the indexer.
func (s *codebuildWebhookLister) List(selector labels.Selector) (ret []*v1alpha1.CodebuildWebhook, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodebuildWebhook))
	})
	return ret, err
}

// CodebuildWebhooks returns an object that can list and get CodebuildWebhooks.
func (s *codebuildWebhookLister) CodebuildWebhooks(namespace string) CodebuildWebhookNamespaceLister {
	return codebuildWebhookNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CodebuildWebhookNamespaceLister helps list and get CodebuildWebhooks.
type CodebuildWebhookNamespaceLister interface {
	// List lists all CodebuildWebhooks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CodebuildWebhook, err error)
	// Get retrieves the CodebuildWebhook from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CodebuildWebhook, error)
	CodebuildWebhookNamespaceListerExpansion
}

// codebuildWebhookNamespaceLister implements the CodebuildWebhookNamespaceLister
// interface.
type codebuildWebhookNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CodebuildWebhooks in the indexer for a given namespace.
func (s codebuildWebhookNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CodebuildWebhook, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CodebuildWebhook))
	})
	return ret, err
}

// Get retrieves the CodebuildWebhook from the indexer for a given namespace and name.
func (s codebuildWebhookNamespaceLister) Get(name string) (*v1alpha1.CodebuildWebhook, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("codebuildwebhook"), name)
	}
	return obj.(*v1alpha1.CodebuildWebhook), nil
}
