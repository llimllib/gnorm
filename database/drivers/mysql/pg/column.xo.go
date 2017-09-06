// Package pg contains the types for schema 'information_schema'.
package pg

// GENERATED BY XO. DO NOT EDIT.

import (
	"github.com/pkg/errors"
)

// ColumnTable is the database name for the table.
const ColumnTable = "information_schema.columns"

// Column represents a row from 'information_schema.columns'.
type Column struct {
	TableCatalog           SQLIdentifier  `yaml:"table_catalog,omitempty"`            // table_catalog
	TableSchema            SQLIdentifier  `yaml:"table_schema,omitempty"`             // table_schema
	TableName              SQLIdentifier  `yaml:"table_name,omitempty"`               // table_name
	ColumnName             SQLIdentifier  `yaml:"column_name,omitempty"`              // column_name
	OrdinalPosition        CardinalNumber `yaml:"ordinal_position,omitempty"`         // ordinal_position
	ColumnDefault          CharacterData  `yaml:"column_default,omitempty"`           // column_default
	IsNullable             YesOrNo        `yaml:"is_nullable,omitempty"`              // is_nullable
	DataType               CharacterData  `yaml:"data_type,omitempty"`                // data_type
	CharacterMaximumLength CardinalNumber `yaml:"character_maximum_length,omitempty"` // character_maximum_length
	CharacterOctetLength   CardinalNumber `yaml:"character_octet_length,omitempty"`   // character_octet_length
	NumericPrecision       CardinalNumber `yaml:"numeric_precision,omitempty"`        // numeric_precision
	NumericScale           CardinalNumber `yaml:"numeric_scale,omitempty"`            // numeric_scale
	DatetimePrecision      CardinalNumber `yaml:"datetime_precision,omitempty"`       // datetime_precision
	CharacterSetName       SQLIdentifier  `yaml:"character_set_name,omitempty"`       // character_set_name
	CollationName          SQLIdentifier  `yaml:"collation_name,omitempty"`           // collation_name
	ColumnType             SQLIdentifier  `yaml:"column_type,omitempty"`              // column_type
}

// Constants defining each column in the table.
const (
	ColumnTableCatalogField           = "table_catalog"
	ColumnTableSchemaField            = "table_schema"
	ColumnTableNameField              = "table_name"
	ColumnColumnNameField             = "column_name"
	ColumnOrdinalPositionField        = "ordinal_position"
	ColumnColumnDefaultField          = "column_default"
	ColumnIsNullableField             = "is_nullable"
	ColumnDataTypeField               = "data_type"
	ColumnCharacterMaximumLengthField = "character_maximum_length"
	ColumnCharacterOctetLengthField   = "character_octet_length"
	ColumnNumericPrecisionField       = "numeric_precision"
	ColumnNumericPrecisionRadixField  = "numeric_precision_radix"
	ColumnNumericScaleField           = "numeric_scale"
	ColumnDatetimePrecisionField      = "datetime_precision"
	ColumnIntervalTypeField           = "interval_type"
	ColumnIntervalPrecisionField      = "interval_precision"
	ColumnCharacterSetCatalogField    = "character_set_catalog"
	ColumnCharacterSetSchemaField     = "character_set_schema"
	ColumnCharacterSetNameField       = "character_set_name"
	ColumnCollationCatalogField       = "collation_catalog"
	ColumnCollationSchemaField        = "collation_schema"
	ColumnCollationNameField          = "collation_name"
	ColumnDomainCatalogField          = "domain_catalog"
	ColumnDomainSchemaField           = "domain_schema"
	ColumnDomainNameField             = "domain_name"
	ColumnUdtCatalogField             = "udt_catalog"
	ColumnUdtSchemaField              = "udt_schema"
	ColumnUdtNameField                = "udt_name"
	ColumnScopeCatalogField           = "scope_catalog"
	ColumnScopeSchemaField            = "scope_schema"
	ColumnScopeNameField              = "scope_name"
	ColumnMaximumCardinalityField     = "maximum_cardinality"
	ColumnDtdIdentifierField          = "dtd_identifier"
	ColumnIsSelfReferencingField      = "is_self_referencing"
	ColumnIsIdentityField             = "is_identity"
	ColumnIdentityGenerationField     = "identity_generation"
	ColumnIdentityStartField          = "identity_start"
	ColumnIdentityIncrementField      = "identity_increment"
	ColumnIdentityMaximumField        = "identity_maximum"
	ColumnIdentityMinimumField        = "identity_minimum"
	ColumnIdentityCycleField          = "identity_cycle"
	ColumnIsGeneratedField            = "is_generated"
	ColumnGenerationExpressionField   = "generation_expression"
	ColumnIsUpdatableField            = "is_updatable"
)

// WhereClauses for every type in Column.
var (
	ColumnTableCatalogWhere           SQLIdentifierField  = "table_catalog"
	ColumnTableSchemaWhere            SQLIdentifierField  = "table_schema"
	ColumnTableNameWhere              SQLIdentifierField  = "table_name"
	ColumnColumnNameWhere             SQLIdentifierField  = "column_name"
	ColumnOrdinalPositionWhere        CardinalNumberField = "ordinal_position"
	ColumnColumnDefaultWhere          CharacterDataField  = "column_default"
	ColumnIsNullableWhere             YesOrNoField        = "is_nullable"
	ColumnDataTypeWhere               CharacterDataField  = "data_type"
	ColumnCharacterMaximumLengthWhere CardinalNumberField = "character_maximum_length"
	ColumnCharacterOctetLengthWhere   CardinalNumberField = "character_octet_length"
	ColumnNumericPrecisionWhere       CardinalNumberField = "numeric_precision"
	ColumnNumericPrecisionRadixWhere  CardinalNumberField = "numeric_precision_radix"
	ColumnNumericScaleWhere           CardinalNumberField = "numeric_scale"
	ColumnDatetimePrecisionWhere      CardinalNumberField = "datetime_precision"
	ColumnIntervalTypeWhere           CharacterDataField  = "interval_type"
	ColumnIntervalPrecisionWhere      CardinalNumberField = "interval_precision"
	ColumnCharacterSetCatalogWhere    SQLIdentifierField  = "character_set_catalog"
	ColumnCharacterSetSchemaWhere     SQLIdentifierField  = "character_set_schema"
	ColumnCharacterSetNameWhere       SQLIdentifierField  = "character_set_name"
	ColumnCollationCatalogWhere       SQLIdentifierField  = "collation_catalog"
	ColumnCollationSchemaWhere        SQLIdentifierField  = "collation_schema"
	ColumnCollationNameWhere          SQLIdentifierField  = "collation_name"
	ColumnDomainCatalogWhere          SQLIdentifierField  = "domain_catalog"
	ColumnDomainSchemaWhere           SQLIdentifierField  = "domain_schema"
	ColumnDomainNameWhere             SQLIdentifierField  = "domain_name"
	ColumnUdtCatalogWhere             SQLIdentifierField  = "udt_catalog"
	ColumnUdtSchemaWhere              SQLIdentifierField  = "udt_schema"
	ColumnUdtNameWhere                SQLIdentifierField  = "udt_name"
	ColumnScopeCatalogWhere           SQLIdentifierField  = "scope_catalog"
	ColumnScopeSchemaWhere            SQLIdentifierField  = "scope_schema"
	ColumnScopeNameWhere              SQLIdentifierField  = "scope_name"
	ColumnMaximumCardinalityWhere     CardinalNumberField = "maximum_cardinality"
	ColumnDtdIdentifierWhere          SQLIdentifierField  = "dtd_identifier"
	ColumnIsSelfReferencingWhere      YesOrNoField        = "is_self_referencing"
	ColumnIsIdentityWhere             YesOrNoField        = "is_identity"
	ColumnIdentityGenerationWhere     CharacterDataField  = "identity_generation"
	ColumnIdentityStartWhere          CharacterDataField  = "identity_start"
	ColumnIdentityIncrementWhere      CharacterDataField  = "identity_increment"
	ColumnIdentityMaximumWhere        CharacterDataField  = "identity_maximum"
	ColumnIdentityMinimumWhere        CharacterDataField  = "identity_minimum"
	ColumnIdentityCycleWhere          YesOrNoField        = "identity_cycle"
	ColumnIsGeneratedWhere            CharacterDataField  = "is_generated"
	ColumnGenerationExpressionWhere   CharacterDataField  = "generation_expression"
	ColumnIsUpdatableWhere            YesOrNoField        = "is_updatable"
)

// QueryOneColumn retrieves a row from 'information_schema.columns' as a Column.
func QueryOneColumn(db XODB, where WhereClause, order OrderBy) (*Column, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, column_name, ordinal_position, column_default, is_nullable, data_type, character_maximum_length, character_octet_length, numeric_precision, numeric_scale, datetime_precision, character_set_name, collation_name, column_type ` +
		`FROM information_schema.columns WHERE (`

	sqlstr := origsqlstr + where.String() + ") " + order.String() + " LIMIT 1"

	c := &Column{}
	err := db.QueryRow(sqlstr, where.Values()...).Scan(&c.TableCatalog, &c.TableSchema, &c.TableName, &c.ColumnName, &c.OrdinalPosition, &c.ColumnDefault, &c.IsNullable, &c.DataType, &c.CharacterMaximumLength, &c.CharacterOctetLength, &c.NumericPrecision, &c.NumericScale, &c.DatetimePrecision, &c.CharacterSetName, &c.CollationName, &c.ColumnType)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return c, nil
}

// QueryColumn retrieves rows from 'information_schema.columns' as a slice of Column.
func QueryColumn(db XODB, where WhereClause, order OrderBy) ([]*Column, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, column_name, ordinal_position, column_default, is_nullable, data_type, character_maximum_length, character_octet_length, numeric_precision, numeric_scale, datetime_precision, character_set_name, collation_name, column_type ` +
		`FROM information_schema.columns WHERE (`

	sqlstr := origsqlstr + where.String() + ") " + order.String()

	var vals []*Column
	q, err := db.Query(sqlstr, where.Values()...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		c := Column{}

		err = q.Scan(&c.TableCatalog, &c.TableSchema, &c.TableName, &c.ColumnName, &c.OrdinalPosition, &c.ColumnDefault, &c.IsNullable, &c.DataType, &c.CharacterMaximumLength, &c.CharacterOctetLength, &c.NumericPrecision, &c.NumericScale, &c.DatetimePrecision, &c.CharacterSetName, &c.CollationName, &c.ColumnType)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &c)
	}
	return vals, nil
}

// AllColumn retrieves all rows from 'information_schema.columns' as a slice of Column.
func AllColumn(db XODB, order OrderBy) ([]*Column, error) {
	const origsqlstr = `SELECT ` +
		`table_catalog, table_schema, table_name, column_name, ordinal_position, column_default, is_nullable, data_type, character_maximum_length, character_octet_length, numeric_precision, numeric_precision_radix, numeric_scale, datetime_precision, interval_type, interval_precision, character_set_catalog, character_set_schema, character_set_name, collation_catalog, collation_schema, collation_name, domain_catalog, domain_schema, domain_name, udt_catalog, udt_schema, udt_name, scope_catalog, scope_schema, scope_name, maximum_cardinality, dtd_identifier, is_self_referencing, is_identity, identity_generation, identity_start, identity_increment, identity_maximum, identity_minimum, identity_cycle, is_generated, generation_expression, is_updatable ` +
		`FROM information_schema.columns`

	sqlstr := origsqlstr + order.String()

	var vals []*Column
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for q.Next() {
		c := Column{}

		err = q.Scan(&c.TableCatalog, &c.TableSchema, &c.TableName, &c.ColumnName, &c.OrdinalPosition, &c.ColumnDefault, &c.IsNullable, &c.DataType, &c.CharacterMaximumLength, &c.CharacterOctetLength, &c.NumericPrecision, &c.NumericScale, &c.DatetimePrecision, &c.CharacterSetName, &c.CollationName, &c.ColumnType)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		vals = append(vals, &c)
	}
	return vals, nil
}