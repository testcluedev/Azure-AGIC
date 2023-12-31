/*
Copyright The Kubernetes Authors.

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
	"context"

	v1beta1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/azureapplicationgatewaybackendpool/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAzureApplicationGatewayBackendPools implements AzureApplicationGatewayBackendPoolInterface
type FakeAzureApplicationGatewayBackendPools struct {
	Fake *FakeAzureapplicationgatewaybackendpoolsV1beta1
}

var azureapplicationgatewaybackendpoolsResource = schema.GroupVersionResource{Group: "azureapplicationgatewaybackendpools.appgw.ingress.azure.io", Version: "v1beta1", Resource: "azureapplicationgatewaybackendpools"}

var azureapplicationgatewaybackendpoolsKind = schema.GroupVersionKind{Group: "azureapplicationgatewaybackendpools.appgw.ingress.azure.io", Version: "v1beta1", Kind: "AzureApplicationGatewayBackendPool"}

// Get takes name of the azureApplicationGatewayBackendPool, and returns the corresponding azureApplicationGatewayBackendPool object, and an error if there is any.
func (c *FakeAzureApplicationGatewayBackendPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.AzureApplicationGatewayBackendPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(azureapplicationgatewaybackendpoolsResource, name), &v1beta1.AzureApplicationGatewayBackendPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AzureApplicationGatewayBackendPool), err
}

// List takes label and field selectors, and returns the list of AzureApplicationGatewayBackendPools that match those selectors.
func (c *FakeAzureApplicationGatewayBackendPools) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.AzureApplicationGatewayBackendPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(azureapplicationgatewaybackendpoolsResource, azureapplicationgatewaybackendpoolsKind, opts), &v1beta1.AzureApplicationGatewayBackendPoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.AzureApplicationGatewayBackendPoolList{ListMeta: obj.(*v1beta1.AzureApplicationGatewayBackendPoolList).ListMeta}
	for _, item := range obj.(*v1beta1.AzureApplicationGatewayBackendPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureApplicationGatewayBackendPools.
func (c *FakeAzureApplicationGatewayBackendPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(azureapplicationgatewaybackendpoolsResource, opts))
}

// Create takes the representation of a azureApplicationGatewayBackendPool and creates it.  Returns the server's representation of the azureApplicationGatewayBackendPool, and an error, if there is any.
func (c *FakeAzureApplicationGatewayBackendPools) Create(ctx context.Context, azureApplicationGatewayBackendPool *v1beta1.AzureApplicationGatewayBackendPool, opts v1.CreateOptions) (result *v1beta1.AzureApplicationGatewayBackendPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(azureapplicationgatewaybackendpoolsResource, azureApplicationGatewayBackendPool), &v1beta1.AzureApplicationGatewayBackendPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AzureApplicationGatewayBackendPool), err
}

// Update takes the representation of a azureApplicationGatewayBackendPool and updates it. Returns the server's representation of the azureApplicationGatewayBackendPool, and an error, if there is any.
func (c *FakeAzureApplicationGatewayBackendPools) Update(ctx context.Context, azureApplicationGatewayBackendPool *v1beta1.AzureApplicationGatewayBackendPool, opts v1.UpdateOptions) (result *v1beta1.AzureApplicationGatewayBackendPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(azureapplicationgatewaybackendpoolsResource, azureApplicationGatewayBackendPool), &v1beta1.AzureApplicationGatewayBackendPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AzureApplicationGatewayBackendPool), err
}

// Delete takes name of the azureApplicationGatewayBackendPool and deletes it. Returns an error if one occurs.
func (c *FakeAzureApplicationGatewayBackendPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(azureapplicationgatewaybackendpoolsResource, name), &v1beta1.AzureApplicationGatewayBackendPool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureApplicationGatewayBackendPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(azureapplicationgatewaybackendpoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.AzureApplicationGatewayBackendPoolList{})
	return err
}

// Patch applies the patch and returns the patched azureApplicationGatewayBackendPool.
func (c *FakeAzureApplicationGatewayBackendPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.AzureApplicationGatewayBackendPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(azureapplicationgatewaybackendpoolsResource, name, pt, data, subresources...), &v1beta1.AzureApplicationGatewayBackendPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.AzureApplicationGatewayBackendPool), err
}
