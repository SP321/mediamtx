package conf

import (
	"encoding/json"
	"fmt"
)

// RTSPRangeType is the type used in the Range header.
type RTSPRangeType int

// supported rtsp range types.
const (
	RTSPRangeTypeUndefined RTSPRangeType = iota
	RTSPRangeTypeClock
	RTSPRangeTypeNPT
	RTSPRangeTypeSMPTE
)

// MarshalJSON implements json.Marshaler.
func (d RTSPRangeType) MarshalJSON() ([]byte, error) {
	var out string

	switch d {
	case RTSPRangeTypeClock:
		out = "clock"

	case RTSPRangeTypeNPT:
		out = "npt"

	case RTSPRangeTypeSMPTE:
		out = "smpte"

	case RTSPRangeTypeUndefined:
		out = ""

	default:
		return nil, fmt.Errorf("invalid rtsp range type: %v", d)
	}

	return json.Marshal(out)
}

// UnmarshalJSON implements json.Unmarshaler.
func (d *RTSPRangeType) UnmarshalJSON(b []byte) error {
	var in string
	if err := json.Unmarshal(b, &in); err != nil {
		return err
	}

	switch in {
	case "clock":
		*d = RTSPRangeTypeClock

	case "npt":
		*d = RTSPRangeTypeNPT

	case "smpte":
		*d = RTSPRangeTypeSMPTE

	case "":
		*d = RTSPRangeTypeUndefined

	default:
		return fmt.Errorf("invalid rtsp range type: '%s'", in)
	}

	return nil
}

// UnmarshalEnv implements envUnmarshaler.
func (d *RTSPRangeType) UnmarshalEnv(s string) error {
	return d.UnmarshalJSON([]byte(`"` + s + `"`))
}
