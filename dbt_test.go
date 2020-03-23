package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbttests = []struct {
	name string
	raw  string
	err  string
	msg  DBT
}{
	{
		name: "good sentence",
		raw:  "$DBDBT,168.0,f,51.21,M,28.0,F*3E",
		msg: DBT{
			DepthFeet:    MustParseDecimal("168.0"),
			DepthMeters:  MustParseDecimal("51.21"),
			DepthFathoms: MustParseDecimal("28.0"),
		},
	},
	{
		name: "bad validity",
		raw:  "$DBDBT,168x.0,f,51.21,M,28.0,F*3E",
		err:  "nmea: sentence checksum mismatch [46 != 3E]",
	},
}

func TestDBT(t *testing.T) {
	for _, tt := range dbttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				dbt := m.(DBT)
				dbt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dbt)
			}
		})
	}
}
