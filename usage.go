package flags

import (
	"flag"
	"fmt"
	"sort"
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

		items := make([]*Flag, 0, len(flags))
		for _, item := range flags {
			index := sort.Search(len(items), func(i int) bool {
				return items[i].name > item.name
			})

			items = append(items, item)
			copy(items[index+1:], items[index:])
			items[index] = item

			if length := len(item.name); length > maxNameLen {
				maxNameLen = length
			}

			if length := len(item.shorthand); length > maxShorthandLen {
				maxShorthandLen = length
			}

			flagType, _ := flag.UnquoteUsage(item.flag)
			if length := len(flagType); length > maxTypeLen {
				maxTypeLen = length
			}
		}

		if maxShorthandLen > 0 {
			maxShorthandLen += 3
		}

		for _, item := range items {
			flagType, usage := flag.UnquoteUsage(item.flag)

			if len(item.shorthand) > 0 {
				_, _ = fmt.Fprintf(output, fmt.Sprintf("  %%-%ds--%%-%ds  %%-%ds  %%s", maxShorthandLen, maxNameLen, maxTypeLen), fmt.Sprintf("-%s, ", item.shorthand), item.name, flagType, usage)
			} else {
				_, _ = fmt.Fprintf(output, fmt.Sprintf("  %%-%ds--%%-%ds  %%-%ds  %%s", maxShorthandLen, maxNameLen, maxTypeLen), "", item.name, flagType, usage)
			}

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
