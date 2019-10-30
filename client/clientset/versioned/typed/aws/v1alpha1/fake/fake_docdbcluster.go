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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDocdbClusters implements DocdbClusterInterface
type FakeDocdbClusters struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var docdbclustersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "docdbclusters"}

var docdbclustersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "DocdbCluster"}

// Get takes name of the docdbCluster, and returns the corresponding docdbCluster object, and an error if there is any.
func (c *FakeDocdbClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.DocdbCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(docdbclustersResource, c.ns, name), &v1alpha1.DocdbCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbCluster), err
}

// List takes label and field selectors, and returns the list of DocdbClusters that match those selectors.
func (c *FakeDocdbClusters) List(opts v1.ListOptions) (result *v1alpha1.DocdbClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(docdbclustersResource, docdbclustersKind, c.ns, opts), &v1alpha1.DocdbClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DocdbClusterList{ListMeta: obj.(*v1alpha1.DocdbClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.DocdbClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested docdbClusters.
func (c *FakeDocdbClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(docdbclustersResource, c.ns, opts))

}

// Create takes the representation of a docdbCluster and creates it.  Returns the server's representation of the docdbCluster, and an error, if there is any.
func (c *FakeDocdbClusters) Create(docdbCluster *v1alpha1.DocdbCluster) (result *v1alpha1.DocdbCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(docdbclustersResource, c.ns, docdbCluster), &v1alpha1.DocdbCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbCluster), err
}

// Update takes the representation of a docdbCluster and updates it. Returns the server's representation of the docdbCluster, and an error, if there is any.
func (c *FakeDocdbClusters) Update(docdbCluster *v1alpha1.DocdbCluster) (result *v1alpha1.DocdbCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(docdbclustersResource, c.ns, docdbCluster), &v1alpha1.DocdbCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDocdbClusters) UpdateStatus(docdbCluster *v1alpha1.DocdbCluster) (*v1alpha1.DocdbCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(docdbclustersResource, "status", c.ns, docdbCluster), &v1alpha1.DocdbCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbCluster), err
}

// Delete takes name of the docdbCluster and deletes it. Returns an error if one occurs.
func (c *FakeDocdbClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(docdbclustersResource, c.ns, name), &v1alpha1.DocdbCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDocdbClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(docdbclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DocdbClusterList{})
	return err
}

// Patch applies the patch and returns the patched docdbCluster.
func (c *FakeDocdbClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DocdbCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(docdbclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.DocdbCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DocdbCluster), err
}
