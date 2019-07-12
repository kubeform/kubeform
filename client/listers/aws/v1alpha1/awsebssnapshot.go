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

// AwsEbsSnapshotLister helps list AwsEbsSnapshots.
type AwsEbsSnapshotLister interface {
	// List lists all AwsEbsSnapshots in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsEbsSnapshot, err error)
	// Get retrieves the AwsEbsSnapshot from the index for a given name.
	Get(name string) (*v1alpha1.AwsEbsSnapshot, error)
	AwsEbsSnapshotListerExpansion
}

// awsEbsSnapshotLister implements the AwsEbsSnapshotLister interface.
type awsEbsSnapshotLister struct {
	indexer cache.Indexer
}

// NewAwsEbsSnapshotLister returns a new AwsEbsSnapshotLister.
func NewAwsEbsSnapshotLister(indexer cache.Indexer) AwsEbsSnapshotLister {
	return &awsEbsSnapshotLister{indexer: indexer}
}

// List lists all AwsEbsSnapshots in the indexer.
func (s *awsEbsSnapshotLister) List(selector labels.Selector) (ret []*v1alpha1.AwsEbsSnapshot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsEbsSnapshot))
	})
	return ret, err
}

// Get retrieves the AwsEbsSnapshot from the index for a given name.
func (s *awsEbsSnapshotLister) Get(name string) (*v1alpha1.AwsEbsSnapshot, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsebssnapshot"), name)
	}
	return obj.(*v1alpha1.AwsEbsSnapshot), nil
}
