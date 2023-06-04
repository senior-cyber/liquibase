package xml

import (
	_xml "encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type DatabaseChangeLog struct {
	XMLName   _xml.Name    `xml:"databaseChangeLog"`
	Property  []*Property  `xml:"property"`
	ChangeSet []*ChangeSet `xml:"changeSet"`
}

type Property struct {
	XMLName _xml.Name `xml:"property"`
	Name    *string   `xml:"name,attr,omitempty"`
	Value   *string   `xml:"value,attr,omitempty"`
	Dbms    *string   `xml:"dbms,attr,omitempty"`
}

type ChangeSet struct {
	XMLName                 _xml.Name                  `xml:"changeSet"`
	Author                  *string                    `xml:"author,attr,omitempty"`
	Id                      *string                    `xml:"id,attr,omitempty"`
	CreateTable             []*CreateTable             `xml:"createTable,omitempty"`
	AddNotNullConstraint    []*AddNotNullConstraint    `xml:"addNotNullConstraint,omitempty"`
	AddForeignKeyConstraint []*AddForeignKeyConstraint `xml:"addForeignKeyConstraint,omitempty"`
	AddUniqueConstraint     []*AddUniqueConstraint     `xml:"addUniqueConstraint,omitempty"`
	CreateIndex             []*CreateIndex             `xml:"createIndex,omitempty"`
	AddColumn               []*AddColumn               `xml:"addColumn,omitempty"`
	Insert                  []*Insert                  `xml:"insert,omitempty"`
	AddPrimaryKey           []*AddPrimaryKey           `xml:"addPrimaryKey,omitempty"`
}

type CreateTable struct {
	XMLName     _xml.Name `xml:"createTable"`
	TableName   *string   `xml:"tableName,attr,omitempty"`
	CatalogName *string   `xml:"catalogName,attr,omitempty"`
	Remarks     *string   `xml:"remarks,attr,omitempty"`
	SchemaName  *string   `xml:"schemaName,attr,omitempty"`
	TableSpace  *string   `xml:"tablespace,attr,omitempty"`
	Column      []*Column `xml:"column,omitempty"`
}

type Column struct {
	XMLName             _xml.Name      `xml:"column"`
	AutoIncrement       *bool          `xml:"autoIncrement,attr,omitempty"`
	Name                *string        `xml:"name,attr,omitempty"`
	Type                *string        `xml:"type,attr,omitempty"`
	DefaultValueBoolean *bool          `xml:"defaultValueBoolean,attr,omitempty"`
	Descending          *bool          `xml:"descending,attr,omitempty"`
	Position            *int           `xml:"position,attr,omitempty"`
	AfterColumn         *string        `xml:"afterColumn,attr,omitempty"`
	Constraints         []*Constraints `xml:"constraints,omitempty"`
	Value               *string        `xml:"value,attr,omitempty"`
}

type Constraints struct {
	XMLName               _xml.Name `xml:"constraints"`
	PrimaryKey            *bool     `xml:"primaryKey,attr,omitempty"`
	PrimaryKeyName        *string   `xml:"primaryKeyName,attr,omitempty"`
	Unique                *bool     `xml:"unique,attr,omitempty"`
	UniqueConstraintName  *string   `xml:"uniqueConstraintName,attr,omitempty"`
	Nullable              *bool     `xml:"nullable,attr,omitempty"`
	NotNullConstraintName *string   `xml:"notNullConstraintName,attr,omitempty"`
}

type AddNotNullConstraint struct {
	XMLName          _xml.Name `xml:"addNotNullConstraint"`
	CatalogName      *string   `xml:"catalogName,attr,omitempty"`
	ColumnDataType   *string   `xml:"columnDataType,attr,omitempty"`
	ColumnName       *string   `xml:"columnName,attr,omitempty"`
	ConstraintName   *string   `xml:"constraintName,attr,omitempty"`
	DefaultNullValue *string   `xml:"defaultNullValue,attr,omitempty"`
	SchemaName       *string   `xml:"schemaName,attr,omitempty"`
	TableName        *string   `xml:"tableName,attr,omitempty"`
	Validate         *bool     `xml:"validate,attr,omitempty"`
}

type AddPrimaryKey struct {
	XMLName     _xml.Name `xml:"addPrimaryKey"`
	CatalogName *string   `xml:"catalogName,attr,omitempty"`

	ColumnNames    *string `xml:"columnNames,attr,omitempty"`
	ConstraintName *string `xml:"constraintName,attr,omitempty"`
	ForIndexName   *string `xml:"forIndexName,attr,omitempty"`
	SchemaName     *string `xml:"schemaName,attr,omitempty"`
	TableName      *string `xml:"tableName,attr,omitempty"`
	TableSpace     *string `xml:"tablespace,attr,omitempty"`

	Clustered *bool `xml:"clustered,attr,omitempty"`
	Validate  *bool `xml:"validate,attr,omitempty"`
}

type AddForeignKeyConstraint struct {
	XMLName                    _xml.Name `xml:"addForeignKeyConstraint"`
	BaseColumnNames            *string   `xml:"baseColumnNames,attr,omitempty"`
	BaseTableCatalogName       *string   `xml:"baseTableCatalogName,attr,omitempty"`
	BaseTableName              *string   `xml:"baseTableName,attr,omitempty"`
	BaseTableSchemaName        *string   `xml:"baseTableSchemaName,attr,omitempty"`
	ConstraintName             *string   `xml:"constraintName,attr,omitempty"`
	Deferrable                 *bool     `xml:"deferrable,attr,omitempty"`
	InitiallyDeferred          *bool     `xml:"initiallyDeferred,attr,omitempty"`
	OnDelete                   *string   `xml:"onDelete,attr,omitempty"`
	OnUpdate                   *string   `xml:"onUpdate,attr,omitempty"`
	ReferencedColumnNames      *string   `xml:"referencedColumnNames,attr,omitempty"`
	ReferencedTableCatalogName *string   `xml:"referencedTableCatalogName,attr,omitempty"`
	ReferencedTableName        *string   `xml:"referencedTableName,attr,omitempty"`
	ReferencedTableSchemaName  *string   `xml:"referencedTableSchemaName,attr,omitempty"`
	Validate                   *bool     `xml:"validate,attr,omitempty"`
}

type AddUniqueConstraint struct {
	XMLName           _xml.Name `xml:"addUniqueConstraint"`
	CatalogName       *string   `xml:"catalogName,attr,omitempty"`
	Clustered         *bool     `xml:"clustered,attr,omitempty"`
	ColumnNames       *string   `xml:"columnNames,attr,omitempty"`
	ConstraintName    *string   `xml:"constraintName,attr,omitempty"`
	Deferrable        *bool     `xml:"deferrable,attr,omitempty"`
	Disabled          *bool     `xml:"disabled,attr,omitempty"`
	ForIndexName      *string   `xml:"forIndexName,attr,omitempty"`
	InitiallyDeferred *bool     `xml:"initiallyDeferred,attr,omitempty"`
	SchemaName        *string   `xml:"schemaName,attr,omitempty"`
	TableName         *string   `xml:"tableName,attr,omitempty"`
	TableSpace        *string   `xml:"tablespace,attr,omitempty"`
	Validate          *bool     `xml:"validate,attr,omitempty"`
}

type CreateIndex struct {
	XMLName    _xml.Name `xml:"createIndex"`
	Clustered  *bool     `xml:"clustered,attr,omitempty"`
	IndexName  *string   `xml:"indexName,attr,omitempty"`
	SchemaName *string   `xml:"schemaName,attr,omitempty"`
	TableName  *string   `xml:"tableName,attr,omitempty"`
	TableSpace *string   `xml:"tablespace,attr,omitempty"`
	Unique     *bool     `xml:"unique,attr,omitempty"`
	Column     []*Column `xml:"column,omitempty"`
}

type AddColumn struct {
	XMLName     _xml.Name `xml:"addColumn"`
	CatalogName *string   `xml:"catalogName,attr,omitempty"`
	SchemaName  *string   `xml:"schemaName,attr,omitempty"`
	TableName   *string   `xml:"tableName,attr,omitempty"`
	Column      []*Column `xml:"column,omitempty"`
}

type Xml struct {
	inputLiquibaseFolder  string
	outputLiquibaseFolder string

	notNullKey  string
	foreignKey  string
	uniqueIndex string
	index       string
	primaryKey  string

	notNullKeyRotation  int
	primaryKeyRotation  int
	uniqueIndexRotation int
	indexRotation       int
	foreignKeyRotation  int
}

type Insert struct {
	XMLName   _xml.Name `xml:"insert"`
	TableName *string   `xml:"tableName,attr,omitempty"`
	Column    []*Column `xml:"column,omitempty"`
}

func New(inputLiquibaseFolder string, outputLiquibaseFolder string) Xml {
	return Xml{
		inputLiquibaseFolder:  inputLiquibaseFolder,
		outputLiquibaseFolder: outputLiquibaseFolder,

		index:       "index_",
		notNullKey:  "notnull_key_",
		foreignKey:  "foreign_key_",
		uniqueIndex: "unique_index_",
		primaryKey:  "primary_key_",

		notNullKeyRotation:  1,
		primaryKeyRotation:  1,
		uniqueIndexRotation: 1,
		indexRotation:       1,
		foreignKeyRotation:  1,
	}
}

func (xml *Xml) ParseXml() {
	inputXmls, inputXmlsError := os.ReadDir(xml.inputLiquibaseFolder)
	if inputXmlsError != nil {
		log.Println(inputXmlsError.Error())
		return
	}

	var names []string

	for _, file := range inputXmls {
		names = append(names, file.Name())
	}
	sort.Strings(names)

	for _, name := range names {
		xmlFile, _ := os.Open(filepath.Join(xml.inputLiquibaseFolder, name))

		inputXmlContent, inputXmlContentError := io.ReadAll(xmlFile)
		if inputXmlContentError != nil {
			log.Println(inputXmlContentError.Error())
			return
		}

		var databaseChangeLog DatabaseChangeLog
		unmarshalError := _xml.Unmarshal(inputXmlContent, &databaseChangeLog)
		if unmarshalError != nil {
			log.Println(unmarshalError.Error())
			return
		}

		var changesets = 0
		for _, changeSet := range databaseChangeLog.ChangeSet {
			changesets = changesets + 1
			if changeSet.Id != nil && *changeSet.Id != "" {
				changesets++
			} else {
				basename := name[0 : len(name)-len(filepath.Ext(name))]
				tempName := fmt.Sprintf("%s_%d", basename, changesets)
				changeSet.Id = &tempName
			}
			if changeSet.AddNotNullConstraint != nil {
				for _, item := range changeSet.AddNotNullConstraint {
					if item.ConstraintName != nil && *item.ConstraintName != "" {
						xml.notNullKeyRotation++
					} else {
						xml.notNullKeyRotation++
						tempName := fmt.Sprintf("%s_%d", xml.notNullKey, xml.notNullKeyRotation)
						item.ConstraintName = &tempName
					}
				}
			}
			if changeSet.AddForeignKeyConstraint != nil {
				for _, item := range changeSet.AddForeignKeyConstraint {
					if item.ConstraintName != nil && *item.ConstraintName != "" {
						xml.foreignKeyRotation++
					} else {
						xml.foreignKeyRotation++
						tempName := fmt.Sprintf("%s_%d", xml.foreignKey, xml.foreignKeyRotation)
						item.ConstraintName = &tempName
					}
				}
			}
			if changeSet.AddUniqueConstraint != nil {
				for _, item := range changeSet.AddUniqueConstraint {
					if item.ConstraintName != nil && *item.ConstraintName != "" {
						xml.uniqueIndexRotation++
					} else {
						xml.uniqueIndexRotation++
						tempName := fmt.Sprintf("%s_%d", xml.uniqueIndex, xml.uniqueIndexRotation)
						item.ConstraintName = &tempName
					}
				}
			}
			if changeSet.AddPrimaryKey != nil {
				for _, item := range changeSet.AddPrimaryKey {
					if item.ConstraintName != nil && *item.ConstraintName != "" {
						xml.primaryKeyRotation++
					} else {
						xml.primaryKeyRotation++
						tempName := fmt.Sprintf("%s_%d", xml.primaryKey, xml.primaryKeyRotation)
						item.ConstraintName = &tempName
					}
				}
			}

			if changeSet.CreateIndex != nil {
				for _, item := range changeSet.CreateIndex {
					if item.IndexName != nil && *item.IndexName != "" {
						xml.indexRotation++
					} else {
						xml.indexRotation++
						tempName := fmt.Sprintf("%s_%d", xml.index, xml.indexRotation)
						item.IndexName = &tempName
					}
				}
			}
			if changeSet.AddColumn != nil {
				for _, addColumn := range changeSet.AddColumn {
					if addColumn.Column != nil {
						for _, column := range addColumn.Column {
							xml.processColumn(column)
						}
					}
				}
			}
			if changeSet.CreateTable != nil {
				for _, createTable := range changeSet.CreateTable {
					if createTable.Column != nil {
						for _, column := range createTable.Column {
							xml.processColumn(column)
						}
					}
				}
			}
		}

		outputContent, _ := _xml.MarshalIndent(databaseChangeLog, "", "   ")

		outputContentText := _xml.Header + string(outputContent)

		outputContentText = strings.ReplaceAll(outputContentText, "<databaseChangeLog>", "<databaseChangeLog xmlns=\"http://www.liquibase.org/xml/ns/dbchangelog\"\n                   xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n                   xsi:schemaLocation=\"http://www.liquibase.org/xml/ns/dbchangelog http://www.liquibase.org/xml/ns/dbchangelog/dbchangelog-4.0.xsd\">")
		outputContentText = strings.ReplaceAll(outputContentText, "></column>", "/>")
		outputContentText = strings.ReplaceAll(outputContentText, "></constraints>", "/>")
		outputContentText = strings.ReplaceAll(outputContentText, "></createIndex>", "/>")
		outputContentText = strings.ReplaceAll(outputContentText, "></addUniqueConstraint>", "/>")
		outputContentText = strings.ReplaceAll(outputContentText, "></addPrimaryKey>", "/>")
		outputContentText = strings.ReplaceAll(outputContentText, "></property>", "/>")
		writeFileError := os.WriteFile(filepath.Join(xml.outputLiquibaseFolder, name), []byte(outputContentText), 0644)
		if writeFileError != nil {
			log.Println(writeFileError.Error())
			return
		}
	}
}

func (xml *Xml) processColumn(column *Column) {
	if column.Type != nil && *column.Type != "" {
		_type := *column.Type
		if strings.HasPrefix(_type, "VARCHAR") && strings.Contains(_type, "(") && strings.Contains(_type, ")") {
			b := strings.Index(_type, "(") + 1
			e := strings.LastIndex(_type, ")")
			size, _ := strconv.Atoi(_type[b:e])
			if size > 4000 {
				tempName := "VARCHAR(4000)"
				column.Type = &tempName
			}
		}
	}
	if column.Constraints != nil {
		for _, constraint := range column.Constraints {
			if constraint.Nullable != nil {
				if constraint.NotNullConstraintName != nil && *constraint.NotNullConstraintName != "" {
					xml.notNullKeyRotation++
				} else {
					xml.notNullKeyRotation++
					tempName := fmt.Sprintf("%s_%d", xml.notNullKey, xml.notNullKeyRotation)
					constraint.NotNullConstraintName = &tempName
				}
			}
			if constraint.PrimaryKey != nil {
				if constraint.PrimaryKeyName != nil && *constraint.PrimaryKeyName != "" {
					xml.primaryKeyRotation++
				} else {
					xml.primaryKeyRotation++
					tempName := fmt.Sprintf("%s_%d", xml.primaryKey, xml.primaryKeyRotation)
					constraint.PrimaryKeyName = &tempName
				}
			}
			if constraint.Unique != nil {
				if constraint.UniqueConstraintName != nil && *constraint.UniqueConstraintName != "" {
					xml.uniqueIndexRotation++
				} else {
					xml.uniqueIndexRotation++
					tempName := fmt.Sprintf("%s_%d", xml.uniqueIndex, xml.uniqueIndexRotation)
					constraint.UniqueConstraintName = &tempName
				}
			}
		}
	}
}
