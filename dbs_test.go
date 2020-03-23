package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbstests = []struct {
	name string
	raw  string
	err  string
	msg  DBS
}{
	{
		name: "good sentence",
		raw:  "$DBDBS,168.0,f,51.21,M,28.0,F*39",
		msg: DBS{
			DepthFeet:    MustParseDecimal("168.0"),
			DepthMeters:  MustParseDecimal("51.21"),
			DepthFathoms: MustParseDecimal("28.0"),
		},
	},
	{
		name: "bad validity",
		raw:  "$DBDBS,168.1,f,51.21,M,28.0,F*39",
		err:  "nmea: sentence checksum mismatch [38 != 39]",
	},
}

func TestDBS(t *testing.T) {
	for _, tt := range dbstests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				dbs := m.(DBS)
				dbs.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dbs)
			}
		})
	}
}
