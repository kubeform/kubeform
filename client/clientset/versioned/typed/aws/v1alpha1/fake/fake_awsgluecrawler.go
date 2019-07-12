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

// FakeAwsGlueCrawlers implements AwsGlueCrawlerInterface
type FakeAwsGlueCrawlers struct {
	Fake *FakeAwsV1alpha1
}

var awsgluecrawlersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsgluecrawlers"}

var awsgluecrawlersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsGlueCrawler"}

// Get takes name of the awsGlueCrawler, and returns the corresponding awsGlueCrawler object, and an error if there is any.
func (c *FakeAwsGlueCrawlers) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsGlueCrawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsgluecrawlersResource, name), &v1alpha1.AwsGlueCrawler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCrawler), err
}

// List takes label and field selectors, and returns the list of AwsGlueCrawlers that match those selectors.
func (c *FakeAwsGlueCrawlers) List(opts v1.ListOptions) (result *v1alpha1.AwsGlueCrawlerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsgluecrawlersResource, awsgluecrawlersKind, opts), &v1alpha1.AwsGlueCrawlerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsGlueCrawlerList{ListMeta: obj.(*v1alpha1.AwsGlueCrawlerList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsGlueCrawlerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsGlueCrawlers.
func (c *FakeAwsGlueCrawlers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsgluecrawlersResource, opts))
}

// Create takes the representation of a awsGlueCrawler and creates it.  Returns the server's representation of the awsGlueCrawler, and an error, if there is any.
func (c *FakeAwsGlueCrawlers) Create(awsGlueCrawler *v1alpha1.AwsGlueCrawler) (result *v1alpha1.AwsGlueCrawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsgluecrawlersResource, awsGlueCrawler), &v1alpha1.AwsGlueCrawler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCrawler), err
}

// Update takes the representation of a awsGlueCrawler and updates it. Returns the server's representation of the awsGlueCrawler, and an error, if there is any.
func (c *FakeAwsGlueCrawlers) Update(awsGlueCrawler *v1alpha1.AwsGlueCrawler) (result *v1alpha1.AwsGlueCrawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsgluecrawlersResource, awsGlueCrawler), &v1alpha1.AwsGlueCrawler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCrawler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsGlueCrawlers) UpdateStatus(awsGlueCrawler *v1alpha1.AwsGlueCrawler) (*v1alpha1.AwsGlueCrawler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsgluecrawlersResource, "status", awsGlueCrawler), &v1alpha1.AwsGlueCrawler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCrawler), err
}

// Delete takes name of the awsGlueCrawler and deletes it. Returns an error if one occurs.
func (c *FakeAwsGlueCrawlers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsgluecrawlersResource, name), &v1alpha1.AwsGlueCrawler{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsGlueCrawlers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsgluecrawlersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsGlueCrawlerList{})
	return err
}

// Patch applies the patch and returns the patched awsGlueCrawler.
func (c *FakeAwsGlueCrawlers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsGlueCrawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsgluecrawlersResource, name, pt, data, subresources...), &v1alpha1.AwsGlueCrawler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsGlueCrawler), err
}
