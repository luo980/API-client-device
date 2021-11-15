/*
Copyright The KubeEdge Authors.

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

package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "restandxorm/scheme"
	v1alpha2 "restandxorm/v1alpha2"
)

// DeviceModelsGetter has a method to return a DeviceModelInterface.
// A group's client should implement this interface.
type DeviceModelsGetter interface {
	DeviceModels(namespace string) DeviceModelInterface
}

// DeviceModelInterface has methods to work with DeviceModel resources.
type DeviceModelInterface interface {
	Create(ctx context.Context, deviceModel *v1alpha2.DeviceModel, opts v1.CreateOptions) (*v1alpha2.DeviceModel, error)
	Update(ctx context.Context, deviceModel *v1alpha2.DeviceModel, opts v1.UpdateOptions) (*v1alpha2.DeviceModel, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.DeviceModel, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.DeviceModelList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.DeviceModel, err error)
	DeviceModelExpansion
}

// deviceModels implements DeviceModelInterface
type deviceModels struct {
	client rest.Interface
	ns     string
}

// newDeviceModels returns a DeviceModels
func newDeviceModels(c *DevicesV1alpha2Client, namespace string) *deviceModels {
	return &deviceModels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deviceModel, and returns the corresponding deviceModel object, and an error if there is any.
func (c *deviceModels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.DeviceModel, err error) {
	result = &v1alpha2.DeviceModel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicemodels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeviceModels that match those selectors.
func (c *deviceModels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.DeviceModelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.DeviceModelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicemodels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deviceModels.
func (c *deviceModels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devicemodels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deviceModel and creates it.  Returns the server's representation of the deviceModel, and an error, if there is any.
func (c *deviceModels) Create(ctx context.Context, deviceModel *v1alpha2.DeviceModel, opts v1.CreateOptions) (result *v1alpha2.DeviceModel, err error) {
	result = &v1alpha2.DeviceModel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devicemodels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deviceModel).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deviceModel and updates it. Returns the server's representation of the deviceModel, and an error, if there is any.
func (c *deviceModels) Update(ctx context.Context, deviceModel *v1alpha2.DeviceModel, opts v1.UpdateOptions) (result *v1alpha2.DeviceModel, err error) {
	result = &v1alpha2.DeviceModel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicemodels").
		Name(deviceModel.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deviceModel).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deviceModel and deletes it. Returns an error if one occurs.
func (c *deviceModels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicemodels").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deviceModels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicemodels").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deviceModel.
func (c *deviceModels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.DeviceModel, err error) {
	result = &v1alpha2.DeviceModel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devicemodels").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
