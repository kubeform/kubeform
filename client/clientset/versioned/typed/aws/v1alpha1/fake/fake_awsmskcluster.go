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

// FakeAwsMskClusters implements AwsMskClusterInterface
type FakeAwsMskClusters struct {
	Fake *FakeAwsV1alpha1
}

var awsmskclustersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "awsmskclusters"}

var awsmskclustersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "AwsMskCluster"}

// Get takes name of the awsMskCluster, and returns the corresponding awsMskCluster object, and an error if there is any.
func (c *FakeAwsMskClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsMskCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsmskclustersResource, name), &v1alpha1.AwsMskCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMskCluster), err
}

// List takes label and field selectors, and returns the list of AwsMskClusters that match those selectors.
func (c *FakeAwsMskClusters) List(opts v1.ListOptions) (result *v1alpha1.AwsMskClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsmskclustersResource, awsmskclustersKind, opts), &v1alpha1.AwsMskClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsMskClusterList{ListMeta: obj.(*v1alpha1.AwsMskClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsMskClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsMskClusters.
func (c *FakeAwsMskClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsmskclustersResource, opts))
}

// Create takes the representation of a awsMskCluster and creates it.  Returns the server's representation of the awsMskCluster, and an error, if there is any.
func (c *FakeAwsMskClusters) Create(awsMskCluster *v1alpha1.AwsMskCluster) (result *v1alpha1.AwsMskCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsmskclustersResource, awsMskCluster), &v1alpha1.AwsMskCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMskCluster), err
}

// Update takes the representation of a awsMskCluster and updates it. Returns the server's representation of the awsMskCluster, and an error, if there is any.
func (c *FakeAwsMskClusters) Update(awsMskCluster *v1alpha1.AwsMskCluster) (result *v1alpha1.AwsMskCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsmskclustersResource, awsMskCluster), &v1alpha1.AwsMskCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMskCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsMskClusters) UpdateStatus(awsMskCluster *v1alpha1.AwsMskCluster) (*v1alpha1.AwsMskCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsmskclustersResource, "status", awsMskCluster), &v1alpha1.AwsMskCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMskCluster), err
}

// Delete takes name of the awsMskCluster and deletes it. Returns an error if one occurs.
func (c *FakeAwsMskClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsmskclustersResource, name), &v1alpha1.AwsMskCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsMskClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsmskclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsMskClusterList{})
	return err
}

// Patch applies the patch and returns the patched awsMskCluster.
func (c *FakeAwsMskClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsMskCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsmskclustersResource, name, pt, data, subresources...), &v1alpha1.AwsMskCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsMskCluster), err
}
