// +build ignore,OMIT

package download

func (s *Server) handleDownload(w http.ResponseWriter, r *http.Request) {
	s.addActiveDownloadTotal(1)
	defer s.addActiveDownloadTotal(-1)
	if !isGetOrHead(w, r) {
		return
	}
	uctx, err := s.newUserContext(r)
	// ...
	pl, cacheable, err := s.chooseValidPayloadToDownload(uctx)
	// ...
	content, err := pl.content()
	// ...
	defer content.Close()
	w.Header().Set("Content-Type", pl.mimeType())
	if etag := pl.etag(); etag != "" {
		w.Header().Set("Etag", strconv.Quote(etag))
	}
	if cacheable {
		w.Header().Set("Expires", pl.expirationTime())
	}
	readSeeker := io.NewSectionReader(content, 0, content.Size())
	http.ServeContent(w, r, "", pl.lastModifiedTime(), readSeeker)
}
