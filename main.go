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

import (
	plsql "plsql-antlr-go/plsql"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// type TreeShapeListener struct {
// 	*parser.BaseJSONListener
// }

// func NewTreeShapeListener() *TreeShapeListener {
// 	return new(TreeShapeListener)
// }

// func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println(ctx.GetText())
// }

func main() {
	// input, _ := antlr.NewFileStream(os.Args[1])
	// lexer := parser.NewJSONLexer(input)
	// stream := antlr.NewCommonTokenStream(lexer, 0)
	// p := parser.NewJSONParser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// p.BuildParseTrees = true
	// tree := p.Json()
	// antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	sql := `
		SELECT
			*
		FROM
			"table"
		WHERE
			"column" = :value
	`

	lexer := plsql.NewPlSqlLexer(antlr.NewInputStream(sql))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	parser := plsql.NewPlSqlParser(stream)
	parser.BuildParseTrees = true

	// Finally parse the expression rule

}
