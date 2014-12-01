// +build OMIT

package before

// START OMIT
func (*ServiceA) HandleRPC(ctx context.Context, a Arg) {
	f(a)
}

func f(a Args) {
	x.M(a)
}

func (x *X) M(a Args) {
	// TODO(sameer): pass a real Context here.
	serviceB.IssueRPC(context.TODO(), a) // HL
}

// END OMIT
