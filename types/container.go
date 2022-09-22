package types

type Iterable[I any] interface {
	Begin() I
	Next(I) I
	End(I) bool
}
