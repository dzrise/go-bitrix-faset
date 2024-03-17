package properties

// Property - структура свойства

type Property struct {
	ID                uint64     `db:"ID" json:"id"`
	TimestampX        nullString `db:"TIMESTAMP_X" json:"timestamp_x"`
	IblockID          uint64     `db:"IBLOCK_ID" json:"iblock_id"`
	Name              string     `db:"NAME" json:"name"`
	Active            bitrixBool `db:"ACTIVE" json:"active"`
	Sort              uint64     `db:"SORT" json:"sort"`
	Code              nullString `db:"CODE" json:"code"`
	DefaultValue      nullString `db:"DEFAULT_VALUE" json:"defaultValue"`
	PreviewPicture    nullString `db:"PREVIEW_PICTURE" json:"preview_picture"`
	DetailPicture     nullString `db:"DETAIL_PICTURE" json:"detail_picture"`
	PreviewText       nullString `db:"PREVIEW_TEXT" json:"preview_text"`
	DetailText        nullString `db:"DETAIL_TEXT" json:"detail_text"`
	XMLID             nullString `db:"XML_ID" json:"xml_id"`
	IblockSectionID   nullInt64  `db:"IBLOCK_SECTION_ID" json:"iblock_section_id"`
	ActiveFrom        nullString `db:"ACTIVE_FROM" json:"active_from"`
	ActiveTo          nullString `db:"ACTIVE_TO" json:"active_to"`
	SearchableContent nullString `db:"SEARCHABLE_CONTENT" json:"searchable_content"`
	DateCreate        nullString `db:"DATE_CREATE" json:"date_create"`
	CreatedBy         uint64     `db:"CREATED_BY" json:"created_by"`
	ModifiedBy        nullInt64  `db:"MODIFIED_BY" json:"modified_by"`
	ShowCounter       nullInt64  `db:"SHOW_COUNTER" json:"show_counter"`
}
