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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
)

// FakeElasticsearchDomains implements ElasticsearchDomainInterface
type FakeElasticsearchDomains struct {
	Fake *FakeAwsV1alpha1
}

var elasticsearchdomainsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "elasticsearchdomains"}

var elasticsearchdomainsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ElasticsearchDomain"}

// Get takes name of the elasticsearchDomain, and returns the corresponding elasticsearchDomain object, and an error if there is any.
func (c *FakeElasticsearchDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.ElasticsearchDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(elasticsearchdomainsResource, name), &v1alpha1.ElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchDomain), err
}

// List takes label and field selectors, and returns the list of ElasticsearchDomains that match those selectors.
func (c *FakeElasticsearchDomains) List(opts v1.ListOptions) (result *v1alpha1.ElasticsearchDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(elasticsearchdomainsResource, elasticsearchdomainsKind, opts), &v1alpha1.ElasticsearchDomainList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElasticsearchDomainList{ListMeta: obj.(*v1alpha1.ElasticsearchDomainList).ListMeta}
	for _, item := range obj.(*v1alpha1.ElasticsearchDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elasticsearchDomains.
func (c *FakeElasticsearchDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(elasticsearchdomainsResource, opts))
}

// Create takes the representation of a elasticsearchDomain and creates it.  Returns the server's representation of the elasticsearchDomain, and an error, if there is any.
func (c *FakeElasticsearchDomains) Create(elasticsearchDomain *v1alpha1.ElasticsearchDomain) (result *v1alpha1.ElasticsearchDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(elasticsearchdomainsResource, elasticsearchDomain), &v1alpha1.ElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchDomain), err
}

// Update takes the representation of a elasticsearchDomain and updates it. Returns the server's representation of the elasticsearchDomain, and an error, if there is any.
func (c *FakeElasticsearchDomains) Update(elasticsearchDomain *v1alpha1.ElasticsearchDomain) (result *v1alpha1.ElasticsearchDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(elasticsearchdomainsResource, elasticsearchDomain), &v1alpha1.ElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchDomain), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeElasticsearchDomains) UpdateStatus(elasticsearchDomain *v1alpha1.ElasticsearchDomain) (*v1alpha1.ElasticsearchDomain, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(elasticsearchdomainsResource, "status", elasticsearchDomain), &v1alpha1.ElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchDomain), err
}

// Delete takes name of the elasticsearchDomain and deletes it. Returns an error if one occurs.
func (c *FakeElasticsearchDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(elasticsearchdomainsResource, name), &v1alpha1.ElasticsearchDomain{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElasticsearchDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(elasticsearchdomainsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElasticsearchDomainList{})
	return err
}

// Patch applies the patch and returns the patched elasticsearchDomain.
func (c *FakeElasticsearchDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElasticsearchDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(elasticsearchdomainsResource, name, pt, data, subresources...), &v1alpha1.ElasticsearchDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElasticsearchDomain), err
}
