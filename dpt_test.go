package nmea

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dpttests = []struct {
	name string
	raw  string
	err  string
	msg  DPT
}{
	{
		name: "good sentence",
		raw:  "$DBDPT,51.21,0.00*6A",
		msg: DPT{
			Depth:  MustParseDecimal("51.21"),
			Offset: MustParseDecimal("0.0"),
		},
	},
	{
		name: "bad validity",
		raw:  "$DBDPT,a51.21,0.00*6A",
		err:  "nmea: sentence checksum mismatch [0B != 6A]",
	},
}

func TestDPT(t *testing.T) {
	for _, tt := range dpttests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := Parse(tt.raw)
			if tt.err != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				dpt := m.(DPT)
				dpt.BaseSentence = BaseSentence{}
				assert.Equal(t, tt.msg, dpt)
			}
		})
	}
}
