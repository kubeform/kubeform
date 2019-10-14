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

// FakeEksClusters implements EksClusterInterface
type FakeEksClusters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var eksclustersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "eksclusters"}

var eksclustersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "EksCluster"}

// Get takes name of the eksCluster, and returns the corresponding eksCluster object, and an error if there is any.
func (c *FakeEksClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.EksCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eksclustersResource, c.ns, name), &v1alpha1.EksCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EksCluster), err
}

// List takes label and field selectors, and returns the list of EksClusters that match those selectors.
func (c *FakeEksClusters) List(opts v1.ListOptions) (result *v1alpha1.EksClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eksclustersResource, eksclustersKind, c.ns, opts), &v1alpha1.EksClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EksClusterList{ListMeta: obj.(*v1alpha1.EksClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.EksClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eksClusters.
func (c *FakeEksClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eksclustersResource, c.ns, opts))

}

// Create takes the representation of a eksCluster and creates it.  Returns the server's representation of the eksCluster, and an error, if there is any.
func (c *FakeEksClusters) Create(eksCluster *v1alpha1.EksCluster) (result *v1alpha1.EksCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eksclustersResource, c.ns, eksCluster), &v1alpha1.EksCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EksCluster), err
}

// Update takes the representation of a eksCluster and updates it. Returns the server's representation of the eksCluster, and an error, if there is any.
func (c *FakeEksClusters) Update(eksCluster *v1alpha1.EksCluster) (result *v1alpha1.EksCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eksclustersResource, c.ns, eksCluster), &v1alpha1.EksCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EksCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEksClusters) UpdateStatus(eksCluster *v1alpha1.EksCluster) (*v1alpha1.EksCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eksclustersResource, "status", c.ns, eksCluster), &v1alpha1.EksCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EksCluster), err
}

// Delete takes name of the eksCluster and deletes it. Returns an error if one occurs.
func (c *FakeEksClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eksclustersResource, c.ns, name), &v1alpha1.EksCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEksClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eksclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EksClusterList{})
	return err
}

// Patch applies the patch and returns the patched eksCluster.
func (c *FakeEksClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EksCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eksclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.EksCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EksCluster), err
}
