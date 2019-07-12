/*
Copyright 2019 The KubeDB Authors.

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

// FakeAwsDatasyncLocationS3s implements AwsDatasyncLocationS3Interface
type FakeAwsDatasyncLocationS3s struct {
	Fake *FakeAwsV1alpha1
}

var awsdatasynclocations3sResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsdatasynclocations3s"}

var awsdatasynclocations3sKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsDatasyncLocationS3"}

// Get takes name of the awsDatasyncLocationS3, and returns the corresponding awsDatasyncLocationS3 object, and an error if there is any.
func (c *FakeAwsDatasyncLocationS3s) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsdatasynclocations3sResource, name), &v1alpha1.AwsDatasyncLocationS3{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationS3), err
}

// List takes label and field selectors, and returns the list of AwsDatasyncLocationS3s that match those selectors.
func (c *FakeAwsDatasyncLocationS3s) List(opts v1.ListOptions) (result *v1alpha1.AwsDatasyncLocationS3List, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsdatasynclocations3sResource, awsdatasynclocations3sKind, opts), &v1alpha1.AwsDatasyncLocationS3List{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsDatasyncLocationS3List{ListMeta: obj.(*v1alpha1.AwsDatasyncLocationS3List).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsDatasyncLocationS3List).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsDatasyncLocationS3s.
func (c *FakeAwsDatasyncLocationS3s) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsdatasynclocations3sResource, opts))
}

// Create takes the representation of a awsDatasyncLocationS3 and creates it.  Returns the server's representation of the awsDatasyncLocationS3, and an error, if there is any.
func (c *FakeAwsDatasyncLocationS3s) Create(awsDatasyncLocationS3 *v1alpha1.AwsDatasyncLocationS3) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsdatasynclocations3sResource, awsDatasyncLocationS3), &v1alpha1.AwsDatasyncLocationS3{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationS3), err
}

// Update takes the representation of a awsDatasyncLocationS3 and updates it. Returns the server's representation of the awsDatasyncLocationS3, and an error, if there is any.
func (c *FakeAwsDatasyncLocationS3s) Update(awsDatasyncLocationS3 *v1alpha1.AwsDatasyncLocationS3) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsdatasynclocations3sResource, awsDatasyncLocationS3), &v1alpha1.AwsDatasyncLocationS3{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationS3), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsDatasyncLocationS3s) UpdateStatus(awsDatasyncLocationS3 *v1alpha1.AwsDatasyncLocationS3) (*v1alpha1.AwsDatasyncLocationS3, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsdatasynclocations3sResource, "status", awsDatasyncLocationS3), &v1alpha1.AwsDatasyncLocationS3{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationS3), err
}

// Delete takes name of the awsDatasyncLocationS3 and deletes it. Returns an error if one occurs.
func (c *FakeAwsDatasyncLocationS3s) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsdatasynclocations3sResource, name), &v1alpha1.AwsDatasyncLocationS3{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsDatasyncLocationS3s) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsdatasynclocations3sResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsDatasyncLocationS3List{})
	return err
}

// Patch applies the patch and returns the patched awsDatasyncLocationS3.
func (c *FakeAwsDatasyncLocationS3s) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsDatasyncLocationS3, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsdatasynclocations3sResource, name, pt, data, subresources...), &v1alpha1.AwsDatasyncLocationS3{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsDatasyncLocationS3), err
}
