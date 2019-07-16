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

// FakeOpsworksPhpAppLayers implements OpsworksPhpAppLayerInterface
type FakeOpsworksPhpAppLayers struct {
	Fake *FakeAwsV1alpha1
}

var opsworksphpapplayersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworksphpapplayers"}

var opsworksphpapplayersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksPhpAppLayer"}

// Get takes name of the opsworksPhpAppLayer, and returns the corresponding opsworksPhpAppLayer object, and an error if there is any.
func (c *FakeOpsworksPhpAppLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(opsworksphpapplayersResource, name), &v1alpha1.OpsworksPhpAppLayer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksPhpAppLayer), err
}

// List takes label and field selectors, and returns the list of OpsworksPhpAppLayers that match those selectors.
func (c *FakeOpsworksPhpAppLayers) List(opts v1.ListOptions) (result *v1alpha1.OpsworksPhpAppLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(opsworksphpapplayersResource, opsworksphpapplayersKind, opts), &v1alpha1.OpsworksPhpAppLayerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksPhpAppLayerList{ListMeta: obj.(*v1alpha1.OpsworksPhpAppLayerList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksPhpAppLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksPhpAppLayers.
func (c *FakeOpsworksPhpAppLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(opsworksphpapplayersResource, opts))
}

// Create takes the representation of a opsworksPhpAppLayer and creates it.  Returns the server's representation of the opsworksPhpAppLayer, and an error, if there is any.
func (c *FakeOpsworksPhpAppLayers) Create(opsworksPhpAppLayer *v1alpha1.OpsworksPhpAppLayer) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(opsworksphpapplayersResource, opsworksPhpAppLayer), &v1alpha1.OpsworksPhpAppLayer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksPhpAppLayer), err
}

// Update takes the representation of a opsworksPhpAppLayer and updates it. Returns the server's representation of the opsworksPhpAppLayer, and an error, if there is any.
func (c *FakeOpsworksPhpAppLayers) Update(opsworksPhpAppLayer *v1alpha1.OpsworksPhpAppLayer) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(opsworksphpapplayersResource, opsworksPhpAppLayer), &v1alpha1.OpsworksPhpAppLayer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksPhpAppLayer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksPhpAppLayers) UpdateStatus(opsworksPhpAppLayer *v1alpha1.OpsworksPhpAppLayer) (*v1alpha1.OpsworksPhpAppLayer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(opsworksphpapplayersResource, "status", opsworksPhpAppLayer), &v1alpha1.OpsworksPhpAppLayer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksPhpAppLayer), err
}

// Delete takes name of the opsworksPhpAppLayer and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksPhpAppLayers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(opsworksphpapplayersResource, name), &v1alpha1.OpsworksPhpAppLayer{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksPhpAppLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(opsworksphpapplayersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksPhpAppLayerList{})
	return err
}

// Patch applies the patch and returns the patched opsworksPhpAppLayer.
func (c *FakeOpsworksPhpAppLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksPhpAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(opsworksphpapplayersResource, name, pt, data, subresources...), &v1alpha1.OpsworksPhpAppLayer{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksPhpAppLayer), err
}
