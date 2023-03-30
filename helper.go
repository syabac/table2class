package main

import "strings"

func getGenerator(dsn string) func(GeneratorParam) string {
	if strings.HasPrefix(dsn, "sqlserver") {
		return GenerateSqlServer
	}

	return func(p GeneratorParam) string {
		return "[UNSUPPORTED DATABASE]"
	}
}

func getSourceGenerator(lang string) func(GeneratorParam, []TableColumnInfo) string {
	if lang == "csharp" {
		return GenerateCSharpCode
	}

	if lang == "go" {
		return GenerateGoCode
	}

	return func(param GeneratorParam, infos []TableColumnInfo) string {
		return "[UNSUPPORTED LANGUAGE]"
	}
}
