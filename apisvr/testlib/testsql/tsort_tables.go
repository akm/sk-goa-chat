package testsql

import (
	"slices"
	"svrlib/sql"
	"testing"

	"github.com/ninedraft/tsort"
	"github.com/stretchr/testify/require"
)

func SortedTableNames(t *testing.T, conn *sql.DB) []string {
	tableNames := []string{}
	{
		rows, err := conn.Query("SHOW TABLES")
		require.NoError(t, err)
		for rows.Next() {
			var tbl string
			err := rows.Scan(&tbl)
			require.NoError(t, err)
			tableNames = append(tableNames, tbl)
		}
	}

	fkMap := map[string][]string{}
	{
		// TABLEN_NAME <> REFERENCED_TABLE_NAME の条件で自己参照を除外しています。
		// 自己参照も循環参照として検出されてしまうため。
		rows, err := conn.Query("SELECT TABLE_NAME, REFERENCED_TABLE_NAME FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE WHERE REFERENCED_TABLE_SCHEMA IS NOT NULL AND TABLE_NAME <> REFERENCED_TABLE_NAME")
		require.NoError(t, err)
		for rows.Next() {
			var tableName, referencedTableName string
			err := rows.Scan(&tableName, &referencedTableName)
			require.NoError(t, err)
			t.Logf("tableName: %s, referencedTableName: %s\n", tableName, referencedTableName)
			if s, ok := fkMap[tableName]; ok {
				fkMap[tableName] = append(s, referencedTableName)
			} else {
				fkMap[tableName] = []string{referencedTableName}
			}
		}
	}

	tableDefs := make(TableDefs, len(tableNames))
	for i, tbl := range tableNames {
		tableDefs[i] = *NewTable(tbl, fkMap[tbl]...)
	}

	// テーブルを追加したら tables.Tables にも定義を追加する必要があります。しないとテストでdropに失敗することがあります。
	names, hasCycle := tableDefs.Graph().Sort()
	require.Falsef(t, hasCycle, "database tables have cycle: %+v", tableDefs)

	return names
}

type TableDef struct {
	Name       string
	References []string
}

func NewTable(Name string, References ...string) *TableDef {
	return &TableDef{Name: Name, References: References}
}

type TableDefs []TableDef

func (s TableDefs) Graph() *TableGraph {
	nodes := []string{}
	dependencyMap := map[string][]string{}
	for _, v := range s {
		if !slices.Contains[[]string, string](nodes, v.Name) {
			nodes = append(nodes, v.Name)
		}
		for _, ref := range v.References {
			if !slices.Contains[[]string, string](nodes, ref) {
				nodes = append(nodes, ref)
			}
		}
		dependencyMap[v.Name] = append(dependencyMap[v.Name], v.References...)
	}
	return &TableGraph{
		Nodes:         nodes,
		DependencyMap: dependencyMap,
	}
}

type TableGraph struct {
	Nodes         []string
	DependencyMap map[string][]string
}

func (g TableGraph) Sort() ([]string, bool) {
	return tsort.Sort(g.Nodes, func(i string) []string { return g.DependencyMap[i] })
}
