package element

import (
	"database/sql"
	"fmt"
)

// Element - структура элемента
type Element struct {
	ID                uint64     `db:"ID" json:"id"`
	Code              nullString `db:"CODE" json:"code"`
	Name              string     `db:"NAME" json:"name"`
	PreviewPicture    nullString `db:"PREVIEW_PICTURE" json:"preview_picture"`
	DetailPicture     nullString `db:"DETAIL_PICTURE" json:"detail_picture"`
	PreviewText       nullString `db:"PREVIEW_TEXT" json:"preview_text"`
	DetailText        nullString `db:"DETAIL_TEXT" json:"detail_text"`
	XMLID             nullString `db:"XML_ID" json:"xml_id"`
	IblockID          uint64     `db:"IBLOCK_ID" json:"iblock_id"`
	IblockSectionID   nullInt64  `db:"IBLOCK_SECTION_ID" json:"iblock_section_id"`
	Active            bitrixBool `db:"ACTIVE" json:"active"`
	ActiveFrom        nullString `db:"ACTIVE_FROM" json:"active_from"`
	ActiveTo          nullString `db:"ACTIVE_TO" json:"active_to"`
	Sort              uint64     `db:"SORT" json:"sort"`
	SearchableContent nullString `db:"SEARCHABLE_CONTENT" json:"searchable_content"`
	DateCreate        nullString `db:"DATE_CREATE" json:"date_create"`
	CreatedBy         uint64     `db:"CREATED_BY" json:"created_by"`
	TimestampX        nullString `db:"TIMESTAMP_X" json:"timestamp_x"`
	ModifiedBy        nullInt64  `db:"MODIFIED_BY" json:"modified_by"`
	ShowCounter       nullInt64  `db:"SHOW_COUNTER" json:"show_counter"`
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

func main() {
	fmt.Println("Start Bitrix Reindex Faset Service")
}
