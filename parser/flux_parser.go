package parser

import (
	"github.com/antlr4-go/antlr/v4"
	fluxCodeObjects "github.com/thanhhuy5902/flux_language/codeobjects"
	"github.com/thanhhuy5902/flux_language/codeobjects/declarations"
	fluxExpr "github.com/thanhhuy5902/flux_language/codeobjects/expression"
	"github.com/thanhhuy5902/flux_language/codeobjects/expression/operators"
	"github.com/thanhhuy5902/flux_language/common"
	"github.com/thanhhuy5902/flux_language/exception"
	fluxIO "github.com/thanhhuy5902/flux_language/io"
	"github.com/thanhhuy5902/flux_language/parsing"
	"github.com/thanhhuy5902/flux_language/utils"
	"strconv"
)

type FluxProgramParser struct {
	logger          fluxIO.Logger
	errorCollector  fluxIO.ErrorCollector
	program         *fluxCodeObjects.Program
	stackStatements *utils.Stack[fluxCodeObjects.CodeObject]
	resultStack     *utils.Stack[interface{}]
	executionCtx    *fluxCodeObjects.ExecutionContext
}

func NewFluxProgramParser(logger fluxIO.Logger, errorCollector fluxIO.ErrorCollector) *FluxProgramParser {
	return &FluxProgramParser{
		logger:          logger,
		errorCollector:  errorCollector,
		stackStatements: utils.NewStack[fluxCodeObjects.CodeObject](),
		resultStack:     utils.NewStack[interface{}](),
	}
}

func (f FluxProgramParser) GetProgram() *fluxCodeObjects.Program {
	return f.program

}

func (f FluxProgramParser) VisitTerminal(node antlr.TerminalNode) {
	f.logger.Infof("Visiting terminal node: %v", node.GetText())
}

func (f FluxProgramParser) VisitErrorNode(node antlr.ErrorNode) {
	f.logger.Errorf("Visiting error node: %v", node.GetText())
}

func (f FluxProgramParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
	f.logger.Infof("Entering rule: %v", ctx.GetText())
}

func (f FluxProgramParser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	f.logger.Infof("Exiting rule: %v", ctx.GetText())

}

func (f FluxProgramParser) EnterProgram(c *parsing.ProgramContext) {
	f.logger.Infof("Entering program")
	prog := fluxCodeObjects.NewProgram()
	f.stackStatements.Push(prog)
}
func (f FluxProgramParser) EnterStatement(c *parsing.StatementContext) {
	f.logger.Infof("Entering statement")
}

func (f FluxProgramParser) EnterExpression(c *parsing.ExpressionContext) {
	f.logger.Infof("Entering expression")
}

func (f FluxProgramParser) EnterBlock(c *parsing.BlockContext) {
	f.logger.Infof("Entering block")
}

func (f FluxProgramParser) EnterLoop_statement(c *parsing.Loop_statementContext) {
	f.logger.Infof("Entering loop statement")
}

func (f FluxProgramParser) EnterIf_statement(c *parsing.If_statementContext) {
	f.logger.Infof("Entering if statement")
}

func (f FluxProgramParser) EnterReturn_statement(c *parsing.Return_statementContext) {
	f.logger.Infof("Entering return statement")
}

func (f FluxProgramParser) EnterData_type(c *parsing.Data_typeContext) {
	f.logger.Infof("Entering data type")
}

func (f FluxProgramParser) EnterFunc_declaration(c *parsing.Func_declarationContext) {
	f.logger.Infof("Entering function declaration: ")
}

func (f FluxProgramParser) EnterVar_type(c *parsing.Var_typeContext) {
	f.logger.Infof("Entering var type")
}

func (f FluxProgramParser) EnterVar_name(c *parsing.Var_nameContext) {
	f.logger.Infof("Entering var name")
}

func (f FluxProgramParser) EnterVar_value(c *parsing.Var_valueContext) {
	f.logger.Infof("Entering var value")
}

func (f FluxProgramParser) EnterOp_one_expression(c *parsing.Op_one_expressionContext) {
	f.logger.Infof("Entering op one expression")
}

func (f FluxProgramParser) EnterOp_one_declaration(c *parsing.Op_one_declarationContext) {
	f.logger.Infof("Entering op one declaration")
}

func (f FluxProgramParser) EnterString_var_declaration(c *parsing.String_var_declarationContext) {
	f.logger.Infof("Entering string var declaration")
}

func (f FluxProgramParser) EnterNumber_var_declaration(c *parsing.Number_var_declarationContext) {
	f.logger.Infof("Entering number var declaration")
	numDecStmt := declarations.NewVarDeclaration(c.GetStart().GetLine(), c.GetStart().GetStart(), 0, "", common.FluxTypeNumber, "", nil)
	f.stackStatements.Push(numDecStmt)
}

func (f FluxProgramParser) EnterBoolean_var_declaration(c *parsing.Boolean_var_declarationContext) {
	f.logger.Infof("Entering boolean var declaration")
}

func (f FluxProgramParser) EnterSingle_var_declaration(c *parsing.Single_var_declarationContext) {
	f.logger.Infof("Entering single var declaration")
}

func (f FluxProgramParser) EnterArray_var_declaration(c *parsing.Array_var_declarationContext) {
	f.logger.Infof("Entering array var declaration")
}

func (f FluxProgramParser) EnterVar_assignment(c *parsing.Var_assignmentContext) {
	f.logger.Infof("Entering var assignment")
}

func (f FluxProgramParser) EnterOp_level1(c *parsing.Op_level1Context) {
	f.logger.Infof("Entering op level 1")
}

func (f FluxProgramParser) EnterOp_level2(c *parsing.Op_level2Context) {
	f.logger.Infof("Entering op level 2")
}

func (f FluxProgramParser) EnterOp_level3(c *parsing.Op_level3Context) {
	f.logger.Infof("Entering op level 3")
}

func (f FluxProgramParser) EnterOp_level4(c *parsing.Op_level4Context) {
	f.logger.Infof("Entering op level 4")
}

func (f FluxProgramParser) EnterOp_level5(c *parsing.Op_level5Context) {
	f.logger.Infof("Entering op level 5")
}

func (f FluxProgramParser) EnterNumberic_expression(c *parsing.Numberic_expressionContext) {
	f.logger.Infof("Entering numeric expression")

	numExpr := &fluxExpr.NumericExpression{
		BaseStatement: &fluxCodeObjects.BaseStatement{
			Line:     c.GetStart().GetLine(),
			StartPos: c.GetStart().GetStart(),
		},
		LeftExpr:  nil,
		Op:        nil,
		RightExpr: nil,
		Value:     0,
	}
	f.stackStatements.Push(numExpr)
}

func (f FluxProgramParser) EnterLogical_expression(c *parsing.Logical_expressionContext) {
	f.logger.Infof("Entering logical expression")
}

func (f FluxProgramParser) EnterComparative_expression(c *parsing.Comparative_expressionContext) {
	f.logger.Infof("Entering comparative expression")
}

func (f FluxProgramParser) EnterGet_variable(c *parsing.Get_variableContext) {
	f.logger.Infof("Entering get variable")
	getVar := &fluxExpr.GetVar{
		BaseStatement: &fluxCodeObjects.BaseStatement{
			Line:     c.GetStart().GetLine(),
			StartPos: c.GetStart().GetStart(),
			EndPos:   0,
		},
		VarName: "",
	}
	f.stackStatements.Push(getVar)

}

func (f FluxProgramParser) EnterMath_expression(c *parsing.Math_expressionContext) {
	f.logger.Infof("Entering math expression")
	mathExpr := &fluxExpr.MathExpression{
		BaseStatement: &fluxCodeObjects.BaseStatement{
			Line:     c.GetStart().GetLine(),
			StartPos: c.GetStart().GetStart(),
			EndPos:   0,
		},
		NumericExpr: nil,
	}
	f.stackStatements.Push(mathExpr)
}

func (f FluxProgramParser) EnterGet_array_element(c *parsing.Get_array_elementContext) {
	f.logger.Infof("Entering get array element")
}

func (f FluxProgramParser) EnterGet_child(c *parsing.Get_childContext) {
	f.logger.Infof("Entering get child")
}

func (f FluxProgramParser) EnterFunction_call(c *parsing.Function_callContext) {
	f.logger.Infof("Entering function call")
}

func (f FluxProgramParser) ExitProgram(c *parsing.ProgramContext) {
	f.logger.Infof("Exiting program")
	if f.stackStatements.Peek() != nil {
		prog := (*f.stackStatements.Pop()).(*fluxCodeObjects.Program)
		f.program = prog
	}
	for !f.resultStack.IsEmpty() {
		statement := (*f.resultStack.Dequeue()).(fluxCodeObjects.CodeObject)
		f.program.Statements = append(f.program.Statements, statement)
	}
}

func (f FluxProgramParser) ExitStatement(c *parsing.StatementContext) {
	f.logger.Infof("Exiting statement")
}

func (f FluxProgramParser) ExitExpression(c *parsing.ExpressionContext) {
	f.logger.Infof("Exiting expression")
}

func (f FluxProgramParser) ExitBlock(c *parsing.BlockContext) {
	f.logger.Infof("Exiting block")
}

func (f FluxProgramParser) ExitLoop_statement(c *parsing.Loop_statementContext) {
	f.logger.Infof("Exiting loop statement")
}

func (f FluxProgramParser) ExitIf_statement(c *parsing.If_statementContext) {
	f.logger.Infof("Exiting if statement")
}

func (f FluxProgramParser) ExitReturn_statement(c *parsing.Return_statementContext) {
	f.logger.Infof("Exiting return statement")
}

func (f FluxProgramParser) ExitData_type(c *parsing.Data_typeContext) {
	f.logger.Infof("Exiting data type")
}

func (f FluxProgramParser) ExitFunc_declaration(c *parsing.Func_declarationContext) {
	f.logger.Infof("Exiting function declaration")
}

func (f FluxProgramParser) ExitVar_type(c *parsing.Var_typeContext) {
	f.logger.Infof("Exiting var type")
}

func (f FluxProgramParser) ExitVar_name(c *parsing.Var_nameContext) {
	f.logger.Infof("Exiting var name")
}

func (f FluxProgramParser) ExitVar_value(c *parsing.Var_valueContext) {
	f.logger.Infof("Exiting var value")
}

func (f FluxProgramParser) ExitOp_one_expression(c *parsing.Op_one_expressionContext) {
	f.logger.Infof("Exiting op one expression")
}

func (f FluxProgramParser) ExitOp_one_declaration(c *parsing.Op_one_declarationContext) {
	f.logger.Infof("Exiting op one declaration")
}

func (f FluxProgramParser) ExitString_var_declaration(c *parsing.String_var_declarationContext) {
	f.logger.Infof("Exiting string var declaration")
}

func (f FluxProgramParser) ExitNumber_var_declaration(c *parsing.Number_var_declarationContext) {
	f.logger.Infof("Exiting number var declaration")
	if f.stackStatements.Peek() != nil {
		if _, ok := (*f.stackStatements.Peek()).(*declarations.VarDeclaration); ok {
			numDecStmt := (*f.stackStatements.Pop()).(*declarations.VarDeclaration)
			if !f.resultStack.IsEmpty() {
				if _, ok := (*f.resultStack.Peek()).(*fluxExpr.NumericExpression); ok != false {
					numDecStmt.Expr = (*f.resultStack.Pop()).(*fluxExpr.MathExpression)
				}
			}

			if c.NUMBER() != nil {
				numDecStmt.RawValue = c.NUMBER().GetText()
			}
			numDecStmt.EndPos = c.GetStop().GetStop()
			numDecStmt.Name = c.Var_name().GetText()
			numDecStmt.Type = common.FluxTypeNumber
			f.resultStack.Push(numDecStmt)
		}
	}

}

func (f FluxProgramParser) ExitBoolean_var_declaration(c *parsing.Boolean_var_declarationContext) {
	f.logger.Infof("Exiting boolean var declaration")
}

func (f FluxProgramParser) ExitSingle_var_declaration(c *parsing.Single_var_declarationContext) {
	f.logger.Infof("Exiting single var declaration")
}

func (f FluxProgramParser) ExitArray_var_declaration(c *parsing.Array_var_declarationContext) {
	f.logger.Infof("Exiting array var declaration")
}

func (f FluxProgramParser) ExitVar_assignment(c *parsing.Var_assignmentContext) {
	f.logger.Infof("Exiting var assignment")
}

func (f FluxProgramParser) ExitOp_level1(c *parsing.Op_level1Context) {
	f.logger.Infof("Exiting op level 1")
}

func (f FluxProgramParser) ExitOp_level2(c *parsing.Op_level2Context) {
	f.logger.Infof("Exiting op level 2")
}

func (f FluxProgramParser) ExitOp_level3(c *parsing.Op_level3Context) {
	f.logger.Infof("Exiting op level 3")
}

func (f FluxProgramParser) ExitOp_level4(c *parsing.Op_level4Context) {
	f.logger.Infof("Exiting op level 4")
}

func (f FluxProgramParser) ExitOp_level5(c *parsing.Op_level5Context) {
	f.logger.Infof("Exiting op level 5")
}

func (f FluxProgramParser) ExitNumberic_expression(c *parsing.Numberic_expressionContext) {
	f.logger.Infof("Exiting numeric expression")
	if f.stackStatements.Peek() != nil {
		if _, ok := (*f.stackStatements.Peek()).(*fluxExpr.NumericExpression); ok != false {
			numExpr := (*f.stackStatements.Pop()).(*fluxExpr.NumericExpression)
			if c.NUMBER() != nil {
				value, err := strconv.ParseFloat(c.NUMBER().GetText(), 64)
				if err != nil {
					f.errorCollector.CollectError(&exception.BaseException{
						MessageFmt: "Error parsing number: %v",
						Args:       []interface{}{c.NUMBER().GetText()},
					})
				} else {
					numExpr.Value = value
				}

			} else if c.Get_variable() != nil {
				if !f.resultStack.IsEmpty() {
					if _, ok := (*f.resultStack.Peek()).(*fluxExpr.GetVar); ok != false {
						numExpr.GetVar = (*f.resultStack.Pop()).(*fluxExpr.GetVar)
					}
				}

			} else if len(c.AllNumberic_expression()) == 2 {
				if !f.resultStack.IsEmpty() {
					if _, ok := (*f.resultStack.Peek()).(*fluxExpr.NumericExpression); ok != false {
						numExpr.RightExpr = (*f.resultStack.Pop()).(*fluxExpr.NumericExpression)
					}

					if _, ok := (*f.resultStack.Peek()).(*fluxExpr.NumericExpression); ok != false {
						numExpr.LeftExpr = (*f.resultStack.Pop()).(*fluxExpr.NumericExpression)

					}

					if c.Op_level2() != nil {
						if c.Op_level2().GetText() == "+" {
							numExpr.Op = operators.NewAdd(c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetStop().GetColumn(), numExpr.LeftExpr, numExpr.RightExpr)
						}
					}
				}

			}
			f.resultStack.Push(numExpr)
		}
	}
}

func (f FluxProgramParser) ExitLogical_expression(c *parsing.Logical_expressionContext) {
	f.logger.Infof("Exiting logical expression")
}

func (f FluxProgramParser) ExitComparative_expression(c *parsing.Comparative_expressionContext) {
	f.logger.Infof("Exiting comparative expression")
}

func (f FluxProgramParser) ExitGet_variable(c *parsing.Get_variableContext) {
	f.logger.Infof("Exiting get variable")
	if f.stackStatements.Peek() != nil {
		if _, ok := (*f.stackStatements.Peek()).(*fluxExpr.GetVar); ok != false {
			getVar := (*f.stackStatements.Pop()).(*fluxExpr.GetVar)
			getVar.VarName = c.GetText()
			getVar.EndPos = c.GetStop().GetStop()
			f.resultStack.Push(getVar)
			// case 1: parent is a MathExpression
			if _, ok := (*f.stackStatements.Peek()).(*fluxExpr.MathExpression); ok != false {
				mathExpr := (*f.stackStatements.Peek()).(*fluxExpr.MathExpression)
				mathExpr.GetVar = getVar
			}
			// case 2: parent is a numeric expression
			if _, ok := (*f.stackStatements.Peek()).(*fluxExpr.NumericExpression); ok != false {
				numExpr := (*f.stackStatements.Peek()).(*fluxExpr.NumericExpression)
				numExpr.GetVar = getVar
			}
		}
	}
}

func (f FluxProgramParser) ExitMath_expression(c *parsing.Math_expressionContext) {
	f.logger.Infof("Exiting math expression")
	if f.stackStatements.Peek() != nil {
		if _, ok := (*f.stackStatements.Peek()).(*fluxExpr.MathExpression); ok != false {
			mathExpr := (*f.stackStatements.Pop()).(*fluxExpr.MathExpression)
			// pop the result stack
			if !f.resultStack.IsEmpty() {
				if _, ok := (*f.resultStack.Peek()).(*fluxExpr.NumericExpression); ok != false {
					mathExpr.NumericExpr = (*f.resultStack.Pop()).(*fluxExpr.NumericExpression)
				} else if _, ok := (*f.resultStack.Peek()).(*fluxExpr.GetVar); ok != false {
					mathExpr.GetVar = (*f.resultStack.Pop()).(*fluxExpr.GetVar)
				}
			}
			f.resultStack.Push(mathExpr)
		}
	}
}

func (f FluxProgramParser) ExitGet_array_element(c *parsing.Get_array_elementContext) {
	f.logger.Infof("Exiting get array element")
}

func (f FluxProgramParser) ExitGet_child(c *parsing.Get_childContext) {
	f.logger.Infof("Exiting get child")
}

func (f FluxProgramParser) ExitFunction_call(c *parsing.Function_callContext) {
	f.logger.Infof("Exiting function call")
}
