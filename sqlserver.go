package main

import (
	_ "github.com/microsoft/go-mssqldb"
	mssql "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

const sqlQuery = `select (case when ao.[type]='U' then 'TABLE' else 'VIEW' end) TABLE_TYPE, 
	ao.name TABLE_NAME, ac.name COLUMN_NAME, 
	tp.name TYPE_NAME, ac.max_length SIZE, ac.[precision] PRECISION, ac.[scale] SCALE, ac.IS_NULLABLE
from sys.all_objects ao 
join sys.all_columns ac 
	on ao.object_id = ac.object_id 
join sys.types tp
	on ac.system_type_id = tp.system_type_id 
	and ac.user_type_id = tp.user_type_id 
where ao.[type] in( 'U', 'V')
	and ao.schema_id = 1
	and ao.name = ?
order by ao.name, ac.column_id `

func GenerateSqlServer(param GeneratorParam) string {
	dialect := mssql.Open(param.DSN)
	db, err := gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	var columns []TableColumnInfo

	db.Raw(sqlQuery, param.TableName).Scan(&columns)

	return getSourceGenerator(param.Language)(param, columns)
}
