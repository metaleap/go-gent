package gentenums

import (
	"strconv"

	. "github.com/go-leap/dev/go/gen"
	"github.com/go-leap/str"
	"github.com/metaleap/go-gent"
)

const (
	DefaultListBothFuncName     = "Wellknown{T}{s}"
	DefaultListBothDocComment   = "{N} returns the `names` and `values` of all {n} well-known `{T}` enumerants."
	DefaultListNamesFuncName    = "Wellknown{T}Names"
	DefaultListNamesDocComment  = "{N} returns the `names` of all {n} well-known `{T}` enumerants."
	DefaultListValuesFuncName   = "Wellknown{T}Values"
	DefaultListValuesDocComment = "{N} returns the `values` of all {n} well-known `{T}` enumerants."
	DefaultListMapFuncName      = "Wellknown{T}NamesAndValues"
	DefaultListMapDocComment    = "{N} returns the `namesToValues` of all {n} well-known `{T}` enumerants."
)

func init() {
	Gents.Listers.ListBoth.DocComment, Gents.Listers.ListBoth.FuncName = DefaultListBothDocComment, DefaultListBothFuncName
	Gents.Listers.ListNames.DocComment, Gents.Listers.ListNames.FuncName = DefaultListNamesDocComment, DefaultListNamesFuncName
	Gents.Listers.ListValues.DocComment, Gents.Listers.ListValues.FuncName = DefaultListValuesDocComment, DefaultListValuesFuncName
	Gents.Listers.ListMap.DocComment, Gents.Listers.ListMap.FuncName = DefaultListMapDocComment, DefaultListMapFuncName
}

// GentListEnumerantsFuncs generates a
// `func WellknownFoos() ([]string, []Foo)`
// for each enum type-def `Foo`.
//
// An instance with illustrative defaults is in `Gents.List`.
type GentListEnumerantsFuncs struct {
	gent.Opts

	ListNames struct {
		DocComment gent.Str
		FuncName   gent.Str
	}
	ListValues struct {
		DocComment gent.Str
		FuncName   gent.Str
	}
	ListMap struct {
		DocComment gent.Str
		FuncName   gent.Str
	}
	ListBoth struct {
		DocComment gent.Str

		// eg. "Wellknown{T}{s}" with `{T}` for type name and
		// `{s}` for pluralization suffix (either "s" or "es")
		FuncName gent.Str
	}

	SkipFirst bool
	Rename    gent.Rename
}

func (this *GentListEnumerantsFuncs) genListNamesFunc(t *gent.Type, funcName string, enumerantNames Syns) *SynFunc {
	return Func(funcName).Ret("names", T.SliceOf.Strings).
		Doc(
			this.ListNames.DocComment.With("N", funcName, "T", t.Name, "n", strconv.Itoa(len(enumerantNames))),
		).
		Code(
			N("names").Set(L(enumerantNames)),
		)
}

func (this *GentListEnumerantsFuncs) genListValuesFunc(t *gent.Type, funcName string, enumerantValues Syns) *SynFunc {
	return Func(funcName).Ret("values", TrSlice(t.G.T)).
		Doc(
			this.ListValues.DocComment.With("N", funcName, "T", t.Name, "n", strconv.Itoa(len(enumerantValues))),
		).
		Code(
			N("values").Set(L(enumerantValues)),
		)
}

func (this *GentListEnumerantsFuncs) genListBothFunc(t *gent.Type, funcName string, funcNameNames string, funcNameValues string, numEnumerants int) *SynFunc {
	return Func(funcName).Ret("names", T.SliceOf.Strings).Ret("values", TrSlice(t.G.T)).
		Doc(
			this.ListBoth.DocComment.With("N", funcName, "T", t.Name, "n", strconv.Itoa(numEnumerants)),
		).
		Code(
			Names("names", "values").Set(Tup(Call(N(funcNameNames)), Call(N(funcNameValues)))),
		)
}

func (this *GentListEnumerantsFuncs) genListMapFunc(t *gent.Type, funcName string, enumerantNames Syns, enumerantValues Syns) *SynFunc {
	maptype := TrMap(T.String, t.G.T)
	return Func(funcName).Ret("namesToValues", maptype).
		Doc(
			this.ListBoth.DocComment.With("N", funcName, "T", t.Name, "n", strconv.Itoa(len(enumerantNames))),
		).
		Code(
			N("namesToValues").Set(B.Make.Of(TrMap(T.String, t.G.T), len(enumerantNames))),
			GEN_FOR(0, len(enumerantNames), 1, func(i int) ISyn {
				return N("namesToValues").At(enumerantNames[i]).Set(enumerantValues[i])
			}),
		)
}

// GenerateTopLevelDecls implements `github.com/metaleap/go-gent.IGent`.
func (this *GentListEnumerantsFuncs) GenerateTopLevelDecls(ctx *gent.Ctx, t *gent.Type) (yield Syns) {
	if t.IsEnumish() {
		names, values := make(Syns, 0, len(t.Enumish.ConstNames)), make(Syns, 0, len(t.Enumish.ConstNames))
		for i, enumerant := range t.Enumish.ConstNames {
			if renamed := enumerant; enumerant != "_" && (i > 0 || !this.SkipFirst) {
				if this.Rename != nil {
					renamed = this.Rename(ctx, t, enumerant)
				}
				if renamed != "" {
					names.Add(L(renamed))
					values.Add(t.G.T.N(enumerant))
				}
			}
		}

		fnameboth := this.ListBoth.FuncName.With("T", t.Name, "s", ustr.If(ustr.Suff(t.Name, "s"), "es", "s"))
		if ustr.Suff(fnameboth, "ys") {
			fnameboth = fnameboth[:len(fnameboth)-2] + "ies"
		}
		fnamevals, fnamenames, fnamemap :=
			this.ListValues.FuncName.With("T", t.Name), this.ListNames.FuncName.With("T", t.Name), this.ListMap.FuncName.With("T", t.Name)

		yield.Add(
			this.genListBothFunc(t, fnameboth, fnamenames, fnamevals, len(names)),
			this.genListNamesFunc(t, fnamenames, names),
			this.genListValuesFunc(t, fnamevals, values),
			this.genListMapFunc(t, fnamemap, names, values),
		)
	}
	return
}
