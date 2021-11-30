package main

//go:generate mkdir -p plsql
//go:generate sh -c "[ -e antlr-4.9.3-complete.jar ] || wget https://www.antlr.org/download/antlr-4.9.3-complete.jar"
//go:generate sh -c "[ -e PlSqlLexer.g4 ] || wget https://github.com/antlr/grammars-v4/raw/master/sql/plsql/PlSqlLexer.g4"
//go:generate sh -c "[ -e PlSqlParser.g4 ] || wget https://github.com/antlr/grammars-v4/raw/master/sql/plsql/PlSqlParser.g4"
//go:generate java -jar antlr-4.9.3-complete.jar -Dlanguage=Go -o plsql/ PlSqlLexer.g4 PlSqlParser.g4
//go:generate sed -i -e "s/self./p./; s/PlSqlLexerBase/PlSqlBaseLexer/" plsql/plsql_lexer.go
//go:generate sed -i -e "s/self./p./; s/PlSqlParserBase/PlSqlBaseParser/" plsql/plsql_parser.go
//go:generate sh -c "[ -e ./plsql/plsql_base_lexer.go ] || (cd plsql && wget https://github.com/antlr/grammars-v4/raw/master/sql/plsql/Go/plsql_base_lexer.go)""
//go:generate sh -c "[ -e ./plsql/plsql_base_parser.go ] || (cd plsql && wget https://github.com/antlr/grammars-v4/raw/master/sql/plsql/Go/plsql_base_parser.go)"
//go:generate sh -c "[ -e SQLLexer.g4 ] || wget https://github.com/pgcodekeeper/pgcodekeeper/raw/master/apgdiff/antlr-src/SQLLexer.g4"
//go:generate sh -c "[ -e SQLParser.g4 ] || wget https://github.com/pgcodekeeper/pgcodekeeper/raw/master/apgdiff/antlr-src/SQLParser.g4"
//go:generate java -jar antlr-4.9.3-complete.jar -Dlanguage=Go -o postgresql/ SQLLexer.g4 SQLParser.g4

import (
	"fmt"
	"io/ioutil"
	"log"
	postgresql "plsql-antlr-go/postgresql"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TreeShapeListener struct {
	*postgresql.BaseSQLParserListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (t *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("ENTER", ctx.GetText())
}

func (t *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("EXIT", ctx.GetStop())
}

func main() {

	// 要求全转换为大写，否则会报 token recognition error at: 'x'
	f, err := ioutil.ReadFile("test/report_query_func_insert_tx_tool_key.sql")
	if err != nil {
		log.Fatalln(err)
	}
	sql := strings.ToUpper(string(f))

	lexer := postgresql.NewSQLLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := postgresql.NewSQLParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true

	// Finally parse the expression rule
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), parser.Sql())

}
