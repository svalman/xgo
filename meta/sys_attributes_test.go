package meta

import (
	"strings"
	"testing"
)

func TestShortSysAttributes_ValidateShortSysAttr(t *testing.T) {
	type fields struct {
		SysGuid   string
		SysGuidfk string
		SysState  int
		SysRev    int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "корректно заполнено всё",
			fields: fields{
				SysGuid:   "1",
				SysGuidfk: "11",
				SysState:  0,
				SysRev:    1,
			},
			wantErr: false,
		},

		{
			name: "пустой sys_guid",
			fields: fields{
				SysGuid:   "",
				SysGuidfk: "11",
				SysState:  0,
				SysRev:    1,
			},
			wantErr: true,
		},

		{
			name: "слишком длинный sys_guid",
			fields: fields{
				SysGuid:  strings.Repeat("1", MaxGuidLength+1),
				SysState: 0,
				SysRev:   1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ShortSysAttributes{
				SysGuid:   tt.fields.SysGuid,
				SysGuidfk: tt.fields.SysGuidfk,
				SysState:  tt.fields.SysState,
				SysRev:    tt.fields.SysRev,
			}
			if err := s.ValidateShortSysAttr(); (err != nil) != tt.wantErr {
				t.Errorf("Fail test `%s` ValidateShortSysAttr() \nerror = %v, wantErr %v",
					tt.name, err, tt.wantErr)
			}
		})
	}
}
