grammar SwiftGrammar; 
// import SwiftLexer; 

options {
  tokenVocab = SwiftLexer;
}

@header {
    import "Server2/interfaces"
    import "Server2/environment"
    import "Server2/expressions"
    import "Server2/instructions"
    import "strings"
}


s returns [[]interface{} code]
: block EOF
    {   
        $code = $block.blk
    }
;

block returns [[]interface{} blk]
@init{
    $blk = []interface{}{}
    var listInt []IInstructionContext
  }
: ins+=instruction+
    {
        listInt = localctx.(*BlockContext).GetIns()
        for _, e := range listInt {
            $blk = append($blk, e.GetInst())
        }
    }
;


instruction returns [interfaces.Instruction inst]
: printstmt { $inst = $printstmt.prnt}
| ifstmt { $inst = $ifstmt.ifinst }
| declarationstmt { $inst = $declarationstmt.dec }
| whilestmt { $inst = $whilestmt.whl }
| assignstmt { $inst = $assignstmt.asg }
| forstmt { $inst = $forstmt.fr }
| guardstmt { $inst = $guardstmt.grd }
| breakstmt { $inst = $breakstmt.brk }
| continuestmt { $inst = $continuestmt.cnt }
// | fnArray { $inst = $fnArray.p }
// | structCreation { $inst = $structCreation.dec }
// | returnstmt { $inst = $returnstmt.ret }
| fnstmt { $inst = $fnstmt.fn }
| callFunction { $inst = $callFunction.cf }
;

printstmt returns [interfaces.Instruction prnt]
: PRINT PARIZQ expr PARDER { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$expr.e)}
| PRINT PARIZQ exprComa PARDER { $prnt = instructions.NewPrint($PRINT.line,$PRINT.pos,$exprComa.t)}
;

ifstmt returns [interfaces.Instruction ifinst]
    : IF expr LLAVEIZQ block LLAVEDER { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, nil) }
    | IF expr LLAVEIZQ e1=block LLAVEDER ELSE LLAVEIZQ e2=block LLAVEDER { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $e1.blk, $e2.blk) }
    | IF expr LLAVEIZQ block LLAVEDER ELSE ifstmt { $ifinst = instructions.NewIf($IF.line, $IF.pos, $expr.e, $block.blk, []interface{}{$ifstmt.ifinst}) }
;

whilestmt returns [interfaces.Instruction whl]
    : WHILE expr LLAVEIZQ block LLAVEDER { $whl = instructions.NewWhile($WHILE.line, $WHILE.pos, $expr.e, $block.blk) }
;

declarationstmt returns [interfaces.Instruction dec]
: VAR ID D_PTS types IG expr  { $dec = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, $types.ty, $expr.e, true) }
| VAR ID IG expr { $dec = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, environment.UNKNOWN, $expr.e, true) }
| VAR ID D_PTS types  { $dec = instructions.NewDeclaration($VAR.line, $VAR.pos, $ID.text, $types.ty, nil, true) }
| LET ID D_PTS types IG expr  { $dec = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, $expr.e, false) }
| LET ID D_PTS types  { $dec = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, $types.ty, nil, false) }
| LET ID IG expr { $dec = instructions.NewDeclaration($LET.line, $LET.pos, $ID.text, environment.UNKNOWN, $expr.e, false) }
;

assignstmt returns [interfaces.Instruction asg]
: ID op=IG expr { $asg = instructions.NewAssignment($ID.line, $ID.pos, $ID.text, $expr.e) }
// | listAccessStruct IG expr { $asg = instructions.NewStructAssign($listAccessStruct.start.GetLine(),$listAccessStruct.start.GetColumn(), $listAccessStruct.l, $expr.e) }
| ID op=(SUB_IG | SUM_IG) expr { $asg = instructions.NewImplicitAssignment($ID.line, $ID.pos, $ID.text, $op.text, $expr.e); }
// | ID listAccessArray IG expr { $asg = instructions.NewArrayAssign($ID.line, $ID.pos, $ID.text, $listAccessArray.l, $expr.e) }
;

forstmt returns [interfaces.Instruction fr]
: FOR ID IN expr LLAVEIZQ block LLAVEDER { $fr = instructions.NewForIn($FOR.line, $FOR.pos, $ID.text, $expr.e, $block.blk) };

guardstmt returns [interfaces.Instruction grd]
: GUARD expr ELSE LLAVEIZQ block LLAVEDER { $grd = instructions.NewGuard($GUARD.line, $GUARD.pos, $expr.e, $block.blk) }
;


breakstmt returns [interfaces.Instruction brk]
: BREAK { $brk = instructions.NewBreak($BREAK.line, $BREAK.pos, nil) }
;

continuestmt returns [interfaces.Instruction cnt]
: CONTINUE { $cnt = instructions.NewContinue($CONTINUE.line, $CONTINUE.pos) }
;

// returnstmt returns [interfaces.Instruction ret]
// : RETURN expr { $ret = instructions.NewReturn($RETURN.line, $RETURN.pos, $expr.e) }
// | RETURN { $ret = instructions.NewReturn($RETURN.line, $RETURN.pos, nil) }
// ;

// fnArray returns[interfaces.Instruction p]
// : ID PUNTO APPEND PARIZQ expr PARDER { $p = instructions.NewAppend($ID.line, $ID.pos, $ID.text, $expr.e) }
// | ID PUNTO REMOVE PARIZQ AT D_PTS expr PARDER { $p = instructions.NewRemoveAt($ID.line, $ID.pos, $ID.text, $expr.e) }
// | ID PUNTO REMOVELAST PARIZQ PARDER { $p = instructions.NewRemoveLast($ID.line, $ID.pos, $ID.text) }
// ;

// structCreation returns[interfaces.Instruction dec]
// : STRUCT ID LLAVEIZQ listStructDec LLAVEDER { $dec = instructions.NewStruct($STRUCT.line, $STRUCT.pos, $ID.text, $listStructDec.l) }
// ;

// listStructDec returns[[]interface{} l]
// : list=listStructDec (COMA)? (VAR|LET) ID D_PTS types {
//                                                 var arr []interface{}
//                                                 newParams := environment.NewStructType($ID.text, $types.ty, "")
//                                                 arr = append($list.l, newParams)
//                                                 $l = arr
//                                             }
// | list=listStructDec (COMA)? (VAR|LET) id1=ID D_PTS id2=ID {
//                                                 var arr []interface{}
//                                                 newParams := environment.NewStructType($id1.text, environment.UNKNOWN, $id2.text)
//                                                 arr = append($list.l, newParams)
//                                                 $l = arr
//                                             } 
// | list=listStructDec (COMA)? (VAR|LET) id1=ID (D_PTS types)? IG expr {
//                                                 var arr []interface{}
//                                                 newParams := environment.NewStructType($id1.text, $types.ty, "")
//                                                 arr = append($list.l, newParams)
//                                                 $l = arr
//                                             }
// | (VAR|LET) ID D_PTS types {
//                         var arr []interface{}
//                         newParams := environment.NewStructType($ID.text, $types.ty, "")
//                         arr = append(arr, newParams)
//                         $l = arr
//                     }
// | (VAR|LET) id1=ID D_PTS id2=ID {
//                                                 var arr []interface{}
//                                                 newParams := environment.NewStructType($id1.text,environment.UNKNOWN , $id2.text)
//                                                 arr = append($list.l, newParams)
//                                                 $l = arr
//                                             } 
// |  { $l = []interface{}{} }
// ;

// listStructExp returns[[]interface{} l]
// : list=listStructExp COMA ID D_PTS expr {
//                                             var arr []interface{}
//                                             StrExp := environment.NewStructContent($ID.text, $expr.e)
//                                             arr = append($list.l, StrExp)
//                                             $l = arr
//                                         }
// | ID D_PTS expr{
//                     var arr []interface{}
//                     StrExp := environment.NewStructContent($ID.text, $expr.e)
//                     arr = append(arr, StrExp)
//                     $l = arr
//                 }
// |   {
//         $l = []interface{}{}
//     }
// ;

// listAccessStruct returns[[]interface{} l]
// : list=listAccessStruct PUNTO ID {
//                                     var arr []interface{}
//                                     arr = append($list.l, $ID.text)
//                                     $l = arr
//                                 }
// | ID { 
//     $l = []interface{}{}
//     $l = append($l, $ID.text) 
//     }
// ;

fnstmt returns[interfaces.Instruction fn]
: FUNC ID PARIZQ listParamsFunc PARDER FLECHA types LLAVEIZQ block LLAVEDER { $fn = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.l, $types.ty, $block.blk) }
| FUNC ID PARIZQ listParamsFunc PARDER LLAVEIZQ block LLAVEDER { $fn = instructions.NewFunction($FUNC.line, $FUNC.pos, $ID.text, $listParamsFunc.l, environment.NIL, $block.blk) }
; 

listParamsFunc returns[[]interface{} l]
: list=listParamsFunc COMA parametro {
                                var arr []interface{}
                                arr = append($list.l, $parametro.p)
                                $l = arr
                            }   
| parametro {
            $l = []interface{}{}
            $l = append($l, $parametro.p)
        }
|   {
        $l = []interface{}{}
    }
;

parametro returns[interfaces.Instruction p]
: ID D_PTS types  { $p = instructions.NewParamsDeclaration($ID.line,$ID.pos,$ID.text, $types.ty)}
| exte=(GUIONBAJO|ID) ID D_PTS types { $p = instructions.NewParamsDeclaration($ID.line,$ID.pos, $ID.text, $types.ty)}
;

callExp returns[interfaces.Expression cfe]
: ID PARIZQ listParamsCall PARDER { $cfe = expressions.NewCallExp($ID.line, $ID.pos, $ID.text, $listParamsCall.l) }
;

callFunction returns[interfaces.Instruction cf]
: ID PARIZQ listParamsCall PARDER { $cf = instructions.NewCallFunc($ID.line, $ID.pos, $ID.text, $listParamsCall.l) }
;


listParamsCall returns[[]interface{} l]
: list=listParamsCall COMA expr {
                                    var arr []interface{}
                                    arr = append($list.l, $expr.e)
                                    $l = arr
                                }
| expr  {
            $l = []interface{}{}
            $l = append($l, $expr.e)
        }
|   {
        $l = []interface{}{}
    }
;

types returns[environment.TipoExpresion ty]
: INT { $ty = environment.INTEGER }
| FLOAT { $ty = environment.FLOAT }
| STR { $ty = environment.STRING }
| BOOL { $ty = environment.BOOLEAN }
| CORIZQ types CORDER { $ty = environment.ARRAY }
| COMILLA STR COMILLA { $ty = environment.STR }
| NIL { $ty = environment.NIL }
| STRUCT { $ty = environment.STRUCT }
| ID { $ty = environment.UNKNOWN }
;

expr returns [interfaces.Expression e]
: SUB opDe=expr {$e = expressions.NewOperation($SUB.line,$SUB.pos,$opDe.e,"NEGACION",nil)}
| expuno=expr PUNTO PUNTO PUNTO expdos=expr { $e = expressions.NewRange($expuno.start.GetLine(),$expuno.start.GetColumn(), $expuno.e, $expdos.e) }
| left=expr op=(SUB_IG|SUM_IG) expr { $e = expressions.NewOperation($op.line, $op.pos, nil, $op.text, $expr.e) }
| left=expr op=(MUL|DIV|MOD) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(ADD|SUB) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MAY_IG|MAYOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(MEN_IG|MENOR) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=(IG_IG|DIF) right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=AND right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| left=expr op=OR right=expr { $e = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.e, $op.text, $right.e) }
| NOT right=expr {$e = expressions.NewOperation($NOT.line, $NOT.pos, $right.e, $NOT.text ,nil)}
| PARIZQ expr PARDER { $e = $expr.e }
| callExp { $e = $callExp.cfe }
// | types PARIZQ expr PARDER { $e = expressions.NewCast($types.start.GetLine(), $types.start.GetColumn(), $types.ty, $expr.e) }
// | ID PARIZQ listStructExp PARDER { $e = expressions.NewStructExp($ID.line, $ID.pos, $ID.text, $listStructExp.l ) }
| CORIZQ CORDER { $e = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, nil) }
| list=listArray { $e = $list.p}
| CORIZQ listParams CORDER { $e = expressions.NewArray($CORIZQ.line, $CORIZQ.pos, $listParams.l) }
| NUMBER
    {
        if (strings.Contains($NUMBER.text,".")){
            num,err := strconv.ParseFloat($NUMBER.text, 64);
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.FLOAT)
        }else{
            num,err := strconv.Atoi($NUMBER.text)
            if err!= nil{
                fmt.Println(err)
            }
            $e = expressions.NewPrimitive($NUMBER.line,$NUMBER.pos,num,environment.INTEGER)
        }
    }
| STRING
    {
        str := $STRING.text
        $e = expressions.NewPrimitive($STRING.line, $STRING.pos, str[1:len(str)-1],environment.STRING)
    }
| TRU { $e = expressions.NewPrimitive($TRU.line, $TRU.pos, true, environment.BOOLEAN) }
| FAL { $e = expressions.NewPrimitive($FAL.line, $FAL.pos, false, environment.BOOLEAN) }
// | ID PUNTO COUNT { $e = expressions.NewCount($ID.line, $ID.pos, $ID.text) }
// | ID PUNTO ISEMPTY { $e = expressions.NewIsEmpty($ID.line, $ID.pos, $ID.text) }
| NIL { $e = expressions.NewPrimitive($NIL.line, $NIL.pos, "nil", environment.NIL) }
;


listParams returns[[]interface{} l]
: list=listParams COMA expr {
                                var arr []interface{}
                                arr = append($list.l, $expr.e)
                                $l = arr
                            }   
| expr {
            $l = []interface{}{}
            $l = append($l, $expr.e)
        }
;

listAccessArray returns[[]interface{} l]
: list=listAccessArray CORIZQ expr CORDER {
                                var arr []interface{}
                                arr = append($list.l, $expr.e)
                                $l = arr
                            }   
|   CORIZQ expr CORDER {
            $l = []interface{}{}
            $l = append($l, $expr.e)
        }
;

listArray returns[interfaces.Expression p]
: list = listArray CORIZQ expr CORDER { $p = expressions.NewArrayAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $expr.e) }
| list = listArray PUNTO ID { $p = expressions.NewStructAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $ID.text)  }
| list = listArray types IG CORIZQ expr CORDER { $p = expressions.NewArrayAccess($list.start.GetLine(), $list.start.GetColumn(), $list.p, $expr.e) }
| ID { $p = expressions.NewCallVar($ID.line, $ID.pos, $ID.text)}
;

exprComa returns[interfaces.Expression t]
: left=exprComa op=COMA right=expr { $t = expressions.NewOperation($left.start.GetLine(), $left.start.GetColumn(), $left.t, $op.text, $right.e) }
| expr { $t = $expr.e }
;