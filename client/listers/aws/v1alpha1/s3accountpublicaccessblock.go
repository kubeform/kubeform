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

// S3AccountPublicAccessBlockLister helps list S3AccountPublicAccessBlocks.
type S3AccountPublicAccessBlockLister interface {
	// List lists all S3AccountPublicAccessBlocks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.S3AccountPublicAccessBlock, err error)
	// S3AccountPublicAccessBlocks returns an object that can list and get S3AccountPublicAccessBlocks.
	S3AccountPublicAccessBlocks(namespace string) S3AccountPublicAccessBlockNamespaceLister
	S3AccountPublicAccessBlockListerExpansion
}

// s3AccountPublicAccessBlockLister implements the S3AccountPublicAccessBlockLister interface.
type s3AccountPublicAccessBlockLister struct {
	indexer cache.Indexer
}

// NewS3AccountPublicAccessBlockLister returns a new S3AccountPublicAccessBlockLister.
func NewS3AccountPublicAccessBlockLister(indexer cache.Indexer) S3AccountPublicAccessBlockLister {
	return &s3AccountPublicAccessBlockLister{indexer: indexer}
}

// List lists all S3AccountPublicAccessBlocks in the indexer.
func (s *s3AccountPublicAccessBlockLister) List(selector labels.Selector) (ret []*v1alpha1.S3AccountPublicAccessBlock, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S3AccountPublicAccessBlock))
	})
	return ret, err
}

// S3AccountPublicAccessBlocks returns an object that can list and get S3AccountPublicAccessBlocks.
func (s *s3AccountPublicAccessBlockLister) S3AccountPublicAccessBlocks(namespace string) S3AccountPublicAccessBlockNamespaceLister {
	return s3AccountPublicAccessBlockNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// S3AccountPublicAccessBlockNamespaceLister helps list and get S3AccountPublicAccessBlocks.
type S3AccountPublicAccessBlockNamespaceLister interface {
	// List lists all S3AccountPublicAccessBlocks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.S3AccountPublicAccessBlock, err error)
	// Get retrieves the S3AccountPublicAccessBlock from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.S3AccountPublicAccessBlock, error)
	S3AccountPublicAccessBlockNamespaceListerExpansion
}

// s3AccountPublicAccessBlockNamespaceLister implements the S3AccountPublicAccessBlockNamespaceLister
// interface.
type s3AccountPublicAccessBlockNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all S3AccountPublicAccessBlocks in the indexer for a given namespace.
func (s s3AccountPublicAccessBlockNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.S3AccountPublicAccessBlock, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S3AccountPublicAccessBlock))
	})
	return ret, err
}

// Get retrieves the S3AccountPublicAccessBlock from the indexer for a given namespace and name.
func (s s3AccountPublicAccessBlockNamespaceLister) Get(name string) (*v1alpha1.S3AccountPublicAccessBlock, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("s3accountpublicaccessblock"), name)
	}
	return obj.(*v1alpha1.S3AccountPublicAccessBlock), nil
}
