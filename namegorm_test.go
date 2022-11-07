package namegorm

import "testing"

func Test_ColumnName(t *testing.T) {
	ns := NamingStrategy{
		ColumnNameCase: CamelCase,
	}

	if got := ns.ColumnName("table", "ID"); got != "id" {
		t.Errorf("ColumnName(\"table\", \"ID\") expected \"id\" got \"%s\"", got)
	}

	if got := ns.ColumnName("table", "RoleID"); got != "roleId" {
		t.Errorf("ColumnName(\"table\", \"RoleID\") expected \"roleId\" got \"%s\"", got)
	}

}
