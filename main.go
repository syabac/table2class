package main

import (
	"fmt"
	"github.com/syabac/table2class/config"
	"os"
	"strings"
)

func getArguments() map[string]string {
	m := make(map[string]string)

	if len(os.Args) < 1 {
		return m
	}

	for i := 1; i < len(os.Args); i++ {
		args := strings.Split(os.Args[i], "=")
		k := strings.TrimLeft(args[0], "-")
		m[k] = ""

		if len(args) > 1 {
			m[k] = args[1]
		}
	}

	return m
}

func printUsage() {
	fmt.Println(" ")
	fmt.Println("Example: ")
	fmt.Println(" ")
	fmt.Println("table2class -n=myConnectionName -t=table_name -c=TableNameClass")
	fmt.Println(" ")
	fmt.Println("     -n  = Connection Name that exists in config.json file")
	fmt.Println("     -t  = Table name")
	fmt.Println("     -c  = Entity Class name")
	fmt.Println("     -d  = Database Name (optional)")
	fmt.Println("     -l  = Programming language (optional), default value is csharp. Supported languages are csharp, go.")
}

func main() {
	cfg := config.LoadConfig()
	args := getArguments()
	conName, _ := args["n"]
	tableName, _ := args["t"]
	className, _ := args["c"]
	dbName, _ := args["d"]
	lang, _ := args["l"]
	isError := false

	if conName == "" {
		fmt.Println("[-n] Connection Name is required!")
		isError = true
	}

	if tableName == "" {
		fmt.Println("[-t] Table Name is required!")
		isError = true
	}

	if className == "" {
		fmt.Println("[-c] Class Name is required! ")
		isError = true
	}

	if isError {
		printUsage()
		return
	}

	if lang == "" {
		lang = "csharp"
	}

	dsn, _ := cfg.Connections[conName]

	if dsn == "" {
		fmt.Println("Invalid connection name:", conName)
		return
	}

	gtor := getGenerator(dsn)
	classCode := gtor(GeneratorParam{
		DSN:       dsn,
		DbName:    dbName,
		TableName: tableName,
		ClassName: className,
		Language:  lang,
	})

	fmt.Println(classCode)
}
