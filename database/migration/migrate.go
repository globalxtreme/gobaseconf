package migration

import (
	"github.com/globalxtreme/gobaseconf/config"
	"gorm.io/gorm"
)

func Migrate(tables []Table, columns []Column) {
	var migration *gorm.DB
	var migrator gorm.Migrator

	for _, table := range tables {
		migration = config.SetMigration(table.Connection)

		if table.CreateTable != nil {
			migrator = migration.Table(table.CreateTable.TableName()).Migrator()
			if !migrator.HasTable(table.CreateTable) {
				migrator.CreateTable(table.CreateTable)
			}
		}

		if len(table.RenameTable.Old) > 0 {
			migrator = migration.Table(table.RenameTable.Old).Migrator()
			if migrator.HasTable(table.RenameTable.Old) {
				migrator.RenameTable(table.RenameTable.Old, table.RenameTable.New)
			}
		}

		if len(table.DropTable) > 0 {
			migrator = migration.Table(table.DropTable).Migrator()
			if migrator.HasTable(table.DropTable) {
				migrator.DropTable(table.DropTable)
			}
		}

	}

	for _, column := range columns {
		migration = config.SetMigration(column.Connection)

		if len(column.RenameColumns) > 0 {
			for _, rename := range column.RenameColumns {
				if migrator.HasColumn(column.Model, rename.Old) {
					migrator.RenameColumn(column.Model, rename.Old, rename.New)
				}
			}
		}

		if len(column.AddColumns) > 0 {
			for _, add := range column.AddColumns {
				if !migrator.HasColumn(column.Model, add) {
					migrator.AddColumn(column.Model, add)
				}
			}
		}

		if len(column.DropColumns) > 0 {
			for _, drop := range column.DropColumns {
				if migrator.HasColumn(column.Model, drop) {
					migrator.DropColumn(column.Model, drop)
				}
			}
		}

		if len(column.AlterColumns) > 0 {
			for _, alter := range column.AlterColumns {
				if migrator.HasColumn(column.Model, alter) {
					migrator.AlterColumn(column.Model, alter)
				}
			}
		}
	}
}
