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

// AwsAppmeshVirtualRouterLister helps list AwsAppmeshVirtualRouters.
type AwsAppmeshVirtualRouterLister interface {
	// List lists all AwsAppmeshVirtualRouters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AwsAppmeshVirtualRouter, err error)
	// Get retrieves the AwsAppmeshVirtualRouter from the index for a given name.
	Get(name string) (*v1alpha1.AwsAppmeshVirtualRouter, error)
	AwsAppmeshVirtualRouterListerExpansion
}

// awsAppmeshVirtualRouterLister implements the AwsAppmeshVirtualRouterLister interface.
type awsAppmeshVirtualRouterLister struct {
	indexer cache.Indexer
}

// NewAwsAppmeshVirtualRouterLister returns a new AwsAppmeshVirtualRouterLister.
func NewAwsAppmeshVirtualRouterLister(indexer cache.Indexer) AwsAppmeshVirtualRouterLister {
	return &awsAppmeshVirtualRouterLister{indexer: indexer}
}

// List lists all AwsAppmeshVirtualRouters in the indexer.
func (s *awsAppmeshVirtualRouterLister) List(selector labels.Selector) (ret []*v1alpha1.AwsAppmeshVirtualRouter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AwsAppmeshVirtualRouter))
	})
	return ret, err
}

// Get retrieves the AwsAppmeshVirtualRouter from the index for a given name.
func (s *awsAppmeshVirtualRouterLister) Get(name string) (*v1alpha1.AwsAppmeshVirtualRouter, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsappmeshvirtualrouter"), name)
	}
	return obj.(*v1alpha1.AwsAppmeshVirtualRouter), nil
}
