// +build OMIT

package sampling

import (
	servicepb "foo/bar/service_proto"
)

type SamplingServer struct {
	// some fields
}

func (server *SamplingServer) SampleMetrics(sampleRequest *servicepb.Request, sampleResponse *servicepb.Response, latency time.Duration) { // HL
	// some code
}
