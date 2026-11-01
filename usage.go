package flags

import (
	"cmp"
	"flag"
	"fmt"
	"maps"
	"slices"
)

type Flag struct {
	flag      *flag.Flag
	name      string
	shorthand string
}

func (f *Flag) AddName(name string) {
	if len(name) < len(f.flag.Name) {
		f.shorthand = name
	} else {
		f.shorthand = f.flag.Name
		f.name = name
	}
}

func Usage(fs *flag.FlagSet) func() {
	return func() {
		flags := make(map[string]*Flag)

		fs.VisitAll(func(f *flag.Flag) {
			usageSha := Sha(f.Usage)

			if exist, ok := flags[usageSha]; ok {
				exist.AddName(f.Name)
			} else {
				flags[usageSha] = &Flag{
					name: f.Name,
					flag: f,
				}
			}
		})

		var (
			maxTypeLen      int
			maxNameLen      int
			maxShorthandLen int
		)

		output := fs.Output()

		if name := fs.Name(); len(name) > 0 {
			_, _ = fmt.Fprintf(output, "Usage of %s:\n", fs.Name())
		} else {
			_, _ = fmt.Fprint(output, "Usage:\n")
		}

		items := slices.SortedFunc(maps.Values(flags), func(a, b *Flag) int {
			return cmp.Compare(a.name, b.name)
		})

		for _, item := range items {
			maxNameLen = max(maxNameLen, len(item.name))
			maxShorthandLen = max(maxShorthandLen, len(item.shorthand))

			flagType, _ := flag.UnquoteUsage(item.flag)
			maxTypeLen = max(maxTypeLen, len(flagType))
		}

		if maxShorthandLen > 0 {
			maxShorthandLen += 3
		}

		for _, item := range items {
			flagType, usage := flag.UnquoteUsage(item.flag)

			var shorthand string
			if len(item.shorthand) > 0 {
				shorthand = "-" + item.shorthand + ", "
			}

			_, _ = fmt.Fprintf(output, "  %-*s--%-*s  %-*s  %s", maxShorthandLen, shorthand, maxNameLen, item.name, maxTypeLen, flagType, usage)

			if defaultValue := item.flag.DefValue; len(defaultValue) > 0 {
				if flagType == "string" {
					_, _ = fmt.Fprintf(output, " (default %q)", defaultValue)
				} else {
					_, _ = fmt.Fprintf(output, " (default %v)", defaultValue)
				}
			}

			_, _ = fmt.Fprint(output, "\n")
		}
	}
}
