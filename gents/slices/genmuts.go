package gentslices

import (
	. "github.com/go-leap/dev/go/gen"
	"github.com/metaleap/go-gent"
)

func init() {
	Gents.Mutators.Append.NameOrSuffix = "Append"
}

type GentMutatorMethods struct {
	gent.Opts

	Append gent.Variant
}

func (this *GentMutatorMethods) genAppendMethod(t *gent.Type) *SynFunc {
	return t.G.Tª.Method(this.Append.NameOrSuffix).Args(ˇ.V.OfType(t.Expr.GenRef.ArrOrSlice.Of)).Spreads().
		Doc().
		Code(
			This.Deref().Set(B.Append.Of(This.Deref(), ˇ.V).Spreads()),
		)
}

// GenerateTopLevelDecls implements `github.com/metaleap/go-gent.IGent`.
func (this *GentMutatorMethods) GenerateTopLevelDecls(ctx *gent.Ctx, t *gent.Type) (yield Syns) {
	if t.IsSlice() {
		if this.Append.Add {
			yield.Add(this.genAppendMethod(t))
		}
	}
	return
}

// EnableOrDisableAllVariantsAndOptionals implements `github.com/metaleap/go-gent.IGent`.
func (this *GentMutatorMethods) EnableOrDisableAllVariantsAndOptionals(enabled bool) {
	this.Append.Add = enabled
}