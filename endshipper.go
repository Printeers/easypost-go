package easypost

import (
	"context"
	"net/http"
)

// ListEndShipperResult holds the results from the list EndShippers API.
type ListEndShipperResult struct {
	EndShippers []*Address `json:"endshippers,omitempty"`
	// HasMore indicates if there are more responses to be fetched. If True,
	// additional responses can be fetched by updating the ListEndShipperOptions
	// parameter's AfterID field with the ID of the last item in this object's
	// EndShippers field.
	HasMore bool `json:"has_more,omitempty"`
}

// CreateEndShipper submits a request to create a new endshipper, and returns the result.
func (c *Client) CreateEndShipper(in *Address) (out *Address, err error) {
	return c.CreateEndShipperWithContext(context.Background(), in)
}

// CreateEndShipperWithContext performs the same operation as CreateEndShipper, but
// allows specifying a context that can interrupt the request.
func (c *Client) CreateEndShipperWithContext(ctx context.Context, in *Address) (out *Address, err error) {
	wrappedParams := map[string]interface{}{
		"address": in,
	}
	err = c.post(ctx, "end_shippers", wrappedParams, &out)
	return
}

// GetEndShipper retrieves a previously created endshipper by its ID.
func (c *Client) GetEndShipper(endshipperID string) (out *Address, err error) {
	return c.GetEndShipperWithContext(context.Background(), endshipperID)
}

// GetEndShipperWithContext performs the same operation as GetEndShipper, but
// allows specifying a context that can interrupt the request.
func (c *Client) GetEndShipperWithContext(ctx context.Context, endshipperID string) (out *Address, err error) {
	err = c.get(ctx, "end_shippers/"+endshipperID, &out)
	return
}

// ListEndShippers provides a paginated result of EndShipper objects.
func (c *Client) ListEndShippers(opts *ListOptions) (out *ListEndShipperResult, err error) {
	return c.ListEndShippersWithContext(context.Background(), opts)
}

// ListEndShippersWithContext performs the same operation as ListEndShippers, but
// allows specifying a context that can interrupt the request.
func (c *Client) ListEndShippersWithContext(ctx context.Context, opts *ListOptions) (out *ListEndShipperResult, err error) {
	err = c.do(ctx, http.MethodGet, "end_shippers", c.convertOptsToURLValues(opts), &out)
	return
}

// UpdateEndShippers updates previously created endshipper
func (c *Client) UpdateEndShippers(in *Address) (out *Address, err error) {
	return c.UpdateEndShippersWithContext(context.Background(), in)
}

// UpdateEndShippersWithContext performs the same operation as UpdateEndShippers, but
// allows specifying a context that can interrupt the request.
func (c *Client) UpdateEndShippersWithContext(ctx context.Context, in *Address) (out *Address, err error) {
	wrappedParams := map[string]interface{}{
		"address": in,
	}

	err = c.do(ctx, http.MethodPut, "end_shippers/"+in.ID, wrappedParams, &out)
	return
}
