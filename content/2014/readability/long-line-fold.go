// +build OMIT

package sampling

import (
	servicepb "foo/bar/service_proto"
)

type SamplingServer struct {
	// some fields
}

func (server *SamplingServer) SampleMetrics( // HL
	sampleRequest *servicepb.Request, sampleResponse *servicepb.Response, // HL
	latency time.Duration) { // HL
	// some code
}
