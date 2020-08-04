package sqlstruct

import (
	"cobra/internal/word"
	"fmt"
	"html/template"
	"os"
)

const structTpl = `type {{.TableName | ToCamelCase }} struct {
{{range .Columns}}{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }} {{ .Name | ToCamelCase }} {{ .Type }} {{ .Tag }} {{ else }} {{.Name | ToCamelCase }} {{end}} {{ $length := len .Comment }} {{ if gt $length 0 }} // {{.Comment}} {{else}} // {{.Name}} {{end}}
{{end}}
}

func (*{{ .TableName | ToCamelCase }}) TableName() string {
	return "{{ .TableName }}"
}

func (*{{ .TableName | ToCamelCase }}) PK() string {
	return "{{ .PK }}"
}
`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     template.HTML
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
	PK        string
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     template.HTML(fmt.Sprintf("`"+`json:"%s" db:"%s"`+"`", column.ColumnName, column.ColumnName)),
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

func (t *StructTemplate) GetPK(tbColumns []*TableColumn) string {
	for _, column := range tbColumns {
		if column.ColumnKey == "PRI" {
			return column.ColumnName
		}
	}

	return "id"
}

func (t *StructTemplate) Generate(tableName, Pk string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
		PK:        Pk,
	}

	return tpl.Execute(os.Stdout, tplDB)
}
