// +build OMIT

package after

// START OMIT
func (*ServiceA) HandleRPC(ctx context.Context, a Arg) {
	f(ctx, a)
}

func f(ctx context.Context, a Args) { // HL
	x.M(ctx, a)
}

func (x *X) M(ctx context.Context, a Args) { // HL
	serviceB.IssueRPC(ctx, a)
}

// END OMIT
