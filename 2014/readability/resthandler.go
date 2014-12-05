// +build OMIT

package resthandler // OMIT

func (h *RESTHandler) finishReq(op *Operation, req *http.Request, w http.ResponseWriter) {
	result, complete := op.StatusOrResult()
	obj := result.Object
	if complete {
		status := http.StatusOK // HL
		if result.Created {
			status = http.StatusCreated // HL
		}
		switch stat := obj.(type) {
		case *api.Status:
			if stat.Code != 0 {
				status = stat.Code // HL
			}
		}
		writeJSON(status, h.codec, obj, w) // HL
	} else {
		writeJSON(http.StatusAccepted, h.codec, obj, w) // HL
	}
}
