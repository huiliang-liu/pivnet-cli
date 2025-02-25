package pivnetversions

import (
	"fmt"
	"io"

	"github.com/olekukonko/tablewriter"
	"github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/pivnet-cli/errorhandler"
	"github.com/pivotal-cf/pivnet-cli/printer"
	"github.com/pivotal-cf/pivnet-cli/semver"
)

//go:generate counterfeiter . PivnetClient
type PivnetClient interface {
	PivnetVersions() (pivnet.PivnetVersions, error)
}

type PivnetVersionsClient struct {
	pivnetClient PivnetClient
	eh           errorhandler.ErrorHandler
	format       string
	outputWriter io.Writer
	printer      printer.Printer
}

func NewPivnetVersionsClient(
	pivnetClient PivnetClient,
	eh errorhandler.ErrorHandler,
	format string,
	outputWriter io.Writer,
	printer printer.Printer,
) *PivnetVersionsClient {
	return &PivnetVersionsClient{
		pivnetClient: pivnetClient,
		eh:           eh,
		format:       format,
		outputWriter: outputWriter,
		printer:      printer,
	}
}

func (c *PivnetVersionsClient) List() error {
	pivnetVersions, err := c.pivnetClient.PivnetVersions()
	if err != nil {
		return c.eh.HandleError(err)
	}

	switch c.format {
	case printer.PrintAsTable:
		table := tablewriter.NewWriter(c.outputWriter)
		table.SetHeader([]string{"Pivnet Component", "Latest Release Version"})
		table.Append([]string{"Pivnet CLI", pivnetVersions.PivnetCliVersion})
		table.Append([]string{"Pivnet Resource", pivnetVersions.PivnetResourceVersion})
		table.Render()
		return nil
	case printer.PrintAsJSON:
		return c.printer.PrintJSON(pivnetVersions)
	case printer.PrintAsYAML:
		return c.printer.PrintYAML(pivnetVersions)
	}

	return nil
}

func (c *PivnetVersionsClient) Warn(currentVersion string) string {
	pivnetVersions, err := c.pivnetClient.PivnetVersions()
	if err == nil {
		comparison, err := semver.Compare(currentVersion, pivnetVersions.PivnetCliVersion)
		if err == nil && comparison < 0 {
			return fmt.Sprintf("Warning: Your version of Pivnet CLI (%s) does not match the currently released version (%s).", currentVersion, pivnetVersions.PivnetCliVersion)
		} else {
			return ""
		}
	} else {
		return ""
	}
}
