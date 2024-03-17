package index

// Index - структура индекса

type Index struct {
	section    int64  `db:"SECTION_ID" json:"section"`
	element    int64  `db:"ELEMENT_ID" json:"element"`
	faset      int64  `db:"FACET_ID" json:"faset"`
	value      int64  `db:"VALUE" json:"value"`
	valueEnum  int64  `db:"VALUE_NUM" json:"valueEnum"`
	subsection string `db:"INCLUDE_SUBSECTIONS" json:"subsection"`
}

// IndexValue - структура индекса

type IndexValue struct {
	id    int64  `db:"ID" json:"id"`
	value string `db:"VALUE" json:"value"`
}
