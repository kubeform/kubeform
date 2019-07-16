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

// FakeRedshiftSnapshotCopyGrants implements RedshiftSnapshotCopyGrantInterface
type FakeRedshiftSnapshotCopyGrants struct {
	Fake *FakeAwsV1alpha1
}

var redshiftsnapshotcopygrantsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "redshiftsnapshotcopygrants"}

var redshiftsnapshotcopygrantsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "RedshiftSnapshotCopyGrant"}

// Get takes name of the redshiftSnapshotCopyGrant, and returns the corresponding redshiftSnapshotCopyGrant object, and an error if there is any.
func (c *FakeRedshiftSnapshotCopyGrants) Get(name string, options v1.GetOptions) (result *v1alpha1.RedshiftSnapshotCopyGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(redshiftsnapshotcopygrantsResource, name), &v1alpha1.RedshiftSnapshotCopyGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftSnapshotCopyGrant), err
}

// List takes label and field selectors, and returns the list of RedshiftSnapshotCopyGrants that match those selectors.
func (c *FakeRedshiftSnapshotCopyGrants) List(opts v1.ListOptions) (result *v1alpha1.RedshiftSnapshotCopyGrantList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(redshiftsnapshotcopygrantsResource, redshiftsnapshotcopygrantsKind, opts), &v1alpha1.RedshiftSnapshotCopyGrantList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RedshiftSnapshotCopyGrantList{ListMeta: obj.(*v1alpha1.RedshiftSnapshotCopyGrantList).ListMeta}
	for _, item := range obj.(*v1alpha1.RedshiftSnapshotCopyGrantList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redshiftSnapshotCopyGrants.
func (c *FakeRedshiftSnapshotCopyGrants) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(redshiftsnapshotcopygrantsResource, opts))
}

// Create takes the representation of a redshiftSnapshotCopyGrant and creates it.  Returns the server's representation of the redshiftSnapshotCopyGrant, and an error, if there is any.
func (c *FakeRedshiftSnapshotCopyGrants) Create(redshiftSnapshotCopyGrant *v1alpha1.RedshiftSnapshotCopyGrant) (result *v1alpha1.RedshiftSnapshotCopyGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(redshiftsnapshotcopygrantsResource, redshiftSnapshotCopyGrant), &v1alpha1.RedshiftSnapshotCopyGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftSnapshotCopyGrant), err
}

// Update takes the representation of a redshiftSnapshotCopyGrant and updates it. Returns the server's representation of the redshiftSnapshotCopyGrant, and an error, if there is any.
func (c *FakeRedshiftSnapshotCopyGrants) Update(redshiftSnapshotCopyGrant *v1alpha1.RedshiftSnapshotCopyGrant) (result *v1alpha1.RedshiftSnapshotCopyGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(redshiftsnapshotcopygrantsResource, redshiftSnapshotCopyGrant), &v1alpha1.RedshiftSnapshotCopyGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftSnapshotCopyGrant), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRedshiftSnapshotCopyGrants) UpdateStatus(redshiftSnapshotCopyGrant *v1alpha1.RedshiftSnapshotCopyGrant) (*v1alpha1.RedshiftSnapshotCopyGrant, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(redshiftsnapshotcopygrantsResource, "status", redshiftSnapshotCopyGrant), &v1alpha1.RedshiftSnapshotCopyGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftSnapshotCopyGrant), err
}

// Delete takes name of the redshiftSnapshotCopyGrant and deletes it. Returns an error if one occurs.
func (c *FakeRedshiftSnapshotCopyGrants) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(redshiftsnapshotcopygrantsResource, name), &v1alpha1.RedshiftSnapshotCopyGrant{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedshiftSnapshotCopyGrants) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(redshiftsnapshotcopygrantsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RedshiftSnapshotCopyGrantList{})
	return err
}

// Patch applies the patch and returns the patched redshiftSnapshotCopyGrant.
func (c *FakeRedshiftSnapshotCopyGrants) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedshiftSnapshotCopyGrant, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(redshiftsnapshotcopygrantsResource, name, pt, data, subresources...), &v1alpha1.RedshiftSnapshotCopyGrant{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RedshiftSnapshotCopyGrant), err
}
