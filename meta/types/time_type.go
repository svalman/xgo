package types

import (
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

type (
	GermanTime struct {
		time.Time
	}
)

var (
	ErrDateFormat = errors.New("Формат даты не соответствует dd.mm.yyyy")
)

const DMYlayout = "02.01.2006" // dd.mm.yyyy  format

func (c *GermanTime) UnmarshalXMLAttr(attr xml.Attr) error {
	parse, err := time.Parse(DMYlayout, attr.Value)
	if err != nil {
		return fmt.Errorf("Ошибка разбора даты %v %w", attr.Value, ErrDateFormat)
	}
	*c = GermanTime{parse}
	return nil
}

func (c *GermanTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string

	if err := d.DecodeElement(&v, &start); err != nil {
		return errors.New("Не извлечь значение тега " + err.Error())
	}

	parse, err := time.Parse(DMYlayout, v)
	if err != nil {
		return fmt.Errorf("Ошибка разбора даты %v %w", v, ErrDateFormat)
	}
	*c = GermanTime{parse}
	return nil
}
