//go:build ignore

package main

import (
	"cmp"
	"compress/gzip"
	_ "embed" // For go:embed.
	"fmt"
	"math"
	"net/http"
	"os"
	"slices"

	"github.com/hortbot/hortbot/internal/pkg/jsonx"
	"github.com/hortbot/hortbot/internal/pkg/must"
)

type entry struct {
	UserAgent      string  `json:"userAgent"`
	DeviceCategory string  `json:"deviceCategory"`
	Weight         float64 `json:"weight"`
}

type choice struct {
	item   string
	weight int
}

const (
	userAgentsURL = "https://unpkg.com/user-agents@1.1.216/src/user-agents.json"
	minThreshold  = 1000
)

func getEntries() []*entry {
	resp, err := http.Get(userAgentsURL)
	must.NilError(err)
	defer resp.Body.Close()

	r := must.Must(gzip.NewReader(resp.Body))

	var entries []*entry
	must.NilError(jsonx.DecodeSingle(r, &entries))
	return entries
}

func main() {
	rawEntries := getEntries()

	entries := make(map[string]float64)

	// The source data has multiple entries depending on other factors, like screen size.
	// Since we're only care about the user agents themselves, group them back together.
	for _, e := range rawEntries {
		if e.DeviceCategory != "desktop" {
			continue
		}
		entries[e.UserAgent] += e.Weight
	}

	minWeight := math.MaxFloat64
	for _, weight := range entries {
		minWeight = min(minWeight, weight)
	}

	choices := make([]choice, 0, len(entries))
	for userAgent, weight := range entries {
		choices = append(choices, choice{userAgent, int(weight / minWeight)})
	}

	// weightedrand does this already, but make startup faster by pre-sorting.
	slices.SortFunc(choices, func(i, j choice) int {
		return cmp.Compare(i.weight, j.weight)
	})

	f, err := os.OpenFile("./useragent_generated.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	must.NilError(err)
	defer f.Close()

	must.Must(f.WriteString("// Code generated by generate.go; DO NOT EDIT.\n\n"))
	must.Must(f.WriteString("package useragent\n\n"))
	must.Must(f.WriteString("import (\n"))
	must.Must(f.WriteString("\t\"github.com/hortbot/hortbot/internal/pkg/must\"\n"))
	must.Must(f.WriteString("\t\"github.com/mroth/weightedrand/v2\"\n"))
	must.Must(f.WriteString(")\n\n"))

	must.Must(f.WriteString("var (\n"))
	must.Must(f.WriteString("\tchoices = []weightedrand.Choice[string, int]{\n"))
	for _, choice := range choices {
		must.Must(f.WriteString(fmt.Sprintf("\t\t{Item: %q, Weight: %d},\n", choice.item, choice.weight)))
	}
	must.Must(f.WriteString("\t}\n\n"))
	must.Must(f.WriteString("\tchooser = must.Must(weightedrand.NewChooser(choices...))\n"))
	must.Must(f.WriteString(")\n"))
}
