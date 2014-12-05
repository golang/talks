// +build OMIT

package sample // OMIT

type LayerExperiment struct{ Layer, Experiment string } // HL

func (t *Layers) Slice() []LayerExperiment { // HL
	return []LayerExperiment{
		{"UI", t.UI},
		{"Launch", t.Launch},
		/* more fields */
	}
}

func sample() { // OMIT
	layers := NewLayers(s.Entries).Slice() // HL
	for _, l := range layers {
		if l.Experiment != "-" {
			eid := &pb.ExperimentId{
				Layer:        proto.String(l.Layer),
				ExperimentId: proto.String(l.Experiment),
			}
			experimentIDs = append(experimentIDs, eid)
		}
	}
} // OMIT
