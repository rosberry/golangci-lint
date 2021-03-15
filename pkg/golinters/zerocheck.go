package golinters

import (
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"github.com/rosberry/go-zerolog-event-sent/analyzer"
	"golang.org/x/tools/go/analysis"
)

func NewZerocheck() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"zerocheck",
		"Simple linter to check that you sent event for zerolog (call Msg, Msgf, Send)",
		[]*analysis.Analyzer{
			analyzer.Analyzer(),
		},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}