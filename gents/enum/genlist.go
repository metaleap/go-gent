package gentenum

import (
	"strconv"
	"strings"

	. "github.com/go-leap/dev/go/gen"
	"github.com/metaleap/go-gent"
)

// GentListEnumerantsFunc generates a
// `func WellknownFoos() ([]string, []Foo)`
// for each enum type-def `Foo`.
type GentListEnumerantsFunc struct {
	Disabled   bool
	DocComment gent.Str

	// eg. "Wellknown{T}{s}" with `{T}` for type name and
	// `{s}` for pluralization suffix (either "s" or "es")
	FuncName gent.Str
}

// GenerateTopLevelDecls implements `github.com/metaleap/go-gent.IGent`.
func (this *GentListEnumerantsFunc) GenerateTopLevelDecls(t *gent.Type) (tlDecls Syns) {
	if (!this.Disabled) && t.SeemsEnumish() {
		num, names, values := 0, make(Syns, 0, len(t.Enumish.ConstNames)), make(Syns, 0, len(t.Enumish.ConstNames))
		for _, enumerant := range t.Enumish.ConstNames {
			if enumerant != "_" {
				num, names, values = num+1, append(names, L(enumerant)), append(values, NT(enumerant, t.CodeGen.Ref))
			}
		}
		pluralsuffix := "s"
		if strings.HasSuffix(t.Name, "s") {
			pluralsuffix = "es"
		}
		fname := this.FuncName.With("{T}", t.Name, "{s}", pluralsuffix)
		if strings.HasSuffix(fname, "ys") {
			fname = fname[:len(fname)-2] + "ies"
		}
		fn := Fn(NoMethodRecv, fname, TdFunc(nil, NT("names", T.Sl.Strings), NT("values", TrSlice(t.CodeGen.Ref))),
			Set(C(N("names"), N("values")), C(L(names), L(values))),
		)
		fn.Doc.Add(this.DocComment.With("{N}", fn.Name, "{T}", t.Name, "{n}", strconv.Itoa(num)))
		tlDecls.Add(fn)
	}
	return
}