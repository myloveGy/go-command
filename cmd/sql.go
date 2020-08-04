package cmd

import (
	"cobra/internal/sqlstruct"
	"github.com/spf13/cobra"
	"log"
)

var (
	username, password, charset, host, dbType, dbName, tableName string
)

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "SQL 转换和处理",
	Long:  "SQL 转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "SQL 转换",
	Long:  "SQL 转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sqlstruct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}

		dbModel := sqlstruct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v \n", err)
		}

		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v \n", err)
		}

		template := sqlstruct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		pk := template.GetPK(columns)
		err = template.Generate(tableName, pk, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v \n", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库用户名称")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库用户密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库字符集")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入数据库表名称")

}
