package flag

import (
	"code.cloudfoundry.org/cli/v9/types"
)

type Buildpack struct {
	types.FilteredString
}

func (b *Buildpack) UnmarshalFlag(val string) error {
	b.ParseValue(val)
	return nil
}
