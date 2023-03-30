# table2class : simple model class generator 

Simple tool to generate class/struct model from database table or view 

## usage
`table2class -n=myConnectionName -t=table_name -c=TableNameClass`

     -n  = Connection Name that exists in config.json file
     -t  = Table name
     -c  = Entity Class name
     -d  = Database Name (optional)
     -l  = Programming language (optional), default value is csharp. Supported languages are csharp, go.

## supported languages
- C#
- Go
## supported databases
- sqlserver