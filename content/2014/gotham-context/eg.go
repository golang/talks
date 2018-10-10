// +build OMIT

package P

import (
	"xpkg"
	"ypkg"

	"golang.org/x/net/context"
)

func before(x xpkg.X, y ypkg.Y) error { // HL
	return x.M(y)
}

func after(x xpkg.X, y ypkg.Y) error { // HL
	return x.MContext(context.TODO(), y)
}
