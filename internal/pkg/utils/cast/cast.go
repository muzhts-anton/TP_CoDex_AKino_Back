package cast

import (
	"encoding/binary"
	"fmt"
	"math"
	"time"

	"github.com/jackc/pgx/pgtype"
)

func ToString(src []byte) string {
	return string(src)
}

func IntToStr(src uint64) string {
	return fmt.Sprint(src)
}

func FlToStr(src float64) string {
	return fmt.Sprintf("%.1f", src)
}

func TimeToStr(src time.Time, withTime bool) string {
	if withTime {
		return src.Format("2006-01-02 15:04:05")
	}
	return src.Format("2006-01-02")
}

func ToUint64(src []byte) uint64 {
	return binary.BigEndian.Uint64(src)
}

func ToFloat64(src []byte) float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(src))
}

func ToTime(src []byte) time.Time {
	tmp := pgtype.Timestamp{}
	tmp.DecodeBinary(nil, src)
	return tmp.Time
}

func ToBool(src []byte) bool {
	tmp := pgtype.Bool{}
	tmp.DecodeBinary(nil, src)
	return tmp.Bool
}
