package jobresume

import (
	"net/http"

	"github.com/yuansuan/ticp/common/openapi-go/common"
	"github.com/yuansuan/ticp/common/openapi-go/utils/xhttp"
	jr "github.com/yuansuan/ticp/common/project-root-api/job/v1/jobtransmitresume"
)

type API func(options ...Option) (*jr.Response, error)

func New(hc *xhttp.Client) API {
	return func(options ...Option) (*jr.Response, error) {
		req, err := NewRequest(options)
		if err != nil {
			return nil, err
		}
		resolver := hc.Prepare(xhttp.NewRequestBuilder().
			Method(http.MethodPatch).
			URI("/api/jobs/" + req.JobID + "/transmit/resume"))

		ret := new(jr.Response)
		err = resolver.Resolve(func(resp *http.Response) error {
			return common.ParseResp(resp, ret)
		})

		return ret, err
	}
}

func NewRequest(options []Option) (*jr.Request, error) {
	req := &jr.Request{}

	for _, option := range options {
		if err := option(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}

type Option func(req *jr.Request) error

func (api API) JobId(JobId string) Option {
	return func(req *jr.Request) error {
		req.JobID = JobId
		return nil
	}
}
