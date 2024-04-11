package declarations

import (
	"context"
	"github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/codeobjects/expression"
	"github.com/thanhhuy5902/flux_language/codeobjects/expression/operators"
	"github.com/thanhhuy5902/flux_language/common"
	"github.com/thanhhuy5902/flux_language/core"
	"testing"
)

func TestVarDeclaration_Execute(t *testing.T) {
	varTable := core.NewVarTable()
	t.Run("Create a new variable", func(t *testing.T) {
		varDeclaration := NewVarDeclaration(1, 1, 1, "testVar", common.FluxTypeNumber, "",
			expression.NewMathExpression(1, 1, 1,
				expression.NewNumericExpression(1, 1, 1,
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(3)),
					operators.NewAdd(1, 1, 1, nil, nil),
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(2)), 0),
			),
		)
		ctx := codeobjects.NewExecutionContext(context.Background(), varTable)
		err := varDeclaration.Execute(ctx)
		if err != nil {
			t.Errorf(err.MessageFmt, err.Args...)
		}
		testVar, except := varTable.GetNumValue("testVar")
		if except != nil {
			t.Errorf("Expected nil, got %v", except)
		}
		if testVar != 5 {
			t.Errorf("Expected 5, got %v", testVar)
		}
	})

	t.Run("Create a new variable", func(t *testing.T) {
		varDeclaration := NewVarDeclaration(1, 1, 1, "testVar", common.FluxTypeNumber, "",
			expression.NewMathExpression(1, 1, 1,
				expression.NewNumericExpression(1, 1, 1,
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(3)),
					operators.NewSub(1, 1, 1, nil, nil),
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(2)), 0),
			),
		)
		ctx := codeobjects.NewExecutionContext(context.Background(), varTable)
		err := varDeclaration.Execute(ctx)
		if err != nil {
			t.Errorf(err.MessageFmt, err.Args...)
		}
		testVar, except := varTable.GetNumValue("testVar")
		if except != nil {
			t.Errorf("Expected nil, got %v", except)
		}
		if testVar != 1 {
			t.Errorf("Expected 5, got %v", testVar)
		}
	})

	t.Run("Create a new variable", func(t *testing.T) {
		varDeclaration := NewVarDeclaration(1, 1, 1, "testVar", common.FluxTypeNumber, "",
			expression.NewMathExpression(1, 1, 1,
				expression.NewNumericExpression(1, 1, 1,
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(3)),
					operators.NewMulti(1, 1, 1, nil, nil),
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(2)), 0),
			),
		)
		ctx := codeobjects.NewExecutionContext(context.Background(), varTable)
		err := varDeclaration.Execute(ctx)
		if err != nil {
			t.Errorf(err.MessageFmt, err.Args...)
		}
		testVar, except := varTable.GetNumValue("testVar")
		if except != nil {
			t.Errorf("Expected nil, got %v", except)
		}
		if testVar != 6 {
			t.Errorf("Expected 5, got %v", testVar)
		}
	})

	t.Run("Create a new variable", func(t *testing.T) {
		varDeclaration := NewVarDeclaration(1, 1, 1, "testVar", common.FluxTypeNumber, "",
			expression.NewMathExpression(1, 1, 1,
				expression.NewNumericExpression(1, 1, 1,
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(3)),
					operators.NewDivision(1, 1, 1, nil, nil),
					expression.NewNumericExpression(1, 1, 1, nil, operators.NullOp, nil, float64(2)), 0),
			),
		)
		ctx := codeobjects.NewExecutionContext(context.Background(), varTable)
		err := varDeclaration.Execute(ctx)
		if err != nil {
			t.Errorf(err.MessageFmt, err.Args...)
		}
		testVar, except := varTable.GetNumValue("testVar")
		if except != nil {
			t.Errorf("Expected nil, got %v", except)
		}
		if testVar != 1.5{
			t.Errorf("Expected 5, got %v", testVar)
		}
	})
}
