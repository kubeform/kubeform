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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PubsubSubscriptionIamMemberLister helps list PubsubSubscriptionIamMembers.
type PubsubSubscriptionIamMemberLister interface {
	// List lists all PubsubSubscriptionIamMembers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubSubscriptionIamMember, err error)
	// PubsubSubscriptionIamMembers returns an object that can list and get PubsubSubscriptionIamMembers.
	PubsubSubscriptionIamMembers(namespace string) PubsubSubscriptionIamMemberNamespaceLister
	PubsubSubscriptionIamMemberListerExpansion
}

// pubsubSubscriptionIamMemberLister implements the PubsubSubscriptionIamMemberLister interface.
type pubsubSubscriptionIamMemberLister struct {
	indexer cache.Indexer
}

// NewPubsubSubscriptionIamMemberLister returns a new PubsubSubscriptionIamMemberLister.
func NewPubsubSubscriptionIamMemberLister(indexer cache.Indexer) PubsubSubscriptionIamMemberLister {
	return &pubsubSubscriptionIamMemberLister{indexer: indexer}
}

// List lists all PubsubSubscriptionIamMembers in the indexer.
func (s *pubsubSubscriptionIamMemberLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubSubscriptionIamMember, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubSubscriptionIamMember))
	})
	return ret, err
}

// PubsubSubscriptionIamMembers returns an object that can list and get PubsubSubscriptionIamMembers.
func (s *pubsubSubscriptionIamMemberLister) PubsubSubscriptionIamMembers(namespace string) PubsubSubscriptionIamMemberNamespaceLister {
	return pubsubSubscriptionIamMemberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PubsubSubscriptionIamMemberNamespaceLister helps list and get PubsubSubscriptionIamMembers.
type PubsubSubscriptionIamMemberNamespaceLister interface {
	// List lists all PubsubSubscriptionIamMembers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PubsubSubscriptionIamMember, err error)
	// Get retrieves the PubsubSubscriptionIamMember from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PubsubSubscriptionIamMember, error)
	PubsubSubscriptionIamMemberNamespaceListerExpansion
}

// pubsubSubscriptionIamMemberNamespaceLister implements the PubsubSubscriptionIamMemberNamespaceLister
// interface.
type pubsubSubscriptionIamMemberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PubsubSubscriptionIamMembers in the indexer for a given namespace.
func (s pubsubSubscriptionIamMemberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PubsubSubscriptionIamMember, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PubsubSubscriptionIamMember))
	})
	return ret, err
}

// Get retrieves the PubsubSubscriptionIamMember from the indexer for a given namespace and name.
func (s pubsubSubscriptionIamMemberNamespaceLister) Get(name string) (*v1alpha1.PubsubSubscriptionIamMember, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pubsubsubscriptioniammember"), name)
	}
	return obj.(*v1alpha1.PubsubSubscriptionIamMember), nil
}
