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

// S3BucketInventoryLister helps list S3BucketInventories.
type S3BucketInventoryLister interface {
	// List lists all S3BucketInventories in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.S3BucketInventory, err error)
	// S3BucketInventories returns an object that can list and get S3BucketInventories.
	S3BucketInventories(namespace string) S3BucketInventoryNamespaceLister
	S3BucketInventoryListerExpansion
}

// s3BucketInventoryLister implements the S3BucketInventoryLister interface.
type s3BucketInventoryLister struct {
	indexer cache.Indexer
}

// NewS3BucketInventoryLister returns a new S3BucketInventoryLister.
func NewS3BucketInventoryLister(indexer cache.Indexer) S3BucketInventoryLister {
	return &s3BucketInventoryLister{indexer: indexer}
}

// List lists all S3BucketInventories in the indexer.
func (s *s3BucketInventoryLister) List(selector labels.Selector) (ret []*v1alpha1.S3BucketInventory, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S3BucketInventory))
	})
	return ret, err
}

// S3BucketInventories returns an object that can list and get S3BucketInventories.
func (s *s3BucketInventoryLister) S3BucketInventories(namespace string) S3BucketInventoryNamespaceLister {
	return s3BucketInventoryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// S3BucketInventoryNamespaceLister helps list and get S3BucketInventories.
type S3BucketInventoryNamespaceLister interface {
	// List lists all S3BucketInventories in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.S3BucketInventory, err error)
	// Get retrieves the S3BucketInventory from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.S3BucketInventory, error)
	S3BucketInventoryNamespaceListerExpansion
}

// s3BucketInventoryNamespaceLister implements the S3BucketInventoryNamespaceLister
// interface.
type s3BucketInventoryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all S3BucketInventories in the indexer for a given namespace.
func (s s3BucketInventoryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.S3BucketInventory, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.S3BucketInventory))
	})
	return ret, err
}

// Get retrieves the S3BucketInventory from the indexer for a given namespace and name.
func (s s3BucketInventoryNamespaceLister) Get(name string) (*v1alpha1.S3BucketInventory, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("s3bucketinventory"), name)
	}
	return obj.(*v1alpha1.S3BucketInventory), nil
}
