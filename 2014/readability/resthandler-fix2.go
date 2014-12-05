// +build OMIT

package resthandler // OMIT

func finishStatus(r Result, complete bool) int {
	if !complete {
		return http.StatusAccepted // HL
	}
	if stat, ok := r.Object.(*api.Status); ok && stat.Code != 0 {
		return stat.Code // HL
	}
	if r.Created {
		return http.StatusCreated // HL
	}
	return http.StatusOK // HL
}

func (h *RESTHandler) finishReq(op *Operation, w http.ResponseWriter, req *http.Request) {
	result, complete := op.StatusOrResult()
	status := finishStatus(result, complete)     // HL
	writeJSON(status, h.codec, result.Object, w) // HL
}
