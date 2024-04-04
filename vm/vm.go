package vm

import (
	"github.com/thanhhuy5902/flux_language/parsing"
	"github.com/thanhhuy5902/flux_language/shared"
	"github.com/thanhhuy5902/flux_language/vm/io"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"time"
)

type FluxVirutalMachine struct {
	traverser parsing.FluxListener
}

func NewFluxVirtualMachine() *FluxVirutalMachine {
	return &FluxVirutalMachine{}

}

func (f *FluxVirutalMachine) Execute(params *shared.ExecutionParams) shared.ExecutionResult {
	elapsedTime := time.Now().UnixMilli()
	absPath, err := os.Getwd()

	if err != nil {
		elapsedTime = time.Now().UnixMilli() - elapsedTime
		return shared.ExecutionResult{
			Error: err.Error(),
			ElapsedTime: elapsedTime,
		}
	}
	params.EntryPoint = absPath + "/" + params.EntryPoint
	mainFileData, err := os.ReadFile(params.EntryPoint)

	if err != nil {
		elapsedTime = time.Now().UnixMilli() - elapsedTime
		return shared.ExecutionResult{
			Error: err.Error(),
			ElapsedTime: elapsedTime,
		}
	}
	input := antlr.NewInputStream(string(mainFileData))
	lexer := parsing.NewPrimitives(input)
	parser := parsing.NewFlux(antlr.NewCommonTokenStream(lexer, 0))
	var logger io.Logger = io.NewEmptyLogger()
	if params.Verbose {
		logger = io.NewBaseLogger()
	}
	f.traverser = NewFluxTraverser(logger)
	parser.AddParseListener(f.traverser)
	parser.Program()
	elapsedTime = time.Now().UnixMilli() - elapsedTime
	return shared.ExecutionResult{
		ElapsedTime: elapsedTime,
	}

}