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

// S3AccountPublicAccessBlockLister helps list S3AccountPublicAccessBlocks.
type S3AccountPublicAccessBlockLister interface {
	// List lists all S3AccountPublicAccessBlocks in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.S3AccountPublicAccessBlock, err error)
	// Get retrieves the S3AccountPublicAccessBlock from the index for a given name.
	Get(name string) (*v1alpha1.S3AccountPublicAccessBlock, error)
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

// Get retrieves the S3AccountPublicAccessBlock from the index for a given name.
func (s *s3AccountPublicAccessBlockLister) Get(name string) (*v1alpha1.S3AccountPublicAccessBlock, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("s3accountpublicaccessblock"), name)
	}
	return obj.(*v1alpha1.S3AccountPublicAccessBlock), nil
}
