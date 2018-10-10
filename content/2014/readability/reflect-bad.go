// +build OMIT

package sample // OMIT

type Layers struct {
	UI, Launch /* more fields */ string
}

func sample() { // OMIT
	layers := NewLayers(s.Entries)
	v := reflect.ValueOf(*layers) // HL
	r := v.Type()                 // type Layers  // HL
	for i := 0; i < r.NumField(); i++ {
		if e := v.Field(i).String(); e != "-" {
			eid := &pb.ExperimentId{
				Layer:        proto.String(r.Field(i).Name()),
				ExperimentId: &e,
			}
			experimentIDs = append(experimentIDs, eid)
		}
	}
} // OMIT
