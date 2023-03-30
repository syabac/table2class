package main

import (
	"fmt"
)

func GenerateGoCode(param GeneratorParam, columns []TableColumnInfo) string {
	code := `
import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)`

	code += fmt.Sprintf("\ntype %s struct {\n", param.ClassName)

	for _, col := range columns {
		nullable := ""
		netType := golangTypeMap(col.TypeName)

		if col.IsNullable {
			nullable = "*"
		}

		if netType == "string" && col.Size != nil && *col.Size == 1 {
			netType = "rune"
		}

		code += fmt.Sprintf("\t%s  %s%s `gorm:\"column:%s\"`\n", CapitalizeWords(col.ColumnName), nullable, netType, col.ColumnName)
	}

	code += "}"
	return code
}

func golangTypeMap(typeName string) string {
	//TODO: fix this
	regs := map[string]string{
		"bigint":           "int64",
		"binary":           "byte",
		"bit":              "bool",
		"char":             "string",
		"date":             "time.Time",
		"datetime":         "time.Time",
		"datetime2":        "time.Time",
		"datetimeoffset":   "time.Time",
		"decimal":          "decimal.Decimal",
		"float":            "float32",
		"geography":        "any",
		"geometry":         "any",
		"hierarchyid":      "any",
		"image":            "any",
		"int":              "int32",
		"money":            "decimal.Decimal",
		"nchar":            "string",
		"ntext":            "string",
		"numeric":          "decimal.Decimal",
		"nvarchar":         "string",
		"real":             "float64",
		"smalldatetime":    "time.Time",
		"smallint":         "int16",
		"smallmoney":       "decimal.Decimal",
		"sql_variant":      "any",
		"sysname":          "any",
		"text":             "string",
		"time":             "any",
		"timestamp":        "time.Time",
		"tinyint":          "int8",
		"uniqueidentifier": "uuid.UUID",
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
