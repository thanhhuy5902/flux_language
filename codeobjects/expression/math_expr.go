package expression

import (
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/exception"
)

type MathExpression struct {
	*codeobjects.BaseStatement
	NumericExpr *NumericExpression
	GetVar      *GetVar
}

func (m MathExpression) Generate(ctx *codeobjects.GenerateContext) string {
	return m.NumericExpr.Generate(ctx)
}

func (m MathExpression) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException {
	if m.NumericExpr != nil {
		return m.NumericExpr.Execute(ctx)
	}
	return nil
}

func NewMathExpression(line int, startPos int, endPos int, numericExpr *NumericExpression) *MathExpression {
	return &MathExpression{BaseStatement: &codeobjects.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, NumericExpr: numericExpr}
}
