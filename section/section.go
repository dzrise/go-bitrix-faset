package section

import (
	"database/sql"
)

// Section - структура категория

type Section struct {
	ID                uint64     `db:"ID" json:"id"`
	Code              nullString `db:"CODE" json:"code"`
	Name              string     `db:"NAME" json:"name"`
	Picture           nullString `db:"PICTURE" json:"picture"`
	Description       nullString `db:"DESCRIPTION" json:"description"`
	XMLID             nullString `db:"XML_ID" json:"xml_id"`
	IblockID          uint64     `db:"IBLOCK_ID" json:"iblock_id"`
	IblockSectionID   nullInt64  `db:"IBLOCK_SECTION_ID" json:"iblock_section_id"`
	Active            bitrixBool `db:"ACTIVE" json:"active"`
	Sort              uint64     `db:"SORT" json:"sort"`
	DepthLevel        uint64     `db:"DEPTH_LEVEL" json:"depth_level"`
	SearchableContent nullString `db:"SEARCHABLE_CONTENT" json:"searchable_content"`
	DateCreate        nullString `db:"DATE_CREATE" json:"date_create"`
	CreatedBy         uint64     `db:"CREATED_BY" json:"created_by"`
	TimestampX        nullString `db:"TIMESTAMP_X" json:"timestamp_x"`
	ModifiedBy        nullInt64  `db:"MODIFIED_BY" json:"modified_by"`
	Elements          []uint64   `json:"elements"`
}

type nullInt64 struct {
	sql.NullInt64
}

type nullFloat64 struct {
	sql.NullFloat64
}

type nullString struct {
	sql.NullString
}

type bitrixBool struct {
	sql.NullString
}
