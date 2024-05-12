package model

import "testing"

func TestMetadata(t *testing.T) {
	var fields []*MetadataField
	err := dbClient.DB().Model(&MetadataField{}).Find(&fields).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range fields {
		f.Name = LcFirst(f.Name)
		err = dbClient.DB().Updates(f).Error
		if err != nil {
			t.Fatal(err)
		}
	}
}
