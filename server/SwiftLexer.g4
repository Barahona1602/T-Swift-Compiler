lexer grammar SwiftLexer;

// --------------- Tokens
// types
INT: 'Int';
FLOAT: 'Float';
BOOL: 'Bool';
STR: 'String';
CHAR: 'Character';



// reserved words
VAR: 'var';
LET: 'let';
VOID: 'void';
TRU: 'true';
FAL: 'false';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
FOR: 'for';
IN: 'in';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
BREAK: 'break';
RETURN: 'return';
CONTINUE: 'continue';
GUARD: 'guard';
FUNC: 'func';
NIL: 'nil';
STRUCT: 'struct';
MUTATING: 'mutating';
SELF: 'self';
INOUT: 'inout';


//FUNCTIONS
APPEND: 'append';
REMOVELAST: 'removeLast';
REMOVE: 'remove';
AT: 'at';
ISEMPTY: 'isEmpty';
COUNT: 'count';

// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"'~["]*'"';
ID: ([a-zA-Z_])[a-zA-Z0-9_]*;

// symbols

DIF:      '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:          '=';
MAY_IG:     '>=';
MEN_IG:     '<=';
SUM_IG:     '+=';
SUB_IG:     '-=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
MOD:            '%';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
D_PTS:          ':';
CORIZQ:         '[';
CORDER:         ']';
COMA:           ',';
PUNTO:          '.';
COMILLA:        '"';
FLECHA:         '->';
GUIONBAJO:      '_';
AMP :           '&';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;

