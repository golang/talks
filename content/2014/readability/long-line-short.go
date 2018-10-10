// +build OMIT

package sampling

import (
	spb "foo/bar/service_proto"
)

type Server struct {
	// some fields
}

func (s *Server) SampleMetrics(req *spb.Request, resp *spb.Response, latency time.Duration) { // HL
	// some code
}
