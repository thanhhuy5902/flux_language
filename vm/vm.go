package vm

import (
	"context"
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/core"
	FluxIO "github.com/thanhhuy5902/flux_language/io"
	fluxParser "github.com/thanhhuy5902/flux_language/parser"
	"github.com/thanhhuy5902/flux_language/shared"
	"os"
	"time"
)

type FluxVirtualMachine struct {
	parser   *fluxParser.FluxProgramParser
	varTable *core.VarTable
}

func NewFluxVirtualMachine() *FluxVirtualMachine {
	return &FluxVirtualMachine{}

}

func (f *FluxVirtualMachine) Execute(params *shared.ExecutionParams) shared.ExecutionResult {
	errorCollector := FluxIO.NewBaseErrCollector()
	elapsedTime := time.Now().UnixMilli()
	absPath, err := os.Getwd()

	if err != nil {
		elapsedTime = time.Now().UnixMilli() - elapsedTime
		return shared.ExecutionResult{
			ErrorCollector: errorCollector,
			ElapsedTime:    elapsedTime,
		}
	}
	mainFileData := ""
	if params.EntryPoint != "" {
		params.EntryPoint = absPath + "/" + params.EntryPoint
		mainFileBytes, err := os.ReadFile(params.EntryPoint)

		if err != nil {
			elapsedTime = time.Now().UnixMilli() - elapsedTime
			return shared.ExecutionResult{
				ErrorCollector: errorCollector,
				ElapsedTime:    elapsedTime,
			}
		}
		mainFileData = string(mainFileBytes)
	} else {
		mainFileData = params.SourceCode
	}

	var logger FluxIO.Logger = FluxIO.NewEmptyLogger()
	if params.Verbose {
		logger = FluxIO.NewBaseLogger()

	}
	varTable := core.NewVarTable()

	program := fluxParser.Parse(mainFileData, errorCollector, logger).GetProgram()

	executionCtx := codeobjects.NewExecutionContext(context.TODO(), varTable)

	except := program.Execute(executionCtx)

	elapsedTime = time.Now().UnixMilli() - elapsedTime
	return shared.ExecutionResult{
		ElapsedTime:      elapsedTime,
		ErrorCollector:   errorCollector,
		RuntimeException: except,
	}

}

func (f *FluxVirtualMachine) GetVarTable() *core.VarTable {
	return f.varTable
}
