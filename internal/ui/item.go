package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/hikkiyomi/passman/internal/databases"
)

type item struct {
	title       string
	description string
	rawContent  databases.Record
}

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return i.description
}

func (i item) FilterValue() string {
	return i.title
}

func MapRecordsToItems(records []databases.Record) []list.Item {
	result := make([]list.Item, 0, len(records))

	for _, record := range records {
		adding := item{
			title:       record.Service,
			description: string(record.Data),
			rawContent:  record,
		}

		result = append(result, adding)
	}

	return result
}

func MapItemToRecord(it item) databases.Record {
	result := it.rawContent

	result.Service = it.title
	result.Data = []byte(it.description)

	return result
}
