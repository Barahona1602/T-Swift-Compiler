// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "Server2/interfaces"
import "Server2/environment"
import "Server2/expressions"
import "Server2/instructions"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SwiftGrammarParser struct {
	*antlr.BaseParser
}

var SwiftGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func swiftgrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'Int'", "'Float'", "'Bool'", "'String'", "'Character'", "'var'",
		"'let'", "'void'", "'true'", "'false'", "'print'", "'if'", "'else'",
		"'while'", "'for'", "'in'", "'switch'", "'case'", "'default'", "'break'",
		"'return'", "'continue'", "'guard'", "'func'", "'nil'", "'struct'",
		"'mutating'", "'self'", "'inout'", "'append'", "'removeLast'", "'remove'",
		"'at'", "'isEmpty'", "'count'", "", "", "", "'!='", "'=='", "'!'", "'||'",
		"'&&'", "'='", "'>='", "'<='", "'+='", "'-='", "'>'", "'<'", "'*'",
		"'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'", "':'", "'['",
		"']'", "','", "'.'", "'\"'", "'->'", "'_'", "'&'",
	}
	staticData.SymbolicNames = []string{
		"", "INT", "FLOAT", "BOOL", "STR", "CHAR", "VAR", "LET", "VOID", "TRU",
		"FAL", "PRINT", "IF", "ELSE", "WHILE", "FOR", "IN", "SWITCH", "CASE",
		"DEFAULT", "BREAK", "RETURN", "CONTINUE", "GUARD", "FUNC", "NIL", "STRUCT",
		"MUTATING", "SELF", "INOUT", "APPEND", "REMOVELAST", "REMOVE", "AT",
		"ISEMPTY", "COUNT", "NUMBER", "STRING", "ID", "DIF", "IG_IG", "NOT",
		"OR", "AND", "IG", "MAY_IG", "MEN_IG", "SUM_IG", "SUB_IG", "MAYOR",
		"MENOR", "MUL", "DIV", "ADD", "SUB", "MOD", "PARIZQ", "PARDER", "LLAVEIZQ",
		"LLAVEDER", "D_PTS", "CORIZQ", "CORDER", "COMA", "PUNTO", "COMILLA",
		"FLECHA", "GUIONBAJO", "AMP", "WHITESPACE", "COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"s", "block", "instruction", "printstmt", "ifstmt", "whilestmt", "declarationstmt",
		"assignstmt", "forstmt", "guardstmt", "breakstmt", "continuestmt", "structCreation",
		"listStructDec", "fnstmt", "listParamsFunc", "parametro", "callExp",
		"callFunction", "listParamsCall", "types", "expr", "listParams", "listAccessArray",
		"listArray", "exprComa",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 561, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 58, 8, 1, 11, 1, 12, 1, 59, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 100, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 114, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 143, 8, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 192,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		204, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 3, 13, 249, 8, 13, 1, 13, 1, 13, 3, 13, 253, 8, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 263, 8, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 272, 8, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 278, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 284, 8,
		13, 10, 13, 12, 13, 287, 9, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 311, 8, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 3, 15, 318, 8, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 5, 15, 325, 8, 15, 10, 15, 12, 15, 328, 9, 15, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 341, 8,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 360, 8, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 367, 8, 19, 10, 19, 12, 19, 370,
		9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 3, 20, 395, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 439, 8, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 5, 21, 488, 8, 21, 10, 21, 12, 21, 491, 9, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 502,
		8, 22, 10, 22, 12, 22, 505, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 519, 8, 23, 10, 23,
		12, 23, 522, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		5, 24, 542, 8, 24, 10, 24, 12, 24, 545, 9, 24, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25, 556, 8, 25, 10, 25, 12, 25,
		559, 9, 25, 1, 25, 0, 8, 26, 30, 38, 42, 44, 46, 48, 50, 26, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 0, 8, 1, 0, 47, 48, 1, 0, 6, 7, 2, 0, 38, 38, 67, 67, 2,
		0, 51, 52, 55, 55, 1, 0, 53, 54, 2, 0, 45, 45, 49, 49, 2, 0, 46, 46, 50,
		50, 1, 0, 39, 40, 604, 0, 52, 1, 0, 0, 0, 2, 57, 1, 0, 0, 0, 4, 99, 1,
		0, 0, 0, 6, 113, 1, 0, 0, 0, 8, 142, 1, 0, 0, 0, 10, 144, 1, 0, 0, 0, 12,
		191, 1, 0, 0, 0, 14, 203, 1, 0, 0, 0, 16, 205, 1, 0, 0, 0, 18, 214, 1,
		0, 0, 0, 20, 222, 1, 0, 0, 0, 22, 225, 1, 0, 0, 0, 24, 228, 1, 0, 0, 0,
		26, 248, 1, 0, 0, 0, 28, 310, 1, 0, 0, 0, 30, 317, 1, 0, 0, 0, 32, 340,
		1, 0, 0, 0, 34, 342, 1, 0, 0, 0, 36, 348, 1, 0, 0, 0, 38, 359, 1, 0, 0,
		0, 40, 394, 1, 0, 0, 0, 42, 438, 1, 0, 0, 0, 44, 492, 1, 0, 0, 0, 46, 506,
		1, 0, 0, 0, 48, 523, 1, 0, 0, 0, 50, 546, 1, 0, 0, 0, 52, 53, 3, 2, 1,
		0, 53, 54, 5, 0, 0, 1, 54, 55, 6, 0, -1, 0, 55, 1, 1, 0, 0, 0, 56, 58,
		3, 4, 2, 0, 57, 56, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0,
		59, 60, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 6, 1, -1, 0, 62, 3, 1,
		0, 0, 0, 63, 64, 3, 6, 3, 0, 64, 65, 6, 2, -1, 0, 65, 100, 1, 0, 0, 0,
		66, 67, 3, 8, 4, 0, 67, 68, 6, 2, -1, 0, 68, 100, 1, 0, 0, 0, 69, 70, 3,
		12, 6, 0, 70, 71, 6, 2, -1, 0, 71, 100, 1, 0, 0, 0, 72, 73, 3, 10, 5, 0,
		73, 74, 6, 2, -1, 0, 74, 100, 1, 0, 0, 0, 75, 76, 3, 14, 7, 0, 76, 77,
		6, 2, -1, 0, 77, 100, 1, 0, 0, 0, 78, 79, 3, 16, 8, 0, 79, 80, 6, 2, -1,
		0, 80, 100, 1, 0, 0, 0, 81, 82, 3, 18, 9, 0, 82, 83, 6, 2, -1, 0, 83, 100,
		1, 0, 0, 0, 84, 85, 3, 20, 10, 0, 85, 86, 6, 2, -1, 0, 86, 100, 1, 0, 0,
		0, 87, 88, 3, 22, 11, 0, 88, 89, 6, 2, -1, 0, 89, 100, 1, 0, 0, 0, 90,
		91, 3, 24, 12, 0, 91, 92, 6, 2, -1, 0, 92, 100, 1, 0, 0, 0, 93, 94, 3,
		28, 14, 0, 94, 95, 6, 2, -1, 0, 95, 100, 1, 0, 0, 0, 96, 97, 3, 36, 18,
		0, 97, 98, 6, 2, -1, 0, 98, 100, 1, 0, 0, 0, 99, 63, 1, 0, 0, 0, 99, 66,
		1, 0, 0, 0, 99, 69, 1, 0, 0, 0, 99, 72, 1, 0, 0, 0, 99, 75, 1, 0, 0, 0,
		99, 78, 1, 0, 0, 0, 99, 81, 1, 0, 0, 0, 99, 84, 1, 0, 0, 0, 99, 87, 1,
		0, 0, 0, 99, 90, 1, 0, 0, 0, 99, 93, 1, 0, 0, 0, 99, 96, 1, 0, 0, 0, 100,
		5, 1, 0, 0, 0, 101, 102, 5, 11, 0, 0, 102, 103, 5, 56, 0, 0, 103, 104,
		3, 42, 21, 0, 104, 105, 5, 57, 0, 0, 105, 106, 6, 3, -1, 0, 106, 114, 1,
		0, 0, 0, 107, 108, 5, 11, 0, 0, 108, 109, 5, 56, 0, 0, 109, 110, 3, 50,
		25, 0, 110, 111, 5, 57, 0, 0, 111, 112, 6, 3, -1, 0, 112, 114, 1, 0, 0,
		0, 113, 101, 1, 0, 0, 0, 113, 107, 1, 0, 0, 0, 114, 7, 1, 0, 0, 0, 115,
		116, 5, 12, 0, 0, 116, 117, 3, 42, 21, 0, 117, 118, 5, 58, 0, 0, 118, 119,
		3, 2, 1, 0, 119, 120, 5, 59, 0, 0, 120, 121, 6, 4, -1, 0, 121, 143, 1,
		0, 0, 0, 122, 123, 5, 12, 0, 0, 123, 124, 3, 42, 21, 0, 124, 125, 5, 58,
		0, 0, 125, 126, 3, 2, 1, 0, 126, 127, 5, 59, 0, 0, 127, 128, 5, 13, 0,
		0, 128, 129, 5, 58, 0, 0, 129, 130, 3, 2, 1, 0, 130, 131, 5, 59, 0, 0,
		131, 132, 6, 4, -1, 0, 132, 143, 1, 0, 0, 0, 133, 134, 5, 12, 0, 0, 134,
		135, 3, 42, 21, 0, 135, 136, 5, 58, 0, 0, 136, 137, 3, 2, 1, 0, 137, 138,
		5, 59, 0, 0, 138, 139, 5, 13, 0, 0, 139, 140, 3, 8, 4, 0, 140, 141, 6,
		4, -1, 0, 141, 143, 1, 0, 0, 0, 142, 115, 1, 0, 0, 0, 142, 122, 1, 0, 0,
		0, 142, 133, 1, 0, 0, 0, 143, 9, 1, 0, 0, 0, 144, 145, 5, 14, 0, 0, 145,
		146, 3, 42, 21, 0, 146, 147, 5, 58, 0, 0, 147, 148, 3, 2, 1, 0, 148, 149,
		5, 59, 0, 0, 149, 150, 6, 5, -1, 0, 150, 11, 1, 0, 0, 0, 151, 152, 5, 6,
		0, 0, 152, 153, 5, 38, 0, 0, 153, 154, 5, 60, 0, 0, 154, 155, 3, 40, 20,
		0, 155, 156, 5, 44, 0, 0, 156, 157, 3, 42, 21, 0, 157, 158, 6, 6, -1, 0,
		158, 192, 1, 0, 0, 0, 159, 160, 5, 6, 0, 0, 160, 161, 5, 38, 0, 0, 161,
		162, 5, 44, 0, 0, 162, 163, 3, 42, 21, 0, 163, 164, 6, 6, -1, 0, 164, 192,
		1, 0, 0, 0, 165, 166, 5, 6, 0, 0, 166, 167, 5, 38, 0, 0, 167, 168, 5, 60,
		0, 0, 168, 169, 3, 40, 20, 0, 169, 170, 6, 6, -1, 0, 170, 192, 1, 0, 0,
		0, 171, 172, 5, 7, 0, 0, 172, 173, 5, 38, 0, 0, 173, 174, 5, 60, 0, 0,
		174, 175, 3, 40, 20, 0, 175, 176, 5, 44, 0, 0, 176, 177, 3, 42, 21, 0,
		177, 178, 6, 6, -1, 0, 178, 192, 1, 0, 0, 0, 179, 180, 5, 7, 0, 0, 180,
		181, 5, 38, 0, 0, 181, 182, 5, 60, 0, 0, 182, 183, 3, 40, 20, 0, 183, 184,
		6, 6, -1, 0, 184, 192, 1, 0, 0, 0, 185, 186, 5, 7, 0, 0, 186, 187, 5, 38,
		0, 0, 187, 188, 5, 44, 0, 0, 188, 189, 3, 42, 21, 0, 189, 190, 6, 6, -1,
		0, 190, 192, 1, 0, 0, 0, 191, 151, 1, 0, 0, 0, 191, 159, 1, 0, 0, 0, 191,
		165, 1, 0, 0, 0, 191, 171, 1, 0, 0, 0, 191, 179, 1, 0, 0, 0, 191, 185,
		1, 0, 0, 0, 192, 13, 1, 0, 0, 0, 193, 194, 5, 38, 0, 0, 194, 195, 5, 44,
		0, 0, 195, 196, 3, 42, 21, 0, 196, 197, 6, 7, -1, 0, 197, 204, 1, 0, 0,
		0, 198, 199, 5, 38, 0, 0, 199, 200, 7, 0, 0, 0, 200, 201, 3, 42, 21, 0,
		201, 202, 6, 7, -1, 0, 202, 204, 1, 0, 0, 0, 203, 193, 1, 0, 0, 0, 203,
		198, 1, 0, 0, 0, 204, 15, 1, 0, 0, 0, 205, 206, 5, 15, 0, 0, 206, 207,
		5, 38, 0, 0, 207, 208, 5, 16, 0, 0, 208, 209, 3, 42, 21, 0, 209, 210, 5,
		58, 0, 0, 210, 211, 3, 2, 1, 0, 211, 212, 5, 59, 0, 0, 212, 213, 6, 8,
		-1, 0, 213, 17, 1, 0, 0, 0, 214, 215, 5, 23, 0, 0, 215, 216, 3, 42, 21,
		0, 216, 217, 5, 13, 0, 0, 217, 218, 5, 58, 0, 0, 218, 219, 3, 2, 1, 0,
		219, 220, 5, 59, 0, 0, 220, 221, 6, 9, -1, 0, 221, 19, 1, 0, 0, 0, 222,
		223, 5, 20, 0, 0, 223, 224, 6, 10, -1, 0, 224, 21, 1, 0, 0, 0, 225, 226,
		5, 22, 0, 0, 226, 227, 6, 11, -1, 0, 227, 23, 1, 0, 0, 0, 228, 229, 5,
		26, 0, 0, 229, 230, 5, 38, 0, 0, 230, 231, 5, 58, 0, 0, 231, 232, 3, 26,
		13, 0, 232, 233, 5, 59, 0, 0, 233, 234, 6, 12, -1, 0, 234, 25, 1, 0, 0,
		0, 235, 236, 6, 13, -1, 0, 236, 237, 7, 1, 0, 0, 237, 238, 5, 38, 0, 0,
		238, 239, 5, 60, 0, 0, 239, 240, 3, 40, 20, 0, 240, 241, 6, 13, -1, 0,
		241, 249, 1, 0, 0, 0, 242, 243, 7, 1, 0, 0, 243, 244, 5, 38, 0, 0, 244,
		245, 5, 60, 0, 0, 245, 246, 5, 38, 0, 0, 246, 249, 6, 13, -1, 0, 247, 249,
		6, 13, -1, 0, 248, 235, 1, 0, 0, 0, 248, 242, 1, 0, 0, 0, 248, 247, 1,
		0, 0, 0, 249, 285, 1, 0, 0, 0, 250, 252, 10, 6, 0, 0, 251, 253, 5, 63,
		0, 0, 252, 251, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0,
		254, 255, 7, 1, 0, 0, 255, 256, 5, 38, 0, 0, 256, 257, 5, 60, 0, 0, 257,
		258, 3, 40, 20, 0, 258, 259, 6, 13, -1, 0, 259, 284, 1, 0, 0, 0, 260, 262,
		10, 5, 0, 0, 261, 263, 5, 63, 0, 0, 262, 261, 1, 0, 0, 0, 262, 263, 1,
		0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265, 7, 1, 0, 0, 265, 266, 5, 38, 0,
		0, 266, 267, 5, 60, 0, 0, 267, 268, 5, 38, 0, 0, 268, 284, 6, 13, -1, 0,
		269, 271, 10, 4, 0, 0, 270, 272, 5, 63, 0, 0, 271, 270, 1, 0, 0, 0, 271,
		272, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 274, 7, 1, 0, 0, 274, 277,
		5, 38, 0, 0, 275, 276, 5, 60, 0, 0, 276, 278, 3, 40, 20, 0, 277, 275, 1,
		0, 0, 0, 277, 278, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279, 280, 5, 44, 0,
		0, 280, 281, 3, 42, 21, 0, 281, 282, 6, 13, -1, 0, 282, 284, 1, 0, 0, 0,
		283, 250, 1, 0, 0, 0, 283, 260, 1, 0, 0, 0, 283, 269, 1, 0, 0, 0, 284,
		287, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 27, 1,
		0, 0, 0, 287, 285, 1, 0, 0, 0, 288, 289, 5, 24, 0, 0, 289, 290, 5, 38,
		0, 0, 290, 291, 5, 56, 0, 0, 291, 292, 3, 30, 15, 0, 292, 293, 5, 57, 0,
		0, 293, 294, 5, 66, 0, 0, 294, 295, 3, 40, 20, 0, 295, 296, 5, 58, 0, 0,
		296, 297, 3, 2, 1, 0, 297, 298, 5, 59, 0, 0, 298, 299, 6, 14, -1, 0, 299,
		311, 1, 0, 0, 0, 300, 301, 5, 24, 0, 0, 301, 302, 5, 38, 0, 0, 302, 303,
		5, 56, 0, 0, 303, 304, 3, 30, 15, 0, 304, 305, 5, 57, 0, 0, 305, 306, 5,
		58, 0, 0, 306, 307, 3, 2, 1, 0, 307, 308, 5, 59, 0, 0, 308, 309, 6, 14,
		-1, 0, 309, 311, 1, 0, 0, 0, 310, 288, 1, 0, 0, 0, 310, 300, 1, 0, 0, 0,
		311, 29, 1, 0, 0, 0, 312, 313, 6, 15, -1, 0, 313, 314, 3, 32, 16, 0, 314,
		315, 6, 15, -1, 0, 315, 318, 1, 0, 0, 0, 316, 318, 6, 15, -1, 0, 317, 312,
		1, 0, 0, 0, 317, 316, 1, 0, 0, 0, 318, 326, 1, 0, 0, 0, 319, 320, 10, 3,
		0, 0, 320, 321, 5, 63, 0, 0, 321, 322, 3, 32, 16, 0, 322, 323, 6, 15, -1,
		0, 323, 325, 1, 0, 0, 0, 324, 319, 1, 0, 0, 0, 325, 328, 1, 0, 0, 0, 326,
		324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 31, 1, 0, 0, 0, 328, 326, 1,
		0, 0, 0, 329, 330, 5, 38, 0, 0, 330, 331, 5, 60, 0, 0, 331, 332, 3, 40,
		20, 0, 332, 333, 6, 16, -1, 0, 333, 341, 1, 0, 0, 0, 334, 335, 7, 2, 0,
		0, 335, 336, 5, 38, 0, 0, 336, 337, 5, 60, 0, 0, 337, 338, 3, 40, 20, 0,
		338, 339, 6, 16, -1, 0, 339, 341, 1, 0, 0, 0, 340, 329, 1, 0, 0, 0, 340,
		334, 1, 0, 0, 0, 341, 33, 1, 0, 0, 0, 342, 343, 5, 38, 0, 0, 343, 344,
		5, 56, 0, 0, 344, 345, 3, 38, 19, 0, 345, 346, 5, 57, 0, 0, 346, 347, 6,
		17, -1, 0, 347, 35, 1, 0, 0, 0, 348, 349, 5, 38, 0, 0, 349, 350, 5, 56,
		0, 0, 350, 351, 3, 38, 19, 0, 351, 352, 5, 57, 0, 0, 352, 353, 6, 18, -1,
		0, 353, 37, 1, 0, 0, 0, 354, 355, 6, 19, -1, 0, 355, 356, 3, 42, 21, 0,
		356, 357, 6, 19, -1, 0, 357, 360, 1, 0, 0, 0, 358, 360, 6, 19, -1, 0, 359,
		354, 1, 0, 0, 0, 359, 358, 1, 0, 0, 0, 360, 368, 1, 0, 0, 0, 361, 362,
		10, 3, 0, 0, 362, 363, 5, 63, 0, 0, 363, 364, 3, 42, 21, 0, 364, 365, 6,
		19, -1, 0, 365, 367, 1, 0, 0, 0, 366, 361, 1, 0, 0, 0, 367, 370, 1, 0,
		0, 0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 39, 1, 0, 0, 0,
		370, 368, 1, 0, 0, 0, 371, 372, 5, 1, 0, 0, 372, 395, 6, 20, -1, 0, 373,
		374, 5, 2, 0, 0, 374, 395, 6, 20, -1, 0, 375, 376, 5, 4, 0, 0, 376, 395,
		6, 20, -1, 0, 377, 378, 5, 3, 0, 0, 378, 395, 6, 20, -1, 0, 379, 380, 5,
		61, 0, 0, 380, 381, 3, 40, 20, 0, 381, 382, 5, 62, 0, 0, 382, 383, 6, 20,
		-1, 0, 383, 395, 1, 0, 0, 0, 384, 385, 5, 65, 0, 0, 385, 386, 5, 4, 0,
		0, 386, 387, 5, 65, 0, 0, 387, 395, 6, 20, -1, 0, 388, 389, 5, 25, 0, 0,
		389, 395, 6, 20, -1, 0, 390, 391, 5, 26, 0, 0, 391, 395, 6, 20, -1, 0,
		392, 393, 5, 38, 0, 0, 393, 395, 6, 20, -1, 0, 394, 371, 1, 0, 0, 0, 394,
		373, 1, 0, 0, 0, 394, 375, 1, 0, 0, 0, 394, 377, 1, 0, 0, 0, 394, 379,
		1, 0, 0, 0, 394, 384, 1, 0, 0, 0, 394, 388, 1, 0, 0, 0, 394, 390, 1, 0,
		0, 0, 394, 392, 1, 0, 0, 0, 395, 41, 1, 0, 0, 0, 396, 397, 6, 21, -1, 0,
		397, 398, 5, 54, 0, 0, 398, 399, 3, 42, 21, 22, 399, 400, 6, 21, -1, 0,
		400, 439, 1, 0, 0, 0, 401, 402, 5, 41, 0, 0, 402, 403, 3, 42, 21, 12, 403,
		404, 6, 21, -1, 0, 404, 439, 1, 0, 0, 0, 405, 406, 5, 56, 0, 0, 406, 407,
		3, 42, 21, 0, 407, 408, 5, 57, 0, 0, 408, 409, 6, 21, -1, 0, 409, 439,
		1, 0, 0, 0, 410, 411, 3, 34, 17, 0, 411, 412, 6, 21, -1, 0, 412, 439, 1,
		0, 0, 0, 413, 414, 5, 61, 0, 0, 414, 415, 5, 62, 0, 0, 415, 439, 6, 21,
		-1, 0, 416, 417, 3, 48, 24, 0, 417, 418, 6, 21, -1, 0, 418, 439, 1, 0,
		0, 0, 419, 420, 5, 61, 0, 0, 420, 421, 3, 44, 22, 0, 421, 422, 5, 62, 0,
		0, 422, 423, 6, 21, -1, 0, 423, 439, 1, 0, 0, 0, 424, 425, 5, 36, 0, 0,
		425, 439, 6, 21, -1, 0, 426, 427, 5, 37, 0, 0, 427, 439, 6, 21, -1, 0,
		428, 429, 5, 9, 0, 0, 429, 439, 6, 21, -1, 0, 430, 431, 5, 10, 0, 0, 431,
		439, 6, 21, -1, 0, 432, 433, 5, 38, 0, 0, 433, 434, 5, 64, 0, 0, 434, 435,
		5, 35, 0, 0, 435, 439, 6, 21, -1, 0, 436, 437, 5, 25, 0, 0, 437, 439, 6,
		21, -1, 0, 438, 396, 1, 0, 0, 0, 438, 401, 1, 0, 0, 0, 438, 405, 1, 0,
		0, 0, 438, 410, 1, 0, 0, 0, 438, 413, 1, 0, 0, 0, 438, 416, 1, 0, 0, 0,
		438, 419, 1, 0, 0, 0, 438, 424, 1, 0, 0, 0, 438, 426, 1, 0, 0, 0, 438,
		428, 1, 0, 0, 0, 438, 430, 1, 0, 0, 0, 438, 432, 1, 0, 0, 0, 438, 436,
		1, 0, 0, 0, 439, 489, 1, 0, 0, 0, 440, 441, 10, 21, 0, 0, 441, 442, 5,
		64, 0, 0, 442, 443, 5, 64, 0, 0, 443, 444, 5, 64, 0, 0, 444, 445, 3, 42,
		21, 22, 445, 446, 6, 21, -1, 0, 446, 488, 1, 0, 0, 0, 447, 448, 10, 20,
		0, 0, 448, 449, 7, 0, 0, 0, 449, 450, 3, 42, 21, 21, 450, 451, 6, 21, -1,
		0, 451, 488, 1, 0, 0, 0, 452, 453, 10, 19, 0, 0, 453, 454, 7, 3, 0, 0,
		454, 455, 3, 42, 21, 20, 455, 456, 6, 21, -1, 0, 456, 488, 1, 0, 0, 0,
		457, 458, 10, 18, 0, 0, 458, 459, 7, 4, 0, 0, 459, 460, 3, 42, 21, 19,
		460, 461, 6, 21, -1, 0, 461, 488, 1, 0, 0, 0, 462, 463, 10, 17, 0, 0, 463,
		464, 7, 5, 0, 0, 464, 465, 3, 42, 21, 18, 465, 466, 6, 21, -1, 0, 466,
		488, 1, 0, 0, 0, 467, 468, 10, 16, 0, 0, 468, 469, 7, 6, 0, 0, 469, 470,
		3, 42, 21, 17, 470, 471, 6, 21, -1, 0, 471, 488, 1, 0, 0, 0, 472, 473,
		10, 15, 0, 0, 473, 474, 7, 7, 0, 0, 474, 475, 3, 42, 21, 16, 475, 476,
		6, 21, -1, 0, 476, 488, 1, 0, 0, 0, 477, 478, 10, 14, 0, 0, 478, 479, 5,
		43, 0, 0, 479, 480, 3, 42, 21, 15, 480, 481, 6, 21, -1, 0, 481, 488, 1,
		0, 0, 0, 482, 483, 10, 13, 0, 0, 483, 484, 5, 42, 0, 0, 484, 485, 3, 42,
		21, 14, 485, 486, 6, 21, -1, 0, 486, 488, 1, 0, 0, 0, 487, 440, 1, 0, 0,
		0, 487, 447, 1, 0, 0, 0, 487, 452, 1, 0, 0, 0, 487, 457, 1, 0, 0, 0, 487,
		462, 1, 0, 0, 0, 487, 467, 1, 0, 0, 0, 487, 472, 1, 0, 0, 0, 487, 477,
		1, 0, 0, 0, 487, 482, 1, 0, 0, 0, 488, 491, 1, 0, 0, 0, 489, 487, 1, 0,
		0, 0, 489, 490, 1, 0, 0, 0, 490, 43, 1, 0, 0, 0, 491, 489, 1, 0, 0, 0,
		492, 493, 6, 22, -1, 0, 493, 494, 3, 42, 21, 0, 494, 495, 6, 22, -1, 0,
		495, 503, 1, 0, 0, 0, 496, 497, 10, 2, 0, 0, 497, 498, 5, 63, 0, 0, 498,
		499, 3, 42, 21, 0, 499, 500, 6, 22, -1, 0, 500, 502, 1, 0, 0, 0, 501, 496,
		1, 0, 0, 0, 502, 505, 1, 0, 0, 0, 503, 501, 1, 0, 0, 0, 503, 504, 1, 0,
		0, 0, 504, 45, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 506, 507, 6, 23, -1, 0,
		507, 508, 5, 61, 0, 0, 508, 509, 3, 42, 21, 0, 509, 510, 5, 62, 0, 0, 510,
		511, 6, 23, -1, 0, 511, 520, 1, 0, 0, 0, 512, 513, 10, 2, 0, 0, 513, 514,
		5, 61, 0, 0, 514, 515, 3, 42, 21, 0, 515, 516, 5, 62, 0, 0, 516, 517, 6,
		23, -1, 0, 517, 519, 1, 0, 0, 0, 518, 512, 1, 0, 0, 0, 519, 522, 1, 0,
		0, 0, 520, 518, 1, 0, 0, 0, 520, 521, 1, 0, 0, 0, 521, 47, 1, 0, 0, 0,
		522, 520, 1, 0, 0, 0, 523, 524, 6, 24, -1, 0, 524, 525, 5, 38, 0, 0, 525,
		526, 6, 24, -1, 0, 526, 543, 1, 0, 0, 0, 527, 528, 10, 3, 0, 0, 528, 529,
		5, 61, 0, 0, 529, 530, 3, 42, 21, 0, 530, 531, 5, 62, 0, 0, 531, 532, 6,
		24, -1, 0, 532, 542, 1, 0, 0, 0, 533, 534, 10, 2, 0, 0, 534, 535, 3, 40,
		20, 0, 535, 536, 5, 44, 0, 0, 536, 537, 5, 61, 0, 0, 537, 538, 3, 42, 21,
		0, 538, 539, 5, 62, 0, 0, 539, 540, 6, 24, -1, 0, 540, 542, 1, 0, 0, 0,
		541, 527, 1, 0, 0, 0, 541, 533, 1, 0, 0, 0, 542, 545, 1, 0, 0, 0, 543,
		541, 1, 0, 0, 0, 543, 544, 1, 0, 0, 0, 544, 49, 1, 0, 0, 0, 545, 543, 1,
		0, 0, 0, 546, 547, 6, 25, -1, 0, 547, 548, 3, 42, 21, 0, 548, 549, 6, 25,
		-1, 0, 549, 557, 1, 0, 0, 0, 550, 551, 10, 2, 0, 0, 551, 552, 5, 63, 0,
		0, 552, 553, 3, 42, 21, 0, 553, 554, 6, 25, -1, 0, 554, 556, 1, 0, 0, 0,
		555, 550, 1, 0, 0, 0, 556, 559, 1, 0, 0, 0, 557, 555, 1, 0, 0, 0, 557,
		558, 1, 0, 0, 0, 558, 51, 1, 0, 0, 0, 559, 557, 1, 0, 0, 0, 28, 59, 99,
		113, 142, 191, 203, 248, 252, 262, 271, 277, 283, 285, 310, 317, 326, 340,
		359, 368, 394, 438, 487, 489, 503, 520, 541, 543, 557,
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

// SwiftGrammarParserInit initializes any static state used to implement SwiftGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSwiftGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SwiftGrammarParserInit() {
	staticData := &SwiftGrammarParserStaticData
	staticData.once.Do(swiftgrammarParserInit)
}

// NewSwiftGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSwiftGrammarParser(input antlr.TokenStream) *SwiftGrammarParser {
	SwiftGrammarParserInit()
	this := new(SwiftGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SwiftGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SwiftGrammar.g4"

	return this
}

// SwiftGrammarParser tokens.
const (
	SwiftGrammarParserEOF          = antlr.TokenEOF
	SwiftGrammarParserINT          = 1
	SwiftGrammarParserFLOAT        = 2
	SwiftGrammarParserBOOL         = 3
	SwiftGrammarParserSTR          = 4
	SwiftGrammarParserCHAR         = 5
	SwiftGrammarParserVAR          = 6
	SwiftGrammarParserLET          = 7
	SwiftGrammarParserVOID         = 8
	SwiftGrammarParserTRU          = 9
	SwiftGrammarParserFAL          = 10
	SwiftGrammarParserPRINT        = 11
	SwiftGrammarParserIF           = 12
	SwiftGrammarParserELSE         = 13
	SwiftGrammarParserWHILE        = 14
	SwiftGrammarParserFOR          = 15
	SwiftGrammarParserIN           = 16
	SwiftGrammarParserSWITCH       = 17
	SwiftGrammarParserCASE         = 18
	SwiftGrammarParserDEFAULT      = 19
	SwiftGrammarParserBREAK        = 20
	SwiftGrammarParserRETURN       = 21
	SwiftGrammarParserCONTINUE     = 22
	SwiftGrammarParserGUARD        = 23
	SwiftGrammarParserFUNC         = 24
	SwiftGrammarParserNIL          = 25
	SwiftGrammarParserSTRUCT       = 26
	SwiftGrammarParserMUTATING     = 27
	SwiftGrammarParserSELF         = 28
	SwiftGrammarParserINOUT        = 29
	SwiftGrammarParserAPPEND       = 30
	SwiftGrammarParserREMOVELAST   = 31
	SwiftGrammarParserREMOVE       = 32
	SwiftGrammarParserAT           = 33
	SwiftGrammarParserISEMPTY      = 34
	SwiftGrammarParserCOUNT        = 35
	SwiftGrammarParserNUMBER       = 36
	SwiftGrammarParserSTRING       = 37
	SwiftGrammarParserID           = 38
	SwiftGrammarParserDIF          = 39
	SwiftGrammarParserIG_IG        = 40
	SwiftGrammarParserNOT          = 41
	SwiftGrammarParserOR           = 42
	SwiftGrammarParserAND          = 43
	SwiftGrammarParserIG           = 44
	SwiftGrammarParserMAY_IG       = 45
	SwiftGrammarParserMEN_IG       = 46
	SwiftGrammarParserSUM_IG       = 47
	SwiftGrammarParserSUB_IG       = 48
	SwiftGrammarParserMAYOR        = 49
	SwiftGrammarParserMENOR        = 50
	SwiftGrammarParserMUL          = 51
	SwiftGrammarParserDIV          = 52
	SwiftGrammarParserADD          = 53
	SwiftGrammarParserSUB          = 54
	SwiftGrammarParserMOD          = 55
	SwiftGrammarParserPARIZQ       = 56
	SwiftGrammarParserPARDER       = 57
	SwiftGrammarParserLLAVEIZQ     = 58
	SwiftGrammarParserLLAVEDER     = 59
	SwiftGrammarParserD_PTS        = 60
	SwiftGrammarParserCORIZQ       = 61
	SwiftGrammarParserCORDER       = 62
	SwiftGrammarParserCOMA         = 63
	SwiftGrammarParserPUNTO        = 64
	SwiftGrammarParserCOMILLA      = 65
	SwiftGrammarParserFLECHA       = 66
	SwiftGrammarParserGUIONBAJO    = 67
	SwiftGrammarParserAMP          = 68
	SwiftGrammarParserWHITESPACE   = 69
	SwiftGrammarParserCOMMENT      = 70
	SwiftGrammarParserLINE_COMMENT = 71
)

// SwiftGrammarParser rules.
const (
	SwiftGrammarParserRULE_s               = 0
	SwiftGrammarParserRULE_block           = 1
	SwiftGrammarParserRULE_instruction     = 2
	SwiftGrammarParserRULE_printstmt       = 3
	SwiftGrammarParserRULE_ifstmt          = 4
	SwiftGrammarParserRULE_whilestmt       = 5
	SwiftGrammarParserRULE_declarationstmt = 6
	SwiftGrammarParserRULE_assignstmt      = 7
	SwiftGrammarParserRULE_forstmt         = 8
	SwiftGrammarParserRULE_guardstmt       = 9
	SwiftGrammarParserRULE_breakstmt       = 10
	SwiftGrammarParserRULE_continuestmt    = 11
	SwiftGrammarParserRULE_structCreation  = 12
	SwiftGrammarParserRULE_listStructDec   = 13
	SwiftGrammarParserRULE_fnstmt          = 14
	SwiftGrammarParserRULE_listParamsFunc  = 15
	SwiftGrammarParserRULE_parametro       = 16
	SwiftGrammarParserRULE_callExp         = 17
	SwiftGrammarParserRULE_callFunction    = 18
	SwiftGrammarParserRULE_listParamsCall  = 19
	SwiftGrammarParserRULE_types           = 20
	SwiftGrammarParserRULE_expr            = 21
	SwiftGrammarParserRULE_listParams      = 22
	SwiftGrammarParserRULE_listAccessArray = 23
	SwiftGrammarParserRULE_listArray       = 24
	SwiftGrammarParserRULE_exprComa        = 25
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetCode returns the code attribute.
	GetCode() []interface{}

	// SetCode sets the code attribute.
	SetCode([]interface{})

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	code   []interface{}
	_block IBlockContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) Get_block() IBlockContext { return s._block }

func (s *SContext) Set_block(v IBlockContext) { s._block = v }

func (s *SContext) GetCode() []interface{} { return s.code }

func (s *SContext) SetCode(v []interface{}) { s.code = v }

func (s *SContext) Block() IBlockContext {
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

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserEOF, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *SwiftGrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SwiftGrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)

		var _x = p.Block()

		localctx.(*SContext)._block = _x
	}
	{
		p.SetState(53)
		p.Match(SwiftGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*SContext).code = localctx.(*SContext).Get_block().GetBlk()

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

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetIns returns the ins rule context list.
	GetIns() []IInstructionContext

	// SetIns sets the ins rule context list.
	SetIns([]IInstructionContext)

	// GetBlk returns the blk attribute.
	GetBlk() []interface{}

	// SetBlk sets the blk attribute.
	SetBlk([]interface{})

	// Getter signatures
	AllInstruction() []IInstructionContext
	Instruction(i int) IInstructionContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	blk          []interface{}
	_instruction IInstructionContext
	ins          []IInstructionContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *BlockContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *BlockContext) GetIns() []IInstructionContext { return s.ins }

func (s *BlockContext) SetIns(v []IInstructionContext) { s.ins = v }

func (s *BlockContext) GetBlk() []interface{} { return s.blk }

func (s *BlockContext) SetBlk(v []interface{}) { s.blk = v }

func (s *BlockContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
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

	return t.(IInstructionContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SwiftGrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SwiftGrammarParserRULE_block)

	localctx.(*BlockContext).blk = []interface{}{}
	var listInt []IInstructionContext

	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&274975480000) != 0) {
		{
			p.SetState(56)

			var _x = p.Instruction()

			localctx.(*BlockContext)._instruction = _x
		}
		localctx.(*BlockContext).ins = append(localctx.(*BlockContext).ins, localctx.(*BlockContext)._instruction)

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	listInt = localctx.(*BlockContext).GetIns()
	for _, e := range listInt {
		localctx.(*BlockContext).blk = append(localctx.(*BlockContext).blk, e.GetInst())
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

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_printstmt returns the _printstmt rule contexts.
	Get_printstmt() IPrintstmtContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Get_declarationstmt returns the _declarationstmt rule contexts.
	Get_declarationstmt() IDeclarationstmtContext

	// Get_whilestmt returns the _whilestmt rule contexts.
	Get_whilestmt() IWhilestmtContext

	// Get_assignstmt returns the _assignstmt rule contexts.
	Get_assignstmt() IAssignstmtContext

	// Get_forstmt returns the _forstmt rule contexts.
	Get_forstmt() IForstmtContext

	// Get_guardstmt returns the _guardstmt rule contexts.
	Get_guardstmt() IGuardstmtContext

	// Get_breakstmt returns the _breakstmt rule contexts.
	Get_breakstmt() IBreakstmtContext

	// Get_continuestmt returns the _continuestmt rule contexts.
	Get_continuestmt() IContinuestmtContext

	// Get_structCreation returns the _structCreation rule contexts.
	Get_structCreation() IStructCreationContext

	// Get_fnstmt returns the _fnstmt rule contexts.
	Get_fnstmt() IFnstmtContext

	// Get_callFunction returns the _callFunction rule contexts.
	Get_callFunction() ICallFunctionContext

	// Set_printstmt sets the _printstmt rule contexts.
	Set_printstmt(IPrintstmtContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// Set_declarationstmt sets the _declarationstmt rule contexts.
	Set_declarationstmt(IDeclarationstmtContext)

	// Set_whilestmt sets the _whilestmt rule contexts.
	Set_whilestmt(IWhilestmtContext)

	// Set_assignstmt sets the _assignstmt rule contexts.
	Set_assignstmt(IAssignstmtContext)

	// Set_forstmt sets the _forstmt rule contexts.
	Set_forstmt(IForstmtContext)

	// Set_guardstmt sets the _guardstmt rule contexts.
	Set_guardstmt(IGuardstmtContext)

	// Set_breakstmt sets the _breakstmt rule contexts.
	Set_breakstmt(IBreakstmtContext)

	// Set_continuestmt sets the _continuestmt rule contexts.
	Set_continuestmt(IContinuestmtContext)

	// Set_structCreation sets the _structCreation rule contexts.
	Set_structCreation(IStructCreationContext)

	// Set_fnstmt sets the _fnstmt rule contexts.
	Set_fnstmt(IFnstmtContext)

	// Set_callFunction sets the _callFunction rule contexts.
	Set_callFunction(ICallFunctionContext)

	// GetInst returns the inst attribute.
	GetInst() interfaces.Instruction

	// SetInst sets the inst attribute.
	SetInst(interfaces.Instruction)

	// Getter signatures
	Printstmt() IPrintstmtContext
	Ifstmt() IIfstmtContext
	Declarationstmt() IDeclarationstmtContext
	Whilestmt() IWhilestmtContext
	Assignstmt() IAssignstmtContext
	Forstmt() IForstmtContext
	Guardstmt() IGuardstmtContext
	Breakstmt() IBreakstmtContext
	Continuestmt() IContinuestmtContext
	StructCreation() IStructCreationContext
	Fnstmt() IFnstmtContext
	CallFunction() ICallFunctionContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser           antlr.Parser
	inst             interfaces.Instruction
	_printstmt       IPrintstmtContext
	_ifstmt          IIfstmtContext
	_declarationstmt IDeclarationstmtContext
	_whilestmt       IWhilestmtContext
	_assignstmt      IAssignstmtContext
	_forstmt         IForstmtContext
	_guardstmt       IGuardstmtContext
	_breakstmt       IBreakstmtContext
	_continuestmt    IContinuestmtContext
	_structCreation  IStructCreationContext
	_fnstmt          IFnstmtContext
	_callFunction    ICallFunctionContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_printstmt() IPrintstmtContext { return s._printstmt }

func (s *InstructionContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *InstructionContext) Get_declarationstmt() IDeclarationstmtContext { return s._declarationstmt }

func (s *InstructionContext) Get_whilestmt() IWhilestmtContext { return s._whilestmt }

func (s *InstructionContext) Get_assignstmt() IAssignstmtContext { return s._assignstmt }

func (s *InstructionContext) Get_forstmt() IForstmtContext { return s._forstmt }

func (s *InstructionContext) Get_guardstmt() IGuardstmtContext { return s._guardstmt }

func (s *InstructionContext) Get_breakstmt() IBreakstmtContext { return s._breakstmt }

func (s *InstructionContext) Get_continuestmt() IContinuestmtContext { return s._continuestmt }

func (s *InstructionContext) Get_structCreation() IStructCreationContext { return s._structCreation }

func (s *InstructionContext) Get_fnstmt() IFnstmtContext { return s._fnstmt }

func (s *InstructionContext) Get_callFunction() ICallFunctionContext { return s._callFunction }

func (s *InstructionContext) Set_printstmt(v IPrintstmtContext) { s._printstmt = v }

func (s *InstructionContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *InstructionContext) Set_declarationstmt(v IDeclarationstmtContext) { s._declarationstmt = v }

func (s *InstructionContext) Set_whilestmt(v IWhilestmtContext) { s._whilestmt = v }

func (s *InstructionContext) Set_assignstmt(v IAssignstmtContext) { s._assignstmt = v }

func (s *InstructionContext) Set_forstmt(v IForstmtContext) { s._forstmt = v }

func (s *InstructionContext) Set_guardstmt(v IGuardstmtContext) { s._guardstmt = v }

func (s *InstructionContext) Set_breakstmt(v IBreakstmtContext) { s._breakstmt = v }

func (s *InstructionContext) Set_continuestmt(v IContinuestmtContext) { s._continuestmt = v }

func (s *InstructionContext) Set_structCreation(v IStructCreationContext) { s._structCreation = v }

func (s *InstructionContext) Set_fnstmt(v IFnstmtContext) { s._fnstmt = v }

func (s *InstructionContext) Set_callFunction(v ICallFunctionContext) { s._callFunction = v }

func (s *InstructionContext) GetInst() interfaces.Instruction { return s.inst }

func (s *InstructionContext) SetInst(v interfaces.Instruction) { s.inst = v }

func (s *InstructionContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *InstructionContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *InstructionContext) Declarationstmt() IDeclarationstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationstmtContext)
}

func (s *InstructionContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *InstructionContext) Assignstmt() IAssignstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignstmtContext)
}

func (s *InstructionContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *InstructionContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
}

func (s *InstructionContext) Breakstmt() IBreakstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakstmtContext)
}

func (s *InstructionContext) Continuestmt() IContinuestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinuestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinuestmtContext)
}

func (s *InstructionContext) StructCreation() IStructCreationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructCreationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructCreationContext)
}

func (s *InstructionContext) Fnstmt() IFnstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnstmtContext)
}

func (s *InstructionContext) CallFunction() ICallFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFunctionContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *SwiftGrammarParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SwiftGrammarParserRULE_instruction)
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)

			var _x = p.Printstmt()

			localctx.(*InstructionContext)._printstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_printstmt().GetPrnt()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)

			var _x = p.Ifstmt()

			localctx.(*InstructionContext)._ifstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_ifstmt().GetIfinst()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)

			var _x = p.Declarationstmt()

			localctx.(*InstructionContext)._declarationstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_declarationstmt().GetDec()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)

			var _x = p.Whilestmt()

			localctx.(*InstructionContext)._whilestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_whilestmt().GetWhl()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)

			var _x = p.Assignstmt()

			localctx.(*InstructionContext)._assignstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_assignstmt().GetAsg()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(78)

			var _x = p.Forstmt()

			localctx.(*InstructionContext)._forstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_forstmt().GetFr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(81)

			var _x = p.Guardstmt()

			localctx.(*InstructionContext)._guardstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_guardstmt().GetGrd()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(84)

			var _x = p.Breakstmt()

			localctx.(*InstructionContext)._breakstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_breakstmt().GetBrk()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(87)

			var _x = p.Continuestmt()

			localctx.(*InstructionContext)._continuestmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_continuestmt().GetCnt()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(90)

			var _x = p.StructCreation()

			localctx.(*InstructionContext)._structCreation = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_structCreation().GetDec()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(93)

			var _x = p.Fnstmt()

			localctx.(*InstructionContext)._fnstmt = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_fnstmt().GetFn()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(96)

			var _x = p.CallFunction()

			localctx.(*InstructionContext)._callFunction = _x
		}
		localctx.(*InstructionContext).inst = localctx.(*InstructionContext).Get_callFunction().GetCf()

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

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_exprComa returns the _exprComa rule contexts.
	Get_exprComa() IExprComaContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_exprComa sets the _exprComa rule contexts.
	Set_exprComa(IExprComaContext)

	// GetPrnt returns the prnt attribute.
	GetPrnt() interfaces.Instruction

	// SetPrnt sets the prnt attribute.
	SetPrnt(interfaces.Instruction)

	// Getter signatures
	PRINT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	Expr() IExprContext
	PARDER() antlr.TerminalNode
	ExprComa() IExprComaContext

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	prnt      interfaces.Instruction
	_PRINT    antlr.Token
	_expr     IExprContext
	_exprComa IExprComaContext
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintstmtContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintstmtContext) Get_expr() IExprContext { return s._expr }

func (s *PrintstmtContext) Get_exprComa() IExprComaContext { return s._exprComa }

func (s *PrintstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *PrintstmtContext) Set_exprComa(v IExprComaContext) { s._exprComa = v }

func (s *PrintstmtContext) GetPrnt() interfaces.Instruction { return s.prnt }

func (s *PrintstmtContext) SetPrnt(v interfaces.Instruction) { s.prnt = v }

func (s *PrintstmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPRINT, 0)
}

func (s *PrintstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *PrintstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrintstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *PrintstmtContext) ExprComa() IExprComaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprComaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprComaContext)
}

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (p *SwiftGrammarParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SwiftGrammarParserRULE_printstmt)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)

			var _m = p.Match(SwiftGrammarParserPRINT)

			localctx.(*PrintstmtContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)

			var _x = p.expr(0)

			localctx.(*PrintstmtContext)._expr = _x
		}
		{
			p.SetState(104)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
			}
		}()), localctx.(*PrintstmtContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)

			var _m = p.Match(SwiftGrammarParserPRINT)

			localctx.(*PrintstmtContext)._PRINT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)

			var _x = p.exprComa(0)

			localctx.(*PrintstmtContext)._exprComa = _x
		}
		{
			p.SetState(110)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*PrintstmtContext).prnt = instructions.NewPrint((func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetLine()
			}
		}()), (func() int {
			if localctx.(*PrintstmtContext).Get_PRINT() == nil {
				return 0
			} else {
				return localctx.(*PrintstmtContext).Get_PRINT().GetColumn()
			}
		}()), localctx.(*PrintstmtContext).Get_exprComa().GetT())

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

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// GetE1 returns the e1 rule contexts.
	GetE1() IBlockContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IBlockContext

	// Get_ifstmt returns the _ifstmt rule contexts.
	Get_ifstmt() IIfstmtContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// SetE1 sets the e1 rule contexts.
	SetE1(IBlockContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IBlockContext)

	// Set_ifstmt sets the _ifstmt rule contexts.
	Set_ifstmt(IIfstmtContext)

	// GetIfinst returns the ifinst attribute.
	GetIfinst() interfaces.Instruction

	// SetIfinst sets the ifinst attribute.
	SetIfinst(interfaces.Instruction)

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllLLAVEIZQ() []antlr.TerminalNode
	LLAVEIZQ(i int) antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	AllLLAVEDER() []antlr.TerminalNode
	LLAVEDER(i int) antlr.TerminalNode
	ELSE() antlr.TerminalNode
	Ifstmt() IIfstmtContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	ifinst  interfaces.Instruction
	_IF     antlr.Token
	_expr   IExprContext
	_block  IBlockContext
	e1      IBlockContext
	e2      IBlockContext
	_ifstmt IIfstmtContext
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Get_IF() antlr.Token { return s._IF }

func (s *IfstmtContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *IfstmtContext) Get_expr() IExprContext { return s._expr }

func (s *IfstmtContext) Get_block() IBlockContext { return s._block }

func (s *IfstmtContext) GetE1() IBlockContext { return s.e1 }

func (s *IfstmtContext) GetE2() IBlockContext { return s.e2 }

func (s *IfstmtContext) Get_ifstmt() IIfstmtContext { return s._ifstmt }

func (s *IfstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *IfstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *IfstmtContext) SetE1(v IBlockContext) { s.e1 = v }

func (s *IfstmtContext) SetE2(v IBlockContext) { s.e2 = v }

func (s *IfstmtContext) Set_ifstmt(v IIfstmtContext) { s._ifstmt = v }

func (s *IfstmtContext) GetIfinst() interfaces.Instruction { return s.ifinst }

func (s *IfstmtContext) SetIfinst(v interfaces.Instruction) { s.ifinst = v }

func (s *IfstmtContext) IF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIF, 0)
}

func (s *IfstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfstmtContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEIZQ)
}

func (s *IfstmtContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, i)
}

func (s *IfstmtContext) AllBlock() []IBlockContext {
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

func (s *IfstmtContext) Block(i int) IBlockContext {
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

func (s *IfstmtContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserLLAVEDER)
}

func (s *IfstmtContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, i)
}

func (s *IfstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *IfstmtContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (p *SwiftGrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SwiftGrammarParserRULE_ifstmt)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(117)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(119)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).Get_block().GetBlk(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(124)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)

			var _x = p.Block()

			localctx.(*IfstmtContext).e1 = _x
		}
		{
			p.SetState(126)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)

			var _x = p.Block()

			localctx.(*IfstmtContext).e2 = _x
		}
		{
			p.SetState(130)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).GetE1().GetBlk(), localctx.(*IfstmtContext).GetE2().GetBlk())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)

			var _m = p.Match(SwiftGrammarParserIF)

			localctx.(*IfstmtContext)._IF = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)

			var _x = p.expr(0)

			localctx.(*IfstmtContext)._expr = _x
		}
		{
			p.SetState(135)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)

			var _x = p.Block()

			localctx.(*IfstmtContext)._block = _x
		}
		{
			p.SetState(137)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(SwiftGrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)

			var _x = p.Ifstmt()

			localctx.(*IfstmtContext)._ifstmt = _x
		}
		localctx.(*IfstmtContext).ifinst = instructions.NewIf((func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetLine()
			}
		}()), (func() int {
			if localctx.(*IfstmtContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*IfstmtContext).Get_IF().GetColumn()
			}
		}()), localctx.(*IfstmtContext).Get_expr().GetE(), localctx.(*IfstmtContext).Get_block().GetBlk(), []interface{}{localctx.(*IfstmtContext).Get_ifstmt().GetIfinst()})

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

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetWhl returns the whl attribute.
	GetWhl() interfaces.Instruction

	// SetWhl sets the whl attribute.
	SetWhl(interfaces.Instruction)

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	whl    interfaces.Instruction
	_WHILE antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *WhilestmtContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *WhilestmtContext) Get_expr() IExprContext { return s._expr }

func (s *WhilestmtContext) Get_block() IBlockContext { return s._block }

func (s *WhilestmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *WhilestmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *WhilestmtContext) GetWhl() interfaces.Instruction { return s.whl }

func (s *WhilestmtContext) SetWhl(v interfaces.Instruction) { s.whl = v }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *WhilestmtContext) Block() IBlockContext {
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

func (s *WhilestmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (p *SwiftGrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SwiftGrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)

		var _m = p.Match(SwiftGrammarParserWHILE)

		localctx.(*WhilestmtContext)._WHILE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)

		var _x = p.expr(0)

		localctx.(*WhilestmtContext)._expr = _x
	}
	{
		p.SetState(146)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)

		var _x = p.Block()

		localctx.(*WhilestmtContext)._block = _x
	}
	{
		p.SetState(148)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*WhilestmtContext).whl = instructions.NewWhile((func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetLine()
		}
	}()), (func() int {
		if localctx.(*WhilestmtContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*WhilestmtContext).Get_WHILE().GetColumn()
		}
	}()), localctx.(*WhilestmtContext).Get_expr().GetE(), localctx.(*WhilestmtContext).Get_block().GetBlk())

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

// IDeclarationstmtContext is an interface to support dynamic dispatch.
type IDeclarationstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VAR returns the _VAR token.
	Get_VAR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Set_VAR sets the _VAR token.
	Set_VAR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	VAR() antlr.TerminalNode
	ID() antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Types() ITypesContext
	IG() antlr.TerminalNode
	Expr() IExprContext
	LET() antlr.TerminalNode

	// IsDeclarationstmtContext differentiates from other interfaces.
	IsDeclarationstmtContext()
}

type DeclarationstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	dec    interfaces.Instruction
	_VAR   antlr.Token
	_ID    antlr.Token
	_types ITypesContext
	_expr  IExprContext
	_LET   antlr.Token
}

func NewEmptyDeclarationstmtContext() *DeclarationstmtContext {
	var p = new(DeclarationstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt
	return p
}

func InitEmptyDeclarationstmtContext(p *DeclarationstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt
}

func (*DeclarationstmtContext) IsDeclarationstmtContext() {}

func NewDeclarationstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationstmtContext {
	var p = new(DeclarationstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_declarationstmt

	return p
}

func (s *DeclarationstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationstmtContext) Get_VAR() antlr.Token { return s._VAR }

func (s *DeclarationstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclarationstmtContext) Get_LET() antlr.Token { return s._LET }

func (s *DeclarationstmtContext) Set_VAR(v antlr.Token) { s._VAR = v }

func (s *DeclarationstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclarationstmtContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *DeclarationstmtContext) Get_types() ITypesContext { return s._types }

func (s *DeclarationstmtContext) Get_expr() IExprContext { return s._expr }

func (s *DeclarationstmtContext) Set_types(v ITypesContext) { s._types = v }

func (s *DeclarationstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *DeclarationstmtContext) GetDec() interfaces.Instruction { return s.dec }

func (s *DeclarationstmtContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *DeclarationstmtContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *DeclarationstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *DeclarationstmtContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *DeclarationstmtContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *DeclarationstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *DeclarationstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclarationstmtContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *DeclarationstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterDeclarationstmt(s)
	}
}

func (s *DeclarationstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitDeclarationstmt(s)
	}
}

func (p *SwiftGrammarParser) Declarationstmt() (localctx IDeclarationstmtContext) {
	localctx = NewDeclarationstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SwiftGrammarParserRULE_declarationstmt)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(155)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationstmtContext).Get_types().GetTy(), localctx.(*DeclarationstmtContext).Get_expr().GetE(), true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), environment.UNKNOWN, localctx.(*DeclarationstmtContext).Get_expr().GetE(), true)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(165)

			var _m = p.Match(SwiftGrammarParserVAR)

			localctx.(*DeclarationstmtContext)._VAR = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_VAR() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_VAR().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationstmtContext).Get_types().GetTy(), nil, true)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(171)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		{
			p.SetState(175)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationstmtContext).Get_types().GetTy(), localctx.(*DeclarationstmtContext).Get_expr().GetE(), false)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(179)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)

			var _x = p.Types()

			localctx.(*DeclarationstmtContext)._types = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclarationstmtContext).Get_types().GetTy(), nil, false)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(185)

			var _m = p.Match(SwiftGrammarParserLET)

			localctx.(*DeclarationstmtContext)._LET = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*DeclarationstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(SwiftGrammarParserIG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)

			var _x = p.expr(0)

			localctx.(*DeclarationstmtContext)._expr = _x
		}
		localctx.(*DeclarationstmtContext).dec = instructions.NewDeclaration((func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetLine()
			}
		}()), (func() int {
			if localctx.(*DeclarationstmtContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclarationstmtContext).Get_LET().GetColumn()
			}
		}()), (func() string {
			if localctx.(*DeclarationstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclarationstmtContext).Get_ID().GetText()
			}
		}()), environment.UNKNOWN, localctx.(*DeclarationstmtContext).Get_expr().GetE(), false)

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

// IAssignstmtContext is an interface to support dynamic dispatch.
type IAssignstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetAsg returns the asg attribute.
	GetAsg() interfaces.Instruction

	// SetAsg sets the asg attribute.
	SetAsg(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext
	IG() antlr.TerminalNode
	SUB_IG() antlr.TerminalNode
	SUM_IG() antlr.TerminalNode

	// IsAssignstmtContext differentiates from other interfaces.
	IsAssignstmtContext()
}

type AssignstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	asg    interfaces.Instruction
	_ID    antlr.Token
	op     antlr.Token
	_expr  IExprContext
}

func NewEmptyAssignstmtContext() *AssignstmtContext {
	var p = new(AssignstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_assignstmt
	return p
}

func InitEmptyAssignstmtContext(p *AssignstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_assignstmt
}

func (*AssignstmtContext) IsAssignstmtContext() {}

func NewAssignstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignstmtContext {
	var p = new(AssignstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_assignstmt

	return p
}

func (s *AssignstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *AssignstmtContext) GetOp() antlr.Token { return s.op }

func (s *AssignstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AssignstmtContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignstmtContext) Get_expr() IExprContext { return s._expr }

func (s *AssignstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *AssignstmtContext) GetAsg() interfaces.Instruction { return s.asg }

func (s *AssignstmtContext) SetAsg(v interfaces.Instruction) { s.asg = v }

func (s *AssignstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *AssignstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignstmtContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *AssignstmtContext) SUB_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB_IG, 0)
}

func (s *AssignstmtContext) SUM_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUM_IG, 0)
}

func (s *AssignstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterAssignstmt(s)
	}
}

func (s *AssignstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitAssignstmt(s)
	}
}

func (p *SwiftGrammarParser) Assignstmt() (localctx IAssignstmtContext) {
	localctx = NewAssignstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SwiftGrammarParserRULE_assignstmt)
	var _la int

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AssignstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)

			var _m = p.Match(SwiftGrammarParserIG)

			localctx.(*AssignstmtContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)

			var _x = p.expr(0)

			localctx.(*AssignstmtContext)._expr = _x
		}
		localctx.(*AssignstmtContext).asg = instructions.NewAssignment((func() int {
			if localctx.(*AssignstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AssignstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*AssignstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AssignstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AssignstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AssignstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*AssignstmtContext).Get_expr().GetE())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*AssignstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AssignstmtContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserSUM_IG || _la == SwiftGrammarParserSUB_IG) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AssignstmtContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(200)

			var _x = p.expr(0)

			localctx.(*AssignstmtContext)._expr = _x
		}
		localctx.(*AssignstmtContext).SetAsg(instructions.NewImplicitAssignment((func() int {
			if localctx.(*AssignstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AssignstmtContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*AssignstmtContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AssignstmtContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*AssignstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AssignstmtContext).Get_ID().GetText()
			}
		}()), (func() string {
			if localctx.(*AssignstmtContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*AssignstmtContext).GetOp().GetText()
			}
		}()), localctx.(*AssignstmtContext).Get_expr().GetE()))

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

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FOR returns the _FOR token.
	Get_FOR() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FOR sets the _FOR token.
	Set_FOR(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetFr returns the fr attribute.
	GetFr() interfaces.Instruction

	// SetFr sets the fr attribute.
	SetFr(interfaces.Instruction)

	// Getter signatures
	FOR() antlr.TerminalNode
	ID() antlr.TerminalNode
	IN() antlr.TerminalNode
	Expr() IExprContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	fr     interfaces.Instruction
	_FOR   antlr.Token
	_ID    antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) Get_FOR() antlr.Token { return s._FOR }

func (s *ForstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *ForstmtContext) Set_FOR(v antlr.Token) { s._FOR = v }

func (s *ForstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ForstmtContext) Get_expr() IExprContext { return s._expr }

func (s *ForstmtContext) Get_block() IBlockContext { return s._block }

func (s *ForstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ForstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *ForstmtContext) GetFr() interfaces.Instruction { return s.fr }

func (s *ForstmtContext) SetFr(v interfaces.Instruction) { s.fr = v }

func (s *ForstmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFOR, 0)
}

func (s *ForstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ForstmtContext) IN() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIN, 0)
}

func (s *ForstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *ForstmtContext) Block() IBlockContext {
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

func (s *ForstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterForstmt(s)
	}
}

func (s *ForstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitForstmt(s)
	}
}

func (p *SwiftGrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SwiftGrammarParserRULE_forstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)

		var _m = p.Match(SwiftGrammarParserFOR)

		localctx.(*ForstmtContext)._FOR = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*ForstmtContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Match(SwiftGrammarParserIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)

		var _x = p.expr(0)

		localctx.(*ForstmtContext)._expr = _x
	}
	{
		p.SetState(209)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)

		var _x = p.Block()

		localctx.(*ForstmtContext)._block = _x
	}
	{
		p.SetState(211)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ForstmtContext).fr = instructions.NewForIn((func() int {
		if localctx.(*ForstmtContext).Get_FOR() == nil {
			return 0
		} else {
			return localctx.(*ForstmtContext).Get_FOR().GetLine()
		}
	}()), (func() int {
		if localctx.(*ForstmtContext).Get_FOR() == nil {
			return 0
		} else {
			return localctx.(*ForstmtContext).Get_FOR().GetColumn()
		}
	}()), (func() string {
		if localctx.(*ForstmtContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ForstmtContext).Get_ID().GetText()
		}
	}()), localctx.(*ForstmtContext).Get_expr().GetE(), localctx.(*ForstmtContext).Get_block().GetBlk())

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

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_GUARD returns the _GUARD token.
	Get_GUARD() antlr.Token

	// Set_GUARD sets the _GUARD token.
	Set_GUARD(antlr.Token)

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetGrd returns the grd attribute.
	GetGrd() interfaces.Instruction

	// SetGrd sets the grd attribute.
	SetGrd(interfaces.Instruction)

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	grd    interfaces.Instruction
	_GUARD antlr.Token
	_expr  IExprContext
	_block IBlockContext
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) Get_GUARD() antlr.Token { return s._GUARD }

func (s *GuardstmtContext) Set_GUARD(v antlr.Token) { s._GUARD = v }

func (s *GuardstmtContext) Get_expr() IExprContext { return s._expr }

func (s *GuardstmtContext) Get_block() IBlockContext { return s._block }

func (s *GuardstmtContext) Set_expr(v IExprContext) { s._expr = v }

func (s *GuardstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *GuardstmtContext) GetGrd() interfaces.Instruction { return s.grd }

func (s *GuardstmtContext) SetGrd(v interfaces.Instruction) { s.grd = v }

func (s *GuardstmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUARD, 0)
}

func (s *GuardstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserELSE, 0)
}

func (s *GuardstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *GuardstmtContext) Block() IBlockContext {
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

func (s *GuardstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterGuardstmt(s)
	}
}

func (s *GuardstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitGuardstmt(s)
	}
}

func (p *SwiftGrammarParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SwiftGrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)

		var _m = p.Match(SwiftGrammarParserGUARD)

		localctx.(*GuardstmtContext)._GUARD = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)

		var _x = p.expr(0)

		localctx.(*GuardstmtContext)._expr = _x
	}
	{
		p.SetState(216)
		p.Match(SwiftGrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)

		var _x = p.Block()

		localctx.(*GuardstmtContext)._block = _x
	}
	{
		p.SetState(219)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*GuardstmtContext).grd = instructions.NewGuard((func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetLine()
		}
	}()), (func() int {
		if localctx.(*GuardstmtContext).Get_GUARD() == nil {
			return 0
		} else {
			return localctx.(*GuardstmtContext).Get_GUARD().GetColumn()
		}
	}()), localctx.(*GuardstmtContext).Get_expr().GetE(), localctx.(*GuardstmtContext).Get_block().GetBlk())

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

// IBreakstmtContext is an interface to support dynamic dispatch.
type IBreakstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// GetBrk returns the brk attribute.
	GetBrk() interfaces.Instruction

	// SetBrk sets the brk attribute.
	SetBrk(interfaces.Instruction)

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakstmtContext differentiates from other interfaces.
	IsBreakstmtContext()
}

type BreakstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	brk    interfaces.Instruction
	_BREAK antlr.Token
}

func NewEmptyBreakstmtContext() *BreakstmtContext {
	var p = new(BreakstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breakstmt
	return p
}

func InitEmptyBreakstmtContext(p *BreakstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_breakstmt
}

func (*BreakstmtContext) IsBreakstmtContext() {}

func NewBreakstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakstmtContext {
	var p = new(BreakstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_breakstmt

	return p
}

func (s *BreakstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakstmtContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *BreakstmtContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *BreakstmtContext) GetBrk() interfaces.Instruction { return s.brk }

func (s *BreakstmtContext) SetBrk(v interfaces.Instruction) { s.brk = v }

func (s *BreakstmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBREAK, 0)
}

func (s *BreakstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterBreakstmt(s)
	}
}

func (s *BreakstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitBreakstmt(s)
	}
}

func (p *SwiftGrammarParser) Breakstmt() (localctx IBreakstmtContext) {
	localctx = NewBreakstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SwiftGrammarParserRULE_breakstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)

		var _m = p.Match(SwiftGrammarParserBREAK)

		localctx.(*BreakstmtContext)._BREAK = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*BreakstmtContext).brk = instructions.NewBreak((func() int {
		if localctx.(*BreakstmtContext).Get_BREAK() == nil {
			return 0
		} else {
			return localctx.(*BreakstmtContext).Get_BREAK().GetLine()
		}
	}()), (func() int {
		if localctx.(*BreakstmtContext).Get_BREAK() == nil {
			return 0
		} else {
			return localctx.(*BreakstmtContext).Get_BREAK().GetColumn()
		}
	}()), nil)

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

// IContinuestmtContext is an interface to support dynamic dispatch.
type IContinuestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// GetCnt returns the cnt attribute.
	GetCnt() interfaces.Instruction

	// SetCnt sets the cnt attribute.
	SetCnt(interfaces.Instruction)

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinuestmtContext differentiates from other interfaces.
	IsContinuestmtContext()
}

type ContinuestmtContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	cnt       interfaces.Instruction
	_CONTINUE antlr.Token
}

func NewEmptyContinuestmtContext() *ContinuestmtContext {
	var p = new(ContinuestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuestmt
	return p
}

func InitEmptyContinuestmtContext(p *ContinuestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_continuestmt
}

func (*ContinuestmtContext) IsContinuestmtContext() {}

func NewContinuestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuestmtContext {
	var p = new(ContinuestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_continuestmt

	return p
}

func (s *ContinuestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuestmtContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *ContinuestmtContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *ContinuestmtContext) GetCnt() interfaces.Instruction { return s.cnt }

func (s *ContinuestmtContext) SetCnt(v interfaces.Instruction) { s.cnt = v }

func (s *ContinuestmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCONTINUE, 0)
}

func (s *ContinuestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterContinuestmt(s)
	}
}

func (s *ContinuestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitContinuestmt(s)
	}
}

func (p *SwiftGrammarParser) Continuestmt() (localctx IContinuestmtContext) {
	localctx = NewContinuestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SwiftGrammarParserRULE_continuestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)

		var _m = p.Match(SwiftGrammarParserCONTINUE)

		localctx.(*ContinuestmtContext)._CONTINUE = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ContinuestmtContext).cnt = instructions.NewContinue((func() int {
		if localctx.(*ContinuestmtContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*ContinuestmtContext).Get_CONTINUE().GetLine()
		}
	}()), (func() int {
		if localctx.(*ContinuestmtContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*ContinuestmtContext).Get_CONTINUE().GetColumn()
		}
	}()))

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

// IStructCreationContext is an interface to support dynamic dispatch.
type IStructCreationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRUCT returns the _STRUCT token.
	Get_STRUCT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_STRUCT sets the _STRUCT token.
	Set_STRUCT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listStructDec returns the _listStructDec rule contexts.
	Get_listStructDec() IListStructDecContext

	// Set_listStructDec sets the _listStructDec rule contexts.
	Set_listStructDec(IListStructDecContext)

	// GetDec returns the dec attribute.
	GetDec() interfaces.Instruction

	// SetDec sets the dec attribute.
	SetDec(interfaces.Instruction)

	// Getter signatures
	STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode
	LLAVEIZQ() antlr.TerminalNode
	ListStructDec() IListStructDecContext
	LLAVEDER() antlr.TerminalNode

	// IsStructCreationContext differentiates from other interfaces.
	IsStructCreationContext()
}

type StructCreationContext struct {
	antlr.BaseParserRuleContext
	parser         antlr.Parser
	dec            interfaces.Instruction
	_STRUCT        antlr.Token
	_ID            antlr.Token
	_listStructDec IListStructDecContext
}

func NewEmptyStructCreationContext() *StructCreationContext {
	var p = new(StructCreationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structCreation
	return p
}

func InitEmptyStructCreationContext(p *StructCreationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_structCreation
}

func (*StructCreationContext) IsStructCreationContext() {}

func NewStructCreationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructCreationContext {
	var p = new(StructCreationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_structCreation

	return p
}

func (s *StructCreationContext) GetParser() antlr.Parser { return s.parser }

func (s *StructCreationContext) Get_STRUCT() antlr.Token { return s._STRUCT }

func (s *StructCreationContext) Get_ID() antlr.Token { return s._ID }

func (s *StructCreationContext) Set_STRUCT(v antlr.Token) { s._STRUCT = v }

func (s *StructCreationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StructCreationContext) Get_listStructDec() IListStructDecContext { return s._listStructDec }

func (s *StructCreationContext) Set_listStructDec(v IListStructDecContext) { s._listStructDec = v }

func (s *StructCreationContext) GetDec() interfaces.Instruction { return s.dec }

func (s *StructCreationContext) SetDec(v interfaces.Instruction) { s.dec = v }

func (s *StructCreationContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRUCT, 0)
}

func (s *StructCreationContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *StructCreationContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *StructCreationContext) ListStructDec() IListStructDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *StructCreationContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *StructCreationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructCreationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructCreationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterStructCreation(s)
	}
}

func (s *StructCreationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitStructCreation(s)
	}
}

func (p *SwiftGrammarParser) StructCreation() (localctx IStructCreationContext) {
	localctx = NewStructCreationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SwiftGrammarParserRULE_structCreation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)

		var _m = p.Match(SwiftGrammarParserSTRUCT)

		localctx.(*StructCreationContext)._STRUCT = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*StructCreationContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(SwiftGrammarParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)

		var _x = p.listStructDec(0)

		localctx.(*StructCreationContext)._listStructDec = _x
	}
	{
		p.SetState(232)
		p.Match(SwiftGrammarParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*StructCreationContext).dec = instructions.NewStruct((func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetLine()
		}
	}()), (func() int {
		if localctx.(*StructCreationContext).Get_STRUCT() == nil {
			return 0
		} else {
			return localctx.(*StructCreationContext).Get_STRUCT().GetColumn()
		}
	}()), (func() string {
		if localctx.(*StructCreationContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*StructCreationContext).Get_ID().GetText()
		}
	}()), localctx.(*StructCreationContext).Get_listStructDec().GetL())

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

// IListStructDecContext is an interface to support dynamic dispatch.
type IListStructDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetId1 returns the id1 token.
	GetId1() antlr.Token

	// GetId2 returns the id2 token.
	GetId2() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetId1 sets the id1 token.
	SetId1(antlr.Token)

	// SetId2 sets the id2 token.
	SetId2(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListStructDecContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// SetList sets the list rule contexts.
	SetList(IListStructDecContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Types() ITypesContext
	VAR() antlr.TerminalNode
	LET() antlr.TerminalNode
	ListStructDec() IListStructDecContext
	COMA() antlr.TerminalNode
	IG() antlr.TerminalNode
	Expr() IExprContext

	// IsListStructDecContext differentiates from other interfaces.
	IsListStructDecContext()
}

type ListStructDecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListStructDecContext
	_ID    antlr.Token
	_types ITypesContext
	id1    antlr.Token
	id2    antlr.Token
}

func NewEmptyListStructDecContext() *ListStructDecContext {
	var p = new(ListStructDecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec
	return p
}

func InitEmptyListStructDecContext(p *ListStructDecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec
}

func (*ListStructDecContext) IsListStructDecContext() {}

func NewListStructDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStructDecContext {
	var p = new(ListStructDecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listStructDec

	return p
}

func (s *ListStructDecContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStructDecContext) Get_ID() antlr.Token { return s._ID }

func (s *ListStructDecContext) GetId1() antlr.Token { return s.id1 }

func (s *ListStructDecContext) GetId2() antlr.Token { return s.id2 }

func (s *ListStructDecContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListStructDecContext) SetId1(v antlr.Token) { s.id1 = v }

func (s *ListStructDecContext) SetId2(v antlr.Token) { s.id2 = v }

func (s *ListStructDecContext) GetList() IListStructDecContext { return s.list }

func (s *ListStructDecContext) Get_types() ITypesContext { return s._types }

func (s *ListStructDecContext) SetList(v IListStructDecContext) { s.list = v }

func (s *ListStructDecContext) Set_types(v ITypesContext) { s._types = v }

func (s *ListStructDecContext) GetL() []interface{} { return s.l }

func (s *ListStructDecContext) SetL(v []interface{}) { s.l = v }

func (s *ListStructDecContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ListStructDecContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ListStructDecContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *ListStructDecContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ListStructDecContext) VAR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserVAR, 0)
}

func (s *ListStructDecContext) LET() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLET, 0)
}

func (s *ListStructDecContext) ListStructDec() IListStructDecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStructDecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStructDecContext)
}

func (s *ListStructDecContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListStructDecContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *ListStructDecContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListStructDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStructDecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStructDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListStructDec(s)
	}
}

func (s *ListStructDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListStructDec(s)
	}
}

func (p *SwiftGrammarParser) ListStructDec() (localctx IListStructDecContext) {
	return p.listStructDec(0)
}

func (p *SwiftGrammarParser) listStructDec(_p int) (localctx IListStructDecContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListStructDecContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListStructDecContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, SwiftGrammarParserRULE_listStructDec, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(236)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(237)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)

			var _x = p.Types()

			localctx.(*ListStructDecContext)._types = _x
		}

		var arr []interface{}
		newParams := environment.NewStructType((func() string {
			if localctx.(*ListStructDecContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ListStructDecContext).Get_ID().GetText()
			}
		}()), localctx.(*ListStructDecContext).Get_types().GetTy(), "")
		arr = append(arr, newParams)
		localctx.(*ListStructDecContext).l = arr

	case 2:
		{
			p.SetState(242)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(243)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext).id1 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(244)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ListStructDecContext).id2 = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		var arr []interface{}
		newParams := environment.NewStructType((func() string {
			if localctx.(*ListStructDecContext).GetId1() == nil {
				return ""
			} else {
				return localctx.(*ListStructDecContext).GetId1().GetText()
			}
		}()), environment.UNKNOWN, (func() string {
			if localctx.(*ListStructDecContext).GetId2() == nil {
				return ""
			} else {
				return localctx.(*ListStructDecContext).GetId2().GetText()
			}
		}()))
		arr = append(localctx.(*ListStructDecContext).GetList().GetL(), newParams)
		localctx.(*ListStructDecContext).l = arr

	case 3:
		localctx.(*ListStructDecContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(283)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListStructDecContext(p, _parentctx, _parentState)
				localctx.(*ListStructDecContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructDec)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				p.SetState(252)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SwiftGrammarParserCOMA {
					{
						p.SetState(251)
						p.Match(SwiftGrammarParserCOMA)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(254)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(255)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext)._ID = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(256)
					p.Match(SwiftGrammarParserD_PTS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(257)

					var _x = p.Types()

					localctx.(*ListStructDecContext)._types = _x
				}

				var arr []interface{}
				newParams := environment.NewStructType((func() string {
					if localctx.(*ListStructDecContext).Get_ID() == nil {
						return ""
					} else {
						return localctx.(*ListStructDecContext).Get_ID().GetText()
					}
				}()), localctx.(*ListStructDecContext).Get_types().GetTy(), "")
				arr = append(localctx.(*ListStructDecContext).GetList().GetL(), newParams)
				localctx.(*ListStructDecContext).l = arr

			case 2:
				localctx = NewListStructDecContext(p, _parentctx, _parentState)
				localctx.(*ListStructDecContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructDec)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				p.SetState(262)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SwiftGrammarParserCOMA {
					{
						p.SetState(261)
						p.Match(SwiftGrammarParserCOMA)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(264)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(265)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext).id1 = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(266)
					p.Match(SwiftGrammarParserD_PTS)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(267)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext).id2 = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				var arr []interface{}
				newParams := environment.NewStructType((func() string {
					if localctx.(*ListStructDecContext).GetId1() == nil {
						return ""
					} else {
						return localctx.(*ListStructDecContext).GetId1().GetText()
					}
				}()), environment.UNKNOWN, (func() string {
					if localctx.(*ListStructDecContext).GetId2() == nil {
						return ""
					} else {
						return localctx.(*ListStructDecContext).GetId2().GetText()
					}
				}()))
				arr = append(localctx.(*ListStructDecContext).GetList().GetL(), newParams)
				localctx.(*ListStructDecContext).l = arr

			case 3:
				localctx = NewListStructDecContext(p, _parentctx, _parentState)
				localctx.(*ListStructDecContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listStructDec)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				p.SetState(271)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SwiftGrammarParserCOMA {
					{
						p.SetState(270)
						p.Match(SwiftGrammarParserCOMA)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(273)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserVAR || _la == SwiftGrammarParserLET) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(274)

					var _m = p.Match(SwiftGrammarParserID)

					localctx.(*ListStructDecContext).id1 = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(277)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == SwiftGrammarParserD_PTS {
					{
						p.SetState(275)
						p.Match(SwiftGrammarParserD_PTS)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}
					{
						p.SetState(276)

						var _x = p.Types()

						localctx.(*ListStructDecContext)._types = _x
					}

				}
				{
					p.SetState(279)
					p.Match(SwiftGrammarParserIG)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(280)
					p.expr(0)
				}

				var arr []interface{}
				newParams := environment.NewStructType((func() string {
					if localctx.(*ListStructDecContext).GetId1() == nil {
						return ""
					} else {
						return localctx.(*ListStructDecContext).GetId1().GetText()
					}
				}()), localctx.(*ListStructDecContext).Get_types().GetTy(), "")
				arr = append(localctx.(*ListStructDecContext).GetList().GetL(), newParams)
				localctx.(*ListStructDecContext).l = arr

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

// IFnstmtContext is an interface to support dynamic dispatch.
type IFnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FUNC returns the _FUNC token.
	Get_FUNC() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FUNC sets the _FUNC token.
	Set_FUNC(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsFunc returns the _listParamsFunc rule contexts.
	Get_listParamsFunc() IListParamsFuncContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Get_block returns the _block rule contexts.
	Get_block() IBlockContext

	// Set_listParamsFunc sets the _listParamsFunc rule contexts.
	Set_listParamsFunc(IListParamsFuncContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// Set_block sets the _block rule contexts.
	Set_block(IBlockContext)

	// GetFn returns the fn attribute.
	GetFn() interfaces.Instruction

	// SetFn sets the fn attribute.
	SetFn(interfaces.Instruction)

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	ListParamsFunc() IListParamsFuncContext
	PARDER() antlr.TerminalNode
	FLECHA() antlr.TerminalNode
	Types() ITypesContext
	LLAVEIZQ() antlr.TerminalNode
	Block() IBlockContext
	LLAVEDER() antlr.TerminalNode

	// IsFnstmtContext differentiates from other interfaces.
	IsFnstmtContext()
}

type FnstmtContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	fn              interfaces.Instruction
	_FUNC           antlr.Token
	_ID             antlr.Token
	_listParamsFunc IListParamsFuncContext
	_types          ITypesContext
	_block          IBlockContext
}

func NewEmptyFnstmtContext() *FnstmtContext {
	var p = new(FnstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_fnstmt
	return p
}

func InitEmptyFnstmtContext(p *FnstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_fnstmt
}

func (*FnstmtContext) IsFnstmtContext() {}

func NewFnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnstmtContext {
	var p = new(FnstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_fnstmt

	return p
}

func (s *FnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FnstmtContext) Get_FUNC() antlr.Token { return s._FUNC }

func (s *FnstmtContext) Get_ID() antlr.Token { return s._ID }

func (s *FnstmtContext) Set_FUNC(v antlr.Token) { s._FUNC = v }

func (s *FnstmtContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FnstmtContext) Get_listParamsFunc() IListParamsFuncContext { return s._listParamsFunc }

func (s *FnstmtContext) Get_types() ITypesContext { return s._types }

func (s *FnstmtContext) Get_block() IBlockContext { return s._block }

func (s *FnstmtContext) Set_listParamsFunc(v IListParamsFuncContext) { s._listParamsFunc = v }

func (s *FnstmtContext) Set_types(v ITypesContext) { s._types = v }

func (s *FnstmtContext) Set_block(v IBlockContext) { s._block = v }

func (s *FnstmtContext) GetFn() interfaces.Instruction { return s.fn }

func (s *FnstmtContext) SetFn(v interfaces.Instruction) { s.fn = v }

func (s *FnstmtContext) FUNC() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFUNC, 0)
}

func (s *FnstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *FnstmtContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *FnstmtContext) ListParamsFunc() IListParamsFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *FnstmtContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *FnstmtContext) FLECHA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLECHA, 0)
}

func (s *FnstmtContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *FnstmtContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEIZQ, 0)
}

func (s *FnstmtContext) Block() IBlockContext {
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

func (s *FnstmtContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserLLAVEDER, 0)
}

func (s *FnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterFnstmt(s)
	}
}

func (s *FnstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitFnstmt(s)
	}
}

func (p *SwiftGrammarParser) Fnstmt() (localctx IFnstmtContext) {
	localctx = NewFnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SwiftGrammarParserRULE_fnstmt)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(288)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FnstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FnstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(291)

			var _x = p.listParamsFunc(0)

			localctx.(*FnstmtContext)._listParamsFunc = _x
		}
		{
			p.SetState(292)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.Match(SwiftGrammarParserFLECHA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(294)

			var _x = p.Types()

			localctx.(*FnstmtContext)._types = _x
		}
		{
			p.SetState(295)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)

			var _x = p.Block()

			localctx.(*FnstmtContext)._block = _x
		}
		{
			p.SetState(297)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FnstmtContext).fn = instructions.NewFunction((func() int {
			if localctx.(*FnstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FnstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FnstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FnstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FnstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FnstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*FnstmtContext).Get_listParamsFunc().GetL(), localctx.(*FnstmtContext).Get_types().GetTy(), localctx.(*FnstmtContext).Get_block().GetBlk())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)

			var _m = p.Match(SwiftGrammarParserFUNC)

			localctx.(*FnstmtContext)._FUNC = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*FnstmtContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)

			var _x = p.listParamsFunc(0)

			localctx.(*FnstmtContext)._listParamsFunc = _x
		}
		{
			p.SetState(304)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(SwiftGrammarParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)

			var _x = p.Block()

			localctx.(*FnstmtContext)._block = _x
		}
		{
			p.SetState(307)
			p.Match(SwiftGrammarParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*FnstmtContext).fn = instructions.NewFunction((func() int {
			if localctx.(*FnstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FnstmtContext).Get_FUNC().GetLine()
			}
		}()), (func() int {
			if localctx.(*FnstmtContext).Get_FUNC() == nil {
				return 0
			} else {
				return localctx.(*FnstmtContext).Get_FUNC().GetColumn()
			}
		}()), (func() string {
			if localctx.(*FnstmtContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FnstmtContext).Get_ID().GetText()
			}
		}()), localctx.(*FnstmtContext).Get_listParamsFunc().GetL(), environment.NIL, localctx.(*FnstmtContext).Get_block().GetBlk())

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

// IListParamsFuncContext is an interface to support dynamic dispatch.
type IListParamsFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsFuncContext

	// Get_parametro returns the _parametro rule contexts.
	Get_parametro() IParametroContext

	// SetList sets the list rule contexts.
	SetList(IListParamsFuncContext)

	// Set_parametro sets the _parametro rule contexts.
	Set_parametro(IParametroContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Parametro() IParametroContext
	COMA() antlr.TerminalNode
	ListParamsFunc() IListParamsFuncContext

	// IsListParamsFuncContext differentiates from other interfaces.
	IsListParamsFuncContext()
}

type ListParamsFuncContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	l          []interface{}
	list       IListParamsFuncContext
	_parametro IParametroContext
}

func NewEmptyListParamsFuncContext() *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc
	return p
}

func InitEmptyListParamsFuncContext(p *ListParamsFuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc
}

func (*ListParamsFuncContext) IsListParamsFuncContext() {}

func NewListParamsFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsFuncContext {
	var p = new(ListParamsFuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParamsFunc

	return p
}

func (s *ListParamsFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsFuncContext) GetList() IListParamsFuncContext { return s.list }

func (s *ListParamsFuncContext) Get_parametro() IParametroContext { return s._parametro }

func (s *ListParamsFuncContext) SetList(v IListParamsFuncContext) { s.list = v }

func (s *ListParamsFuncContext) Set_parametro(v IParametroContext) { s._parametro = v }

func (s *ListParamsFuncContext) GetL() []interface{} { return s.l }

func (s *ListParamsFuncContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsFuncContext) Parametro() IParametroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *ListParamsFuncContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsFuncContext) ListParamsFunc() IListParamsFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsFuncContext)
}

func (s *ListParamsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParamsFunc(s)
	}
}

func (s *ListParamsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParamsFunc(s)
	}
}

func (p *SwiftGrammarParser) ListParamsFunc() (localctx IListParamsFuncContext) {
	return p.listParamsFunc(0)
}

func (p *SwiftGrammarParser) listParamsFunc(_p int) (localctx IListParamsFuncContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsFuncContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsFuncContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, SwiftGrammarParserRULE_listParamsFunc, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(313)

			var _x = p.Parametro()

			localctx.(*ListParamsFuncContext)._parametro = _x
		}

		localctx.(*ListParamsFuncContext).l = []interface{}{}
		localctx.(*ListParamsFuncContext).l = append(localctx.(*ListParamsFuncContext).l, localctx.(*ListParamsFuncContext).Get_parametro().GetP())

	case 2:

		localctx.(*ListParamsFuncContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsFuncContext(p, _parentctx, _parentState)
			localctx.(*ListParamsFuncContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsFunc)
			p.SetState(319)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(320)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(321)

				var _x = p.Parametro()

				localctx.(*ListParamsFuncContext)._parametro = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsFuncContext).GetList().GetL(), localctx.(*ListParamsFuncContext).Get_parametro().GetP())
			localctx.(*ListParamsFuncContext).l = arr

		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
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

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetExte returns the exte token.
	GetExte() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetExte sets the exte token.
	SetExte(antlr.Token)

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// GetP returns the p attribute.
	GetP() interfaces.Instruction

	// SetP sets the p attribute.
	SetP(interfaces.Instruction)

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	D_PTS() antlr.TerminalNode
	Types() ITypesContext
	GUIONBAJO() antlr.TerminalNode

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.Instruction
	_ID    antlr.Token
	_types ITypesContext
	exte   antlr.Token
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_parametro
	return p
}

func InitEmptyParametroContext(p *ParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_parametro
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) Get_ID() antlr.Token { return s._ID }

func (s *ParametroContext) GetExte() antlr.Token { return s.exte }

func (s *ParametroContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ParametroContext) SetExte(v antlr.Token) { s.exte = v }

func (s *ParametroContext) Get_types() ITypesContext { return s._types }

func (s *ParametroContext) Set_types(v ITypesContext) { s._types = v }

func (s *ParametroContext) GetP() interfaces.Instruction { return s.p }

func (s *ParametroContext) SetP(v interfaces.Instruction) { s.p = v }

func (s *ParametroContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserID)
}

func (s *ParametroContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, i)
}

func (s *ParametroContext) D_PTS() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserD_PTS, 0)
}

func (s *ParametroContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ParametroContext) GUIONBAJO() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserGUIONBAJO, 0)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (p *SwiftGrammarParser) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SwiftGrammarParserRULE_parametro)
	var _la int

	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ParametroContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)

			var _x = p.Types()

			localctx.(*ParametroContext)._types = _x
		}
		localctx.(*ParametroContext).p = instructions.NewParamsDeclaration((func() int {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParametroContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParametroContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParametroContext).Get_ID().GetText()
			}
		}()), localctx.(*ParametroContext).Get_types().GetTy())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ParametroContext).exte = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SwiftGrammarParserID || _la == SwiftGrammarParserGUIONBAJO) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ParametroContext).exte = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(335)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ParametroContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(SwiftGrammarParserD_PTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)

			var _x = p.Types()

			localctx.(*ParametroContext)._types = _x
		}
		localctx.(*ParametroContext).p = instructions.NewParamsDeclaration((func() int {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParametroContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParametroContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParametroContext).Get_ID().GetText()
			}
		}()), localctx.(*ParametroContext).Get_types().GetTy())

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

// ICallExpContext is an interface to support dynamic dispatch.
type ICallExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsCall returns the _listParamsCall rule contexts.
	Get_listParamsCall() IListParamsCallContext

	// Set_listParamsCall sets the _listParamsCall rule contexts.
	Set_listParamsCall(IListParamsCallContext)

	// GetCfe returns the cfe attribute.
	GetCfe() interfaces.Expression

	// SetCfe sets the cfe attribute.
	SetCfe(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	ListParamsCall() IListParamsCallContext
	PARDER() antlr.TerminalNode

	// IsCallExpContext differentiates from other interfaces.
	IsCallExpContext()
}

type CallExpContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	cfe             interfaces.Expression
	_ID             antlr.Token
	_listParamsCall IListParamsCallContext
}

func NewEmptyCallExpContext() *CallExpContext {
	var p = new(CallExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callExp
	return p
}

func InitEmptyCallExpContext(p *CallExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callExp
}

func (*CallExpContext) IsCallExpContext() {}

func NewCallExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallExpContext {
	var p = new(CallExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_callExp

	return p
}

func (s *CallExpContext) GetParser() antlr.Parser { return s.parser }

func (s *CallExpContext) Get_ID() antlr.Token { return s._ID }

func (s *CallExpContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallExpContext) Get_listParamsCall() IListParamsCallContext { return s._listParamsCall }

func (s *CallExpContext) Set_listParamsCall(v IListParamsCallContext) { s._listParamsCall = v }

func (s *CallExpContext) GetCfe() interfaces.Expression { return s.cfe }

func (s *CallExpContext) SetCfe(v interfaces.Expression) { s.cfe = v }

func (s *CallExpContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CallExpContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *CallExpContext) ListParamsCall() IListParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *CallExpContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *CallExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCallExp(s)
	}
}

func (s *CallExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCallExp(s)
	}
}

func (p *SwiftGrammarParser) CallExp() (localctx ICallExpContext) {
	localctx = NewCallExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SwiftGrammarParserRULE_callExp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CallExpContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(343)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)

		var _x = p.listParamsCall(0)

		localctx.(*CallExpContext)._listParamsCall = _x
	}
	{
		p.SetState(345)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CallExpContext).cfe = expressions.NewCallExp((func() int {
		if localctx.(*CallExpContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallExpContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*CallExpContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallExpContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CallExpContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*CallExpContext).Get_ID().GetText()
		}
	}()), localctx.(*CallExpContext).Get_listParamsCall().GetL())

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

// ICallFunctionContext is an interface to support dynamic dispatch.
type ICallFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listParamsCall returns the _listParamsCall rule contexts.
	Get_listParamsCall() IListParamsCallContext

	// Set_listParamsCall sets the _listParamsCall rule contexts.
	Set_listParamsCall(IListParamsCallContext)

	// GetCf returns the cf attribute.
	GetCf() interfaces.Instruction

	// SetCf sets the cf attribute.
	SetCf(interfaces.Instruction)

	// Getter signatures
	ID() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	ListParamsCall() IListParamsCallContext
	PARDER() antlr.TerminalNode

	// IsCallFunctionContext differentiates from other interfaces.
	IsCallFunctionContext()
}

type CallFunctionContext struct {
	antlr.BaseParserRuleContext
	parser          antlr.Parser
	cf              interfaces.Instruction
	_ID             antlr.Token
	_listParamsCall IListParamsCallContext
}

func NewEmptyCallFunctionContext() *CallFunctionContext {
	var p = new(CallFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFunction
	return p
}

func InitEmptyCallFunctionContext(p *CallFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_callFunction
}

func (*CallFunctionContext) IsCallFunctionContext() {}

func NewCallFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFunctionContext {
	var p = new(CallFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_callFunction

	return p
}

func (s *CallFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *CallFunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *CallFunctionContext) Get_listParamsCall() IListParamsCallContext { return s._listParamsCall }

func (s *CallFunctionContext) Set_listParamsCall(v IListParamsCallContext) { s._listParamsCall = v }

func (s *CallFunctionContext) GetCf() interfaces.Instruction { return s.cf }

func (s *CallFunctionContext) SetCf(v interfaces.Instruction) { s.cf = v }

func (s *CallFunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *CallFunctionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *CallFunctionContext) ListParamsCall() IListParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *CallFunctionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *CallFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterCallFunction(s)
	}
}

func (s *CallFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitCallFunction(s)
	}
}

func (p *SwiftGrammarParser) CallFunction() (localctx ICallFunctionContext) {
	localctx = NewCallFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SwiftGrammarParserRULE_callFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(348)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*CallFunctionContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Match(SwiftGrammarParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(350)

		var _x = p.listParamsCall(0)

		localctx.(*CallFunctionContext)._listParamsCall = _x
	}
	{
		p.SetState(351)
		p.Match(SwiftGrammarParserPARDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*CallFunctionContext).cf = instructions.NewCallFunc((func() int {
		if localctx.(*CallFunctionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFunctionContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*CallFunctionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*CallFunctionContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*CallFunctionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*CallFunctionContext).Get_ID().GetText()
		}
	}()), localctx.(*CallFunctionContext).Get_listParamsCall().GetL())

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

// IListParamsCallContext is an interface to support dynamic dispatch.
type IListParamsCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsCallContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsCallContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParamsCall() IListParamsCallContext

	// IsListParamsCallContext differentiates from other interfaces.
	IsListParamsCallContext()
}

type ListParamsCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsCallContext
	_expr  IExprContext
}

func NewEmptyListParamsCallContext() *ListParamsCallContext {
	var p = new(ListParamsCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall
	return p
}

func InitEmptyListParamsCallContext(p *ListParamsCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall
}

func (*ListParamsCallContext) IsListParamsCallContext() {}

func NewListParamsCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsCallContext {
	var p = new(ListParamsCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParamsCall

	return p
}

func (s *ListParamsCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsCallContext) GetList() IListParamsCallContext { return s.list }

func (s *ListParamsCallContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsCallContext) SetList(v IListParamsCallContext) { s.list = v }

func (s *ListParamsCallContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsCallContext) GetL() []interface{} { return s.l }

func (s *ListParamsCallContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsCallContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsCallContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsCallContext) ListParamsCall() IListParamsCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsCallContext)
}

func (s *ListParamsCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParamsCall(s)
	}
}

func (s *ListParamsCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParamsCall(s)
	}
}

func (p *SwiftGrammarParser) ListParamsCall() (localctx IListParamsCallContext) {
	return p.listParamsCall(0)
}

func (p *SwiftGrammarParser) listParamsCall(_p int) (localctx IListParamsCallContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsCallContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsCallContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, SwiftGrammarParserRULE_listParamsCall, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(355)

			var _x = p.expr(0)

			localctx.(*ListParamsCallContext)._expr = _x
		}

		localctx.(*ListParamsCallContext).l = []interface{}{}
		localctx.(*ListParamsCallContext).l = append(localctx.(*ListParamsCallContext).l, localctx.(*ListParamsCallContext).Get_expr().GetE())

	case 2:

		localctx.(*ListParamsCallContext).l = []interface{}{}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsCallContext(p, _parentctx, _parentState)
			localctx.(*ListParamsCallContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParamsCall)
			p.SetState(361)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(362)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(363)

				var _x = p.expr(0)

				localctx.(*ListParamsCallContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsCallContext).GetList().GetL(), localctx.(*ListParamsCallContext).Get_expr().GetE())
			localctx.(*ListParamsCallContext).l = arr

		}
		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
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

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTy returns the ty attribute.
	GetTy() environment.TipoExpresion

	// SetTy sets the ty attribute.
	SetTy(environment.TipoExpresion)

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	STR() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Types() ITypesContext
	CORDER() antlr.TerminalNode
	AllCOMILLA() []antlr.TerminalNode
	COMILLA(i int) antlr.TerminalNode
	NIL() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     environment.TipoExpresion
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_types
	return p
}

func InitEmptyTypesContext(p *TypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_types
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) GetTy() environment.TipoExpresion { return s.ty }

func (s *TypesContext) SetTy(v environment.TipoExpresion) { s.ty = v }

func (s *TypesContext) INT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserINT, 0)
}

func (s *TypesContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFLOAT, 0)
}

func (s *TypesContext) STR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTR, 0)
}

func (s *TypesContext) BOOL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserBOOL, 0)
}

func (s *TypesContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *TypesContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *TypesContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *TypesContext) AllCOMILLA() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserCOMILLA)
}

func (s *TypesContext) COMILLA(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMILLA, i)
}

func (s *TypesContext) NIL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNIL, 0)
}

func (s *TypesContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRUCT, 0)
}

func (s *TypesContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *SwiftGrammarParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SwiftGrammarParserRULE_types)
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SwiftGrammarParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.Match(SwiftGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.INTEGER

	case SwiftGrammarParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.Match(SwiftGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.FLOAT

	case SwiftGrammarParserSTR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(375)
			p.Match(SwiftGrammarParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.STRING

	case SwiftGrammarParserBOOL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(377)
			p.Match(SwiftGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.BOOLEAN

	case SwiftGrammarParserCORIZQ:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(379)
			p.Match(SwiftGrammarParserCORIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Types()
		}
		{
			p.SetState(381)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.ARRAY

	case SwiftGrammarParserCOMILLA:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(384)
			p.Match(SwiftGrammarParserCOMILLA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Match(SwiftGrammarParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.Match(SwiftGrammarParserCOMILLA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.STR

	case SwiftGrammarParserNIL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(388)
			p.Match(SwiftGrammarParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.NIL

	case SwiftGrammarParserSTRUCT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(390)
			p.Match(SwiftGrammarParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.STRUCT

	case SwiftGrammarParserID:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(392)
			p.Match(SwiftGrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*TypesContext).ty = environment.UNKNOWN

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SUB returns the _SUB token.
	Get_SUB() antlr.Token

	// Get_NOT returns the _NOT token.
	Get_NOT() antlr.Token

	// Get_CORIZQ returns the _CORIZQ token.
	Get_CORIZQ() antlr.Token

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRU returns the _TRU token.
	Get_TRU() antlr.Token

	// Get_FAL returns the _FAL token.
	Get_FAL() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_NIL returns the _NIL token.
	Get_NIL() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_SUB sets the _SUB token.
	Set_SUB(antlr.Token)

	// Set_NOT sets the _NOT token.
	Set_NOT(antlr.Token)

	// Set_CORIZQ sets the _CORIZQ token.
	Set_CORIZQ(antlr.Token)

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRU sets the _TRU token.
	Set_TRU(antlr.Token)

	// Set_FAL sets the _FAL token.
	Set_FAL(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_NIL sets the _NIL token.
	Set_NIL(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetExpuno returns the expuno rule contexts.
	GetExpuno() IExprContext

	// GetLeft returns the left rule contexts.
	GetLeft() IExprContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// Get_callExp returns the _callExp rule contexts.
	Get_callExp() ICallExpContext

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// GetExpdos returns the expdos rule contexts.
	GetExpdos() IExprContext

	// SetExpuno sets the expuno rule contexts.
	SetExpuno(IExprContext)

	// SetLeft sets the left rule contexts.
	SetLeft(IExprContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// Set_callExp sets the _callExp rule contexts.
	Set_callExp(ICallExpContext)

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// SetExpdos sets the expdos rule contexts.
	SetExpdos(IExprContext)

	// GetE returns the e attribute.
	GetE() interfaces.Expression

	// SetE sets the e attribute.
	SetE(interfaces.Expression)

	// Getter signatures
	SUB() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	NOT() antlr.TerminalNode
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	CallExp() ICallExpContext
	CORIZQ() antlr.TerminalNode
	CORDER() antlr.TerminalNode
	ListArray() IListArrayContext
	ListParams() IListParamsContext
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TRU() antlr.TerminalNode
	FAL() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllPUNTO() []antlr.TerminalNode
	PUNTO(i int) antlr.TerminalNode
	COUNT() antlr.TerminalNode
	NIL() antlr.TerminalNode
	SUB_IG() antlr.TerminalNode
	SUM_IG() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	ADD() antlr.TerminalNode
	MAY_IG() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MEN_IG() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	IG_IG() antlr.TerminalNode
	DIF() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	e           interfaces.Expression
	expuno      IExprContext
	left        IExprContext
	_SUB        antlr.Token
	opDe        IExprContext
	_expr       IExprContext
	_NOT        antlr.Token
	right       IExprContext
	_callExp    ICallExpContext
	_CORIZQ     antlr.Token
	list        IListArrayContext
	_listParams IListParamsContext
	_NUMBER     antlr.Token
	_STRING     antlr.Token
	_TRU        antlr.Token
	_FAL        antlr.Token
	_ID         antlr.Token
	_NIL        antlr.Token
	expdos      IExprContext
	op          antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Get_SUB() antlr.Token { return s._SUB }

func (s *ExprContext) Get_NOT() antlr.Token { return s._NOT }

func (s *ExprContext) Get_CORIZQ() antlr.Token { return s._CORIZQ }

func (s *ExprContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ExprContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ExprContext) Get_TRU() antlr.Token { return s._TRU }

func (s *ExprContext) Get_FAL() antlr.Token { return s._FAL }

func (s *ExprContext) Get_ID() antlr.Token { return s._ID }

func (s *ExprContext) Get_NIL() antlr.Token { return s._NIL }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) Set_SUB(v antlr.Token) { s._SUB = v }

func (s *ExprContext) Set_NOT(v antlr.Token) { s._NOT = v }

func (s *ExprContext) Set_CORIZQ(v antlr.Token) { s._CORIZQ = v }

func (s *ExprContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ExprContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ExprContext) Set_TRU(v antlr.Token) { s._TRU = v }

func (s *ExprContext) Set_FAL(v antlr.Token) { s._FAL = v }

func (s *ExprContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExprContext) Set_NIL(v antlr.Token) { s._NIL = v }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) GetExpuno() IExprContext { return s.expuno }

func (s *ExprContext) GetLeft() IExprContext { return s.left }

func (s *ExprContext) GetOpDe() IExprContext { return s.opDe }

func (s *ExprContext) Get_expr() IExprContext { return s._expr }

func (s *ExprContext) GetRight() IExprContext { return s.right }

func (s *ExprContext) Get_callExp() ICallExpContext { return s._callExp }

func (s *ExprContext) GetList() IListArrayContext { return s.list }

func (s *ExprContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *ExprContext) GetExpdos() IExprContext { return s.expdos }

func (s *ExprContext) SetExpuno(v IExprContext) { s.expuno = v }

func (s *ExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *ExprContext) SetOpDe(v IExprContext) { s.opDe = v }

func (s *ExprContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprContext) Set_callExp(v ICallExpContext) { s._callExp = v }

func (s *ExprContext) SetList(v IListArrayContext) { s.list = v }

func (s *ExprContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *ExprContext) SetExpdos(v IExprContext) { s.expdos = v }

func (s *ExprContext) GetE() interfaces.Expression { return s.e }

func (s *ExprContext) SetE(v interfaces.Expression) { s.e = v }

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNOT, 0)
}

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARIZQ, 0)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPARDER, 0)
}

func (s *ExprContext) CallExp() ICallExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallExpContext)
}

func (s *ExprContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ExprContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ExprContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ExprContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNUMBER, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSTRING, 0)
}

func (s *ExprContext) TRU() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserTRU, 0)
}

func (s *ExprContext) FAL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserFAL, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ExprContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(SwiftGrammarParserPUNTO)
}

func (s *ExprContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserPUNTO, i)
}

func (s *ExprContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOUNT, 0)
}

func (s *ExprContext) NIL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserNIL, 0)
}

func (s *ExprContext) SUB_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUB_IG, 0)
}

func (s *ExprContext) SUM_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserSUM_IG, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIV, 0)
}

func (s *ExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMOD, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserADD, 0)
}

func (s *ExprContext) MAY_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAY_IG, 0)
}

func (s *ExprContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMAYOR, 0)
}

func (s *ExprContext) MEN_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMEN_IG, 0)
}

func (s *ExprContext) MENOR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserMENOR, 0)
}

func (s *ExprContext) IG_IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG_IG, 0)
}

func (s *ExprContext) DIF() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserDIF, 0)
}

func (s *ExprContext) AND() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserAND, 0)
}

func (s *ExprContext) OR() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserOR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SwiftGrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *SwiftGrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, SwiftGrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(397)

			var _m = p.Match(SwiftGrammarParserSUB)

			localctx.(*ExprContext)._SUB = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)

			var _x = p.expr(22)

			localctx.(*ExprContext).opDe = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation((func() int {
			if localctx.(*ExprContext).Get_SUB() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_SUB().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_SUB() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_SUB().GetColumn()
			}
		}()), localctx.(*ExprContext).GetOpDe().GetE(), "NEGACION", nil)

	case 2:
		{
			p.SetState(401)

			var _m = p.Match(SwiftGrammarParserNOT)

			localctx.(*ExprContext)._NOT = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)

			var _x = p.expr(12)

			localctx.(*ExprContext).right = _x
			localctx.(*ExprContext)._expr = _x
		}
		localctx.(*ExprContext).e = expressions.NewOperation((func() int {
			if localctx.(*ExprContext).Get_NOT() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NOT().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_NOT() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NOT().GetColumn()
			}
		}()), localctx.(*ExprContext).GetRight().GetE(), (func() string {
			if localctx.(*ExprContext).Get_NOT() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NOT().GetText()
			}
		}()), nil)

	case 3:
		{
			p.SetState(405)
			p.Match(SwiftGrammarParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)

			var _x = p.expr(0)

			localctx.(*ExprContext)._expr = _x
		}
		{
			p.SetState(407)
			p.Match(SwiftGrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_expr().GetE()

	case 4:
		{
			p.SetState(410)

			var _x = p.CallExp()

			localctx.(*ExprContext)._callExp = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).Get_callExp().GetCfe()

	case 5:
		{
			p.SetState(413)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*ExprContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(414)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewArray((func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetColumn()
			}
		}()), nil)

	case 6:
		{
			p.SetState(416)

			var _x = p.listArray(0)

			localctx.(*ExprContext).list = _x
		}
		localctx.(*ExprContext).e = localctx.(*ExprContext).GetList().GetP()

	case 7:
		{
			p.SetState(419)

			var _m = p.Match(SwiftGrammarParserCORIZQ)

			localctx.(*ExprContext)._CORIZQ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)

			var _x = p.listParams(0)

			localctx.(*ExprContext)._listParams = _x
		}
		{
			p.SetState(421)
			p.Match(SwiftGrammarParserCORDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewArray((func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_CORIZQ() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_CORIZQ().GetColumn()
			}
		}()), localctx.(*ExprContext).Get_listParams().GetL())

	case 8:
		{
			p.SetState(424)

			var _m = p.Match(SwiftGrammarParserNUMBER)

			localctx.(*ExprContext)._NUMBER = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		if strings.Contains((func() string {
			if localctx.(*ExprContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_NUMBER().GetText()
			}
		}()), ".") {
			num, err := strconv.ParseFloat((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()), 64)
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.FLOAT)
		} else {
			num, err := strconv.Atoi((func() string {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return ""
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetText()
				}
			}()))
			if err != nil {
				fmt.Println(err)
			}
			localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetLine()
				}
			}()), (func() int {
				if localctx.(*ExprContext).Get_NUMBER() == nil {
					return 0
				} else {
					return localctx.(*ExprContext).Get_NUMBER().GetColumn()
				}
			}()), num, environment.INTEGER)
		}

	case 9:
		{
			p.SetState(426)

			var _m = p.Match(SwiftGrammarParserSTRING)

			localctx.(*ExprContext)._STRING = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		str := (func() string {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_STRING().GetText()
			}
		}())
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_STRING().GetColumn()
			}
		}()), str[1:len(str)-1], environment.STRING)

	case 10:
		{
			p.SetState(428)

			var _m = p.Match(SwiftGrammarParserTRU)

			localctx.(*ExprContext)._TRU = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_TRU() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_TRU().GetColumn()
			}
		}()), true, environment.BOOLEAN)

	case 11:
		{
			p.SetState(430)

			var _m = p.Match(SwiftGrammarParserFAL)

			localctx.(*ExprContext)._FAL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_FAL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_FAL().GetColumn()
			}
		}()), false, environment.BOOLEAN)

	case 12:
		{
			p.SetState(432)

			var _m = p.Match(SwiftGrammarParserID)

			localctx.(*ExprContext)._ID = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.Match(SwiftGrammarParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Match(SwiftGrammarParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewCount((func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_ID().GetColumn()
			}
		}()), (func() string {
			if localctx.(*ExprContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).Get_ID().GetText()
			}
		}()))

	case 13:
		{
			p.SetState(436)

			var _m = p.Match(SwiftGrammarParserNIL)

			localctx.(*ExprContext)._NIL = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).e = expressions.NewPrimitive((func() int {
			if localctx.(*ExprContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NIL().GetLine()
			}
		}()), (func() int {
			if localctx.(*ExprContext).Get_NIL() == nil {
				return 0
			} else {
				return localctx.(*ExprContext).Get_NIL().GetColumn()
			}
		}()), "nil", environment.NIL)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(487)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).expuno = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(440)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(441)
					p.Match(SwiftGrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(442)
					p.Match(SwiftGrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(443)
					p.Match(SwiftGrammarParserPUNTO)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(444)

					var _x = p.expr(22)

					localctx.(*ExprContext).expdos = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewRange((func() antlr.Token {
					if localctx.(*ExprContext).GetExpuno() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetExpuno().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetExpuno() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetExpuno().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetExpuno().GetE(), localctx.(*ExprContext).GetExpdos().GetE())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(447)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(448)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserSUM_IG || _la == SwiftGrammarParserSUB_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(449)

					var _x = p.expr(21)

					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() int {
					if localctx.(*ExprContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExprContext).GetOp().GetLine()
					}
				}()), (func() int {
					if localctx.(*ExprContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExprContext).GetOp().GetColumn()
					}
				}()), nil, (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).Get_expr().GetE())

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(452)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(453)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&42784196460019712) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(454)

					var _x = p.expr(20)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(457)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(458)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserADD || _la == SwiftGrammarParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(459)

					var _x = p.expr(19)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(462)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(463)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMAY_IG || _la == SwiftGrammarParserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(464)

					var _x = p.expr(18)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 6:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(467)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(468)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserMEN_IG || _la == SwiftGrammarParserMENOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(469)

					var _x = p.expr(17)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 7:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(472)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(473)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SwiftGrammarParserDIF || _la == SwiftGrammarParserIG_IG) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(474)

					var _x = p.expr(16)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 8:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(477)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(478)

					var _m = p.Match(SwiftGrammarParserAND)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(479)

					var _x = p.expr(15)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case 9:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_expr)
				p.SetState(482)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(483)

					var _m = p.Match(SwiftGrammarParserOR)

					localctx.(*ExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(484)

					var _x = p.expr(14)

					localctx.(*ExprContext).right = _x
					localctx.(*ExprContext)._expr = _x
				}
				localctx.(*ExprContext).e = expressions.NewOperation((func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ExprContext).GetLeft() == nil {
						return nil
					} else {
						return localctx.(*ExprContext).GetLeft().GetStart()
					}
				}()).GetColumn(), localctx.(*ExprContext).GetLeft().GetE(), (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()), localctx.(*ExprContext).GetRight().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
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

// IListParamsContext is an interface to support dynamic dispatch.
type IListParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListParamsContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListParamsContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	Expr() IExprContext
	COMA() antlr.TerminalNode
	ListParams() IListParamsContext

	// IsListParamsContext differentiates from other interfaces.
	IsListParamsContext()
}

type ListParamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListParamsContext
	_expr  IExprContext
}

func NewEmptyListParamsContext() *ListParamsContext {
	var p = new(ListParamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
	return p
}

func InitEmptyListParamsContext(p *ListParamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listParams
}

func (*ListParamsContext) IsListParamsContext() {}

func NewListParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsContext {
	var p = new(ListParamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listParams

	return p
}

func (s *ListParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsContext) GetList() IListParamsContext { return s.list }

func (s *ListParamsContext) Get_expr() IExprContext { return s._expr }

func (s *ListParamsContext) SetList(v IListParamsContext) { s.list = v }

func (s *ListParamsContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListParamsContext) GetL() []interface{} { return s.l }

func (s *ListParamsContext) SetL(v []interface{}) { s.l = v }

func (s *ListParamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListParamsContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ListParamsContext) ListParams() IListParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ListParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListParams(s)
	}
}

func (s *ListParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListParams(s)
	}
}

func (p *SwiftGrammarParser) ListParams() (localctx IListParamsContext) {
	return p.listParams(0)
}

func (p *SwiftGrammarParser) listParams(_p int) (localctx IListParamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListParamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, SwiftGrammarParserRULE_listParams, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)

		var _x = p.expr(0)

		localctx.(*ListParamsContext)._expr = _x
	}

	localctx.(*ListParamsContext).l = []interface{}{}
	localctx.(*ListParamsContext).l = append(localctx.(*ListParamsContext).l, localctx.(*ListParamsContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listParams)
			p.SetState(496)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(497)
				p.Match(SwiftGrammarParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(498)

				var _x = p.expr(0)

				localctx.(*ListParamsContext)._expr = _x
			}

			var arr []interface{}
			arr = append(localctx.(*ListParamsContext).GetList().GetL(), localctx.(*ListParamsContext).Get_expr().GetE())
			localctx.(*ListParamsContext).l = arr

		}
		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
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

// IListAccessArrayContext is an interface to support dynamic dispatch.
type IListAccessArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListAccessArrayContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListAccessArrayContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetL returns the l attribute.
	GetL() []interface{}

	// SetL sets the l attribute.
	SetL([]interface{})

	// Getter signatures
	CORIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORDER() antlr.TerminalNode
	ListAccessArray() IListAccessArrayContext

	// IsListAccessArrayContext differentiates from other interfaces.
	IsListAccessArrayContext()
}

type ListAccessArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	l      []interface{}
	list   IListAccessArrayContext
	_expr  IExprContext
}

func NewEmptyListAccessArrayContext() *ListAccessArrayContext {
	var p = new(ListAccessArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listAccessArray
	return p
}

func InitEmptyListAccessArrayContext(p *ListAccessArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listAccessArray
}

func (*ListAccessArrayContext) IsListAccessArrayContext() {}

func NewListAccessArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListAccessArrayContext {
	var p = new(ListAccessArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listAccessArray

	return p
}

func (s *ListAccessArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListAccessArrayContext) GetList() IListAccessArrayContext { return s.list }

func (s *ListAccessArrayContext) Get_expr() IExprContext { return s._expr }

func (s *ListAccessArrayContext) SetList(v IListAccessArrayContext) { s.list = v }

func (s *ListAccessArrayContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListAccessArrayContext) GetL() []interface{} { return s.l }

func (s *ListAccessArrayContext) SetL(v []interface{}) { s.l = v }

func (s *ListAccessArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ListAccessArrayContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListAccessArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ListAccessArrayContext) ListAccessArray() IListAccessArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListAccessArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListAccessArrayContext)
}

func (s *ListAccessArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListAccessArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListAccessArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListAccessArray(s)
	}
}

func (s *ListAccessArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListAccessArray(s)
	}
}

func (p *SwiftGrammarParser) ListAccessArray() (localctx IListAccessArrayContext) {
	return p.listAccessArray(0)
}

func (p *SwiftGrammarParser) listAccessArray(_p int) (localctx IListAccessArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListAccessArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListAccessArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 46
	p.EnterRecursionRule(localctx, 46, SwiftGrammarParserRULE_listAccessArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		p.Match(SwiftGrammarParserCORIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(508)

		var _x = p.expr(0)

		localctx.(*ListAccessArrayContext)._expr = _x
	}
	{
		p.SetState(509)
		p.Match(SwiftGrammarParserCORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*ListAccessArrayContext).l = []interface{}{}
	localctx.(*ListAccessArrayContext).l = append(localctx.(*ListAccessArrayContext).l, localctx.(*ListAccessArrayContext).Get_expr().GetE())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(520)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListAccessArrayContext(p, _parentctx, _parentState)
			localctx.(*ListAccessArrayContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listAccessArray)
			p.SetState(512)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(513)
				p.Match(SwiftGrammarParserCORIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(514)

				var _x = p.expr(0)

				localctx.(*ListAccessArrayContext)._expr = _x
			}
			{
				p.SetState(515)
				p.Match(SwiftGrammarParserCORDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			var arr []interface{}
			arr = append(localctx.(*ListAccessArrayContext).GetList().GetL(), localctx.(*ListAccessArrayContext).Get_expr().GetE())
			localctx.(*ListAccessArrayContext).l = arr

		}
		p.SetState(522)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
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

// IListArrayContext is an interface to support dynamic dispatch.
type IListArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expression

	// SetP sets the p attribute.
	SetP(interfaces.Expression)

	// Getter signatures
	ID() antlr.TerminalNode
	CORIZQ() antlr.TerminalNode
	Expr() IExprContext
	CORDER() antlr.TerminalNode
	ListArray() IListArrayContext
	Types() ITypesContext
	IG() antlr.TerminalNode

	// IsListArrayContext differentiates from other interfaces.
	IsListArrayContext()
}

type ListArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.Expression
	list   IListArrayContext
	_ID    antlr.Token
	_expr  IExprContext
}

func NewEmptyListArrayContext() *ListArrayContext {
	var p = new(ListArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
	return p
}

func InitEmptyListArrayContext(p *ListArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_listArray
}

func (*ListArrayContext) IsListArrayContext() {}

func NewListArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArrayContext {
	var p = new(ListArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_listArray

	return p
}

func (s *ListArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArrayContext) Get_ID() antlr.Token { return s._ID }

func (s *ListArrayContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListArrayContext) GetList() IListArrayContext { return s.list }

func (s *ListArrayContext) Get_expr() IExprContext { return s._expr }

func (s *ListArrayContext) SetList(v IListArrayContext) { s.list = v }

func (s *ListArrayContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ListArrayContext) GetP() interfaces.Expression { return s.p }

func (s *ListArrayContext) SetP(v interfaces.Expression) { s.p = v }

func (s *ListArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserID, 0)
}

func (s *ListArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORIZQ, 0)
}

func (s *ListArrayContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ListArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCORDER, 0)
}

func (s *ListArrayContext) ListArray() IListArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ListArrayContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *ListArrayContext) IG() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserIG, 0)
}

func (s *ListArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterListArray(s)
	}
}

func (s *ListArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitListArray(s)
	}
}

func (p *SwiftGrammarParser) ListArray() (localctx IListArrayContext) {
	return p.listArray(0)
}

func (p *SwiftGrammarParser) listArray(_p int) (localctx IListArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, SwiftGrammarParserRULE_listArray, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)

		var _m = p.Match(SwiftGrammarParserID)

		localctx.(*ListArrayContext)._ID = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	localctx.(*ListArrayContext).p = expressions.NewCallVar((func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetLine()
		}
	}()), (func() int {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetColumn()
		}
	}()), (func() string {
		if localctx.(*ListArrayContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListArrayContext).Get_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(541)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
			case 1:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(527)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(528)
					p.Match(SwiftGrammarParserCORIZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(529)

					var _x = p.expr(0)

					localctx.(*ListArrayContext)._expr = _x
				}
				{
					p.SetState(530)
					p.Match(SwiftGrammarParserCORDER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				localctx.(*ListArrayContext).p = expressions.NewArrayAccess((func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), localctx.(*ListArrayContext).Get_expr().GetE())

			case 2:
				localctx = NewListArrayContext(p, _parentctx, _parentState)
				localctx.(*ListArrayContext).list = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_listArray)
				p.SetState(533)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(534)
					p.Types()
				}
				{
					p.SetState(535)
					p.Match(SwiftGrammarParserIG)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(536)
					p.Match(SwiftGrammarParserCORIZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(537)

					var _x = p.expr(0)

					localctx.(*ListArrayContext)._expr = _x
				}
				{
					p.SetState(538)
					p.Match(SwiftGrammarParserCORDER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				localctx.(*ListArrayContext).p = expressions.NewArrayAccess((func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetLine(), (func() antlr.Token {
					if localctx.(*ListArrayContext).GetList() == nil {
						return nil
					} else {
						return localctx.(*ListArrayContext).GetList().GetStart()
					}
				}()).GetColumn(), localctx.(*ListArrayContext).GetList().GetP(), localctx.(*ListArrayContext).Get_expr().GetE())

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(545)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
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

// IExprComaContext is an interface to support dynamic dispatch.
type IExprComaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IExprComaContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetRight returns the right rule contexts.
	GetRight() IExprContext

	// SetLeft sets the left rule contexts.
	SetLeft(IExprComaContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetRight sets the right rule contexts.
	SetRight(IExprContext)

	// GetT returns the t attribute.
	GetT() interfaces.Expression

	// SetT sets the t attribute.
	SetT(interfaces.Expression)

	// Getter signatures
	Expr() IExprContext
	ExprComa() IExprComaContext
	COMA() antlr.TerminalNode

	// IsExprComaContext differentiates from other interfaces.
	IsExprComaContext()
}

type ExprComaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	t      interfaces.Expression
	left   IExprComaContext
	_expr  IExprContext
	op     antlr.Token
	right  IExprContext
}

func NewEmptyExprComaContext() *ExprComaContext {
	var p = new(ExprComaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_exprComa
	return p
}

func InitEmptyExprComaContext(p *ExprComaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SwiftGrammarParserRULE_exprComa
}

func (*ExprComaContext) IsExprComaContext() {}

func NewExprComaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprComaContext {
	var p = new(ExprComaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SwiftGrammarParserRULE_exprComa

	return p
}

func (s *ExprComaContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprComaContext) GetOp() antlr.Token { return s.op }

func (s *ExprComaContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprComaContext) GetLeft() IExprComaContext { return s.left }

func (s *ExprComaContext) Get_expr() IExprContext { return s._expr }

func (s *ExprComaContext) GetRight() IExprContext { return s.right }

func (s *ExprComaContext) SetLeft(v IExprComaContext) { s.left = v }

func (s *ExprComaContext) Set_expr(v IExprContext) { s._expr = v }

func (s *ExprComaContext) SetRight(v IExprContext) { s.right = v }

func (s *ExprComaContext) GetT() interfaces.Expression { return s.t }

func (s *ExprComaContext) SetT(v interfaces.Expression) { s.t = v }

func (s *ExprComaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprComaContext) ExprComa() IExprComaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprComaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprComaContext)
}

func (s *ExprComaContext) COMA() antlr.TerminalNode {
	return s.GetToken(SwiftGrammarParserCOMA, 0)
}

func (s *ExprComaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprComaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprComaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.EnterExprComa(s)
	}
}

func (s *ExprComaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SwiftGrammarListener); ok {
		listenerT.ExitExprComa(s)
	}
}

func (p *SwiftGrammarParser) ExprComa() (localctx IExprComaContext) {
	return p.exprComa(0)
}

func (p *SwiftGrammarParser) exprComa(_p int) (localctx IExprComaContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprComaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprComaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 50
	p.EnterRecursionRule(localctx, 50, SwiftGrammarParserRULE_exprComa, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(547)

		var _x = p.expr(0)

		localctx.(*ExprComaContext)._expr = _x
	}
	localctx.(*ExprComaContext).t = localctx.(*ExprComaContext).Get_expr().GetE()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprComaContext(p, _parentctx, _parentState)
			localctx.(*ExprComaContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SwiftGrammarParserRULE_exprComa)
			p.SetState(550)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(551)

				var _m = p.Match(SwiftGrammarParserCOMA)

				localctx.(*ExprComaContext).op = _m
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(552)

				var _x = p.expr(0)

				localctx.(*ExprComaContext).right = _x
				localctx.(*ExprComaContext)._expr = _x
			}
			localctx.(*ExprComaContext).t = expressions.NewOperation((func() antlr.Token {
				if localctx.(*ExprComaContext).GetLeft() == nil {
					return nil
				} else {
					return localctx.(*ExprComaContext).GetLeft().GetStart()
				}
			}()).GetLine(), (func() antlr.Token {
				if localctx.(*ExprComaContext).GetLeft() == nil {
					return nil
				} else {
					return localctx.(*ExprComaContext).GetLeft().GetStart()
				}
			}()).GetColumn(), localctx.(*ExprComaContext).GetLeft().GetT(), (func() string {
				if localctx.(*ExprComaContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*ExprComaContext).GetOp().GetText()
				}
			}()), localctx.(*ExprComaContext).GetRight().GetE())

		}
		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
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

func (p *SwiftGrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *ListStructDecContext = nil
		if localctx != nil {
			t = localctx.(*ListStructDecContext)
		}
		return p.ListStructDec_Sempred(t, predIndex)

	case 15:
		var t *ListParamsFuncContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsFuncContext)
		}
		return p.ListParamsFunc_Sempred(t, predIndex)

	case 19:
		var t *ListParamsCallContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsCallContext)
		}
		return p.ListParamsCall_Sempred(t, predIndex)

	case 21:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 22:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 23:
		var t *ListAccessArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListAccessArrayContext)
		}
		return p.ListAccessArray_Sempred(t, predIndex)

	case 24:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	case 25:
		var t *ExprComaContext = nil
		if localctx != nil {
			t = localctx.(*ExprComaContext)
		}
		return p.ExprComa_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SwiftGrammarParser) ListStructDec_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParamsFunc_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParamsCall_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListAccessArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 16:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SwiftGrammarParser) ExprComa_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 18:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
