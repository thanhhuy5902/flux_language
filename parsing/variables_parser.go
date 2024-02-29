// Code generated from Variables.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Variables
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Variables struct {
	*antlr.BaseParser
}

var VariablesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func variablesParserInit() {
	staticData := &VariablesParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'text'", "", "'boolean'", "", "", "", "", "'num'", "", "",
		"'ipv4'", "", "", "", "", "'{'", "'}'", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "DIGIT", "NUMBER",
		"UNIT", "SI_POSTFIX", "NUMBER_TYPE", "OCTET", "IPV4", "IPV4_TYPE", "VAR_IDENTIFIER",
		"COMMON_IDENTIFIER", "ARITHMETIC_OPERATOR", "LOGICAL_OPERATOR", "L_BLOCK",
		"R_BLOCK", "AT", "WS", "OP_INCREMENT", "OP_DECREMENT", "OP_ADD_ASSIGN",
		"OP_SUB_ASSIGN", "OP_MUL_ASSIGN", "OP_DIV_ASSIGN", "OP_EQUAL", "NUMBE",
		"OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "NEWLINE",
		"COMMA", "OP_MULTIPLY", "OP_DIVIDE", "OP_POWER", "OP_PLUS", "OP_MINUS",
		"OP_AND", "OP_OR", "OP_XOR", "OP_NOT_EQUAL", "OP_LESS", "OP_LESS_EQUAL",
		"OP_GREATER", "OP_GREATER_EQUAL", "OP_NOT", "L_PAREN", "R_PAREN", "L_SQUARE",
		"R_SQUARE", "DOT",
	}
	staticData.RuleNames = []string{
		"var_type", "var_name", "var_value", "op_one_expression", "op_one_declaration",
		"op_two_expression", "string_var_declaration", "number_var_declaration",
		"boolean_var_declaration", "single_var_declaration", "array_var_declaration",
		"var_assignment", "op_level1", "op_level2", "op_level3", "op_level4",
		"op_level5", "numberic_expression", "logical_expression", "comparative_expression",
		"get_variable", "math_expression", "get_array_element", "get_child",
		"function_call",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 456, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 98, 8, 4, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 106, 8, 6, 10, 6, 12, 6, 109, 9, 6, 1,
		6, 1, 6, 5, 6, 113, 8, 6, 10, 6, 12, 6, 116, 9, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 124, 8, 7, 10, 7, 12, 7, 127, 9, 7, 1, 7, 1, 7, 5,
		7, 131, 8, 7, 10, 7, 12, 7, 134, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 5, 7, 142, 8, 7, 10, 7, 12, 7, 145, 9, 7, 1, 7, 1, 7, 5, 7, 149, 8,
		7, 10, 7, 12, 7, 152, 9, 7, 1, 7, 1, 7, 3, 7, 156, 8, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 5, 8, 162, 8, 8, 10, 8, 12, 8, 165, 9, 8, 1, 8, 1, 8, 5, 8, 169,
		8, 8, 10, 8, 12, 8, 172, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8,
		180, 8, 8, 10, 8, 12, 8, 183, 9, 8, 1, 8, 1, 8, 5, 8, 187, 8, 8, 10, 8,
		12, 8, 190, 9, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 198, 8, 8,
		10, 8, 12, 8, 201, 9, 8, 1, 8, 1, 8, 5, 8, 205, 8, 8, 10, 8, 12, 8, 208,
		9, 8, 1, 8, 1, 8, 3, 8, 212, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 5, 9, 221, 8, 9, 10, 9, 12, 9, 224, 9, 9, 1, 9, 1, 9, 3, 9, 228,
		8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 234, 8, 10, 10, 10, 12, 10, 237,
		9, 10, 1, 10, 1, 10, 1, 10, 5, 10, 242, 8, 10, 10, 10, 12, 10, 245, 9,
		10, 1, 10, 5, 10, 248, 8, 10, 10, 10, 12, 10, 251, 9, 10, 1, 10, 5, 10,
		254, 8, 10, 10, 10, 12, 10, 257, 9, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 5, 10, 265, 8, 10, 10, 10, 12, 10, 268, 9, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 273, 8, 10, 10, 10, 12, 10, 276, 9, 10, 1, 10, 5, 10, 279,
		8, 10, 10, 10, 12, 10, 282, 9, 10, 1, 10, 1, 10, 1, 10, 3, 10, 287, 8,
		10, 1, 11, 1, 11, 1, 11, 5, 11, 292, 8, 11, 10, 11, 12, 11, 295, 9, 11,
		1, 11, 1, 11, 5, 11, 299, 8, 11, 10, 11, 12, 11, 302, 9, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 5, 11, 309, 8, 11, 10, 11, 12, 11, 312, 9, 11,
		1, 11, 1, 11, 5, 11, 316, 8, 11, 10, 11, 12, 11, 319, 9, 11, 1, 11, 1,
		11, 3, 11, 323, 8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 3, 17, 345, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 5, 17, 355, 8, 17, 10, 17, 12, 17, 358, 9, 17, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18,
		370, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5,
		18, 380, 8, 18, 10, 18, 12, 18, 383, 9, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3,
		19, 399, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 405, 8, 19, 10, 19,
		12, 19, 408, 9, 19, 1, 20, 1, 20, 1, 20, 3, 20, 413, 8, 20, 1, 21, 1, 21,
		1, 21, 3, 21, 418, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 440, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 5, 24, 447, 8, 24, 10, 24, 12, 24, 450, 9, 24, 3, 24, 452, 8, 24, 1,
		24, 1, 24, 1, 24, 0, 3, 34, 36, 38, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 8, 4,
		0, 2, 2, 4, 4, 9, 9, 12, 12, 4, 0, 1, 1, 3, 3, 6, 6, 11, 11, 1, 0, 21,
		26, 1, 0, 29, 34, 2, 0, 33, 33, 37, 39, 1, 0, 40, 41, 1, 0, 42, 44, 2,
		0, 27, 27, 45, 49, 496, 0, 50, 1, 0, 0, 0, 2, 52, 1, 0, 0, 0, 4, 54, 1,
		0, 0, 0, 6, 56, 1, 0, 0, 0, 8, 97, 1, 0, 0, 0, 10, 99, 1, 0, 0, 0, 12,
		101, 1, 0, 0, 0, 14, 155, 1, 0, 0, 0, 16, 211, 1, 0, 0, 0, 18, 227, 1,
		0, 0, 0, 20, 286, 1, 0, 0, 0, 22, 322, 1, 0, 0, 0, 24, 324, 1, 0, 0, 0,
		26, 326, 1, 0, 0, 0, 28, 328, 1, 0, 0, 0, 30, 330, 1, 0, 0, 0, 32, 332,
		1, 0, 0, 0, 34, 344, 1, 0, 0, 0, 36, 369, 1, 0, 0, 0, 38, 398, 1, 0, 0,
		0, 40, 412, 1, 0, 0, 0, 42, 417, 1, 0, 0, 0, 44, 419, 1, 0, 0, 0, 46, 439,
		1, 0, 0, 0, 48, 441, 1, 0, 0, 0, 50, 51, 7, 0, 0, 0, 51, 1, 1, 0, 0, 0,
		52, 53, 5, 13, 0, 0, 53, 3, 1, 0, 0, 0, 54, 55, 7, 1, 0, 0, 55, 5, 1, 0,
		0, 0, 56, 57, 7, 2, 0, 0, 57, 7, 1, 0, 0, 0, 58, 59, 3, 2, 1, 0, 59, 60,
		5, 21, 0, 0, 60, 98, 1, 0, 0, 0, 61, 62, 3, 2, 1, 0, 62, 63, 5, 22, 0,
		0, 63, 98, 1, 0, 0, 0, 64, 65, 5, 9, 0, 0, 65, 66, 3, 2, 1, 0, 66, 67,
		5, 21, 0, 0, 67, 98, 1, 0, 0, 0, 68, 69, 5, 9, 0, 0, 69, 70, 3, 2, 1, 0,
		70, 71, 5, 22, 0, 0, 71, 98, 1, 0, 0, 0, 72, 73, 5, 9, 0, 0, 73, 74, 3,
		2, 1, 0, 74, 75, 5, 27, 0, 0, 75, 76, 5, 6, 0, 0, 76, 98, 1, 0, 0, 0, 77,
		78, 5, 9, 0, 0, 78, 79, 3, 2, 1, 0, 79, 80, 5, 23, 0, 0, 80, 81, 5, 6,
		0, 0, 81, 98, 1, 0, 0, 0, 82, 83, 5, 9, 0, 0, 83, 84, 3, 2, 1, 0, 84, 85,
		5, 24, 0, 0, 85, 86, 5, 6, 0, 0, 86, 98, 1, 0, 0, 0, 87, 88, 5, 9, 0, 0,
		88, 89, 3, 2, 1, 0, 89, 90, 5, 25, 0, 0, 90, 91, 5, 6, 0, 0, 91, 98, 1,
		0, 0, 0, 92, 93, 5, 9, 0, 0, 93, 94, 3, 2, 1, 0, 94, 95, 5, 26, 0, 0, 95,
		96, 5, 28, 0, 0, 96, 98, 1, 0, 0, 0, 97, 58, 1, 0, 0, 0, 97, 61, 1, 0,
		0, 0, 97, 64, 1, 0, 0, 0, 97, 68, 1, 0, 0, 0, 97, 72, 1, 0, 0, 0, 97, 77,
		1, 0, 0, 0, 97, 82, 1, 0, 0, 0, 97, 87, 1, 0, 0, 0, 97, 92, 1, 0, 0, 0,
		98, 9, 1, 0, 0, 0, 99, 100, 7, 3, 0, 0, 100, 11, 1, 0, 0, 0, 101, 102,
		5, 2, 0, 0, 102, 103, 3, 2, 1, 0, 103, 107, 5, 17, 0, 0, 104, 106, 5, 35,
		0, 0, 105, 104, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0,
		107, 108, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110,
		114, 5, 1, 0, 0, 111, 113, 5, 35, 0, 0, 112, 111, 1, 0, 0, 0, 113, 116,
		1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0,
		0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 18, 0, 0, 118, 13, 1, 0, 0, 0,
		119, 120, 5, 9, 0, 0, 120, 121, 3, 2, 1, 0, 121, 125, 5, 17, 0, 0, 122,
		124, 5, 35, 0, 0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 128, 1, 0, 0, 0, 127, 125, 1, 0,
		0, 0, 128, 132, 5, 6, 0, 0, 129, 131, 5, 35, 0, 0, 130, 129, 1, 0, 0, 0,
		131, 134, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133,
		135, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 135, 136, 5, 18, 0, 0, 136, 156,
		1, 0, 0, 0, 137, 138, 5, 9, 0, 0, 138, 139, 3, 2, 1, 0, 139, 143, 5, 17,
		0, 0, 140, 142, 5, 35, 0, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0,
		143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145,
		143, 1, 0, 0, 0, 146, 150, 3, 42, 21, 0, 147, 149, 5, 35, 0, 0, 148, 147,
		1, 0, 0, 0, 149, 152, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0,
		0, 0, 151, 153, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 153, 154, 5, 18, 0, 0,
		154, 156, 1, 0, 0, 0, 155, 119, 1, 0, 0, 0, 155, 137, 1, 0, 0, 0, 156,
		15, 1, 0, 0, 0, 157, 158, 5, 4, 0, 0, 158, 159, 3, 2, 1, 0, 159, 163, 5,
		17, 0, 0, 160, 162, 5, 35, 0, 0, 161, 160, 1, 0, 0, 0, 162, 165, 1, 0,
		0, 0, 163, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0,
		165, 163, 1, 0, 0, 0, 166, 170, 5, 3, 0, 0, 167, 169, 5, 35, 0, 0, 168,
		167, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171,
		1, 0, 0, 0, 171, 173, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 174, 5, 18,
		0, 0, 174, 212, 1, 0, 0, 0, 175, 176, 5, 4, 0, 0, 176, 177, 3, 2, 1, 0,
		177, 181, 5, 17, 0, 0, 178, 180, 5, 35, 0, 0, 179, 178, 1, 0, 0, 0, 180,
		183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 184,
		1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 188, 3, 36, 18, 0, 185, 187, 5,
		35, 0, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0, 0,
		0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0, 191,
		192, 5, 18, 0, 0, 192, 212, 1, 0, 0, 0, 193, 194, 5, 4, 0, 0, 194, 195,
		3, 2, 1, 0, 195, 199, 5, 17, 0, 0, 196, 198, 5, 35, 0, 0, 197, 196, 1,
		0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0,
		0, 200, 202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 206, 3, 38, 19, 0,
		203, 205, 5, 35, 0, 0, 204, 203, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206,
		204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 209, 1, 0, 0, 0, 208, 206,
		1, 0, 0, 0, 209, 210, 5, 18, 0, 0, 210, 212, 1, 0, 0, 0, 211, 157, 1, 0,
		0, 0, 211, 175, 1, 0, 0, 0, 211, 193, 1, 0, 0, 0, 212, 17, 1, 0, 0, 0,
		213, 228, 3, 12, 6, 0, 214, 228, 3, 14, 7, 0, 215, 228, 3, 16, 8, 0, 216,
		217, 3, 0, 0, 0, 217, 218, 3, 2, 1, 0, 218, 222, 5, 17, 0, 0, 219, 221,
		5, 35, 0, 0, 220, 219, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0,
		0, 0, 222, 223, 1, 0, 0, 0, 223, 225, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0,
		225, 226, 5, 18, 0, 0, 226, 228, 1, 0, 0, 0, 227, 213, 1, 0, 0, 0, 227,
		214, 1, 0, 0, 0, 227, 215, 1, 0, 0, 0, 227, 216, 1, 0, 0, 0, 228, 19, 1,
		0, 0, 0, 229, 230, 3, 0, 0, 0, 230, 231, 3, 2, 1, 0, 231, 235, 5, 17, 0,
		0, 232, 234, 5, 35, 0, 0, 233, 232, 1, 0, 0, 0, 234, 237, 1, 0, 0, 0, 235,
		233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 238, 1, 0, 0, 0, 237, 235,
		1, 0, 0, 0, 238, 249, 3, 4, 2, 0, 239, 243, 5, 36, 0, 0, 240, 242, 5, 35,
		0, 0, 241, 240, 1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0,
		243, 244, 1, 0, 0, 0, 244, 246, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246,
		248, 3, 4, 2, 0, 247, 239, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247,
		1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 255, 1, 0, 0, 0, 251, 249, 1, 0,
		0, 0, 252, 254, 5, 35, 0, 0, 253, 252, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0,
		255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 258, 1, 0, 0, 0, 257,
		255, 1, 0, 0, 0, 258, 259, 5, 18, 0, 0, 259, 287, 1, 0, 0, 0, 260, 261,
		3, 0, 0, 0, 261, 262, 3, 2, 1, 0, 262, 266, 5, 17, 0, 0, 263, 265, 5, 35,
		0, 0, 264, 263, 1, 0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0,
		266, 267, 1, 0, 0, 0, 267, 269, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269,
		280, 3, 42, 21, 0, 270, 274, 5, 36, 0, 0, 271, 273, 5, 35, 0, 0, 272, 271,
		1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 275, 1, 0,
		0, 0, 275, 277, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 277, 279, 3, 42, 21,
		0, 278, 270, 1, 0, 0, 0, 279, 282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280,
		281, 1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 284,
		5, 35, 0, 0, 284, 285, 5, 18, 0, 0, 285, 287, 1, 0, 0, 0, 286, 229, 1,
		0, 0, 0, 286, 260, 1, 0, 0, 0, 287, 21, 1, 0, 0, 0, 288, 289, 3, 2, 1,
		0, 289, 293, 5, 17, 0, 0, 290, 292, 5, 35, 0, 0, 291, 290, 1, 0, 0, 0,
		292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294,
		296, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 300, 3, 4, 2, 0, 297, 299,
		5, 35, 0, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0,
		0, 0, 300, 301, 1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0,
		303, 304, 5, 18, 0, 0, 304, 323, 1, 0, 0, 0, 305, 306, 3, 2, 1, 0, 306,
		310, 5, 17, 0, 0, 307, 309, 5, 35, 0, 0, 308, 307, 1, 0, 0, 0, 309, 312,
		1, 0, 0, 0, 310, 308, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 313, 1, 0,
		0, 0, 312, 310, 1, 0, 0, 0, 313, 317, 3, 42, 21, 0, 314, 316, 5, 35, 0,
		0, 315, 314, 1, 0, 0, 0, 316, 319, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 317,
		318, 1, 0, 0, 0, 318, 320, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 320, 321,
		5, 18, 0, 0, 321, 323, 1, 0, 0, 0, 322, 288, 1, 0, 0, 0, 322, 305, 1, 0,
		0, 0, 323, 23, 1, 0, 0, 0, 324, 325, 7, 4, 0, 0, 325, 25, 1, 0, 0, 0, 326,
		327, 7, 5, 0, 0, 327, 27, 1, 0, 0, 0, 328, 329, 7, 6, 0, 0, 329, 29, 1,
		0, 0, 0, 330, 331, 7, 7, 0, 0, 331, 31, 1, 0, 0, 0, 332, 333, 5, 50, 0,
		0, 333, 33, 1, 0, 0, 0, 334, 335, 6, 17, -1, 0, 335, 336, 5, 51, 0, 0,
		336, 337, 3, 34, 17, 0, 337, 338, 5, 52, 0, 0, 338, 345, 1, 0, 0, 0, 339,
		345, 5, 6, 0, 0, 340, 345, 3, 40, 20, 0, 341, 345, 3, 48, 24, 0, 342, 343,
		5, 41, 0, 0, 343, 345, 3, 34, 17, 1, 344, 334, 1, 0, 0, 0, 344, 339, 1,
		0, 0, 0, 344, 340, 1, 0, 0, 0, 344, 341, 1, 0, 0, 0, 344, 342, 1, 0, 0,
		0, 345, 356, 1, 0, 0, 0, 346, 347, 10, 6, 0, 0, 347, 348, 3, 24, 12, 0,
		348, 349, 3, 34, 17, 7, 349, 355, 1, 0, 0, 0, 350, 351, 10, 5, 0, 0, 351,
		352, 3, 26, 13, 0, 352, 353, 3, 34, 17, 6, 353, 355, 1, 0, 0, 0, 354, 346,
		1, 0, 0, 0, 354, 350, 1, 0, 0, 0, 355, 358, 1, 0, 0, 0, 356, 354, 1, 0,
		0, 0, 356, 357, 1, 0, 0, 0, 357, 35, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0,
		359, 360, 6, 18, -1, 0, 360, 361, 5, 51, 0, 0, 361, 362, 3, 36, 18, 0,
		362, 363, 5, 52, 0, 0, 363, 370, 1, 0, 0, 0, 364, 365, 5, 50, 0, 0, 365,
		370, 3, 36, 18, 4, 366, 370, 5, 3, 0, 0, 367, 370, 3, 40, 20, 0, 368, 370,
		3, 48, 24, 0, 369, 359, 1, 0, 0, 0, 369, 364, 1, 0, 0, 0, 369, 366, 1,
		0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 368, 1, 0, 0, 0, 370, 381, 1, 0, 0,
		0, 371, 372, 10, 6, 0, 0, 372, 373, 3, 28, 14, 0, 373, 374, 3, 36, 18,
		7, 374, 380, 1, 0, 0, 0, 375, 376, 10, 5, 0, 0, 376, 377, 3, 30, 15, 0,
		377, 378, 3, 36, 18, 6, 378, 380, 1, 0, 0, 0, 379, 371, 1, 0, 0, 0, 379,
		375, 1, 0, 0, 0, 380, 383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382,
		1, 0, 0, 0, 382, 37, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384, 385, 6, 19,
		-1, 0, 385, 386, 3, 34, 17, 0, 386, 387, 3, 30, 15, 0, 387, 388, 3, 34,
		17, 0, 388, 399, 1, 0, 0, 0, 389, 390, 3, 36, 18, 0, 390, 391, 3, 28, 14,
		0, 391, 392, 3, 36, 18, 0, 392, 399, 1, 0, 0, 0, 393, 394, 3, 32, 16, 0,
		394, 395, 5, 51, 0, 0, 395, 396, 3, 38, 19, 0, 396, 397, 5, 52, 0, 0, 397,
		399, 1, 0, 0, 0, 398, 384, 1, 0, 0, 0, 398, 389, 1, 0, 0, 0, 398, 393,
		1, 0, 0, 0, 399, 406, 1, 0, 0, 0, 400, 401, 10, 2, 0, 0, 401, 402, 3, 28,
		14, 0, 402, 403, 3, 36, 18, 0, 403, 405, 1, 0, 0, 0, 404, 400, 1, 0, 0,
		0, 405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407,
		39, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 413, 5, 13, 0, 0, 410, 413,
		3, 44, 22, 0, 411, 413, 3, 46, 23, 0, 412, 409, 1, 0, 0, 0, 412, 410, 1,
		0, 0, 0, 412, 411, 1, 0, 0, 0, 413, 41, 1, 0, 0, 0, 414, 418, 3, 40, 20,
		0, 415, 418, 3, 34, 17, 0, 416, 418, 3, 36, 18, 0, 417, 414, 1, 0, 0, 0,
		417, 415, 1, 0, 0, 0, 417, 416, 1, 0, 0, 0, 418, 43, 1, 0, 0, 0, 419, 420,
		5, 13, 0, 0, 420, 421, 5, 53, 0, 0, 421, 422, 3, 34, 17, 0, 422, 423, 5,
		54, 0, 0, 423, 45, 1, 0, 0, 0, 424, 425, 5, 13, 0, 0, 425, 426, 5, 55,
		0, 0, 426, 440, 5, 13, 0, 0, 427, 428, 5, 13, 0, 0, 428, 429, 5, 55, 0,
		0, 429, 440, 3, 44, 22, 0, 430, 431, 5, 13, 0, 0, 431, 432, 5, 55, 0, 0,
		432, 440, 3, 46, 23, 0, 433, 434, 3, 44, 22, 0, 434, 435, 5, 55, 0, 0,
		435, 436, 3, 46, 23, 0, 436, 440, 1, 0, 0, 0, 437, 440, 3, 44, 22, 0, 438,
		440, 5, 13, 0, 0, 439, 424, 1, 0, 0, 0, 439, 427, 1, 0, 0, 0, 439, 430,
		1, 0, 0, 0, 439, 433, 1, 0, 0, 0, 439, 437, 1, 0, 0, 0, 439, 438, 1, 0,
		0, 0, 440, 47, 1, 0, 0, 0, 441, 442, 5, 13, 0, 0, 442, 451, 5, 51, 0, 0,
		443, 448, 3, 42, 21, 0, 444, 445, 5, 36, 0, 0, 445, 447, 3, 42, 21, 0,
		446, 444, 1, 0, 0, 0, 447, 450, 1, 0, 0, 0, 448, 446, 1, 0, 0, 0, 448,
		449, 1, 0, 0, 0, 449, 452, 1, 0, 0, 0, 450, 448, 1, 0, 0, 0, 451, 443,
		1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 454, 5, 52,
		0, 0, 454, 49, 1, 0, 0, 0, 43, 97, 107, 114, 125, 132, 143, 150, 155, 163,
		170, 181, 188, 199, 206, 211, 222, 227, 235, 243, 249, 255, 266, 274, 280,
		286, 293, 300, 310, 317, 322, 344, 354, 356, 369, 379, 381, 398, 406, 412,
		417, 439, 448, 451,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VariablesInit initializes any static state used to implement Variables. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVariables(). You can call this function if you wish to initialize the static state ahead
// of time.
func VariablesInit() {
	staticData := &VariablesParserStaticData
	staticData.once.Do(variablesParserInit)
}

// NewVariables produces a new parser instance for the optional input antlr.TokenStream.
func NewVariables(input antlr.TokenStream) *Variables {
	VariablesInit()
	this := new(Variables)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VariablesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Variables.g4"

	return this
}

// Variables tokens.
const (
	VariablesEOF                 = antlr.TokenEOF
	VariablesTEXT                = 1
	VariablesTEXT_TYPE           = 2
	VariablesBOOLEAN             = 3
	VariablesBOOLEAN_TYPE        = 4
	VariablesDIGIT               = 5
	VariablesNUMBER              = 6
	VariablesUNIT                = 7
	VariablesSI_POSTFIX          = 8
	VariablesNUMBER_TYPE         = 9
	VariablesOCTET               = 10
	VariablesIPV4                = 11
	VariablesIPV4_TYPE           = 12
	VariablesVAR_IDENTIFIER      = 13
	VariablesCOMMON_IDENTIFIER   = 14
	VariablesARITHMETIC_OPERATOR = 15
	VariablesLOGICAL_OPERATOR    = 16
	VariablesL_BLOCK             = 17
	VariablesR_BLOCK             = 18
	VariablesAT                  = 19
	VariablesWS                  = 20
	VariablesOP_INCREMENT        = 21
	VariablesOP_DECREMENT        = 22
	VariablesOP_ADD_ASSIGN       = 23
	VariablesOP_SUB_ASSIGN       = 24
	VariablesOP_MUL_ASSIGN       = 25
	VariablesOP_DIV_ASSIGN       = 26
	VariablesOP_EQUAL            = 27
	VariablesNUMBE               = 28
	VariablesOP_ADD              = 29
	VariablesOP_SUB              = 30
	VariablesOP_MUL              = 31
	VariablesOP_DIV              = 32
	VariablesOP_MOD              = 33
	VariablesOP_POW              = 34
	VariablesNEWLINE             = 35
	VariablesCOMMA               = 36
	VariablesOP_MULTIPLY         = 37
	VariablesOP_DIVIDE           = 38
	VariablesOP_POWER            = 39
	VariablesOP_PLUS             = 40
	VariablesOP_MINUS            = 41
	VariablesOP_AND              = 42
	VariablesOP_OR               = 43
	VariablesOP_XOR              = 44
	VariablesOP_NOT_EQUAL        = 45
	VariablesOP_LESS             = 46
	VariablesOP_LESS_EQUAL       = 47
	VariablesOP_GREATER          = 48
	VariablesOP_GREATER_EQUAL    = 49
	VariablesOP_NOT              = 50
	VariablesL_PAREN             = 51
	VariablesR_PAREN             = 52
	VariablesL_SQUARE            = 53
	VariablesR_SQUARE            = 54
	VariablesDOT                 = 55
)

// Variables rules.
const (
	VariablesRULE_var_type                = 0
	VariablesRULE_var_name                = 1
	VariablesRULE_var_value               = 2
	VariablesRULE_op_one_expression       = 3
	VariablesRULE_op_one_declaration      = 4
	VariablesRULE_op_two_expression       = 5
	VariablesRULE_string_var_declaration  = 6
	VariablesRULE_number_var_declaration  = 7
	VariablesRULE_boolean_var_declaration = 8
	VariablesRULE_single_var_declaration  = 9
	VariablesRULE_array_var_declaration   = 10
	VariablesRULE_var_assignment          = 11
	VariablesRULE_op_level1               = 12
	VariablesRULE_op_level2               = 13
	VariablesRULE_op_level3               = 14
	VariablesRULE_op_level4               = 15
	VariablesRULE_op_level5               = 16
	VariablesRULE_numberic_expression     = 17
	VariablesRULE_logical_expression      = 18
	VariablesRULE_comparative_expression  = 19
	VariablesRULE_get_variable            = 20
	VariablesRULE_math_expression         = 21
	VariablesRULE_get_array_element       = 22
	VariablesRULE_get_child               = 23
	VariablesRULE_function_call           = 24
)

// IVar_typeContext is an interface to support dynamic dispatch.
type IVar_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	IPV4_TYPE() antlr.TerminalNode

	// IsVar_typeContext differentiates from other interfaces.
	IsVar_typeContext()
}

type Var_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_typeContext() *Var_typeContext {
	var p = new(Var_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesTEXT_TYPE, 0)
}

func (s *Var_typeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER_TYPE, 0)
}

func (s *Var_typeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN_TYPE, 0)
}

func (s *Var_typeContext) IPV4_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesIPV4_TYPE, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (p *Variables) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VariablesRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4628) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_nameContext is an interface to support dynamic dispatch.
type IVar_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode

	// IsVar_nameContext differentiates from other interfaces.
	IsVar_nameContext()
}

type Var_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_nameContext() *Var_nameContext {
	var p = new(Var_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_name
	return p
}

func InitEmptyVar_nameContext(p *Var_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_name
}

func (*Var_nameContext) IsVar_nameContext() {}

func NewVar_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_nameContext {
	var p = new(Var_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_var_name

	return p
}

func (s *Var_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_nameContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VariablesVAR_IDENTIFIER, 0)
}

func (s *Var_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterVar_name(s)
	}
}

func (s *Var_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitVar_name(s)
	}
}

func (p *Variables) Var_name() (localctx IVar_nameContext) {
	localctx = NewVar_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VariablesRULE_var_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(VariablesVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_valueContext is an interface to support dynamic dispatch.
type IVar_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IPV4() antlr.TerminalNode
	TEXT() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode

	// IsVar_valueContext differentiates from other interfaces.
	IsVar_valueContext()
}

type Var_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_valueContext() *Var_valueContext {
	var p = new(Var_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_value
	return p
}

func InitEmptyVar_valueContext(p *Var_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_value
}

func (*Var_valueContext) IsVar_valueContext() {}

func NewVar_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_valueContext {
	var p = new(Var_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_var_value

	return p
}

func (s *Var_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_valueContext) IPV4() antlr.TerminalNode {
	return s.GetToken(VariablesIPV4, 0)
}

func (s *Var_valueContext) TEXT() antlr.TerminalNode {
	return s.GetToken(VariablesTEXT, 0)
}

func (s *Var_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER, 0)
}

func (s *Var_valueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN, 0)
}

func (s *Var_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterVar_value(s)
	}
}

func (s *Var_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitVar_value(s)
	}
}

func (p *Variables) Var_value() (localctx IVar_valueContext) {
	localctx = NewVar_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, VariablesRULE_var_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2122) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_one_expressionContext is an interface to support dynamic dispatch.
type IOp_one_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_INCREMENT() antlr.TerminalNode
	OP_DECREMENT() antlr.TerminalNode
	OP_ADD_ASSIGN() antlr.TerminalNode
	OP_SUB_ASSIGN() antlr.TerminalNode
	OP_MUL_ASSIGN() antlr.TerminalNode
	OP_DIV_ASSIGN() antlr.TerminalNode

	// IsOp_one_expressionContext differentiates from other interfaces.
	IsOp_one_expressionContext()
}

type Op_one_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_one_expressionContext() *Op_one_expressionContext {
	var p = new(Op_one_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_one_expression
	return p
}

func InitEmptyOp_one_expressionContext(p *Op_one_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_one_expression
}

func (*Op_one_expressionContext) IsOp_one_expressionContext() {}

func NewOp_one_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_one_expressionContext {
	var p = new(Op_one_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_op_one_expression

	return p
}

func (s *Op_one_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_one_expressionContext) OP_INCREMENT() antlr.TerminalNode {
	return s.GetToken(VariablesOP_INCREMENT, 0)
}

func (s *Op_one_expressionContext) OP_DECREMENT() antlr.TerminalNode {
	return s.GetToken(VariablesOP_DECREMENT, 0)
}

func (s *Op_one_expressionContext) OP_ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VariablesOP_ADD_ASSIGN, 0)
}

func (s *Op_one_expressionContext) OP_SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VariablesOP_SUB_ASSIGN, 0)
}

func (s *Op_one_expressionContext) OP_MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VariablesOP_MUL_ASSIGN, 0)
}

func (s *Op_one_expressionContext) OP_DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VariablesOP_DIV_ASSIGN, 0)
}

func (s *Op_one_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_one_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_one_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterOp_one_expression(s)
	}
}

func (s *Op_one_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitOp_one_expression(s)
	}
}

func (p *Variables) Op_one_expression() (localctx IOp_one_expressionContext) {
	localctx = NewOp_one_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, VariablesRULE_op_one_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&132120576) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_one_declarationContext is an interface to support dynamic dispatch.
type IOp_one_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_name() IVar_nameContext
	OP_INCREMENT() antlr.TerminalNode
	OP_DECREMENT() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode
	OP_EQUAL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	OP_ADD_ASSIGN() antlr.TerminalNode
	OP_SUB_ASSIGN() antlr.TerminalNode
	OP_MUL_ASSIGN() antlr.TerminalNode
	OP_DIV_ASSIGN() antlr.TerminalNode
	NUMBE() antlr.TerminalNode

	// IsOp_one_declarationContext differentiates from other interfaces.
	IsOp_one_declarationContext()
}

type Op_one_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_one_declarationContext() *Op_one_declarationContext {
	var p = new(Op_one_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_one_declaration
	return p
}

func InitEmptyOp_one_declarationContext(p *Op_one_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_one_declaration
}

func (*Op_one_declarationContext) IsOp_one_declarationContext() {}

func NewOp_one_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_one_declarationContext {
	var p = new(Op_one_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_op_one_declaration

	return p
}

func (s *Op_one_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_one_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Op_one_declarationContext) OP_INCREMENT() antlr.TerminalNode {
	return s.GetToken(VariablesOP_INCREMENT, 0)
}

func (s *Op_one_declarationContext) OP_DECREMENT() antlr.TerminalNode {
	return s.GetToken(VariablesOP_DECREMENT, 0)
}

func (s *Op_one_declarationContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER_TYPE, 0)
}

func (s *Op_one_declarationContext) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(VariablesOP_EQUAL, 0)
}

func (s *Op_one_declarationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER, 0)
}

func (s *Op_one_declarationContext) OP_ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VariablesOP_ADD_ASSIGN, 0)
}

func (s *Op_one_declarationContext) OP_SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VariablesOP_SUB_ASSIGN, 0)
}

func (s *Op_one_declarationContext) OP_MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VariablesOP_MUL_ASSIGN, 0)
}

func (s *Op_one_declarationContext) OP_DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(VariablesOP_DIV_ASSIGN, 0)
}

func (s *Op_one_declarationContext) NUMBE() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBE, 0)
}

func (s *Op_one_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_one_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_one_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterOp_one_declaration(s)
	}
}

func (s *Op_one_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitOp_one_declaration(s)
	}
}

func (p *Variables) Op_one_declaration() (localctx IOp_one_declarationContext) {
	localctx = NewOp_one_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, VariablesRULE_op_one_declaration)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Var_name()
		}
		{
			p.SetState(59)
			p.Match(VariablesOP_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Var_name()
		}
		{
			p.SetState(62)
			p.Match(VariablesOP_DECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Var_name()
		}
		{
			p.SetState(66)
			p.Match(VariablesOP_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)
			p.Var_name()
		}
		{
			p.SetState(70)
			p.Match(VariablesOP_DECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(72)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Var_name()
		}
		{
			p.SetState(74)
			p.Match(VariablesOP_EQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Match(VariablesNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(78)
			p.Var_name()
		}
		{
			p.SetState(79)
			p.Match(VariablesOP_ADD_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(VariablesNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(82)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Var_name()
		}
		{
			p.SetState(84)
			p.Match(VariablesOP_SUB_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(VariablesNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(87)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.Var_name()
		}
		{
			p.SetState(89)
			p.Match(VariablesOP_MUL_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(90)
			p.Match(VariablesNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(92)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(93)
			p.Var_name()
		}
		{
			p.SetState(94)
			p.Match(VariablesOP_DIV_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Match(VariablesNUMBE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_two_expressionContext is an interface to support dynamic dispatch.
type IOp_two_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_ADD() antlr.TerminalNode
	OP_SUB() antlr.TerminalNode
	OP_MUL() antlr.TerminalNode
	OP_DIV() antlr.TerminalNode
	OP_MOD() antlr.TerminalNode
	OP_POW() antlr.TerminalNode

	// IsOp_two_expressionContext differentiates from other interfaces.
	IsOp_two_expressionContext()
}

type Op_two_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_two_expressionContext() *Op_two_expressionContext {
	var p = new(Op_two_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_two_expression
	return p
}

func InitEmptyOp_two_expressionContext(p *Op_two_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_two_expression
}

func (*Op_two_expressionContext) IsOp_two_expressionContext() {}

func NewOp_two_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_two_expressionContext {
	var p = new(Op_two_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_op_two_expression

	return p
}

func (s *Op_two_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_two_expressionContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(VariablesOP_ADD, 0)
}

func (s *Op_two_expressionContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(VariablesOP_SUB, 0)
}

func (s *Op_two_expressionContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(VariablesOP_MUL, 0)
}

func (s *Op_two_expressionContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(VariablesOP_DIV, 0)
}

func (s *Op_two_expressionContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(VariablesOP_MOD, 0)
}

func (s *Op_two_expressionContext) OP_POW() antlr.TerminalNode {
	return s.GetToken(VariablesOP_POW, 0)
}

func (s *Op_two_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_two_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_two_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterOp_two_expression(s)
	}
}

func (s *Op_two_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitOp_two_expression(s)
	}
}

func (p *Variables) Op_two_expression() (localctx IOp_two_expressionContext) {
	localctx = NewOp_two_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, VariablesRULE_op_two_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33822867456) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_var_declarationContext is an interface to support dynamic dispatch.
type IString_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	TEXT() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsString_var_declarationContext differentiates from other interfaces.
	IsString_var_declarationContext()
}

type String_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_var_declarationContext() *String_var_declarationContext {
	var p = new(String_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_string_var_declaration
	return p
}

func InitEmptyString_var_declarationContext(p *String_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_string_var_declaration
}

func (*String_var_declarationContext) IsString_var_declarationContext() {}

func NewString_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_var_declarationContext {
	var p = new(String_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_string_var_declaration

	return p
}

func (s *String_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *String_var_declarationContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesTEXT_TYPE, 0)
}

func (s *String_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *String_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *String_var_declarationContext) TEXT() antlr.TerminalNode {
	return s.GetToken(VariablesTEXT, 0)
}

func (s *String_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *String_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VariablesNEWLINE)
}

func (s *String_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VariablesNEWLINE, i)
}

func (s *String_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterString_var_declaration(s)
	}
}

func (s *String_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitString_var_declaration(s)
	}
}

func (p *Variables) String_var_declaration() (localctx IString_var_declarationContext) {
	localctx = NewString_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, VariablesRULE_string_var_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(VariablesTEXT_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Var_name()
	}
	{
		p.SetState(103)
		p.Match(VariablesL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VariablesNEWLINE {
		{
			p.SetState(104)
			p.Match(VariablesNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)
		p.Match(VariablesTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == VariablesNEWLINE {
		{
			p.SetState(111)
			p.Match(VariablesNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(VariablesR_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumber_var_declarationContext is an interface to support dynamic dispatch.
type INumber_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Math_expression() IMath_expressionContext

	// IsNumber_var_declarationContext differentiates from other interfaces.
	IsNumber_var_declarationContext()
}

type Number_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumber_var_declarationContext() *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_number_var_declaration
	return p
}

func InitEmptyNumber_var_declarationContext(p *Number_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_number_var_declaration
}

func (*Number_var_declarationContext) IsNumber_var_declarationContext() {}

func NewNumber_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_number_var_declaration

	return p
}

func (s *Number_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_var_declarationContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER_TYPE, 0)
}

func (s *Number_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Number_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Number_var_declarationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER, 0)
}

func (s *Number_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Number_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VariablesNEWLINE)
}

func (s *Number_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VariablesNEWLINE, i)
}

func (s *Number_var_declarationContext) Math_expression() IMath_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Number_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Number_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Number_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterNumber_var_declaration(s)
	}
}

func (s *Number_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitNumber_var_declaration(s)
	}
}

func (p *Variables) Number_var_declaration() (localctx INumber_var_declarationContext) {
	localctx = NewNumber_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, VariablesRULE_number_var_declaration)
	var _la int

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Var_name()
		}
		{
			p.SetState(121)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(122)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(128)
			p.Match(VariablesNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(129)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(135)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(VariablesNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Var_name()
		}
		{
			p.SetState(139)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(140)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(146)
			p.Math_expression()
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(147)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(153)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolean_var_declarationContext is an interface to support dynamic dispatch.
type IBoolean_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN_TYPE() antlr.TerminalNode
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Logical_expression() ILogical_expressionContext
	Comparative_expression() IComparative_expressionContext

	// IsBoolean_var_declarationContext differentiates from other interfaces.
	IsBoolean_var_declarationContext()
}

type Boolean_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_var_declarationContext() *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_boolean_var_declaration
	return p
}

func InitEmptyBoolean_var_declarationContext(p *Boolean_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_boolean_var_declaration
}

func (*Boolean_var_declarationContext) IsBoolean_var_declarationContext() {}

func NewBoolean_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_boolean_var_declaration

	return p
}

func (s *Boolean_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_var_declarationContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN_TYPE, 0)
}

func (s *Boolean_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Boolean_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN, 0)
}

func (s *Boolean_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VariablesNEWLINE)
}

func (s *Boolean_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VariablesNEWLINE, i)
}

func (s *Boolean_var_declarationContext) Logical_expression() ILogical_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Boolean_var_declarationContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Boolean_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterBoolean_var_declaration(s)
	}
}

func (s *Boolean_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitBoolean_var_declaration(s)
	}
}

func (p *Variables) Boolean_var_declaration() (localctx IBoolean_var_declarationContext) {
	localctx = NewBoolean_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, VariablesRULE_boolean_var_declaration)
	var _la int

	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(VariablesBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Var_name()
		}
		{
			p.SetState(159)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(160)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(166)
			p.Match(VariablesBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(167)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(172)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(173)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(VariablesBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Var_name()
		}
		{
			p.SetState(177)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(178)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(184)
			p.logical_expression(0)
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(185)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(191)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.Match(VariablesBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Var_name()
		}
		{
			p.SetState(195)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(196)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(202)
			p.comparative_expression(0)
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(203)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(209)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingle_var_declarationContext is an interface to support dynamic dispatch.
type ISingle_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_var_declaration() IString_var_declarationContext
	Number_var_declaration() INumber_var_declarationContext
	Boolean_var_declaration() IBoolean_var_declarationContext
	Var_type() IVar_typeContext
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsSingle_var_declarationContext differentiates from other interfaces.
	IsSingle_var_declarationContext()
}

type Single_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_var_declarationContext() *Single_var_declarationContext {
	var p = new(Single_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_single_var_declaration
	return p
}

func InitEmptySingle_var_declarationContext(p *Single_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_single_var_declaration
}

func (*Single_var_declarationContext) IsSingle_var_declarationContext() {}

func NewSingle_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_var_declarationContext {
	var p = new(Single_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_single_var_declaration

	return p
}

func (s *Single_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_var_declarationContext) String_var_declaration() IString_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_var_declarationContext)
}

func (s *Single_var_declarationContext) Number_var_declaration() INumber_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumber_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumber_var_declarationContext)
}

func (s *Single_var_declarationContext) Boolean_var_declaration() IBoolean_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolean_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolean_var_declarationContext)
}

func (s *Single_var_declarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Single_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Single_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Single_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Single_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VariablesNEWLINE)
}

func (s *Single_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VariablesNEWLINE, i)
}

func (s *Single_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterSingle_var_declaration(s)
	}
}

func (s *Single_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitSingle_var_declaration(s)
	}
}

func (p *Variables) Single_var_declaration() (localctx ISingle_var_declarationContext) {
	localctx = NewSingle_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, VariablesRULE_single_var_declaration)
	var _la int

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)
			p.String_var_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Number_var_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(215)
			p.Boolean_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(216)
			p.Var_type()
		}
		{
			p.SetState(217)
			p.Var_name()
		}
		{
			p.SetState(218)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(219)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(224)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(225)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_var_declarationContext is an interface to support dynamic dispatch.
type IArray_var_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_type() IVar_typeContext
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	AllVar_value() []IVar_valueContext
	Var_value(i int) IVar_valueContext
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllMath_expression() []IMath_expressionContext
	Math_expression(i int) IMath_expressionContext

	// IsArray_var_declarationContext differentiates from other interfaces.
	IsArray_var_declarationContext()
}

type Array_var_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_var_declarationContext() *Array_var_declarationContext {
	var p = new(Array_var_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_array_var_declaration
	return p
}

func InitEmptyArray_var_declarationContext(p *Array_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_array_var_declaration
}

func (*Array_var_declarationContext) IsArray_var_declarationContext() {}

func NewArray_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_var_declarationContext {
	var p = new(Array_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_array_var_declaration

	return p
}

func (s *Array_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_var_declarationContext) Var_type() IVar_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_typeContext)
}

func (s *Array_var_declarationContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Array_var_declarationContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Array_var_declarationContext) AllVar_value() []IVar_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVar_valueContext); ok {
			len++
		}
	}

	tst := make([]IVar_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVar_valueContext); ok {
			tst[i] = t.(IVar_valueContext)
			i++
		}
	}

	return tst
}

func (s *Array_var_declarationContext) Var_value(i int) IVar_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_valueContext)
}

func (s *Array_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Array_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VariablesNEWLINE)
}

func (s *Array_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VariablesNEWLINE, i)
}

func (s *Array_var_declarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VariablesCOMMA)
}

func (s *Array_var_declarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VariablesCOMMA, i)
}

func (s *Array_var_declarationContext) AllMath_expression() []IMath_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMath_expressionContext); ok {
			len++
		}
	}

	tst := make([]IMath_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMath_expressionContext); ok {
			tst[i] = t.(IMath_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Array_var_declarationContext) Math_expression(i int) IMath_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Array_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterArray_var_declaration(s)
	}
}

func (s *Array_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitArray_var_declaration(s)
	}
}

func (p *Variables) Array_var_declaration() (localctx IArray_var_declarationContext) {
	localctx = NewArray_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, VariablesRULE_array_var_declaration)
	var _la int

	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Var_type()
		}
		{
			p.SetState(230)
			p.Var_name()
		}
		{
			p.SetState(231)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(232)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(237)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(238)
			p.Var_value()
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesCOMMA {
			{
				p.SetState(239)
				p.Match(VariablesCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(243)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VariablesNEWLINE {
				{
					p.SetState(240)
					p.Match(VariablesNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(245)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(246)
				p.Var_value()
			}

			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(252)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(258)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Var_type()
		}
		{
			p.SetState(261)
			p.Var_name()
		}
		{
			p.SetState(262)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(263)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(268)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(269)
			p.Math_expression()
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesCOMMA {
			{
				p.SetState(270)
				p.Match(VariablesCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == VariablesNEWLINE {
				{
					p.SetState(271)
					p.Match(VariablesNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(276)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(277)
				p.Math_expression()
			}

			p.SetState(282)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(283)
			p.Match(VariablesNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_assignmentContext is an interface to support dynamic dispatch.
type IVar_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_name() IVar_nameContext
	L_BLOCK() antlr.TerminalNode
	Var_value() IVar_valueContext
	R_BLOCK() antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode
	Math_expression() IMath_expressionContext

	// IsVar_assignmentContext differentiates from other interfaces.
	IsVar_assignmentContext()
}

type Var_assignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_assignmentContext() *Var_assignmentContext {
	var p = new(Var_assignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_assignment
	return p
}

func InitEmptyVar_assignmentContext(p *Var_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_var_assignment
}

func (*Var_assignmentContext) IsVar_assignmentContext() {}

func NewVar_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignmentContext {
	var p = new(Var_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_var_assignment

	return p
}

func (s *Var_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_assignmentContext) Var_name() IVar_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_nameContext)
}

func (s *Var_assignmentContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesL_BLOCK, 0)
}

func (s *Var_assignmentContext) Var_value() IVar_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_valueContext)
}

func (s *Var_assignmentContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(VariablesR_BLOCK, 0)
}

func (s *Var_assignmentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(VariablesNEWLINE)
}

func (s *Var_assignmentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(VariablesNEWLINE, i)
}

func (s *Var_assignmentContext) Math_expression() IMath_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Var_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterVar_assignment(s)
	}
}

func (s *Var_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitVar_assignment(s)
	}
}

func (p *Variables) Var_assignment() (localctx IVar_assignmentContext) {
	localctx = NewVar_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, VariablesRULE_var_assignment)
	var _la int

	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)
			p.Var_name()
		}
		{
			p.SetState(289)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(290)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(296)
			p.Var_value()
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(297)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(303)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Var_name()
		}
		{
			p.SetState(306)
			p.Match(VariablesL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(307)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(312)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(313)
			p.Math_expression()
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesNEWLINE {
			{
				p.SetState(314)
				p.Match(VariablesNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(319)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(320)
			p.Match(VariablesR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level1Context is an interface to support dynamic dispatch.
type IOp_level1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_MULTIPLY() antlr.TerminalNode
	OP_DIVIDE() antlr.TerminalNode
	OP_MOD() antlr.TerminalNode
	OP_POWER() antlr.TerminalNode

	// IsOp_level1Context differentiates from other interfaces.
	IsOp_level1Context()
}

type Op_level1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level1Context() *Op_level1Context {
	var p = new(Op_level1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level1
	return p
}

func InitEmptyOp_level1Context(p *Op_level1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level1
}

func (*Op_level1Context) IsOp_level1Context() {}

func NewOp_level1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level1Context {
	var p = new(Op_level1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_op_level1

	return p
}

func (s *Op_level1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level1Context) OP_MULTIPLY() antlr.TerminalNode {
	return s.GetToken(VariablesOP_MULTIPLY, 0)
}

func (s *Op_level1Context) OP_DIVIDE() antlr.TerminalNode {
	return s.GetToken(VariablesOP_DIVIDE, 0)
}

func (s *Op_level1Context) OP_MOD() antlr.TerminalNode {
	return s.GetToken(VariablesOP_MOD, 0)
}

func (s *Op_level1Context) OP_POWER() antlr.TerminalNode {
	return s.GetToken(VariablesOP_POWER, 0)
}

func (s *Op_level1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterOp_level1(s)
	}
}

func (s *Op_level1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitOp_level1(s)
	}
}

func (p *Variables) Op_level1() (localctx IOp_level1Context) {
	localctx = NewOp_level1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, VariablesRULE_op_level1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&970662608896) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level2Context is an interface to support dynamic dispatch.
type IOp_level2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_PLUS() antlr.TerminalNode
	OP_MINUS() antlr.TerminalNode

	// IsOp_level2Context differentiates from other interfaces.
	IsOp_level2Context()
}

type Op_level2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level2Context() *Op_level2Context {
	var p = new(Op_level2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level2
	return p
}

func InitEmptyOp_level2Context(p *Op_level2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level2
}

func (*Op_level2Context) IsOp_level2Context() {}

func NewOp_level2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level2Context {
	var p = new(Op_level2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_op_level2

	return p
}

func (s *Op_level2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level2Context) OP_PLUS() antlr.TerminalNode {
	return s.GetToken(VariablesOP_PLUS, 0)
}

func (s *Op_level2Context) OP_MINUS() antlr.TerminalNode {
	return s.GetToken(VariablesOP_MINUS, 0)
}

func (s *Op_level2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterOp_level2(s)
	}
}

func (s *Op_level2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitOp_level2(s)
	}
}

func (p *Variables) Op_level2() (localctx IOp_level2Context) {
	localctx = NewOp_level2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, VariablesRULE_op_level2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		_la = p.GetTokenStream().LA(1)

		if !(_la == VariablesOP_PLUS || _la == VariablesOP_MINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level3Context is an interface to support dynamic dispatch.
type IOp_level3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_AND() antlr.TerminalNode
	OP_OR() antlr.TerminalNode
	OP_XOR() antlr.TerminalNode

	// IsOp_level3Context differentiates from other interfaces.
	IsOp_level3Context()
}

type Op_level3Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level3Context() *Op_level3Context {
	var p = new(Op_level3Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level3
	return p
}

func InitEmptyOp_level3Context(p *Op_level3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level3
}

func (*Op_level3Context) IsOp_level3Context() {}

func NewOp_level3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level3Context {
	var p = new(Op_level3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_op_level3

	return p
}

func (s *Op_level3Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level3Context) OP_AND() antlr.TerminalNode {
	return s.GetToken(VariablesOP_AND, 0)
}

func (s *Op_level3Context) OP_OR() antlr.TerminalNode {
	return s.GetToken(VariablesOP_OR, 0)
}

func (s *Op_level3Context) OP_XOR() antlr.TerminalNode {
	return s.GetToken(VariablesOP_XOR, 0)
}

func (s *Op_level3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterOp_level3(s)
	}
}

func (s *Op_level3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitOp_level3(s)
	}
}

func (p *Variables) Op_level3() (localctx IOp_level3Context) {
	localctx = NewOp_level3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, VariablesRULE_op_level3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30786325577728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level4Context is an interface to support dynamic dispatch.
type IOp_level4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_EQUAL() antlr.TerminalNode
	OP_NOT_EQUAL() antlr.TerminalNode
	OP_LESS() antlr.TerminalNode
	OP_LESS_EQUAL() antlr.TerminalNode
	OP_GREATER() antlr.TerminalNode
	OP_GREATER_EQUAL() antlr.TerminalNode

	// IsOp_level4Context differentiates from other interfaces.
	IsOp_level4Context()
}

type Op_level4Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level4Context() *Op_level4Context {
	var p = new(Op_level4Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level4
	return p
}

func InitEmptyOp_level4Context(p *Op_level4Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level4
}

func (*Op_level4Context) IsOp_level4Context() {}

func NewOp_level4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level4Context {
	var p = new(Op_level4Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_op_level4

	return p
}

func (s *Op_level4Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level4Context) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(VariablesOP_EQUAL, 0)
}

func (s *Op_level4Context) OP_NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(VariablesOP_NOT_EQUAL, 0)
}

func (s *Op_level4Context) OP_LESS() antlr.TerminalNode {
	return s.GetToken(VariablesOP_LESS, 0)
}

func (s *Op_level4Context) OP_LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(VariablesOP_LESS_EQUAL, 0)
}

func (s *Op_level4Context) OP_GREATER() antlr.TerminalNode {
	return s.GetToken(VariablesOP_GREATER, 0)
}

func (s *Op_level4Context) OP_GREATER_EQUAL() antlr.TerminalNode {
	return s.GetToken(VariablesOP_GREATER_EQUAL, 0)
}

func (s *Op_level4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterOp_level4(s)
	}
}

func (s *Op_level4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitOp_level4(s)
	}
}

func (p *Variables) Op_level4() (localctx IOp_level4Context) {
	localctx = NewOp_level4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, VariablesRULE_op_level4)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1090715668971520) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOp_level5Context is an interface to support dynamic dispatch.
type IOp_level5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_NOT() antlr.TerminalNode

	// IsOp_level5Context differentiates from other interfaces.
	IsOp_level5Context()
}

type Op_level5Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOp_level5Context() *Op_level5Context {
	var p = new(Op_level5Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level5
	return p
}

func InitEmptyOp_level5Context(p *Op_level5Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_op_level5
}

func (*Op_level5Context) IsOp_level5Context() {}

func NewOp_level5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level5Context {
	var p = new(Op_level5Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_op_level5

	return p
}

func (s *Op_level5Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level5Context) OP_NOT() antlr.TerminalNode {
	return s.GetToken(VariablesOP_NOT, 0)
}

func (s *Op_level5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterOp_level5(s)
	}
}

func (s *Op_level5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitOp_level5(s)
	}
}

func (p *Variables) Op_level5() (localctx IOp_level5Context) {
	localctx = NewOp_level5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, VariablesRULE_op_level5)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(VariablesOP_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumberic_expressionContext is an interface to support dynamic dispatch.
type INumberic_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllNumberic_expression() []INumberic_expressionContext
	Numberic_expression(i int) INumberic_expressionContext
	R_PAREN() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Get_variable() IGet_variableContext
	Function_call() IFunction_callContext
	OP_MINUS() antlr.TerminalNode
	Op_level1() IOp_level1Context
	Op_level2() IOp_level2Context

	// IsNumberic_expressionContext differentiates from other interfaces.
	IsNumberic_expressionContext()
}

type Numberic_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberic_expressionContext() *Numberic_expressionContext {
	var p = new(Numberic_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_numberic_expression
	return p
}

func InitEmptyNumberic_expressionContext(p *Numberic_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_numberic_expression
}

func (*Numberic_expressionContext) IsNumberic_expressionContext() {}

func NewNumberic_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numberic_expressionContext {
	var p = new(Numberic_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_numberic_expression

	return p
}

func (s *Numberic_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Numberic_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(VariablesL_PAREN, 0)
}

func (s *Numberic_expressionContext) AllNumberic_expression() []INumberic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			len++
		}
	}

	tst := make([]INumberic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberic_expressionContext); ok {
			tst[i] = t.(INumberic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Numberic_expressionContext) Numberic_expression(i int) INumberic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Numberic_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(VariablesR_PAREN, 0)
}

func (s *Numberic_expressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(VariablesNUMBER, 0)
}

func (s *Numberic_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Numberic_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Numberic_expressionContext) OP_MINUS() antlr.TerminalNode {
	return s.GetToken(VariablesOP_MINUS, 0)
}

func (s *Numberic_expressionContext) Op_level1() IOp_level1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level1Context)
}

func (s *Numberic_expressionContext) Op_level2() IOp_level2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level2Context)
}

func (s *Numberic_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numberic_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numberic_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterNumberic_expression(s)
	}
}

func (s *Numberic_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitNumberic_expression(s)
	}
}

func (p *Variables) Numberic_expression() (localctx INumberic_expressionContext) {
	return p.numberic_expression(0)
}

func (p *Variables) numberic_expression(_p int) (localctx INumberic_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewNumberic_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumberic_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, VariablesRULE_numberic_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(335)
			p.Match(VariablesL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.numberic_expression(0)
		}
		{
			p.SetState(337)
			p.Match(VariablesR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(339)
			p.Match(VariablesNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(340)
			p.Get_variable()
		}

	case 4:
		{
			p.SetState(341)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(342)
			p.Match(VariablesOP_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.numberic_expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(354)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewNumberic_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VariablesRULE_numberic_expression)
				p.SetState(346)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(347)
					p.Op_level1()
				}
				{
					p.SetState(348)
					p.numberic_expression(7)
				}

			case 2:
				localctx = NewNumberic_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VariablesRULE_numberic_expression)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(351)
					p.Op_level2()
				}
				{
					p.SetState(352)
					p.numberic_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogical_expressionContext is an interface to support dynamic dispatch.
type ILogical_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllLogical_expression() []ILogical_expressionContext
	Logical_expression(i int) ILogical_expressionContext
	R_PAREN() antlr.TerminalNode
	OP_NOT() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	Get_variable() IGet_variableContext
	Function_call() IFunction_callContext
	Op_level3() IOp_level3Context
	Op_level4() IOp_level4Context

	// IsLogical_expressionContext differentiates from other interfaces.
	IsLogical_expressionContext()
}

type Logical_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_expressionContext() *Logical_expressionContext {
	var p = new(Logical_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_logical_expression
	return p
}

func InitEmptyLogical_expressionContext(p *Logical_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_logical_expression
}

func (*Logical_expressionContext) IsLogical_expressionContext() {}

func NewLogical_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_expressionContext {
	var p = new(Logical_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_logical_expression

	return p
}

func (s *Logical_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(VariablesL_PAREN, 0)
}

func (s *Logical_expressionContext) AllLogical_expression() []ILogical_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			len++
		}
	}

	tst := make([]ILogical_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogical_expressionContext); ok {
			tst[i] = t.(ILogical_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Logical_expressionContext) Logical_expression(i int) ILogical_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Logical_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(VariablesR_PAREN, 0)
}

func (s *Logical_expressionContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(VariablesOP_NOT, 0)
}

func (s *Logical_expressionContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(VariablesBOOLEAN, 0)
}

func (s *Logical_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Logical_expressionContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Logical_expressionContext) Op_level3() IOp_level3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level3Context)
}

func (s *Logical_expressionContext) Op_level4() IOp_level4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level4Context)
}

func (s *Logical_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterLogical_expression(s)
	}
}

func (s *Logical_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitLogical_expression(s)
	}
}

func (p *Variables) Logical_expression() (localctx ILogical_expressionContext) {
	return p.logical_expression(0)
}

func (p *Variables) logical_expression(_p int) (localctx ILogical_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogical_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogical_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, VariablesRULE_logical_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(360)
			p.Match(VariablesL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.logical_expression(0)
		}
		{
			p.SetState(362)
			p.Match(VariablesR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(364)
			p.Match(VariablesOP_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(365)
			p.logical_expression(4)
		}

	case 3:
		{
			p.SetState(366)
			p.Match(VariablesBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(367)
			p.Get_variable()
		}

	case 5:
		{
			p.SetState(368)
			p.Function_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(379)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VariablesRULE_logical_expression)
				p.SetState(371)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(372)
					p.Op_level3()
				}
				{
					p.SetState(373)
					p.logical_expression(7)
				}

			case 2:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, VariablesRULE_logical_expression)
				p.SetState(375)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(376)
					p.Op_level4()
				}
				{
					p.SetState(377)
					p.logical_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparative_expressionContext is an interface to support dynamic dispatch.
type IComparative_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumberic_expression() []INumberic_expressionContext
	Numberic_expression(i int) INumberic_expressionContext
	Op_level4() IOp_level4Context
	AllLogical_expression() []ILogical_expressionContext
	Logical_expression(i int) ILogical_expressionContext
	Op_level3() IOp_level3Context
	Op_level5() IOp_level5Context
	L_PAREN() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	R_PAREN() antlr.TerminalNode

	// IsComparative_expressionContext differentiates from other interfaces.
	IsComparative_expressionContext()
}

type Comparative_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparative_expressionContext() *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_comparative_expression
	return p
}

func InitEmptyComparative_expressionContext(p *Comparative_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_comparative_expression
}

func (*Comparative_expressionContext) IsComparative_expressionContext() {}

func NewComparative_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_comparative_expression

	return p
}

func (s *Comparative_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparative_expressionContext) AllNumberic_expression() []INumberic_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			len++
		}
	}

	tst := make([]INumberic_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberic_expressionContext); ok {
			tst[i] = t.(INumberic_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparative_expressionContext) Numberic_expression(i int) INumberic_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Comparative_expressionContext) Op_level4() IOp_level4Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level4Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level4Context)
}

func (s *Comparative_expressionContext) AllLogical_expression() []ILogical_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			len++
		}
	}

	tst := make([]ILogical_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogical_expressionContext); ok {
			tst[i] = t.(ILogical_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Comparative_expressionContext) Logical_expression(i int) ILogical_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Comparative_expressionContext) Op_level3() IOp_level3Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level3Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level3Context)
}

func (s *Comparative_expressionContext) Op_level5() IOp_level5Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOp_level5Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOp_level5Context)
}

func (s *Comparative_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(VariablesL_PAREN, 0)
}

func (s *Comparative_expressionContext) Comparative_expression() IComparative_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparative_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparative_expressionContext)
}

func (s *Comparative_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(VariablesR_PAREN, 0)
}

func (s *Comparative_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparative_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparative_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterComparative_expression(s)
	}
}

func (s *Comparative_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitComparative_expression(s)
	}
}

func (p *Variables) Comparative_expression() (localctx IComparative_expressionContext) {
	return p.comparative_expression(0)
}

func (p *Variables) comparative_expression(_p int) (localctx IComparative_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewComparative_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparative_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, VariablesRULE_comparative_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(385)
			p.numberic_expression(0)
		}
		{
			p.SetState(386)
			p.Op_level4()
		}
		{
			p.SetState(387)
			p.numberic_expression(0)
		}

	case 2:
		{
			p.SetState(389)
			p.logical_expression(0)
		}
		{
			p.SetState(390)
			p.Op_level3()
		}
		{
			p.SetState(391)
			p.logical_expression(0)
		}

	case 3:
		{
			p.SetState(393)
			p.Op_level5()
		}
		{
			p.SetState(394)
			p.Match(VariablesL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.comparative_expression(0)
		}
		{
			p.SetState(396)
			p.Match(VariablesR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewComparative_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, VariablesRULE_comparative_expression)
			p.SetState(400)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(401)
				p.Op_level3()
			}
			{
				p.SetState(402)
				p.logical_expression(0)
			}

		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_variableContext is an interface to support dynamic dispatch.
type IGet_variableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	Get_array_element() IGet_array_elementContext
	Get_child() IGet_childContext

	// IsGet_variableContext differentiates from other interfaces.
	IsGet_variableContext()
}

type Get_variableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_variableContext() *Get_variableContext {
	var p = new(Get_variableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_get_variable
	return p
}

func InitEmptyGet_variableContext(p *Get_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_get_variable
}

func (*Get_variableContext) IsGet_variableContext() {}

func NewGet_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_variableContext {
	var p = new(Get_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_get_variable

	return p
}

func (s *Get_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_variableContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VariablesVAR_IDENTIFIER, 0)
}

func (s *Get_variableContext) Get_array_element() IGet_array_elementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_array_elementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_array_elementContext)
}

func (s *Get_variableContext) Get_child() IGet_childContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_childContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_childContext)
}

func (s *Get_variableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_variableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_variableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterGet_variable(s)
	}
}

func (s *Get_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitGet_variable(s)
	}
}

func (p *Variables) Get_variable() (localctx IGet_variableContext) {
	localctx = NewGet_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, VariablesRULE_get_variable)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.Match(VariablesVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
			p.Get_child()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMath_expressionContext is an interface to support dynamic dispatch.
type IMath_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Get_variable() IGet_variableContext
	Numberic_expression() INumberic_expressionContext
	Logical_expression() ILogical_expressionContext

	// IsMath_expressionContext differentiates from other interfaces.
	IsMath_expressionContext()
}

type Math_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMath_expressionContext() *Math_expressionContext {
	var p = new(Math_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_math_expression
	return p
}

func InitEmptyMath_expressionContext(p *Math_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_math_expression
}

func (*Math_expressionContext) IsMath_expressionContext() {}

func NewMath_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_expressionContext {
	var p = new(Math_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_math_expression

	return p
}

func (s *Math_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Math_expressionContext) Get_variable() IGet_variableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_variableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_variableContext)
}

func (s *Math_expressionContext) Numberic_expression() INumberic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Math_expressionContext) Logical_expression() ILogical_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_expressionContext)
}

func (s *Math_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Math_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Math_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterMath_expression(s)
	}
}

func (s *Math_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitMath_expression(s)
	}
}

func (p *Variables) Math_expression() (localctx IMath_expressionContext) {
	localctx = NewMath_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, VariablesRULE_math_expression)
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			p.Get_variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(415)
			p.numberic_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(416)
			p.logical_expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_array_elementContext is an interface to support dynamic dispatch.
type IGet_array_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	L_SQUARE() antlr.TerminalNode
	Numberic_expression() INumberic_expressionContext
	R_SQUARE() antlr.TerminalNode

	// IsGet_array_elementContext differentiates from other interfaces.
	IsGet_array_elementContext()
}

type Get_array_elementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_array_elementContext() *Get_array_elementContext {
	var p = new(Get_array_elementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_get_array_element
	return p
}

func InitEmptyGet_array_elementContext(p *Get_array_elementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_get_array_element
}

func (*Get_array_elementContext) IsGet_array_elementContext() {}

func NewGet_array_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_array_elementContext {
	var p = new(Get_array_elementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_get_array_element

	return p
}

func (s *Get_array_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_array_elementContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VariablesVAR_IDENTIFIER, 0)
}

func (s *Get_array_elementContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(VariablesL_SQUARE, 0)
}

func (s *Get_array_elementContext) Numberic_expression() INumberic_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberic_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberic_expressionContext)
}

func (s *Get_array_elementContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(VariablesR_SQUARE, 0)
}

func (s *Get_array_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_array_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_array_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterGet_array_element(s)
	}
}

func (s *Get_array_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitGet_array_element(s)
	}
}

func (p *Variables) Get_array_element() (localctx IGet_array_elementContext) {
	localctx = NewGet_array_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, VariablesRULE_get_array_element)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(VariablesVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(420)
		p.Match(VariablesL_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(421)
		p.numberic_expression(0)
	}
	{
		p.SetState(422)
		p.Match(VariablesR_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGet_childContext is an interface to support dynamic dispatch.
type IGet_childContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVAR_IDENTIFIER() []antlr.TerminalNode
	VAR_IDENTIFIER(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode
	Get_array_element() IGet_array_elementContext
	Get_child() IGet_childContext

	// IsGet_childContext differentiates from other interfaces.
	IsGet_childContext()
}

type Get_childContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGet_childContext() *Get_childContext {
	var p = new(Get_childContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_get_child
	return p
}

func InitEmptyGet_childContext(p *Get_childContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_get_child
}

func (*Get_childContext) IsGet_childContext() {}

func NewGet_childContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_childContext {
	var p = new(Get_childContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_get_child

	return p
}

func (s *Get_childContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_childContext) AllVAR_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(VariablesVAR_IDENTIFIER)
}

func (s *Get_childContext) VAR_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(VariablesVAR_IDENTIFIER, i)
}

func (s *Get_childContext) DOT() antlr.TerminalNode {
	return s.GetToken(VariablesDOT, 0)
}

func (s *Get_childContext) Get_array_element() IGet_array_elementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_array_elementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_array_elementContext)
}

func (s *Get_childContext) Get_child() IGet_childContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGet_childContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGet_childContext)
}

func (s *Get_childContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_childContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_childContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterGet_child(s)
	}
}

func (s *Get_childContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitGet_child(s)
	}
}

func (p *Variables) Get_child() (localctx IGet_childContext) {
	localctx = NewGet_childContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, VariablesRULE_get_child)
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(424)
			p.Match(VariablesVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.Match(VariablesDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.Match(VariablesVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(427)
			p.Match(VariablesVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(VariablesDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(430)
			p.Match(VariablesVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.Match(VariablesDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Get_child()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(433)
			p.Get_array_element()
		}
		{
			p.SetState(434)
			p.Match(VariablesDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Get_child()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(437)
			p.Get_array_element()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(438)
			p.Match(VariablesVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_IDENTIFIER() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllMath_expression() []IMath_expressionContext
	Math_expression(i int) IMath_expressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VariablesRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VariablesRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(VariablesVAR_IDENTIFIER, 0)
}

func (s *Function_callContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(VariablesL_PAREN, 0)
}

func (s *Function_callContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(VariablesR_PAREN, 0)
}

func (s *Function_callContext) AllMath_expression() []IMath_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMath_expressionContext); ok {
			len++
		}
	}

	tst := make([]IMath_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMath_expressionContext); ok {
			tst[i] = t.(IMath_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Math_expression(i int) IMath_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMath_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMath_expressionContext)
}

func (s *Function_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(VariablesCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(VariablesCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VariablesListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *Variables) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, VariablesRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.Match(VariablesVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)
		p.Match(VariablesL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3379898743791688) != 0 {
		{
			p.SetState(443)
			p.Math_expression()
		}
		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == VariablesCOMMA {
			{
				p.SetState(444)
				p.Match(VariablesCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(445)
				p.Math_expression()
			}

			p.SetState(450)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(453)
		p.Match(VariablesR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *Variables) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *Numberic_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Numberic_expressionContext)
		}
		return p.Numberic_expression_Sempred(t, predIndex)

	case 18:
		var t *Logical_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Logical_expressionContext)
		}
		return p.Logical_expression_Sempred(t, predIndex)

	case 19:
		var t *Comparative_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Comparative_expressionContext)
		}
		return p.Comparative_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Variables) Numberic_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Variables) Logical_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Variables) Comparative_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
