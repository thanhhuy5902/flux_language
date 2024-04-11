package shared

import (
	"github.com/thanhhuy5902/flux_language/exception"
	"github.com/thanhhuy5902/flux_language/io"
)

type ExecutionParams struct {
	EntryPoint string `json:"entryPoint"`
	Verbose    bool   `json:"verbose"`
	SourceCode string `json:"sourceCode"`
}

type ExecutionResult struct {
	ErrorCollector   io.ErrorCollector
	ElapsedTime      int64                    `json:"elapsedTime"`
	RuntimeException *exception.BaseException `json:"runtimeException"`
}
