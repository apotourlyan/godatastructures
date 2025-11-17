package panics

import (
	"fmt"

	"github.com/apotourlyan/godatastructures/internal/utilities/constraints"
)

func CatchPanic(f func()) (panicked bool, message string) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			message = fmt.Sprint(r)
		}
	}()
	f()
	return false, ""
}

func RequireNonNegative[T constraints.Numeric](pval T, pname string) {
	if pval < 0 {
		panic(fmt.Sprintf("%q must be >= 0, got %v", pname, pval))
	}
}

func RequireEqualTo[T constraints.Numeric](pval T, limit T, pname string) {
	if pval != limit {
		panic(fmt.Sprintf("%q must be == %v, got %v", pname, limit, pval))
	}
}

func RequireLessThan[T constraints.Numeric](pval T, limit T, pname string) {
	if pval >= limit {
		panic(fmt.Sprintf("%q must be < %v, got %v", pname, limit, pval))
	}
}

func RequireLessThanOrEqualTo[T constraints.Numeric](pval T, limit T, pname string) {
	if pval > limit {
		panic(fmt.Sprintf("%q must be <= %v, got %v", pname, limit, pval))
	}
}
