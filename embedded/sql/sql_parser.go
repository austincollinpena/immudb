// Code generated by goyacc -l -o sql_parser.go sql_grammar.y. DO NOT EDIT.
package sql

import __yyfmt__ "fmt"

import "fmt"

func setResult(l yyLexer, stmts []SQLStmt) {
	l.(*lexer).result = stmts
}

type yySymType struct {
	yys      int
	stmts    []SQLStmt
	stmt     SQLStmt
	colsSpec []*ColSpec
	colSpec  *ColSpec
	cols     []*ColSelector
	rows     []*RowSpec
	row      *RowSpec
	values   []ValueExp
	value    ValueExp
	id       string
	number   uint64
	str      string
	boolean  bool
	blob     []byte
	sqlType  SQLValueType
	aggFn    AggregateFn
	ids      []string
	col      *ColSelector
	sel      Selector
	sels     []Selector
	distinct bool
	ds       DataSource
	tableRef *tableRef
	joins    []*JoinSpec
	join     *JoinSpec
	joinType JoinType
	boolExp  ValueExp
	binExp   ValueExp
	err      error
	ordcols  []*OrdCol
	opt_ord  Order
	logicOp  LogicOperator
	cmpOp    CmpOperator
	pparam   int
}

const CREATE = 57346
const USE = 57347
const DATABASE = 57348
const SNAPSHOT = 57349
const SINCE = 57350
const UP = 57351
const TO = 57352
const TABLE = 57353
const UNIQUE = 57354
const INDEX = 57355
const ON = 57356
const ALTER = 57357
const ADD = 57358
const COLUMN = 57359
const PRIMARY = 57360
const KEY = 57361
const BEGIN = 57362
const TRANSACTION = 57363
const COMMIT = 57364
const INSERT = 57365
const UPSERT = 57366
const INTO = 57367
const VALUES = 57368
const SELECT = 57369
const DISTINCT = 57370
const FROM = 57371
const BEFORE = 57372
const TX = 57373
const JOIN = 57374
const HAVING = 57375
const WHERE = 57376
const GROUP = 57377
const BY = 57378
const LIMIT = 57379
const ORDER = 57380
const ASC = 57381
const DESC = 57382
const AS = 57383
const NOT = 57384
const LIKE = 57385
const IF = 57386
const EXISTS = 57387
const AUTO_INCREMENT = 57388
const NULL = 57389
const NPARAM = 57390
const PPARAM = 57391
const JOINTYPE = 57392
const LOP = 57393
const CMPOP = 57394
const IDENTIFIER = 57395
const TYPE = 57396
const NUMBER = 57397
const VARCHAR = 57398
const BOOLEAN = 57399
const BLOB = 57400
const AGGREGATE_FUNC = 57401
const ERROR = 57402
const STMT_SEPARATOR = 57403

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CREATE",
	"USE",
	"DATABASE",
	"SNAPSHOT",
	"SINCE",
	"UP",
	"TO",
	"TABLE",
	"UNIQUE",
	"INDEX",
	"ON",
	"ALTER",
	"ADD",
	"COLUMN",
	"PRIMARY",
	"KEY",
	"BEGIN",
	"TRANSACTION",
	"COMMIT",
	"INSERT",
	"UPSERT",
	"INTO",
	"VALUES",
	"SELECT",
	"DISTINCT",
	"FROM",
	"BEFORE",
	"TX",
	"JOIN",
	"HAVING",
	"WHERE",
	"GROUP",
	"BY",
	"LIMIT",
	"ORDER",
	"ASC",
	"DESC",
	"AS",
	"NOT",
	"LIKE",
	"IF",
	"EXISTS",
	"AUTO_INCREMENT",
	"NULL",
	"NPARAM",
	"PPARAM",
	"JOINTYPE",
	"LOP",
	"CMPOP",
	"IDENTIFIER",
	"TYPE",
	"NUMBER",
	"VARCHAR",
	"BOOLEAN",
	"BLOB",
	"AGGREGATE_FUNC",
	"ERROR",
	"','",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'.'",
	"STMT_SEPARATOR",
	"'('",
	"')'",
	"'['",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 280

var yyAct = [...]int{
	229, 38, 58, 199, 90, 129, 131, 153, 152, 4,
	104, 74, 66, 95, 133, 75, 205, 136, 111, 144,
	142, 143, 150, 40, 111, 141, 224, 137, 138, 139,
	140, 39, 151, 214, 111, 134, 163, 164, 213, 79,
	135, 200, 122, 111, 207, 50, 52, 159, 160, 162,
	161, 112, 163, 164, 195, 197, 201, 164, 173, 61,
	51, 170, 80, 159, 160, 162, 161, 159, 160, 162,
	161, 118, 110, 101, 170, 76, 154, 93, 169, 107,
	100, 159, 160, 162, 161, 84, 99, 102, 82, 106,
	72, 70, 98, 60, 55, 18, 16, 162, 161, 109,
	71, 61, 40, 228, 212, 208, 178, 120, 39, 111,
	115, 117, 123, 35, 57, 144, 142, 143, 5, 194,
	146, 182, 130, 137, 138, 139, 140, 40, 145, 32,
	226, 148, 177, 39, 108, 155, 87, 147, 33, 166,
	167, 168, 121, 37, 40, 7, 91, 171, 125, 119,
	105, 92, 85, 81, 78, 64, 62, 51, 49, 46,
	187, 181, 41, 185, 97, 188, 189, 190, 191, 192,
	193, 51, 105, 216, 33, 176, 204, 83, 198, 196,
	77, 203, 43, 165, 63, 73, 206, 230, 231, 184,
	59, 219, 211, 158, 128, 15, 114, 157, 116, 86,
	17, 68, 67, 56, 21, 7, 215, 126, 124, 30,
	29, 53, 19, 222, 223, 217, 10, 11, 174, 88,
	2, 225, 10, 11, 69, 227, 220, 12, 172, 54,
	233, 232, 6, 12, 234, 13, 14, 31, 22, 7,
	65, 13, 14, 23, 25, 24, 210, 44, 45, 28,
	48, 26, 27, 202, 175, 42, 209, 183, 221, 149,
	218, 127, 132, 156, 113, 96, 94, 47, 20, 36,
	34, 179, 180, 186, 89, 103, 9, 8, 3, 1,
}

var yyPact = [...]int{
	212, -1000, -1000, 29, 28, -1000, 191, 176, -1000, -1000,
	232, 245, 238, 185, 184, -1000, 212, -1000, -1000, 218,
	49, -1000, 109, 138, 233, 235, 106, 242, 105, 104,
	104, -1000, 189, 27, 174, -1000, 53, 149, -1000, 25,
	35, -1000, 103, 142, 102, 226, -1000, 172, 170, 208,
	23, 34, 22, -1000, -1000, 218, 7, 74, -1000, 101,
	-30, 100, 20, 132, 17, 99, -1000, 168, 81, 202,
	93, 98, 93, -1000, 114, -1000, 118, 149, -1000, -1000,
	4, 21, 97, -1000, 93, 11, 79, -1000, 97, 3,
	48, -1000, -1000, -18, 162, -1000, 114, 166, 172, 2,
	-1000, -1000, 96, 46, -1000, 88, -27, 93, -1000, -1000,
	182, 95, 181, 159, -28, -1000, 7, 149, -1000, -1000,
	119, -48, -1000, -37, 8, -1000, 8, 164, 157, 1,
	140, -1000, -1000, -28, -28, -28, 10, -1000, -1000, -1000,
	-1000, -7, 94, -1000, -1000, 214, -11, 199, -1000, 129,
	77, -1000, 45, -1000, 68, 45, 151, -28, 91, -28,
	-28, -28, -28, -28, -28, 63, 5, 33, -15, 178,
	-14, -1000, -28, -1000, -12, 134, -1000, -55, 8, -25,
	44, -1000, 6, 241, 156, 1, 43, -1000, 33, 33,
	-1000, -1000, 5, 19, -1000, -1000, -31, -1000, 1, -36,
	-1000, 93, -1000, -1000, 126, -1000, -1000, -1000, 68, 154,
	213, 91, 91, -1000, -1000, -43, -1000, -1000, 149, 75,
	211, 42, 148, -1000, -1000, -1000, -1000, -12, 91, -1000,
	-1000, -1000, -1000, 148, -1000,
}

var yyPgo = [...]int{
	0, 279, 220, 129, 278, 118, 277, 276, 9, 275,
	10, 4, 3, 274, 273, 8, 7, 272, 271, 6,
	122, 270, 269, 1, 268, 11, 15, 267, 12, 266,
	13, 265, 5, 264, 263, 262, 261, 260, 259, 2,
	258, 257, 0, 256, 255, 254, 253, 195,
}

var yyR1 = [...]int{
	0, 1, 2, 2, 2, 47, 47, 4, 4, 5,
	5, 3, 3, 6, 6, 6, 6, 6, 6, 6,
	27, 27, 44, 44, 12, 12, 7, 7, 13, 13,
	15, 15, 16, 11, 11, 14, 14, 18, 18, 17,
	17, 19, 19, 19, 19, 19, 19, 19, 19, 9,
	9, 10, 38, 38, 45, 45, 46, 46, 46, 8,
	24, 24, 21, 21, 22, 22, 20, 20, 20, 23,
	23, 23, 25, 25, 25, 26, 26, 28, 28, 29,
	29, 30, 30, 31, 33, 33, 36, 36, 34, 34,
	37, 37, 41, 41, 43, 43, 40, 40, 42, 42,
	42, 39, 39, 32, 32, 32, 32, 32, 32, 32,
	32, 35, 35, 35, 35, 35, 35,
}

var yyR2 = [...]int{
	0, 1, 2, 2, 3, 0, 1, 1, 4, 1,
	1, 2, 3, 3, 3, 4, 11, 7, 8, 6,
	0, 3, 0, 3, 1, 3, 8, 8, 0, 1,
	1, 3, 3, 1, 3, 1, 3, 0, 1, 1,
	3, 1, 1, 1, 1, 3, 2, 1, 1, 1,
	3, 5, 0, 3, 0, 1, 0, 1, 2, 13,
	0, 1, 1, 1, 2, 4, 1, 3, 4, 1,
	3, 5, 1, 5, 3, 1, 3, 0, 3, 0,
	1, 1, 2, 5, 0, 2, 0, 3, 0, 2,
	0, 2, 0, 3, 0, 4, 2, 4, 0, 1,
	1, 0, 2, 1, 1, 1, 2, 2, 3, 3,
	4, 3, 3, 3, 3, 3, 3,
}

var yyChk = [...]int{
	-1000, -1, -2, -4, -8, -5, 20, 27, -6, -7,
	4, 5, 15, 23, 24, -47, 67, -47, 67, 21,
	-24, 28, 6, 11, 13, 12, 6, 7, 11, 25,
	25, -2, -3, -5, -21, 64, -22, -20, -23, 59,
	53, 53, -44, 44, 14, 13, 53, -27, 8, 53,
	-26, 53, -26, 22, -47, 67, 29, 61, -39, 41,
	68, 66, 53, 42, 53, 14, -28, 30, 31, 16,
	68, 66, 68, -3, -25, -26, 68, -20, 53, 69,
	-23, 53, 68, 45, 68, 53, 31, 55, 17, -13,
	-11, 53, 53, -11, -29, -30, -31, 50, -26, -8,
	-39, 69, 66, -9, -10, 53, -11, 68, 55, -10,
	69, 61, 69, -33, 34, -30, 32, -28, 69, 53,
	61, 54, 69, -11, 26, 53, 26, -36, 35, -32,
	-20, -19, -35, 42, 63, 68, 45, 55, 56, 57,
	58, 53, 48, 49, 47, -25, -39, 18, -10, -38,
	70, 69, -15, -16, 68, -15, -34, 33, 36, 62,
	63, 65, 64, 51, 52, 43, -32, -32, -32, 68,
	68, 53, 14, 69, 19, -45, 46, 55, 61, -18,
	-17, -19, 53, -41, 38, -32, -14, -23, -32, -32,
	-32, -32, -32, -32, 56, 69, -8, 69, -32, -12,
	53, 68, -46, 47, 42, 71, -16, 69, 61, -43,
	5, 36, 61, 69, 69, -11, 47, -19, -37, 37,
	13, -40, -23, -23, 69, -39, 55, 14, 61, -42,
	39, 40, -12, -23, -42,
}

var yyDef = [...]int{
	0, -2, 1, 5, 5, 7, 0, 60, 9, 10,
	0, 0, 0, 0, 0, 2, 6, 3, 6, 0,
	0, 61, 0, 22, 0, 0, 0, 20, 0, 0,
	0, 4, 0, 5, 0, 62, 63, 101, 66, 0,
	69, 13, 0, 0, 0, 0, 14, 77, 0, 0,
	0, 75, 0, 8, 11, 6, 0, 0, 64, 0,
	0, 0, 0, 0, 0, 0, 15, 0, 0, 0,
	28, 0, 0, 12, 79, 72, 0, 101, 102, 67,
	0, 70, 0, 23, 0, 0, 0, 21, 0, 0,
	29, 33, 76, 0, 84, 80, 81, 0, 77, 0,
	65, 68, 0, 0, 49, 0, 0, 0, 78, 19,
	0, 0, 0, 86, 0, 82, 0, 101, 74, 71,
	0, 52, 17, 0, 0, 34, 0, 88, 0, 85,
	103, 104, 105, 0, 0, 0, 0, 41, 42, 43,
	44, 69, 0, 47, 48, 0, 0, 0, 50, 54,
	0, 18, 26, 30, 37, 27, 92, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 106, 107, 0, 0,
	0, 46, 0, 73, 0, 56, 55, 0, 0, 0,
	38, 39, 0, 94, 0, 89, 87, 35, 111, 112,
	113, 114, 115, 116, 109, 108, 0, 45, 83, 0,
	24, 0, 51, 57, 0, 53, 31, 32, 0, 90,
	0, 0, 0, 110, 16, 0, 58, 40, 101, 0,
	0, 93, 98, 36, 25, 59, 91, 0, 0, 96,
	99, 100, 95, 98, 97,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	68, 69, 64, 62, 61, 63, 66, 65, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 70, 3, 71,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 67,
}

var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmts = yyDollar[1].stmts
			setResult(yylex, yyDollar[1].stmts)
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 8:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &TxStmt{stmts: yyDollar[3].stmts}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.stmts = []SQLStmt{yyDollar[1].stmt}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmts = append([]SQLStmt{yyDollar[1].stmt}, yyDollar[3].stmts...)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &CreateDatabaseStmt{DB: yyDollar[3].id}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.stmt = &UseDatabaseStmt{DB: yyDollar[3].id}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.stmt = &UseSnapshotStmt{sinceTx: yyDollar[3].number, asBefore: yyDollar[4].number}
		}
	case 16:
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			yyVAL.stmt = &CreateTableStmt{ifNotExists: yyDollar[3].boolean, table: yyDollar[4].id, colsSpec: yyDollar[6].colsSpec, pkColNames: yyDollar[10].ids}
		}
	case 17:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{table: yyDollar[4].id, cols: yyDollar[6].ids}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &CreateIndexStmt{unique: true, table: yyDollar[5].id, cols: yyDollar[7].ids}
		}
	case 19:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.stmt = &AddColumnStmt{table: yyDollar[3].id, colSpec: yyDollar[6].colSpec}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = yyDollar[2].ids
		}
	case 26:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{isInsert: true, tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 27:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.stmt = &UpsertIntoStmt{tableRef: yyDollar[3].tableRef, cols: yyDollar[5].ids, rows: yyDollar[8].rows}
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = yyDollar[1].ids
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.rows = []*RowSpec{yyDollar[1].row}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.rows = append(yyDollar[1].rows, yyDollar[3].row)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.row = &RowSpec{Values: yyDollar[2].values}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ids = []string{yyDollar[1].id}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ids = append(yyDollar[1].ids, yyDollar[3].id)
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.cols = []*ColSelector{yyDollar[1].col}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = append(yyDollar[1].cols, yyDollar[3].col)
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.values = nil
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = yyDollar[1].values
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.values = []ValueExp{yyDollar[1].value}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].value)
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Number{val: int64(yyDollar[1].number)}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Varchar{val: yyDollar[1].str}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Bool{val: yyDollar[1].boolean}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Blob{val: yyDollar[1].blob}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.value = &SysFn{fn: yyDollar[1].id}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.value = &Param{id: yyDollar[2].id}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &Param{id: fmt.Sprintf("param%d", yyDollar[1].pparam), pos: yyDollar[1].pparam}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.value = &NullValue{t: AnyType}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.colsSpec = []*ColSpec{yyDollar[1].colSpec}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.colsSpec = append(yyDollar[1].colsSpec, yyDollar[3].colSpec)
		}
	case 51:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.colSpec = &ColSpec{colName: yyDollar[1].id, colType: yyDollar[2].sqlType, maxLen: int(yyDollar[3].number), autoIncrement: yyDollar[4].boolean, notNull: yyDollar[5].boolean}
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolean = false
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolean = true
		}
	case 59:
		yyDollar = yyS[yypt-13 : yypt+1]
		{
			yyVAL.stmt = &SelectStmt{
				distinct:  yyDollar[2].distinct,
				selectors: yyDollar[3].sels,
				ds:        yyDollar[5].ds,
				joins:     yyDollar[6].joins,
				where:     yyDollar[7].boolExp,
				groupBy:   yyDollar[8].cols,
				having:    yyDollar[9].boolExp,
				orderBy:   yyDollar[10].ordcols,
				indexOn:   yyDollar[11].ids,
				limit:     yyDollar[12].number,
				as:        yyDollar[13].id,
			}
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.distinct = false
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.distinct = true
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = nil
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sels = yyDollar[1].sels
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyDollar[1].sel.setAlias(yyDollar[2].id)
			yyVAL.sels = []Selector{yyDollar[1].sel}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyDollar[3].sel.setAlias(yyDollar[4].id)
			yyVAL.sels = append(yyDollar[1].sels, yyDollar[3].sel)
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.sel = yyDollar[1].col
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, col: "*"}
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.sel = &AggColSelector{aggFn: yyDollar[1].aggFn, db: yyDollar[3].col.db, table: yyDollar[3].col.table, col: yyDollar[3].col.col}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.col = &ColSelector{col: yyDollar[1].id}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.col = &ColSelector{table: yyDollar[1].id, col: yyDollar[3].id}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.col = &ColSelector{db: yyDollar[1].id, table: yyDollar[3].id, col: yyDollar[5].id}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ds = yyDollar[1].tableRef
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyDollar[2].tableRef.asBefore = yyDollar[3].number
			yyDollar[2].tableRef.as = yyDollar[4].id
			yyVAL.ds = yyDollar[2].tableRef
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ds = yyDollar[2].stmt.(*SelectStmt)
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.tableRef = &tableRef{table: yyDollar[1].id}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.tableRef = &tableRef{db: yyDollar[1].id, table: yyDollar[3].id}
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.number = yyDollar[3].number
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.joins = nil
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = yyDollar[1].joins
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.joins = []*JoinSpec{yyDollar[1].join}
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.joins = append([]*JoinSpec{yyDollar[1].join}, yyDollar[2].joins...)
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			yyVAL.join = &JoinSpec{joinType: yyDollar[1].joinType, ds: yyDollar[3].ds, cond: yyDollar[5].boolExp}
		}
	case 84:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolExp = nil
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.cols = nil
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.cols = yyDollar[3].cols
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.boolExp = nil
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 90:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.number = 0
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.number = yyDollar[2].number
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ordcols = nil
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.ordcols = yyDollar[3].ordcols
		}
	case 94:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.ids = nil
		}
	case 95:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ids = yyDollar[4].ids
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.ordcols = []*OrdCol{{sel: yyDollar[1].col, order: yyDollar[2].opt_ord}}
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.ordcols = append(yyDollar[1].ordcols, &OrdCol{sel: yyDollar[3].col, order: yyDollar[4].opt_ord})
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.opt_ord = AscOrder
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = AscOrder
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.opt_ord = DescOrder
		}
	case 101:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.id = ""
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.id = yyDollar[2].id
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[1].sel
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[1].value
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[1].binExp
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = &NotBoolExp{exp: yyDollar[2].boolExp}
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.boolExp = &NumExp{left: &Number{val: 0}, op: SUBSOP, right: yyDollar[2].boolExp}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = yyDollar[2].boolExp
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.boolExp = &LikeBoolExp{sel: yyDollar[1].sel, pattern: yyDollar[3].str}
		}
	case 110:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.boolExp = &ExistsBoolExp{q: (yyDollar[3].stmt).(*SelectStmt)}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].boolExp, op: ADDOP, right: yyDollar[3].boolExp}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].boolExp, op: SUBSOP, right: yyDollar[3].boolExp}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].boolExp, op: DIVOP, right: yyDollar[3].boolExp}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &NumExp{left: yyDollar[1].boolExp, op: MULTOP, right: yyDollar[3].boolExp}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &BinBoolExp{left: yyDollar[1].boolExp, op: yyDollar[2].logicOp, right: yyDollar[3].boolExp}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.binExp = &CmpBoolExp{left: yyDollar[1].boolExp, op: yyDollar[2].cmpOp, right: yyDollar[3].boolExp}
		}
	}
	goto yystack /* stack new state and value */
}
