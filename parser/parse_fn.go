package parser

import (
	"github.com/antlr4-go/antlr/v4"
	fluxIO "github.com/thanhhuy5902/flux_language/io"
	"github.com/thanhhuy5902/flux_language/parsing"
)

func Parse(data string, errorCollector fluxIO.ErrorCollector, logger fluxIO.Logger) *FluxProgramParser {
	input := antlr.NewInputStream(string(data))

	lexer := parsing.NewPrimitives(input)

	parser := parsing.NewFlux(antlr.NewCommonTokenStream(lexer, 0))

	fluxProgramParser := NewFluxProgramParser(logger, errorCollector)

	parser.AddParseListener(fluxProgramParser)

	parser.Program()

	return fluxProgramParser
}
