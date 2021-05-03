// Code generated by sdkgen. DO NOT EDIT.

//nolint
package containerregistry

import (
	"context"

	"google.golang.org/grpc"

	containerregistry "github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ImageServiceClient is a containerregistry.ImageServiceClient with
// lazy GRPC connection initialization.
type ImageServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Delete implements containerregistry.ImageServiceClient
func (c *ImageServiceClient) Delete(ctx context.Context, in *containerregistry.DeleteImageRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return containerregistry.NewImageServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements containerregistry.ImageServiceClient
func (c *ImageServiceClient) Get(ctx context.Context, in *containerregistry.GetImageRequest, opts ...grpc.CallOption) (*containerregistry.Image, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return containerregistry.NewImageServiceClient(conn).Get(ctx, in, opts...)
}

// List implements containerregistry.ImageServiceClient
func (c *ImageServiceClient) List(ctx context.Context, in *containerregistry.ListImagesRequest, opts ...grpc.CallOption) (*containerregistry.ListImagesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return containerregistry.NewImageServiceClient(conn).List(ctx, in, opts...)
}

type ImageIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *ImageServiceClient
	request *containerregistry.ListImagesRequest

	items []*containerregistry.Image
}

func (c *ImageServiceClient) ImageIterator(ctx context.Context, folderId string, opts ...grpc.CallOption) *ImageIterator {
	return &ImageIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &containerregistry.ListImagesRequest{
			FolderId: folderId,
			PageSize: 1000,
		},
	}
}

func (it *ImageIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Images
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *ImageIterator) Value() *containerregistry.Image {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *ImageIterator) Error() error {
	return it.err
}