package output

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/go-utils/gopackages"
	"github.com/labstack/echo/v4"
	"github.com/olekukonko/tablewriter"
)

// Do - Display echo routing in a table
func Do(routes []*echo.Route) {
	var rootPackage string

	{
		var goModPath string
		pwd, err := os.Getwd()
		if err != nil {
			goto STEP
		}

		goModPath, err = gopackages.GetGoModPath(pwd)
		if err != nil {
			goto STEP
		}

		rootPackage, err = gopackages.GetGoModule(goModPath)
		if err != nil {
			goto STEP
		}

		rootPackage += "/"
	}

STEP:
	sort.SliceStable(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.Append([]string{"Method", "Path", "FunctionName"})

	var cnt int
	for _, r := range routes {
		if r.Name == "github.com/labstack/echo/v4.glob..func1" {
			continue
		}
		cnt++

		name := strings.TrimPrefix(r.Name, rootPackage)
		if strings.Contains(name, ".func1") {
			name = strings.Split(name, ".func1")[0]
		}
		column := []string{r.Method, r.Path, name}

		table.Append(column)
	}

	totalEndpoints := fmt.Sprintf("%d Endpoints", cnt)
	table.Append([]string{"TOTAL", totalEndpoints, ""})

	table.SetAutoMergeCells(true)
	table.SetRowLine(true)

	table.Render()
}
