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

// CloudwatchEventPermissionLister helps list CloudwatchEventPermissions.
type CloudwatchEventPermissionLister interface {
	// List lists all CloudwatchEventPermissions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CloudwatchEventPermission, err error)
	// CloudwatchEventPermissions returns an object that can list and get CloudwatchEventPermissions.
	CloudwatchEventPermissions(namespace string) CloudwatchEventPermissionNamespaceLister
	CloudwatchEventPermissionListerExpansion
}

// cloudwatchEventPermissionLister implements the CloudwatchEventPermissionLister interface.
type cloudwatchEventPermissionLister struct {
	indexer cache.Indexer
}

// NewCloudwatchEventPermissionLister returns a new CloudwatchEventPermissionLister.
func NewCloudwatchEventPermissionLister(indexer cache.Indexer) CloudwatchEventPermissionLister {
	return &cloudwatchEventPermissionLister{indexer: indexer}
}

// List lists all CloudwatchEventPermissions in the indexer.
func (s *cloudwatchEventPermissionLister) List(selector labels.Selector) (ret []*v1alpha1.CloudwatchEventPermission, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudwatchEventPermission))
	})
	return ret, err
}

// CloudwatchEventPermissions returns an object that can list and get CloudwatchEventPermissions.
func (s *cloudwatchEventPermissionLister) CloudwatchEventPermissions(namespace string) CloudwatchEventPermissionNamespaceLister {
	return cloudwatchEventPermissionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudwatchEventPermissionNamespaceLister helps list and get CloudwatchEventPermissions.
type CloudwatchEventPermissionNamespaceLister interface {
	// List lists all CloudwatchEventPermissions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CloudwatchEventPermission, err error)
	// Get retrieves the CloudwatchEventPermission from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CloudwatchEventPermission, error)
	CloudwatchEventPermissionNamespaceListerExpansion
}

// cloudwatchEventPermissionNamespaceLister implements the CloudwatchEventPermissionNamespaceLister
// interface.
type cloudwatchEventPermissionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudwatchEventPermissions in the indexer for a given namespace.
func (s cloudwatchEventPermissionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudwatchEventPermission, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudwatchEventPermission))
	})
	return ret, err
}

// Get retrieves the CloudwatchEventPermission from the indexer for a given namespace and name.
func (s cloudwatchEventPermissionNamespaceLister) Get(name string) (*v1alpha1.CloudwatchEventPermission, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudwatcheventpermission"), name)
	}
	return obj.(*v1alpha1.CloudwatchEventPermission), nil
}