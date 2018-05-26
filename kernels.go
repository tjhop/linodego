package linodego

import (
	"fmt"

	"github.com/go-resty/resty"
)

// LinodeKernel represents a linode kernel object
type LinodeKernel struct {
	ID           string
	Label        string
	Version      string
	KVM          bool
	XEN          bool
	Architecture string
	PVOPS        bool
}

// LinodeKernelsPagedResponse represents a linode kernels API response for listing
type LinodeKernelsPagedResponse struct {
	*PagedResponse
	Data []*LinodeKernel
}

// ListKernels lists linode kernels
func (c *Client) ListKernels(opts *ListOptions) ([]*LinodeKernel, error) {
	response := LinodeKernelsPagedResponse{}
	err := c.listHelper(&response, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (LinodeKernelsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Kernels.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

func (resp *LinodeKernelsPagedResponse) appendData(r *LinodeKernelsPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

func (LinodeKernelsPagedResponse) setResult(r *resty.Request) {
	r.SetResult(LinodeKernelsPagedResponse{})
}

// GetKernel gets the kernel with the provided ID
func (c *Client) GetKernel(kernelID string) (*LinodeKernel, error) {
	e, err := c.Kernels.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%s", e, kernelID)
	r, err := c.R().
		SetResult(&LinodeKernel{}).
		Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*LinodeKernel), nil
}
