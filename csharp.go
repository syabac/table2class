package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func GenerateCSharpCode(param GeneratorParam, columns []TableColumnInfo) string {
	code := `using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

`
	code += fmt.Sprintf("[Table(\"%s\")]\n", param.TableName)
	code += fmt.Sprintf("public class %s{\n", param.ClassName)

	for _, col := range columns {
		nullable := ""
		netType := csharpTypeMap(col.TypeName)

		if col.IsNullable {
			nullable = "?"
		} else {
			code += "\t[Required]\n"
		}

		if netType == "string" && col.Size != nil && *col.Size == 1 {
			netType = "char"
		}

		if netType == "string" && col.Size != nil && *col.Size > 1 {
			code += fmt.Sprintf("\t[MaxLength(%d)]\n", *col.Size)
		}

		code += fmt.Sprintf("\t[Column(\"%s\")]\n", col.ColumnName)
		code += fmt.Sprintf("\tpublic %s%s %s { get; set; }\n", netType, nullable, CapitalizeWords(col.ColumnName))
	}

	code += "}"
	return code
}

func csharpTypeMap(typeName string) string {
	regs := map[string]string{
		"bigint":           "long",
		"binary":           "byte",
		"bit":              "bool",
		"char":             "string",
		"date":             "DateTime",
		"datetime":         "DateTime",
		"datetime2":        "DateTime",
		"datetimeoffset":   "DateTime",
		"decimal":          "decimal",
		"float":            "float",
		"geography":        "object",
		"geometry":         "object",
		"hierarchyid":      "object",
		"image":            "object",
		"int":              "int",
		"money":            "decimal",
		"nchar":            "string",
		"ntext":            "string",
		"numeric":          "decimal",
		"nvarchar":         "string",
		"real":             "double",
		"smalldatetime":    "DateTime",
		"smallint":         "short",
		"smallmoney":       "decimal",
		"sql_variant":      "object",
		"sysname":          "object",
		"text":             "string",
		"time":             "TimeSpan",
		"timestamp":        "DateTime",
		"tinyint":          "byte",
		"uniqueidentifier": "Guid",
		"varbinary":        "byte",
		"varchar":          "string",
		"xml":              "object",
	}

	n, x := regs[typeName]

	if x {
		return n
	}

	return "object"
}

func CapitalizeWords(str string) string {
	str = strings.ReplaceAll(str, "_", " ")
	str = cases.Title(language.English, cases.Compact).String(str)

	return strings.ReplaceAll(str, " ", "")
}
