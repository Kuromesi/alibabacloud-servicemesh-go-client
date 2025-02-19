// Copyright 2025 Alibaba Cloud Service Mesh
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	scheme "istio.io/client-go/asm/pkg/clientset/scheme"
	"context"
	"time"

	v1 "istio.io/api/alibabacloudservicemesh/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LLMProvidersGetter has a method to return a LLMProviderInterface.
// A group's client should implement this interface.
type LLMProvidersGetter interface {
	LLMProviders(namespace string) LLMProviderInterface
}

// LLMProviderInterface has methods to work with LLMProvider resources.
type LLMProviderInterface interface {
	Create(ctx context.Context, lLMProvider *v1.LLMProvider, opts metav1.CreateOptions) (*v1.LLMProvider, error)
	Update(ctx context.Context, lLMProvider *v1.LLMProvider, opts metav1.UpdateOptions) (*v1.LLMProvider, error)
	UpdateStatus(ctx context.Context, lLMProvider *v1.LLMProvider, opts metav1.UpdateOptions) (*v1.LLMProvider, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LLMProvider, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LLMProviderList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LLMProvider, err error)
	LLMProviderExpansion
}

// lLMProviders implements LLMProviderInterface
type lLMProviders struct {
	client rest.Interface
	ns     string
}

// newLLMProviders returns a LLMProviders
func newLLMProviders(c *IstioV1Client, namespace string) *lLMProviders {
	return &lLMProviders{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lLMProvider, and returns the corresponding lLMProvider object, and an error if there is any.
func (c *lLMProviders) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LLMProvider, err error) {
	result = &v1.LLMProvider{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("llmproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LLMProviders that match those selectors.
func (c *lLMProviders) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LLMProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LLMProviderList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("llmproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested lLMProviders.
func (c *lLMProviders) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("llmproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a lLMProvider and creates it.  Returns the server's representation of the lLMProvider, and an error, if there is any.
func (c *lLMProviders) Create(ctx context.Context, lLMProvider *v1.LLMProvider, opts metav1.CreateOptions) (result *v1.LLMProvider, err error) {
	result = &v1.LLMProvider{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("llmproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lLMProvider).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a lLMProvider and updates it. Returns the server's representation of the lLMProvider, and an error, if there is any.
func (c *lLMProviders) Update(ctx context.Context, lLMProvider *v1.LLMProvider, opts metav1.UpdateOptions) (result *v1.LLMProvider, err error) {
	result = &v1.LLMProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("llmproviders").
		Name(lLMProvider.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lLMProvider).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *lLMProviders) UpdateStatus(ctx context.Context, lLMProvider *v1.LLMProvider, opts metav1.UpdateOptions) (result *v1.LLMProvider, err error) {
	result = &v1.LLMProvider{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("llmproviders").
		Name(lLMProvider.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lLMProvider).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the lLMProvider and deletes it. Returns an error if one occurs.
func (c *lLMProviders) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("llmproviders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *lLMProviders) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("llmproviders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched lLMProvider.
func (c *lLMProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LLMProvider, err error) {
	result = &v1.LLMProvider{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("llmproviders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
