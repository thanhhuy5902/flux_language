package declarations

import (
	"fmt"
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/codeobjects/expression"
	"github.com/thanhhuy5902/flux_language/common"
	"github.com/thanhhuy5902/flux_language/exception"
	"strconv"
)

type VarDeclaration struct {
	*codeobjects.BaseStatement
	Name     string
	Type     common.FluxType
	RawValue string
	Expr     *expression.MathExpression
}

func (v VarDeclaration) Generate(ctx *codeobjects.GenerateContext) string {
	// to golang type
	varType := "double"
	if v.Type == common.FluxTypeString {
		varType = "string"
	} else if v.Type == common.FluxTypeBool {
		varType = "bool"
	} else if v.Type == common.FluxTypeNumber {
		varType = "double"
	}
	if v.Expr != nil {
		return fmt.Sprintf("%v %v = %v;", varType, v.Name, v.Expr.Generate(ctx))
	} else if v.RawValue != "" {
		return fmt.Sprintf("%v %v = %v;", varType, v.Name, v.RawValue)
	}
	return fmt.Sprintf("%v %v;", v.Name, varType)
}
func (v VarDeclaration) Execute(ctx *codeobjects.ExecutionContext) *exception.BaseException {
	exprCtx := ctx.Clone()
	if v.Expr != nil {
		err := v.Expr.Execute(exprCtx)
		if err != nil {
			return err
		}
		if v.Type == common.FluxTypeNumber {
			err := ctx.VarTable.SetNum(v.Name, exprCtx.NumericValue)
			if err != nil {
				return err
			}
		}
	} else if v.RawValue != "" {
		if v.Type == common.FluxTypeNumber {
			value, parseErr := strconv.ParseFloat(v.RawValue, 64)
			if parseErr != nil {
				return &exception.BaseException{
					MessageFmt: "Invalid number value: %s",
					Args:       []interface{}{v.RawValue},
					Line:       v.GetLine(),
					StartPos:   v.GetStartPos(),
					EndPos:     v.GetEndPos(),
				}
			}
			err := ctx.VarTable.SetNum(v.Name, value)
			if err != nil {
				return err
			}
		} else {
			// set default value
			if v.Type == common.FluxTypeNumber {
				err := ctx.VarTable.SetNum(v.Name, 0)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}

func NewVarDeclaration(line int, startPos int, endPos int, name string, type_ common.FluxType, rawValue string, expr *expression.MathExpression) *VarDeclaration {
	return &VarDeclaration{BaseStatement: &codeobjects.BaseStatement{Line: line, StartPos: startPos, EndPos: endPos}, Name: name, Type: type_, RawValue: rawValue, Expr: expr}
}
