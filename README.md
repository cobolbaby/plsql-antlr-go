# postgresql-antlr-go

经验证，发现 ANTLR 生成的 GO 版本解释器无法工作:

```bash
$ java -jar antlr-4.9.3-complete.jar -Dlanguage=Go -o postgresql/ SQLLexer.g4 SQLParser.g4
error(134): SQLParser.g4:1747:0: symbol action conflicts with generated code in target language or runtime
error(134): SQLParser.g4:3513:0: symbol var conflicts with generated code in target language or runtime
error(134): SQLParser.g4:3518:6: symbol var conflicts with generated code in target language or runtime
error(134): SQLParser.g4:3535:6: symbol var conflicts with generated code in target language or runtime
error(134): SQLParser.g4:3551:11: symbol var conflicts with generated code in target language or runtime
error(134): SQLParser.g4:3552:11: symbol var conflicts with generated code in target language or runtime
error(134): SQLParser.g4:3553:47: symbol var conflicts with generated code in target language or runtime
error(134): SQLParser.g4:3554:46: symbol var conflicts with generated code in target language or runtime
error(134): SQLParser.g4:3555:12: symbol var conflicts with generated code in target language or runtime
error(134): SQLParser.g4:984:33: symbol type conflicts with generated code in target language or runtime
error(134): SQLParser.g4:994:16: symbol type conflicts with generated code in target language or runtime
error(134): SQLParser.g4:1287:9: symbol type conflicts with generated code in target language or runtime
error(134): SQLParser.g4:1675:64: symbol action conflicts with generated code in target language or runtime
error(134): SQLParser.g4:736:45: symbol type conflicts with generated code in target language or runtime
```

另外原生的 Java 版本也无法做到对 PL/pgSQL UDF 的拆解，现阶段建议用 pganalyze/libpg_query 来实现。

