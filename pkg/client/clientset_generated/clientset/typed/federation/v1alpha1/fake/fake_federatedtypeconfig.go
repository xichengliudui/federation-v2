/*
Copyright 2018 The Kubernetes Authors.

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
package fake

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedTypeConfigs implements FederatedTypeConfigInterface
type FakeFederatedTypeConfigs struct {
	Fake *FakeFederationV1alpha1
}

var federatedtypeconfigsResource = schema.GroupVersionResource{Group: "federation.k8s.io", Version: "v1alpha1", Resource: "federatedtypeconfigs"}

var federatedtypeconfigsKind = schema.GroupVersionKind{Group: "federation.k8s.io", Version: "v1alpha1", Kind: "FederatedTypeConfig"}

// Get takes name of the federatedTypeConfig, and returns the corresponding federatedTypeConfig object, and an error if there is any.
func (c *FakeFederatedTypeConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedTypeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(federatedtypeconfigsResource, name), &v1alpha1.FederatedTypeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedTypeConfig), err
}

// List takes label and field selectors, and returns the list of FederatedTypeConfigs that match those selectors.
func (c *FakeFederatedTypeConfigs) List(opts v1.ListOptions) (result *v1alpha1.FederatedTypeConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(federatedtypeconfigsResource, federatedtypeconfigsKind, opts), &v1alpha1.FederatedTypeConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FederatedTypeConfigList{}
	for _, item := range obj.(*v1alpha1.FederatedTypeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedTypeConfigs.
func (c *FakeFederatedTypeConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(federatedtypeconfigsResource, opts))
}

// Create takes the representation of a federatedTypeConfig and creates it.  Returns the server's representation of the federatedTypeConfig, and an error, if there is any.
func (c *FakeFederatedTypeConfigs) Create(federatedTypeConfig *v1alpha1.FederatedTypeConfig) (result *v1alpha1.FederatedTypeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(federatedtypeconfigsResource, federatedTypeConfig), &v1alpha1.FederatedTypeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedTypeConfig), err
}

// Update takes the representation of a federatedTypeConfig and updates it. Returns the server's representation of the federatedTypeConfig, and an error, if there is any.
func (c *FakeFederatedTypeConfigs) Update(federatedTypeConfig *v1alpha1.FederatedTypeConfig) (result *v1alpha1.FederatedTypeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(federatedtypeconfigsResource, federatedTypeConfig), &v1alpha1.FederatedTypeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedTypeConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedTypeConfigs) UpdateStatus(federatedTypeConfig *v1alpha1.FederatedTypeConfig) (*v1alpha1.FederatedTypeConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(federatedtypeconfigsResource, "status", federatedTypeConfig), &v1alpha1.FederatedTypeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedTypeConfig), err
}

// Delete takes name of the federatedTypeConfig and deletes it. Returns an error if one occurs.
func (c *FakeFederatedTypeConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(federatedtypeconfigsResource, name), &v1alpha1.FederatedTypeConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedTypeConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(federatedtypeconfigsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FederatedTypeConfigList{})
	return err
}

// Patch applies the patch and returns the patched federatedTypeConfig.
func (c *FakeFederatedTypeConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedTypeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(federatedtypeconfigsResource, name, data, subresources...), &v1alpha1.FederatedTypeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedTypeConfig), err
}