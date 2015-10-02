package cmd

type Generate struct {
  noSymbols bool
  passLength int
}

func NewGenerate() (*Generate) {
  // set defaults
  passLength := 15
  noSymbols := false
  return &Generate{noSymbols, passLength}
}
