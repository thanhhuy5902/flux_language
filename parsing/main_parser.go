// Code generated from Main.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parsing // Main
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

type Main struct {
	*antlr.BaseParser
}

var MainParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mainParserInit() {
	staticData := &MainParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'text'", "", "'boolean'", "", "", "", "", "'num'", "", "",
		"'ipv4'", "", "", "", "", "'{'", "'}'", "'@'",
	}
	staticData.SymbolicNames = []string{
		"", "TEXT", "TEXT_TYPE", "BOOLEAN", "BOOLEAN_TYPE", "DIGIT", "NUMBER",
		"UNIT", "SI_POSTFIX", "NUMBER_TYPE", "OCTET", "IPV4", "IPV4_TYPE", "VAR_IDENTIFIER",
		"COMMON_IDENTIFIER", "ARITHMETIC_OPERATOR", "LOGICAL_OPERATOR", "L_BLOCK",
		"R_BLOCK", "AT", "WS", "NEWLINE", "LOOP", "IF", "ELSE", "RETURN", "FUNC",
		"L_PAREN", "COMMA", "R_PAREN", "RETURN_TYPE", "OP_INCREMENT", "OP_DECREMENT",
		"OP_ADD_ASSIGN", "OP_SUB_ASSIGN", "OP_MUL_ASSIGN", "OP_DIV_ASSIGN",
		"OP_EQUAL", "NUMBE", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "OP_MOD",
		"OP_POW", "OP_MULTIPLY", "OP_DIVIDE", "OP_POWER", "OP_PLUS", "OP_MINUS",
		"OP_AND", "OP_OR", "OP_XOR", "OP_NOT_EQUAL", "OP_LESS", "OP_LESS_EQUAL",
		"OP_GREATER", "OP_GREATER_EQUAL", "OP_NOT", "L_SQUARE", "R_SQUARE",
		"DOT",
	}
	staticData.RuleNames = []string{
		"program", "statement", "expression", "block", "loop_statement", "if_statement",
		"return_statement", "data_type", "func_declaration", "var_type", "var_name",
		"var_value", "op_one_expression", "op_one_declaration", "op_two_expression",
		"string_var_declaration", "number_var_declaration", "boolean_var_declaration",
		"single_var_declaration", "array_var_declaration", "var_assignment",
		"op_level1", "op_level2", "op_level3", "op_level4", "op_level5", "numberic_expression",
		"logical_expression", "comparative_expression", "get_variable", "math_expression",
		"get_array_element", "get_child", "function_call",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 61, 547, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 5, 0, 70, 8, 0, 10, 0, 12, 0, 73,
		9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 83, 8, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 90, 8, 2, 1, 3, 1, 3, 5, 3, 94, 8, 3,
		10, 3, 12, 3, 97, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 113, 8, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 128,
		8, 8, 10, 8, 12, 8, 131, 9, 8, 3, 8, 133, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		138, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 189, 8,
		13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 197, 8, 15, 10, 15,
		12, 15, 200, 9, 15, 1, 15, 1, 15, 5, 15, 204, 8, 15, 10, 15, 12, 15, 207,
		9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 215, 8, 16, 10,
		16, 12, 16, 218, 9, 16, 1, 16, 1, 16, 5, 16, 222, 8, 16, 10, 16, 12, 16,
		225, 9, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 233, 8, 16,
		10, 16, 12, 16, 236, 9, 16, 1, 16, 1, 16, 5, 16, 240, 8, 16, 10, 16, 12,
		16, 243, 9, 16, 1, 16, 1, 16, 3, 16, 247, 8, 16, 1, 17, 1, 17, 1, 17, 1,
		17, 5, 17, 253, 8, 17, 10, 17, 12, 17, 256, 9, 17, 1, 17, 1, 17, 5, 17,
		260, 8, 17, 10, 17, 12, 17, 263, 9, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 5, 17, 271, 8, 17, 10, 17, 12, 17, 274, 9, 17, 1, 17, 1, 17,
		5, 17, 278, 8, 17, 10, 17, 12, 17, 281, 9, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 5, 17, 289, 8, 17, 10, 17, 12, 17, 292, 9, 17, 1, 17,
		1, 17, 5, 17, 296, 8, 17, 10, 17, 12, 17, 299, 9, 17, 1, 17, 1, 17, 3,
		17, 303, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18,
		312, 8, 18, 10, 18, 12, 18, 315, 9, 18, 1, 18, 1, 18, 3, 18, 319, 8, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 325, 8, 19, 10, 19, 12, 19, 328, 9,
		19, 1, 19, 1, 19, 1, 19, 5, 19, 333, 8, 19, 10, 19, 12, 19, 336, 9, 19,
		1, 19, 5, 19, 339, 8, 19, 10, 19, 12, 19, 342, 9, 19, 1, 19, 5, 19, 345,
		8, 19, 10, 19, 12, 19, 348, 9, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 5, 19, 356, 8, 19, 10, 19, 12, 19, 359, 9, 19, 1, 19, 1, 19, 1, 19,
		5, 19, 364, 8, 19, 10, 19, 12, 19, 367, 9, 19, 1, 19, 5, 19, 370, 8, 19,
		10, 19, 12, 19, 373, 9, 19, 1, 19, 1, 19, 1, 19, 3, 19, 378, 8, 19, 1,
		20, 1, 20, 1, 20, 5, 20, 383, 8, 20, 10, 20, 12, 20, 386, 9, 20, 1, 20,
		1, 20, 5, 20, 390, 8, 20, 10, 20, 12, 20, 393, 9, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 5, 20, 400, 8, 20, 10, 20, 12, 20, 403, 9, 20, 1, 20,
		1, 20, 5, 20, 407, 8, 20, 10, 20, 12, 20, 410, 9, 20, 1, 20, 1, 20, 3,
		20, 414, 8, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 3, 26, 436, 8, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 5, 26, 446, 8, 26, 10, 26, 12, 26, 449, 9, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 461,
		8, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 471,
		8, 27, 10, 27, 12, 27, 474, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 490,
		8, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 496, 8, 28, 10, 28, 12, 28, 499,
		9, 28, 1, 29, 1, 29, 1, 29, 3, 29, 504, 8, 29, 1, 30, 1, 30, 1, 30, 3,
		30, 509, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 531, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 5, 33,
		538, 8, 33, 10, 33, 12, 33, 541, 9, 33, 3, 33, 543, 8, 33, 1, 33, 1, 33,
		1, 33, 0, 3, 52, 54, 56, 34, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 0, 9, 4, 0, 2, 2, 4, 4, 9, 9, 14, 14, 4, 0, 2, 2, 4, 4,
		9, 9, 12, 12, 4, 0, 1, 1, 3, 3, 6, 6, 11, 11, 1, 0, 31, 36, 1, 0, 39, 44,
		2, 0, 43, 43, 45, 47, 1, 0, 48, 49, 1, 0, 50, 52, 2, 0, 37, 37, 53, 57,
		593, 0, 71, 1, 0, 0, 0, 2, 82, 1, 0, 0, 0, 4, 89, 1, 0, 0, 0, 6, 91, 1,
		0, 0, 0, 8, 100, 1, 0, 0, 0, 10, 105, 1, 0, 0, 0, 12, 114, 1, 0, 0, 0,
		14, 118, 1, 0, 0, 0, 16, 120, 1, 0, 0, 0, 18, 141, 1, 0, 0, 0, 20, 143,
		1, 0, 0, 0, 22, 145, 1, 0, 0, 0, 24, 147, 1, 0, 0, 0, 26, 188, 1, 0, 0,
		0, 28, 190, 1, 0, 0, 0, 30, 192, 1, 0, 0, 0, 32, 246, 1, 0, 0, 0, 34, 302,
		1, 0, 0, 0, 36, 318, 1, 0, 0, 0, 38, 377, 1, 0, 0, 0, 40, 413, 1, 0, 0,
		0, 42, 415, 1, 0, 0, 0, 44, 417, 1, 0, 0, 0, 46, 419, 1, 0, 0, 0, 48, 421,
		1, 0, 0, 0, 50, 423, 1, 0, 0, 0, 52, 435, 1, 0, 0, 0, 54, 460, 1, 0, 0,
		0, 56, 489, 1, 0, 0, 0, 58, 503, 1, 0, 0, 0, 60, 508, 1, 0, 0, 0, 62, 510,
		1, 0, 0, 0, 64, 530, 1, 0, 0, 0, 66, 532, 1, 0, 0, 0, 68, 70, 3, 2, 1,
		0, 69, 68, 1, 0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72,
		1, 0, 0, 0, 72, 74, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 5, 0, 0, 1,
		75, 1, 1, 0, 0, 0, 76, 83, 3, 4, 2, 0, 77, 83, 3, 8, 4, 0, 78, 83, 3, 10,
		5, 0, 79, 83, 3, 16, 8, 0, 80, 83, 3, 12, 6, 0, 81, 83, 5, 21, 0, 0, 82,
		76, 1, 0, 0, 0, 82, 77, 1, 0, 0, 0, 82, 78, 1, 0, 0, 0, 82, 79, 1, 0, 0,
		0, 82, 80, 1, 0, 0, 0, 82, 81, 1, 0, 0, 0, 83, 3, 1, 0, 0, 0, 84, 90, 3,
		66, 33, 0, 85, 90, 3, 40, 20, 0, 86, 90, 3, 36, 18, 0, 87, 90, 3, 38, 19,
		0, 88, 90, 3, 58, 29, 0, 89, 84, 1, 0, 0, 0, 89, 85, 1, 0, 0, 0, 89, 86,
		1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0, 90, 5, 1, 0, 0, 0,
		91, 95, 5, 17, 0, 0, 92, 94, 3, 2, 1, 0, 93, 92, 1, 0, 0, 0, 94, 97, 1,
		0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0, 97,
		95, 1, 0, 0, 0, 98, 99, 5, 18, 0, 0, 99, 7, 1, 0, 0, 0, 100, 101, 5, 19,
		0, 0, 101, 102, 5, 22, 0, 0, 102, 103, 3, 56, 28, 0, 103, 104, 3, 6, 3,
		0, 104, 9, 1, 0, 0, 0, 105, 106, 5, 19, 0, 0, 106, 107, 5, 23, 0, 0, 107,
		108, 3, 56, 28, 0, 108, 112, 3, 6, 3, 0, 109, 110, 5, 19, 0, 0, 110, 111,
		5, 24, 0, 0, 111, 113, 3, 6, 3, 0, 112, 109, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 11, 1, 0, 0, 0, 114, 115, 5, 19, 0, 0, 115, 116, 5, 25, 0, 0,
		116, 117, 3, 4, 2, 0, 117, 13, 1, 0, 0, 0, 118, 119, 7, 0, 0, 0, 119, 15,
		1, 0, 0, 0, 120, 121, 5, 19, 0, 0, 121, 122, 5, 26, 0, 0, 122, 123, 5,
		13, 0, 0, 123, 132, 5, 27, 0, 0, 124, 129, 5, 13, 0, 0, 125, 126, 5, 28,
		0, 0, 126, 128, 5, 13, 0, 0, 127, 125, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0,
		129, 127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131,
		129, 1, 0, 0, 0, 132, 124, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134,
		1, 0, 0, 0, 134, 137, 5, 29, 0, 0, 135, 136, 5, 30, 0, 0, 136, 138, 3,
		14, 7, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0,
		0, 139, 140, 3, 6, 3, 0, 140, 17, 1, 0, 0, 0, 141, 142, 7, 1, 0, 0, 142,
		19, 1, 0, 0, 0, 143, 144, 5, 13, 0, 0, 144, 21, 1, 0, 0, 0, 145, 146, 7,
		2, 0, 0, 146, 23, 1, 0, 0, 0, 147, 148, 7, 3, 0, 0, 148, 25, 1, 0, 0, 0,
		149, 150, 3, 20, 10, 0, 150, 151, 5, 31, 0, 0, 151, 189, 1, 0, 0, 0, 152,
		153, 3, 20, 10, 0, 153, 154, 5, 32, 0, 0, 154, 189, 1, 0, 0, 0, 155, 156,
		5, 9, 0, 0, 156, 157, 3, 20, 10, 0, 157, 158, 5, 31, 0, 0, 158, 189, 1,
		0, 0, 0, 159, 160, 5, 9, 0, 0, 160, 161, 3, 20, 10, 0, 161, 162, 5, 32,
		0, 0, 162, 189, 1, 0, 0, 0, 163, 164, 5, 9, 0, 0, 164, 165, 3, 20, 10,
		0, 165, 166, 5, 37, 0, 0, 166, 167, 5, 6, 0, 0, 167, 189, 1, 0, 0, 0, 168,
		169, 5, 9, 0, 0, 169, 170, 3, 20, 10, 0, 170, 171, 5, 33, 0, 0, 171, 172,
		5, 6, 0, 0, 172, 189, 1, 0, 0, 0, 173, 174, 5, 9, 0, 0, 174, 175, 3, 20,
		10, 0, 175, 176, 5, 34, 0, 0, 176, 177, 5, 6, 0, 0, 177, 189, 1, 0, 0,
		0, 178, 179, 5, 9, 0, 0, 179, 180, 3, 20, 10, 0, 180, 181, 5, 35, 0, 0,
		181, 182, 5, 6, 0, 0, 182, 189, 1, 0, 0, 0, 183, 184, 5, 9, 0, 0, 184,
		185, 3, 20, 10, 0, 185, 186, 5, 36, 0, 0, 186, 187, 5, 38, 0, 0, 187, 189,
		1, 0, 0, 0, 188, 149, 1, 0, 0, 0, 188, 152, 1, 0, 0, 0, 188, 155, 1, 0,
		0, 0, 188, 159, 1, 0, 0, 0, 188, 163, 1, 0, 0, 0, 188, 168, 1, 0, 0, 0,
		188, 173, 1, 0, 0, 0, 188, 178, 1, 0, 0, 0, 188, 183, 1, 0, 0, 0, 189,
		27, 1, 0, 0, 0, 190, 191, 7, 4, 0, 0, 191, 29, 1, 0, 0, 0, 192, 193, 5,
		2, 0, 0, 193, 194, 3, 20, 10, 0, 194, 198, 5, 17, 0, 0, 195, 197, 5, 21,
		0, 0, 196, 195, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0,
		198, 199, 1, 0, 0, 0, 199, 201, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 201,
		205, 5, 1, 0, 0, 202, 204, 5, 21, 0, 0, 203, 202, 1, 0, 0, 0, 204, 207,
		1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 208, 1, 0,
		0, 0, 207, 205, 1, 0, 0, 0, 208, 209, 5, 18, 0, 0, 209, 31, 1, 0, 0, 0,
		210, 211, 5, 9, 0, 0, 211, 212, 3, 20, 10, 0, 212, 216, 5, 17, 0, 0, 213,
		215, 5, 21, 0, 0, 214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214,
		1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 216, 1, 0,
		0, 0, 219, 223, 5, 6, 0, 0, 220, 222, 5, 21, 0, 0, 221, 220, 1, 0, 0, 0,
		222, 225, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224,
		226, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 226, 227, 5, 18, 0, 0, 227, 247,
		1, 0, 0, 0, 228, 229, 5, 9, 0, 0, 229, 230, 3, 20, 10, 0, 230, 234, 5,
		17, 0, 0, 231, 233, 5, 21, 0, 0, 232, 231, 1, 0, 0, 0, 233, 236, 1, 0,
		0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 237, 1, 0, 0, 0,
		236, 234, 1, 0, 0, 0, 237, 241, 3, 60, 30, 0, 238, 240, 5, 21, 0, 0, 239,
		238, 1, 0, 0, 0, 240, 243, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 241, 242,
		1, 0, 0, 0, 242, 244, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 244, 245, 5, 18,
		0, 0, 245, 247, 1, 0, 0, 0, 246, 210, 1, 0, 0, 0, 246, 228, 1, 0, 0, 0,
		247, 33, 1, 0, 0, 0, 248, 249, 5, 4, 0, 0, 249, 250, 3, 20, 10, 0, 250,
		254, 5, 17, 0, 0, 251, 253, 5, 21, 0, 0, 252, 251, 1, 0, 0, 0, 253, 256,
		1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 1, 0,
		0, 0, 256, 254, 1, 0, 0, 0, 257, 261, 5, 3, 0, 0, 258, 260, 5, 21, 0, 0,
		259, 258, 1, 0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261,
		262, 1, 0, 0, 0, 262, 264, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 264, 265,
		5, 18, 0, 0, 265, 303, 1, 0, 0, 0, 266, 267, 5, 4, 0, 0, 267, 268, 3, 20,
		10, 0, 268, 272, 5, 17, 0, 0, 269, 271, 5, 21, 0, 0, 270, 269, 1, 0, 0,
		0, 271, 274, 1, 0, 0, 0, 272, 270, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273,
		275, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 275, 279, 3, 54, 27, 0, 276, 278,
		5, 21, 0, 0, 277, 276, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0,
		0, 0, 279, 280, 1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0,
		282, 283, 5, 18, 0, 0, 283, 303, 1, 0, 0, 0, 284, 285, 5, 4, 0, 0, 285,
		286, 3, 20, 10, 0, 286, 290, 5, 17, 0, 0, 287, 289, 5, 21, 0, 0, 288, 287,
		1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 290, 291, 1, 0,
		0, 0, 291, 293, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 297, 3, 56, 28,
		0, 294, 296, 5, 21, 0, 0, 295, 294, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297,
		295, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 1, 0, 0, 0, 299, 297,
		1, 0, 0, 0, 300, 301, 5, 18, 0, 0, 301, 303, 1, 0, 0, 0, 302, 248, 1, 0,
		0, 0, 302, 266, 1, 0, 0, 0, 302, 284, 1, 0, 0, 0, 303, 35, 1, 0, 0, 0,
		304, 319, 3, 30, 15, 0, 305, 319, 3, 32, 16, 0, 306, 319, 3, 34, 17, 0,
		307, 308, 3, 18, 9, 0, 308, 309, 3, 20, 10, 0, 309, 313, 5, 17, 0, 0, 310,
		312, 5, 21, 0, 0, 311, 310, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311,
		1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 313, 1, 0,
		0, 0, 316, 317, 5, 18, 0, 0, 317, 319, 1, 0, 0, 0, 318, 304, 1, 0, 0, 0,
		318, 305, 1, 0, 0, 0, 318, 306, 1, 0, 0, 0, 318, 307, 1, 0, 0, 0, 319,
		37, 1, 0, 0, 0, 320, 321, 3, 18, 9, 0, 321, 322, 3, 20, 10, 0, 322, 326,
		5, 17, 0, 0, 323, 325, 5, 21, 0, 0, 324, 323, 1, 0, 0, 0, 325, 328, 1,
		0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 329, 1, 0, 0,
		0, 328, 326, 1, 0, 0, 0, 329, 340, 3, 22, 11, 0, 330, 334, 5, 28, 0, 0,
		331, 333, 5, 21, 0, 0, 332, 331, 1, 0, 0, 0, 333, 336, 1, 0, 0, 0, 334,
		332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 337, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 337, 339, 3, 22, 11, 0, 338, 330, 1, 0, 0, 0, 339, 342, 1,
		0, 0, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 346, 1, 0, 0,
		0, 342, 340, 1, 0, 0, 0, 343, 345, 5, 21, 0, 0, 344, 343, 1, 0, 0, 0, 345,
		348, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 349,
		1, 0, 0, 0, 348, 346, 1, 0, 0, 0, 349, 350, 5, 18, 0, 0, 350, 378, 1, 0,
		0, 0, 351, 352, 3, 18, 9, 0, 352, 353, 3, 20, 10, 0, 353, 357, 5, 17, 0,
		0, 354, 356, 5, 21, 0, 0, 355, 354, 1, 0, 0, 0, 356, 359, 1, 0, 0, 0, 357,
		355, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 360, 1, 0, 0, 0, 359, 357,
		1, 0, 0, 0, 360, 371, 3, 60, 30, 0, 361, 365, 5, 28, 0, 0, 362, 364, 5,
		21, 0, 0, 363, 362, 1, 0, 0, 0, 364, 367, 1, 0, 0, 0, 365, 363, 1, 0, 0,
		0, 365, 366, 1, 0, 0, 0, 366, 368, 1, 0, 0, 0, 367, 365, 1, 0, 0, 0, 368,
		370, 3, 60, 30, 0, 369, 361, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369,
		1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 374, 1, 0, 0, 0, 373, 371, 1, 0,
		0, 0, 374, 375, 5, 21, 0, 0, 375, 376, 5, 18, 0, 0, 376, 378, 1, 0, 0,
		0, 377, 320, 1, 0, 0, 0, 377, 351, 1, 0, 0, 0, 378, 39, 1, 0, 0, 0, 379,
		380, 3, 20, 10, 0, 380, 384, 5, 17, 0, 0, 381, 383, 5, 21, 0, 0, 382, 381,
		1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0,
		0, 0, 385, 387, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 391, 3, 22, 11,
		0, 388, 390, 5, 21, 0, 0, 389, 388, 1, 0, 0, 0, 390, 393, 1, 0, 0, 0, 391,
		389, 1, 0, 0, 0, 391, 392, 1, 0, 0, 0, 392, 394, 1, 0, 0, 0, 393, 391,
		1, 0, 0, 0, 394, 395, 5, 18, 0, 0, 395, 414, 1, 0, 0, 0, 396, 397, 3, 20,
		10, 0, 397, 401, 5, 17, 0, 0, 398, 400, 5, 21, 0, 0, 399, 398, 1, 0, 0,
		0, 400, 403, 1, 0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 402, 1, 0, 0, 0, 402,
		404, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 404, 408, 3, 60, 30, 0, 405, 407,
		5, 21, 0, 0, 406, 405, 1, 0, 0, 0, 407, 410, 1, 0, 0, 0, 408, 406, 1, 0,
		0, 0, 408, 409, 1, 0, 0, 0, 409, 411, 1, 0, 0, 0, 410, 408, 1, 0, 0, 0,
		411, 412, 5, 18, 0, 0, 412, 414, 1, 0, 0, 0, 413, 379, 1, 0, 0, 0, 413,
		396, 1, 0, 0, 0, 414, 41, 1, 0, 0, 0, 415, 416, 7, 5, 0, 0, 416, 43, 1,
		0, 0, 0, 417, 418, 7, 6, 0, 0, 418, 45, 1, 0, 0, 0, 419, 420, 7, 7, 0,
		0, 420, 47, 1, 0, 0, 0, 421, 422, 7, 8, 0, 0, 422, 49, 1, 0, 0, 0, 423,
		424, 5, 58, 0, 0, 424, 51, 1, 0, 0, 0, 425, 426, 6, 26, -1, 0, 426, 427,
		5, 27, 0, 0, 427, 428, 3, 52, 26, 0, 428, 429, 5, 29, 0, 0, 429, 436, 1,
		0, 0, 0, 430, 436, 5, 6, 0, 0, 431, 436, 3, 58, 29, 0, 432, 436, 3, 66,
		33, 0, 433, 434, 5, 49, 0, 0, 434, 436, 3, 52, 26, 1, 435, 425, 1, 0, 0,
		0, 435, 430, 1, 0, 0, 0, 435, 431, 1, 0, 0, 0, 435, 432, 1, 0, 0, 0, 435,
		433, 1, 0, 0, 0, 436, 447, 1, 0, 0, 0, 437, 438, 10, 6, 0, 0, 438, 439,
		3, 42, 21, 0, 439, 440, 3, 52, 26, 7, 440, 446, 1, 0, 0, 0, 441, 442, 10,
		5, 0, 0, 442, 443, 3, 44, 22, 0, 443, 444, 3, 52, 26, 6, 444, 446, 1, 0,
		0, 0, 445, 437, 1, 0, 0, 0, 445, 441, 1, 0, 0, 0, 446, 449, 1, 0, 0, 0,
		447, 445, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 53, 1, 0, 0, 0, 449, 447,
		1, 0, 0, 0, 450, 451, 6, 27, -1, 0, 451, 452, 5, 27, 0, 0, 452, 453, 3,
		54, 27, 0, 453, 454, 5, 29, 0, 0, 454, 461, 1, 0, 0, 0, 455, 456, 5, 58,
		0, 0, 456, 461, 3, 54, 27, 4, 457, 461, 5, 3, 0, 0, 458, 461, 3, 58, 29,
		0, 459, 461, 3, 66, 33, 0, 460, 450, 1, 0, 0, 0, 460, 455, 1, 0, 0, 0,
		460, 457, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 460, 459, 1, 0, 0, 0, 461,
		472, 1, 0, 0, 0, 462, 463, 10, 6, 0, 0, 463, 464, 3, 46, 23, 0, 464, 465,
		3, 54, 27, 7, 465, 471, 1, 0, 0, 0, 466, 467, 10, 5, 0, 0, 467, 468, 3,
		48, 24, 0, 468, 469, 3, 54, 27, 6, 469, 471, 1, 0, 0, 0, 470, 462, 1, 0,
		0, 0, 470, 466, 1, 0, 0, 0, 471, 474, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0,
		472, 473, 1, 0, 0, 0, 473, 55, 1, 0, 0, 0, 474, 472, 1, 0, 0, 0, 475, 476,
		6, 28, -1, 0, 476, 477, 3, 52, 26, 0, 477, 478, 3, 48, 24, 0, 478, 479,
		3, 52, 26, 0, 479, 490, 1, 0, 0, 0, 480, 481, 3, 54, 27, 0, 481, 482, 3,
		46, 23, 0, 482, 483, 3, 54, 27, 0, 483, 490, 1, 0, 0, 0, 484, 485, 3, 50,
		25, 0, 485, 486, 5, 27, 0, 0, 486, 487, 3, 56, 28, 0, 487, 488, 5, 29,
		0, 0, 488, 490, 1, 0, 0, 0, 489, 475, 1, 0, 0, 0, 489, 480, 1, 0, 0, 0,
		489, 484, 1, 0, 0, 0, 490, 497, 1, 0, 0, 0, 491, 492, 10, 2, 0, 0, 492,
		493, 3, 46, 23, 0, 493, 494, 3, 54, 27, 0, 494, 496, 1, 0, 0, 0, 495, 491,
		1, 0, 0, 0, 496, 499, 1, 0, 0, 0, 497, 495, 1, 0, 0, 0, 497, 498, 1, 0,
		0, 0, 498, 57, 1, 0, 0, 0, 499, 497, 1, 0, 0, 0, 500, 504, 5, 13, 0, 0,
		501, 504, 3, 62, 31, 0, 502, 504, 3, 64, 32, 0, 503, 500, 1, 0, 0, 0, 503,
		501, 1, 0, 0, 0, 503, 502, 1, 0, 0, 0, 504, 59, 1, 0, 0, 0, 505, 509, 3,
		58, 29, 0, 506, 509, 3, 52, 26, 0, 507, 509, 3, 54, 27, 0, 508, 505, 1,
		0, 0, 0, 508, 506, 1, 0, 0, 0, 508, 507, 1, 0, 0, 0, 509, 61, 1, 0, 0,
		0, 510, 511, 5, 13, 0, 0, 511, 512, 5, 59, 0, 0, 512, 513, 3, 52, 26, 0,
		513, 514, 5, 60, 0, 0, 514, 63, 1, 0, 0, 0, 515, 516, 5, 13, 0, 0, 516,
		517, 5, 61, 0, 0, 517, 531, 5, 13, 0, 0, 518, 519, 5, 13, 0, 0, 519, 520,
		5, 61, 0, 0, 520, 531, 3, 62, 31, 0, 521, 522, 5, 13, 0, 0, 522, 523, 5,
		61, 0, 0, 523, 531, 3, 64, 32, 0, 524, 525, 3, 62, 31, 0, 525, 526, 5,
		61, 0, 0, 526, 527, 3, 64, 32, 0, 527, 531, 1, 0, 0, 0, 528, 531, 3, 62,
		31, 0, 529, 531, 5, 13, 0, 0, 530, 515, 1, 0, 0, 0, 530, 518, 1, 0, 0,
		0, 530, 521, 1, 0, 0, 0, 530, 524, 1, 0, 0, 0, 530, 528, 1, 0, 0, 0, 530,
		529, 1, 0, 0, 0, 531, 65, 1, 0, 0, 0, 532, 533, 5, 13, 0, 0, 533, 542,
		5, 27, 0, 0, 534, 539, 3, 60, 30, 0, 535, 536, 5, 28, 0, 0, 536, 538, 3,
		60, 30, 0, 537, 535, 1, 0, 0, 0, 538, 541, 1, 0, 0, 0, 539, 537, 1, 0,
		0, 0, 539, 540, 1, 0, 0, 0, 540, 543, 1, 0, 0, 0, 541, 539, 1, 0, 0, 0,
		542, 534, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544,
		545, 5, 29, 0, 0, 545, 67, 1, 0, 0, 0, 51, 71, 82, 89, 95, 112, 129, 132,
		137, 188, 198, 205, 216, 223, 234, 241, 246, 254, 261, 272, 279, 290, 297,
		302, 313, 318, 326, 334, 340, 346, 357, 365, 371, 377, 384, 391, 401, 408,
		413, 435, 445, 447, 460, 470, 472, 489, 497, 503, 508, 530, 539, 542,
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

// MainInit initializes any static state used to implement Main. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMain(). You can call this function if you wish to initialize the static state ahead
// of time.
func MainInit() {
	staticData := &MainParserStaticData
	staticData.once.Do(mainParserInit)
}

// NewMain produces a new parser instance for the optional input antlr.TokenStream.
func NewMain(input antlr.TokenStream) *Main {
	MainInit()
	this := new(Main)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MainParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Main.g4"

	return this
}

// Main tokens.
const (
	MainEOF                 = antlr.TokenEOF
	MainTEXT                = 1
	MainTEXT_TYPE           = 2
	MainBOOLEAN             = 3
	MainBOOLEAN_TYPE        = 4
	MainDIGIT               = 5
	MainNUMBER              = 6
	MainUNIT                = 7
	MainSI_POSTFIX          = 8
	MainNUMBER_TYPE         = 9
	MainOCTET               = 10
	MainIPV4                = 11
	MainIPV4_TYPE           = 12
	MainVAR_IDENTIFIER      = 13
	MainCOMMON_IDENTIFIER   = 14
	MainARITHMETIC_OPERATOR = 15
	MainLOGICAL_OPERATOR    = 16
	MainL_BLOCK             = 17
	MainR_BLOCK             = 18
	MainAT                  = 19
	MainWS                  = 20
	MainNEWLINE             = 21
	MainLOOP                = 22
	MainIF                  = 23
	MainELSE                = 24
	MainRETURN              = 25
	MainFUNC                = 26
	MainL_PAREN             = 27
	MainCOMMA               = 28
	MainR_PAREN             = 29
	MainRETURN_TYPE         = 30
	MainOP_INCREMENT        = 31
	MainOP_DECREMENT        = 32
	MainOP_ADD_ASSIGN       = 33
	MainOP_SUB_ASSIGN       = 34
	MainOP_MUL_ASSIGN       = 35
	MainOP_DIV_ASSIGN       = 36
	MainOP_EQUAL            = 37
	MainNUMBE               = 38
	MainOP_ADD              = 39
	MainOP_SUB              = 40
	MainOP_MUL              = 41
	MainOP_DIV              = 42
	MainOP_MOD              = 43
	MainOP_POW              = 44
	MainOP_MULTIPLY         = 45
	MainOP_DIVIDE           = 46
	MainOP_POWER            = 47
	MainOP_PLUS             = 48
	MainOP_MINUS            = 49
	MainOP_AND              = 50
	MainOP_OR               = 51
	MainOP_XOR              = 52
	MainOP_NOT_EQUAL        = 53
	MainOP_LESS             = 54
	MainOP_LESS_EQUAL       = 55
	MainOP_GREATER          = 56
	MainOP_GREATER_EQUAL    = 57
	MainOP_NOT              = 58
	MainL_SQUARE            = 59
	MainR_SQUARE            = 60
	MainDOT                 = 61
)

// Main rules.
const (
	MainRULE_program                 = 0
	MainRULE_statement               = 1
	MainRULE_expression              = 2
	MainRULE_block                   = 3
	MainRULE_loop_statement          = 4
	MainRULE_if_statement            = 5
	MainRULE_return_statement        = 6
	MainRULE_data_type               = 7
	MainRULE_func_declaration        = 8
	MainRULE_var_type                = 9
	MainRULE_var_name                = 10
	MainRULE_var_value               = 11
	MainRULE_op_one_expression       = 12
	MainRULE_op_one_declaration      = 13
	MainRULE_op_two_expression       = 14
	MainRULE_string_var_declaration  = 15
	MainRULE_number_var_declaration  = 16
	MainRULE_boolean_var_declaration = 17
	MainRULE_single_var_declaration  = 18
	MainRULE_array_var_declaration   = 19
	MainRULE_var_assignment          = 20
	MainRULE_op_level1               = 21
	MainRULE_op_level2               = 22
	MainRULE_op_level3               = 23
	MainRULE_op_level4               = 24
	MainRULE_op_level5               = 25
	MainRULE_numberic_expression     = 26
	MainRULE_logical_expression      = 27
	MainRULE_comparative_expression  = 28
	MainRULE_get_variable            = 29
	MainRULE_math_expression         = 30
	MainRULE_get_array_element       = 31
	MainRULE_get_child               = 32
	MainRULE_function_call           = 33
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(MainEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *Main) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MainRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2634260) != 0 {
		{
			p.SetState(68)
			p.Statement()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
		p.Match(MainEOF)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Loop_statement() ILoop_statementContext
	If_statement() IIf_statementContext
	Func_declaration() IFunc_declarationContext
	Return_statement() IReturn_statementContext
	NEWLINE() antlr.TerminalNode

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) Loop_statement() ILoop_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoop_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoop_statementContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) Func_declaration() IFunc_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_declarationContext)
}

func (s *StatementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *StatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(MainNEWLINE, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *Main) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MainRULE_statement)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Loop_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.If_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.Func_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(80)
			p.Return_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(81)
			p.Match(MainNEWLINE)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_call() IFunction_callContext
	Var_assignment() IVar_assignmentContext
	Single_var_declaration() ISingle_var_declarationContext
	Array_var_declaration() IArray_var_declarationContext
	Get_variable() IGet_variableContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Function_call() IFunction_callContext {
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

func (s *ExpressionContext) Var_assignment() IVar_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_assignmentContext)
}

func (s *ExpressionContext) Single_var_declaration() ISingle_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingle_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingle_var_declarationContext)
}

func (s *ExpressionContext) Array_var_declaration() IArray_var_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_var_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_var_declarationContext)
}

func (s *ExpressionContext) Get_variable() IGet_variableContext {
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

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Main) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MainRULE_expression)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Function_call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Var_assignment()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Single_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.Array_var_declaration()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(88)
			p.Get_variable()
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BLOCK() antlr.TerminalNode
	R_BLOCK() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_BLOCK() antlr.TerminalNode {
	return s.GetToken(MainL_BLOCK, 0)
}

func (s *BlockContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(MainR_BLOCK, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *Main) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MainRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(MainL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2634260) != 0 {
		{
			p.SetState(92)
			p.Statement()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
		p.Match(MainR_BLOCK)
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

// ILoop_statementContext is an interface to support dynamic dispatch.
type ILoop_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	LOOP() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	Block() IBlockContext

	// IsLoop_statementContext differentiates from other interfaces.
	IsLoop_statementContext()
}

type Loop_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoop_statementContext() *Loop_statementContext {
	var p = new(Loop_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_loop_statement
	return p
}

func InitEmptyLoop_statementContext(p *Loop_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_loop_statement
}

func (*Loop_statementContext) IsLoop_statementContext() {}

func NewLoop_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loop_statementContext {
	var p = new(Loop_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_loop_statement

	return p
}

func (s *Loop_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Loop_statementContext) AT() antlr.TerminalNode {
	return s.GetToken(MainAT, 0)
}

func (s *Loop_statementContext) LOOP() antlr.TerminalNode {
	return s.GetToken(MainLOOP, 0)
}

func (s *Loop_statementContext) Comparative_expression() IComparative_expressionContext {
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

func (s *Loop_statementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Loop_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loop_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loop_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterLoop_statement(s)
	}
}

func (s *Loop_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitLoop_statement(s)
	}
}

func (p *Main) Loop_statement() (localctx ILoop_statementContext) {
	localctx = NewLoop_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MainRULE_loop_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(MainAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.Match(MainLOOP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.comparative_expression(0)
	}
	{
		p.SetState(103)
		p.Block()
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

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAT() []antlr.TerminalNode
	AT(i int) antlr.TerminalNode
	IF() antlr.TerminalNode
	Comparative_expression() IComparative_expressionContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	ELSE() antlr.TerminalNode

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) AllAT() []antlr.TerminalNode {
	return s.GetTokens(MainAT)
}

func (s *If_statementContext) AT(i int) antlr.TerminalNode {
	return s.GetToken(MainAT, i)
}

func (s *If_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(MainIF, 0)
}

func (s *If_statementContext) Comparative_expression() IComparative_expressionContext {
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

func (s *If_statementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *If_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(MainELSE, 0)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (p *Main) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MainRULE_if_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(MainAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(MainIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.comparative_expression(0)
	}
	{
		p.SetState(108)
		p.Block()
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(109)
			p.Match(MainAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(MainELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Block()
		}

	} else if p.HasError() { // JIM
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

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) AT() antlr.TerminalNode {
	return s.GetToken(MainAT, 0)
}

func (s *Return_statementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(MainRETURN, 0)
}

func (s *Return_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterReturn_statement(s)
	}
}

func (s *Return_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitReturn_statement(s)
	}
}

func (p *Main) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MainRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(MainAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(MainRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Expression()
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

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TEXT_TYPE() antlr.TerminalNode
	NUMBER_TYPE() antlr.TerminalNode
	BOOLEAN_TYPE() antlr.TerminalNode
	COMMON_IDENTIFIER() antlr.TerminalNode

	// IsData_typeContext differentiates from other interfaces.
	IsData_typeContext()
}

type Data_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_data_type
	return p
}

func InitEmptyData_typeContext(p *Data_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_data_type
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(MainTEXT_TYPE, 0)
}

func (s *Data_typeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(MainNUMBER_TYPE, 0)
}

func (s *Data_typeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(MainBOOLEAN_TYPE, 0)
}

func (s *Data_typeContext) COMMON_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MainCOMMON_IDENTIFIER, 0)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterData_type(s)
	}
}

func (s *Data_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitData_type(s)
	}
}

func (p *Main) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MainRULE_data_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16916) != 0) {
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

// IFunc_declarationContext is an interface to support dynamic dispatch.
type IFunc_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AT() antlr.TerminalNode
	FUNC() antlr.TerminalNode
	AllVAR_IDENTIFIER() []antlr.TerminalNode
	VAR_IDENTIFIER(i int) antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	Block() IBlockContext
	RETURN_TYPE() antlr.TerminalNode
	Data_type() IData_typeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunc_declarationContext differentiates from other interfaces.
	IsFunc_declarationContext()
}

type Func_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_declarationContext() *Func_declarationContext {
	var p = new(Func_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_func_declaration
	return p
}

func InitEmptyFunc_declarationContext(p *Func_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_func_declaration
}

func (*Func_declarationContext) IsFunc_declarationContext() {}

func NewFunc_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_declarationContext {
	var p = new(Func_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_func_declaration

	return p
}

func (s *Func_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_declarationContext) AT() antlr.TerminalNode {
	return s.GetToken(MainAT, 0)
}

func (s *Func_declarationContext) FUNC() antlr.TerminalNode {
	return s.GetToken(MainFUNC, 0)
}

func (s *Func_declarationContext) AllVAR_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MainVAR_IDENTIFIER)
}

func (s *Func_declarationContext) VAR_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MainVAR_IDENTIFIER, i)
}

func (s *Func_declarationContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MainL_PAREN, 0)
}

func (s *Func_declarationContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MainR_PAREN, 0)
}

func (s *Func_declarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Func_declarationContext) RETURN_TYPE() antlr.TerminalNode {
	return s.GetToken(MainRETURN_TYPE, 0)
}

func (s *Func_declarationContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Func_declarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MainCOMMA)
}

func (s *Func_declarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MainCOMMA, i)
}

func (s *Func_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Func_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterFunc_declaration(s)
	}
}

func (s *Func_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitFunc_declaration(s)
	}
}

func (p *Main) Func_declaration() (localctx IFunc_declarationContext) {
	localctx = NewFunc_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MainRULE_func_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(MainAT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Match(MainFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(MainVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Match(MainL_PAREN)
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

	if _la == MainVAR_IDENTIFIER {
		{
			p.SetState(124)
			p.Match(MainVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainCOMMA {
			{
				p.SetState(125)
				p.Match(MainCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(126)
				p.Match(MainVAR_IDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(134)
		p.Match(MainR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MainRETURN_TYPE {
		{
			p.SetState(135)
			p.Match(MainRETURN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Data_type()
		}

	}
	{
		p.SetState(139)
		p.Block()
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
	p.RuleIndex = MainRULE_var_type
	return p
}

func InitEmptyVar_typeContext(p *Var_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_var_type
}

func (*Var_typeContext) IsVar_typeContext() {}

func NewVar_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_typeContext {
	var p = new(Var_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_var_type

	return p
}

func (s *Var_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_typeContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(MainTEXT_TYPE, 0)
}

func (s *Var_typeContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(MainNUMBER_TYPE, 0)
}

func (s *Var_typeContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(MainBOOLEAN_TYPE, 0)
}

func (s *Var_typeContext) IPV4_TYPE() antlr.TerminalNode {
	return s.GetToken(MainIPV4_TYPE, 0)
}

func (s *Var_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterVar_type(s)
	}
}

func (s *Var_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitVar_type(s)
	}
}

func (p *Main) Var_type() (localctx IVar_typeContext) {
	localctx = NewVar_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MainRULE_var_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
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
	p.RuleIndex = MainRULE_var_name
	return p
}

func InitEmptyVar_nameContext(p *Var_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_var_name
}

func (*Var_nameContext) IsVar_nameContext() {}

func NewVar_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_nameContext {
	var p = new(Var_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_var_name

	return p
}

func (s *Var_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_nameContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MainVAR_IDENTIFIER, 0)
}

func (s *Var_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterVar_name(s)
	}
}

func (s *Var_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitVar_name(s)
	}
}

func (p *Main) Var_name() (localctx IVar_nameContext) {
	localctx = NewVar_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MainRULE_var_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(MainVAR_IDENTIFIER)
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
	p.RuleIndex = MainRULE_var_value
	return p
}

func InitEmptyVar_valueContext(p *Var_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_var_value
}

func (*Var_valueContext) IsVar_valueContext() {}

func NewVar_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_valueContext {
	var p = new(Var_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_var_value

	return p
}

func (s *Var_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_valueContext) IPV4() antlr.TerminalNode {
	return s.GetToken(MainIPV4, 0)
}

func (s *Var_valueContext) TEXT() antlr.TerminalNode {
	return s.GetToken(MainTEXT, 0)
}

func (s *Var_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MainNUMBER, 0)
}

func (s *Var_valueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(MainBOOLEAN, 0)
}

func (s *Var_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterVar_value(s)
	}
}

func (s *Var_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitVar_value(s)
	}
}

func (p *Main) Var_value() (localctx IVar_valueContext) {
	localctx = NewVar_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MainRULE_var_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
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
	p.RuleIndex = MainRULE_op_one_expression
	return p
}

func InitEmptyOp_one_expressionContext(p *Op_one_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_op_one_expression
}

func (*Op_one_expressionContext) IsOp_one_expressionContext() {}

func NewOp_one_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_one_expressionContext {
	var p = new(Op_one_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_op_one_expression

	return p
}

func (s *Op_one_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_one_expressionContext) OP_INCREMENT() antlr.TerminalNode {
	return s.GetToken(MainOP_INCREMENT, 0)
}

func (s *Op_one_expressionContext) OP_DECREMENT() antlr.TerminalNode {
	return s.GetToken(MainOP_DECREMENT, 0)
}

func (s *Op_one_expressionContext) OP_ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MainOP_ADD_ASSIGN, 0)
}

func (s *Op_one_expressionContext) OP_SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MainOP_SUB_ASSIGN, 0)
}

func (s *Op_one_expressionContext) OP_MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MainOP_MUL_ASSIGN, 0)
}

func (s *Op_one_expressionContext) OP_DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MainOP_DIV_ASSIGN, 0)
}

func (s *Op_one_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_one_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_one_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterOp_one_expression(s)
	}
}

func (s *Op_one_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitOp_one_expression(s)
	}
}

func (p *Main) Op_one_expression() (localctx IOp_one_expressionContext) {
	localctx = NewOp_one_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MainRULE_op_one_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135291469824) != 0) {
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
	p.RuleIndex = MainRULE_op_one_declaration
	return p
}

func InitEmptyOp_one_declarationContext(p *Op_one_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_op_one_declaration
}

func (*Op_one_declarationContext) IsOp_one_declarationContext() {}

func NewOp_one_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_one_declarationContext {
	var p = new(Op_one_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_op_one_declaration

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
	return s.GetToken(MainOP_INCREMENT, 0)
}

func (s *Op_one_declarationContext) OP_DECREMENT() antlr.TerminalNode {
	return s.GetToken(MainOP_DECREMENT, 0)
}

func (s *Op_one_declarationContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(MainNUMBER_TYPE, 0)
}

func (s *Op_one_declarationContext) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(MainOP_EQUAL, 0)
}

func (s *Op_one_declarationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MainNUMBER, 0)
}

func (s *Op_one_declarationContext) OP_ADD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MainOP_ADD_ASSIGN, 0)
}

func (s *Op_one_declarationContext) OP_SUB_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MainOP_SUB_ASSIGN, 0)
}

func (s *Op_one_declarationContext) OP_MUL_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MainOP_MUL_ASSIGN, 0)
}

func (s *Op_one_declarationContext) OP_DIV_ASSIGN() antlr.TerminalNode {
	return s.GetToken(MainOP_DIV_ASSIGN, 0)
}

func (s *Op_one_declarationContext) NUMBE() antlr.TerminalNode {
	return s.GetToken(MainNUMBE, 0)
}

func (s *Op_one_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_one_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_one_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterOp_one_declaration(s)
	}
}

func (s *Op_one_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitOp_one_declaration(s)
	}
}

func (p *Main) Op_one_declaration() (localctx IOp_one_declarationContext) {
	localctx = NewOp_one_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MainRULE_op_one_declaration)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Var_name()
		}
		{
			p.SetState(150)
			p.Match(MainOP_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Var_name()
		}
		{
			p.SetState(153)
			p.Match(MainOP_DECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Var_name()
		}
		{
			p.SetState(157)
			p.Match(MainOP_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(159)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Var_name()
		}
		{
			p.SetState(161)
			p.Match(MainOP_DECREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(163)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Var_name()
		}
		{
			p.SetState(165)
			p.Match(MainOP_EQUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Match(MainNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(168)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Var_name()
		}
		{
			p.SetState(170)
			p.Match(MainOP_ADD_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(MainNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(173)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Var_name()
		}
		{
			p.SetState(175)
			p.Match(MainOP_SUB_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(MainNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(178)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Var_name()
		}
		{
			p.SetState(180)
			p.Match(MainOP_MUL_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(MainNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(183)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Var_name()
		}
		{
			p.SetState(185)
			p.Match(MainOP_DIV_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(MainNUMBE)
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
	p.RuleIndex = MainRULE_op_two_expression
	return p
}

func InitEmptyOp_two_expressionContext(p *Op_two_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_op_two_expression
}

func (*Op_two_expressionContext) IsOp_two_expressionContext() {}

func NewOp_two_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_two_expressionContext {
	var p = new(Op_two_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_op_two_expression

	return p
}

func (s *Op_two_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Op_two_expressionContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(MainOP_ADD, 0)
}

func (s *Op_two_expressionContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(MainOP_SUB, 0)
}

func (s *Op_two_expressionContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(MainOP_MUL, 0)
}

func (s *Op_two_expressionContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(MainOP_DIV, 0)
}

func (s *Op_two_expressionContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(MainOP_MOD, 0)
}

func (s *Op_two_expressionContext) OP_POW() antlr.TerminalNode {
	return s.GetToken(MainOP_POW, 0)
}

func (s *Op_two_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_two_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_two_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterOp_two_expression(s)
	}
}

func (s *Op_two_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitOp_two_expression(s)
	}
}

func (p *Main) Op_two_expression() (localctx IOp_two_expressionContext) {
	localctx = NewOp_two_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MainRULE_op_two_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34634616274944) != 0) {
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
	p.RuleIndex = MainRULE_string_var_declaration
	return p
}

func InitEmptyString_var_declarationContext(p *String_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_string_var_declaration
}

func (*String_var_declarationContext) IsString_var_declarationContext() {}

func NewString_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_var_declarationContext {
	var p = new(String_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_string_var_declaration

	return p
}

func (s *String_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *String_var_declarationContext) TEXT_TYPE() antlr.TerminalNode {
	return s.GetToken(MainTEXT_TYPE, 0)
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
	return s.GetToken(MainL_BLOCK, 0)
}

func (s *String_var_declarationContext) TEXT() antlr.TerminalNode {
	return s.GetToken(MainTEXT, 0)
}

func (s *String_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(MainR_BLOCK, 0)
}

func (s *String_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MainNEWLINE)
}

func (s *String_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MainNEWLINE, i)
}

func (s *String_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterString_var_declaration(s)
	}
}

func (s *String_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitString_var_declaration(s)
	}
}

func (p *Main) String_var_declaration() (localctx IString_var_declarationContext) {
	localctx = NewString_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MainRULE_string_var_declaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(MainTEXT_TYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Var_name()
	}
	{
		p.SetState(194)
		p.Match(MainL_BLOCK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MainNEWLINE {
		{
			p.SetState(195)
			p.Match(MainNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)
		p.Match(MainTEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MainNEWLINE {
		{
			p.SetState(202)
			p.Match(MainNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(208)
		p.Match(MainR_BLOCK)
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
	p.RuleIndex = MainRULE_number_var_declaration
	return p
}

func InitEmptyNumber_var_declarationContext(p *Number_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_number_var_declaration
}

func (*Number_var_declarationContext) IsNumber_var_declarationContext() {}

func NewNumber_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Number_var_declarationContext {
	var p = new(Number_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_number_var_declaration

	return p
}

func (s *Number_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Number_var_declarationContext) NUMBER_TYPE() antlr.TerminalNode {
	return s.GetToken(MainNUMBER_TYPE, 0)
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
	return s.GetToken(MainL_BLOCK, 0)
}

func (s *Number_var_declarationContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MainNUMBER, 0)
}

func (s *Number_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(MainR_BLOCK, 0)
}

func (s *Number_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MainNEWLINE)
}

func (s *Number_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MainNEWLINE, i)
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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterNumber_var_declaration(s)
	}
}

func (s *Number_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitNumber_var_declaration(s)
	}
}

func (p *Main) Number_var_declaration() (localctx INumber_var_declarationContext) {
	localctx = NewNumber_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MainRULE_number_var_declaration)
	var _la int

	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Var_name()
		}
		{
			p.SetState(212)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(213)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(219)
			p.Match(MainNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(220)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(225)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(226)
			p.Match(MainR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Match(MainNUMBER_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Var_name()
		}
		{
			p.SetState(230)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(231)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(236)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(237)
			p.Math_expression()
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(238)
				p.Match(MainNEWLINE)
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
		}
		{
			p.SetState(244)
			p.Match(MainR_BLOCK)
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
	p.RuleIndex = MainRULE_boolean_var_declaration
	return p
}

func InitEmptyBoolean_var_declarationContext(p *Boolean_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_boolean_var_declaration
}

func (*Boolean_var_declarationContext) IsBoolean_var_declarationContext() {}

func NewBoolean_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_var_declarationContext {
	var p = new(Boolean_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_boolean_var_declaration

	return p
}

func (s *Boolean_var_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_var_declarationContext) BOOLEAN_TYPE() antlr.TerminalNode {
	return s.GetToken(MainBOOLEAN_TYPE, 0)
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
	return s.GetToken(MainL_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(MainBOOLEAN, 0)
}

func (s *Boolean_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(MainR_BLOCK, 0)
}

func (s *Boolean_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MainNEWLINE)
}

func (s *Boolean_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MainNEWLINE, i)
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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterBoolean_var_declaration(s)
	}
}

func (s *Boolean_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitBoolean_var_declaration(s)
	}
}

func (p *Main) Boolean_var_declaration() (localctx IBoolean_var_declarationContext) {
	localctx = NewBoolean_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, MainRULE_boolean_var_declaration)
	var _la int

	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Match(MainBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Var_name()
		}
		{
			p.SetState(250)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(251)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(257)
			p.Match(MainBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(258)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(264)
			p.Match(MainR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Match(MainBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Var_name()
		}
		{
			p.SetState(268)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(269)
				p.Match(MainNEWLINE)
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
		}
		{
			p.SetState(275)
			p.logical_expression(0)
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(276)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(281)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(282)
			p.Match(MainR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(284)
			p.Match(MainBOOLEAN_TYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Var_name()
		}
		{
			p.SetState(286)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(287)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(293)
			p.comparative_expression(0)
		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(294)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(300)
			p.Match(MainR_BLOCK)
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
	p.RuleIndex = MainRULE_single_var_declaration
	return p
}

func InitEmptySingle_var_declarationContext(p *Single_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_single_var_declaration
}

func (*Single_var_declarationContext) IsSingle_var_declarationContext() {}

func NewSingle_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_var_declarationContext {
	var p = new(Single_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_single_var_declaration

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
	return s.GetToken(MainL_BLOCK, 0)
}

func (s *Single_var_declarationContext) R_BLOCK() antlr.TerminalNode {
	return s.GetToken(MainR_BLOCK, 0)
}

func (s *Single_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MainNEWLINE)
}

func (s *Single_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MainNEWLINE, i)
}

func (s *Single_var_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_var_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_var_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterSingle_var_declaration(s)
	}
}

func (s *Single_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitSingle_var_declaration(s)
	}
}

func (p *Main) Single_var_declaration() (localctx ISingle_var_declarationContext) {
	localctx = NewSingle_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MainRULE_single_var_declaration)
	var _la int

	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.String_var_declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Number_var_declaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(306)
			p.Boolean_var_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(307)
			p.Var_type()
		}
		{
			p.SetState(308)
			p.Var_name()
		}
		{
			p.SetState(309)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(313)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(310)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(316)
			p.Match(MainR_BLOCK)
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
	p.RuleIndex = MainRULE_array_var_declaration
	return p
}

func InitEmptyArray_var_declarationContext(p *Array_var_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_array_var_declaration
}

func (*Array_var_declarationContext) IsArray_var_declarationContext() {}

func NewArray_var_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_var_declarationContext {
	var p = new(Array_var_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_array_var_declaration

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
	return s.GetToken(MainL_BLOCK, 0)
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
	return s.GetToken(MainR_BLOCK, 0)
}

func (s *Array_var_declarationContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MainNEWLINE)
}

func (s *Array_var_declarationContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MainNEWLINE, i)
}

func (s *Array_var_declarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MainCOMMA)
}

func (s *Array_var_declarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MainCOMMA, i)
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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterArray_var_declaration(s)
	}
}

func (s *Array_var_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitArray_var_declaration(s)
	}
}

func (p *Main) Array_var_declaration() (localctx IArray_var_declarationContext) {
	localctx = NewArray_var_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MainRULE_array_var_declaration)
	var _la int

	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.Var_type()
		}
		{
			p.SetState(321)
			p.Var_name()
		}
		{
			p.SetState(322)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(323)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(328)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(329)
			p.Var_value()
		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainCOMMA {
			{
				p.SetState(330)
				p.Match(MainCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(334)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MainNEWLINE {
				{
					p.SetState(331)
					p.Match(MainNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(336)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(337)
				p.Var_value()
			}

			p.SetState(342)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(343)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(348)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(349)
			p.Match(MainR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(351)
			p.Var_type()
		}
		{
			p.SetState(352)
			p.Var_name()
		}
		{
			p.SetState(353)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(354)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(359)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(360)
			p.Math_expression()
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainCOMMA {
			{
				p.SetState(361)
				p.Match(MainCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(365)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == MainNEWLINE {
				{
					p.SetState(362)
					p.Match(MainNEWLINE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(367)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(368)
				p.Math_expression()
			}

			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(374)
			p.Match(MainNEWLINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Match(MainR_BLOCK)
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
	p.RuleIndex = MainRULE_var_assignment
	return p
}

func InitEmptyVar_assignmentContext(p *Var_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_var_assignment
}

func (*Var_assignmentContext) IsVar_assignmentContext() {}

func NewVar_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_assignmentContext {
	var p = new(Var_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_var_assignment

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
	return s.GetToken(MainL_BLOCK, 0)
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
	return s.GetToken(MainR_BLOCK, 0)
}

func (s *Var_assignmentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(MainNEWLINE)
}

func (s *Var_assignmentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(MainNEWLINE, i)
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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterVar_assignment(s)
	}
}

func (s *Var_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitVar_assignment(s)
	}
}

func (p *Main) Var_assignment() (localctx IVar_assignmentContext) {
	localctx = NewVar_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MainRULE_var_assignment)
	var _la int

	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.Var_name()
		}
		{
			p.SetState(380)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(381)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(386)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(387)
			p.Var_value()
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(388)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(393)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(394)
			p.Match(MainR_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.Var_name()
		}
		{
			p.SetState(397)
			p.Match(MainL_BLOCK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(398)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(403)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(404)
			p.Math_expression()
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainNEWLINE {
			{
				p.SetState(405)
				p.Match(MainNEWLINE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(410)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(411)
			p.Match(MainR_BLOCK)
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
	p.RuleIndex = MainRULE_op_level1
	return p
}

func InitEmptyOp_level1Context(p *Op_level1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_op_level1
}

func (*Op_level1Context) IsOp_level1Context() {}

func NewOp_level1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level1Context {
	var p = new(Op_level1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_op_level1

	return p
}

func (s *Op_level1Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level1Context) OP_MULTIPLY() antlr.TerminalNode {
	return s.GetToken(MainOP_MULTIPLY, 0)
}

func (s *Op_level1Context) OP_DIVIDE() antlr.TerminalNode {
	return s.GetToken(MainOP_DIVIDE, 0)
}

func (s *Op_level1Context) OP_MOD() antlr.TerminalNode {
	return s.GetToken(MainOP_MOD, 0)
}

func (s *Op_level1Context) OP_POWER() antlr.TerminalNode {
	return s.GetToken(MainOP_POWER, 0)
}

func (s *Op_level1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterOp_level1(s)
	}
}

func (s *Op_level1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitOp_level1(s)
	}
}

func (p *Main) Op_level1() (localctx IOp_level1Context) {
	localctx = NewOp_level1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, MainRULE_op_level1)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(415)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&255086697644032) != 0) {
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
	p.RuleIndex = MainRULE_op_level2
	return p
}

func InitEmptyOp_level2Context(p *Op_level2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_op_level2
}

func (*Op_level2Context) IsOp_level2Context() {}

func NewOp_level2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level2Context {
	var p = new(Op_level2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_op_level2

	return p
}

func (s *Op_level2Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level2Context) OP_PLUS() antlr.TerminalNode {
	return s.GetToken(MainOP_PLUS, 0)
}

func (s *Op_level2Context) OP_MINUS() antlr.TerminalNode {
	return s.GetToken(MainOP_MINUS, 0)
}

func (s *Op_level2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterOp_level2(s)
	}
}

func (s *Op_level2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitOp_level2(s)
	}
}

func (p *Main) Op_level2() (localctx IOp_level2Context) {
	localctx = NewOp_level2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, MainRULE_op_level2)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)
		_la = p.GetTokenStream().LA(1)

		if !(_la == MainOP_PLUS || _la == MainOP_MINUS) {
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
	p.RuleIndex = MainRULE_op_level3
	return p
}

func InitEmptyOp_level3Context(p *Op_level3Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_op_level3
}

func (*Op_level3Context) IsOp_level3Context() {}

func NewOp_level3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level3Context {
	var p = new(Op_level3Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_op_level3

	return p
}

func (s *Op_level3Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level3Context) OP_AND() antlr.TerminalNode {
	return s.GetToken(MainOP_AND, 0)
}

func (s *Op_level3Context) OP_OR() antlr.TerminalNode {
	return s.GetToken(MainOP_OR, 0)
}

func (s *Op_level3Context) OP_XOR() antlr.TerminalNode {
	return s.GetToken(MainOP_XOR, 0)
}

func (s *Op_level3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterOp_level3(s)
	}
}

func (s *Op_level3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitOp_level3(s)
	}
}

func (p *Main) Op_level3() (localctx IOp_level3Context) {
	localctx = NewOp_level3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, MainRULE_op_level3)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7881299347898368) != 0) {
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
	p.RuleIndex = MainRULE_op_level4
	return p
}

func InitEmptyOp_level4Context(p *Op_level4Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_op_level4
}

func (*Op_level4Context) IsOp_level4Context() {}

func NewOp_level4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level4Context {
	var p = new(Op_level4Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_op_level4

	return p
}

func (s *Op_level4Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level4Context) OP_EQUAL() antlr.TerminalNode {
	return s.GetToken(MainOP_EQUAL, 0)
}

func (s *Op_level4Context) OP_NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(MainOP_NOT_EQUAL, 0)
}

func (s *Op_level4Context) OP_LESS() antlr.TerminalNode {
	return s.GetToken(MainOP_LESS, 0)
}

func (s *Op_level4Context) OP_LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(MainOP_LESS_EQUAL, 0)
}

func (s *Op_level4Context) OP_GREATER() antlr.TerminalNode {
	return s.GetToken(MainOP_GREATER, 0)
}

func (s *Op_level4Context) OP_GREATER_EQUAL() antlr.TerminalNode {
	return s.GetToken(MainOP_GREATER_EQUAL, 0)
}

func (s *Op_level4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterOp_level4(s)
	}
}

func (s *Op_level4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitOp_level4(s)
	}
}

func (p *Main) Op_level4() (localctx IOp_level4Context) {
	localctx = NewOp_level4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, MainRULE_op_level4)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&279223314335924224) != 0) {
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
	p.RuleIndex = MainRULE_op_level5
	return p
}

func InitEmptyOp_level5Context(p *Op_level5Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_op_level5
}

func (*Op_level5Context) IsOp_level5Context() {}

func NewOp_level5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Op_level5Context {
	var p = new(Op_level5Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_op_level5

	return p
}

func (s *Op_level5Context) GetParser() antlr.Parser { return s.parser }

func (s *Op_level5Context) OP_NOT() antlr.TerminalNode {
	return s.GetToken(MainOP_NOT, 0)
}

func (s *Op_level5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_level5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Op_level5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterOp_level5(s)
	}
}

func (s *Op_level5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitOp_level5(s)
	}
}

func (p *Main) Op_level5() (localctx IOp_level5Context) {
	localctx = NewOp_level5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, MainRULE_op_level5)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.Match(MainOP_NOT)
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
	p.RuleIndex = MainRULE_numberic_expression
	return p
}

func InitEmptyNumberic_expressionContext(p *Numberic_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_numberic_expression
}

func (*Numberic_expressionContext) IsNumberic_expressionContext() {}

func NewNumberic_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numberic_expressionContext {
	var p = new(Numberic_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_numberic_expression

	return p
}

func (s *Numberic_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Numberic_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MainL_PAREN, 0)
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
	return s.GetToken(MainR_PAREN, 0)
}

func (s *Numberic_expressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(MainNUMBER, 0)
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
	return s.GetToken(MainOP_MINUS, 0)
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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterNumberic_expression(s)
	}
}

func (s *Numberic_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitNumberic_expression(s)
	}
}

func (p *Main) Numberic_expression() (localctx INumberic_expressionContext) {
	return p.numberic_expression(0)
}

func (p *Main) numberic_expression(_p int) (localctx INumberic_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewNumberic_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumberic_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, MainRULE_numberic_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(426)
			p.Match(MainL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.numberic_expression(0)
		}
		{
			p.SetState(428)
			p.Match(MainR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(430)
			p.Match(MainNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(431)
			p.Get_variable()
		}

	case 4:
		{
			p.SetState(432)
			p.Function_call()
		}

	case 5:
		{
			p.SetState(433)
			p.Match(MainOP_MINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.numberic_expression(1)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(445)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
			case 1:
				localctx = NewNumberic_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MainRULE_numberic_expression)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(438)
					p.Op_level1()
				}
				{
					p.SetState(439)
					p.numberic_expression(7)
				}

			case 2:
				localctx = NewNumberic_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MainRULE_numberic_expression)
				p.SetState(441)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(442)
					p.Op_level2()
				}
				{
					p.SetState(443)
					p.numberic_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
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
	p.RuleIndex = MainRULE_logical_expression
	return p
}

func InitEmptyLogical_expressionContext(p *Logical_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_logical_expression
}

func (*Logical_expressionContext) IsLogical_expressionContext() {}

func NewLogical_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_expressionContext {
	var p = new(Logical_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_logical_expression

	return p
}

func (s *Logical_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MainL_PAREN, 0)
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
	return s.GetToken(MainR_PAREN, 0)
}

func (s *Logical_expressionContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(MainOP_NOT, 0)
}

func (s *Logical_expressionContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(MainBOOLEAN, 0)
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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterLogical_expression(s)
	}
}

func (s *Logical_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitLogical_expression(s)
	}
}

func (p *Main) Logical_expression() (localctx ILogical_expressionContext) {
	return p.logical_expression(0)
}

func (p *Main) logical_expression(_p int) (localctx ILogical_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogical_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogical_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, MainRULE_logical_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(451)
			p.Match(MainL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.logical_expression(0)
		}
		{
			p.SetState(453)
			p.Match(MainR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(455)
			p.Match(MainOP_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)
			p.logical_expression(4)
		}

	case 3:
		{
			p.SetState(457)
			p.Match(MainBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(458)
			p.Get_variable()
		}

	case 5:
		{
			p.SetState(459)
			p.Function_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(470)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MainRULE_logical_expression)
				p.SetState(462)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(463)
					p.Op_level3()
				}
				{
					p.SetState(464)
					p.logical_expression(7)
				}

			case 2:
				localctx = NewLogical_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, MainRULE_logical_expression)
				p.SetState(466)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(467)
					p.Op_level4()
				}
				{
					p.SetState(468)
					p.logical_expression(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
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
	p.RuleIndex = MainRULE_comparative_expression
	return p
}

func InitEmptyComparative_expressionContext(p *Comparative_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_comparative_expression
}

func (*Comparative_expressionContext) IsComparative_expressionContext() {}

func NewComparative_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparative_expressionContext {
	var p = new(Comparative_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_comparative_expression

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
	return s.GetToken(MainL_PAREN, 0)
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
	return s.GetToken(MainR_PAREN, 0)
}

func (s *Comparative_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparative_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparative_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterComparative_expression(s)
	}
}

func (s *Comparative_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitComparative_expression(s)
	}
}

func (p *Main) Comparative_expression() (localctx IComparative_expressionContext) {
	return p.comparative_expression(0)
}

func (p *Main) comparative_expression(_p int) (localctx IComparative_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewComparative_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComparative_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, MainRULE_comparative_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(476)
			p.numberic_expression(0)
		}
		{
			p.SetState(477)
			p.Op_level4()
		}
		{
			p.SetState(478)
			p.numberic_expression(0)
		}

	case 2:
		{
			p.SetState(480)
			p.logical_expression(0)
		}
		{
			p.SetState(481)
			p.Op_level3()
		}
		{
			p.SetState(482)
			p.logical_expression(0)
		}

	case 3:
		{
			p.SetState(484)
			p.Op_level5()
		}
		{
			p.SetState(485)
			p.Match(MainL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(486)
			p.comparative_expression(0)
		}
		{
			p.SetState(487)
			p.Match(MainR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
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
			p.PushNewRecursionContext(localctx, _startState, MainRULE_comparative_expression)
			p.SetState(491)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(492)
				p.Op_level3()
			}
			{
				p.SetState(493)
				p.logical_expression(0)
			}

		}
		p.SetState(499)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
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
	p.RuleIndex = MainRULE_get_variable
	return p
}

func InitEmptyGet_variableContext(p *Get_variableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_get_variable
}

func (*Get_variableContext) IsGet_variableContext() {}

func NewGet_variableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_variableContext {
	var p = new(Get_variableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_get_variable

	return p
}

func (s *Get_variableContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_variableContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MainVAR_IDENTIFIER, 0)
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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterGet_variable(s)
	}
}

func (s *Get_variableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitGet_variable(s)
	}
}

func (p *Main) Get_variable() (localctx IGet_variableContext) {
	localctx = NewGet_variableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, MainRULE_get_variable)
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(500)
			p.Match(MainVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(501)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(502)
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
	p.RuleIndex = MainRULE_math_expression
	return p
}

func InitEmptyMath_expressionContext(p *Math_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_math_expression
}

func (*Math_expressionContext) IsMath_expressionContext() {}

func NewMath_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Math_expressionContext {
	var p = new(Math_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_math_expression

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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterMath_expression(s)
	}
}

func (s *Math_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitMath_expression(s)
	}
}

func (p *Main) Math_expression() (localctx IMath_expressionContext) {
	localctx = NewMath_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, MainRULE_math_expression)
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(505)
			p.Get_variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(506)
			p.numberic_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(507)
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
	p.RuleIndex = MainRULE_get_array_element
	return p
}

func InitEmptyGet_array_elementContext(p *Get_array_elementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_get_array_element
}

func (*Get_array_elementContext) IsGet_array_elementContext() {}

func NewGet_array_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_array_elementContext {
	var p = new(Get_array_elementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_get_array_element

	return p
}

func (s *Get_array_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_array_elementContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MainVAR_IDENTIFIER, 0)
}

func (s *Get_array_elementContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(MainL_SQUARE, 0)
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
	return s.GetToken(MainR_SQUARE, 0)
}

func (s *Get_array_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Get_array_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Get_array_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterGet_array_element(s)
	}
}

func (s *Get_array_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitGet_array_element(s)
	}
}

func (p *Main) Get_array_element() (localctx IGet_array_elementContext) {
	localctx = NewGet_array_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, MainRULE_get_array_element)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.Match(MainVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(511)
		p.Match(MainL_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(512)
		p.numberic_expression(0)
	}
	{
		p.SetState(513)
		p.Match(MainR_SQUARE)
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
	p.RuleIndex = MainRULE_get_child
	return p
}

func InitEmptyGet_childContext(p *Get_childContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_get_child
}

func (*Get_childContext) IsGet_childContext() {}

func NewGet_childContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Get_childContext {
	var p = new(Get_childContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_get_child

	return p
}

func (s *Get_childContext) GetParser() antlr.Parser { return s.parser }

func (s *Get_childContext) AllVAR_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MainVAR_IDENTIFIER)
}

func (s *Get_childContext) VAR_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MainVAR_IDENTIFIER, i)
}

func (s *Get_childContext) DOT() antlr.TerminalNode {
	return s.GetToken(MainDOT, 0)
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
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterGet_child(s)
	}
}

func (s *Get_childContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitGet_child(s)
	}
}

func (p *Main) Get_child() (localctx IGet_childContext) {
	localctx = NewGet_childContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, MainRULE_get_child)
	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(515)
			p.Match(MainVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(516)
			p.Match(MainDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(517)
			p.Match(MainVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(518)
			p.Match(MainVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(519)
			p.Match(MainDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)
			p.Get_array_element()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(521)
			p.Match(MainVAR_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)
			p.Match(MainDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(523)
			p.Get_child()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(524)
			p.Get_array_element()
		}
		{
			p.SetState(525)
			p.Match(MainDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(526)
			p.Get_child()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(528)
			p.Get_array_element()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(529)
			p.Match(MainVAR_IDENTIFIER)
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
	p.RuleIndex = MainRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MainRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MainRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) VAR_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MainVAR_IDENTIFIER, 0)
}

func (s *Function_callContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(MainL_PAREN, 0)
}

func (s *Function_callContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(MainR_PAREN, 0)
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
	return s.GetTokens(MainCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MainCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MainListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *Main) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, MainRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(532)
		p.Match(MainVAR_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(533)
		p.Match(MainL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(542)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&288793326239359048) != 0 {
		{
			p.SetState(534)
			p.Math_expression()
		}
		p.SetState(539)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MainCOMMA {
			{
				p.SetState(535)
				p.Match(MainCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(536)
				p.Math_expression()
			}

			p.SetState(541)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(544)
		p.Match(MainR_PAREN)
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

func (p *Main) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 26:
		var t *Numberic_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Numberic_expressionContext)
		}
		return p.Numberic_expression_Sempred(t, predIndex)

	case 27:
		var t *Logical_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Logical_expressionContext)
		}
		return p.Logical_expression_Sempred(t, predIndex)

	case 28:
		var t *Comparative_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Comparative_expressionContext)
		}
		return p.Comparative_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Main) Numberic_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Main) Logical_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Main) Comparative_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
