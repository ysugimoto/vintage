package function

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/ysugimoto/vintage/errors"
	"github.com/ysugimoto/vintage/runtime/core"
)

const Strftime_Name = "strftime"

// Golang's time format of AM treats noon as AM,
// but it should be PM in VCL.
func Strftime_ampm(t time.Time) string {
	if t.Hour() == 12 {
		return "PM"
	}
	return t.Format("AM")
}

var Strftime_weeks = map[string]int{
	"Mon": 1,
	"Tue": 2,
	"Wed": 3,
	"Thu": 4,
	"Fri": 5,
	"Sat": 6,
	"Sun": 7,
}

// Fastly built-in function implementation of strftime
// Arguments may be:
// - STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/strftime/
func Strftime[T core.EdgeRuntime](
	ctx *core.Runtime[T],
	format string,
	t time.Time,
) (string, error) {
	var formatted string
	for i := 0; i < len(format); i++ {
		v := format[i]
		if v != 0x25 { // %
			formatted += string(v)
			continue
		}
		if i+1 > len(format)-1 {
			return "", errors.FunctionError(
				Strftime_Name, "Invalid format string: %s", format,
			)
		}
		i++
		vv := format[i]
		switch vv {
		case 0x61: // %a
			formatted += t.Format("Mon")
		case 0x41: // %A
			formatted += t.Format("Monday")
		case 0x62: // %b
			formatted += t.Format("Jan")
		case 0x42: // %B
			formatted += t.Format("January")
		case 0x63: // %c
			formatted += t.Format("Mon Jan _2 15:04:05 2001")
		case 0x43: // %C
			formatted += fmt.Sprint(t.Year() / 100)
		case 0x64: // %d
			formatted += t.Format("02")
		case 0x44: // %D
			formatted += t.Format("01/02/06")
		case 0x65: // %e
			formatted += t.Format("_2")
		case 0x46: // %F
			formatted += t.Format("2006-01-02")
		case 0x67: // %g
			formatted += t.Format("06")
		case 0x47: // %G
			formatted += t.Format("2006")
		case 0x68: // %h
			formatted += t.Format("Jan")
		case 0x48: // %H
			formatted += t.Format("15")
		case 0x49: // %I
			formatted += t.Format("03")
		case 0x6A: // %j
			formatted += t.Format("002")
		case 0x6B: // %k
			formatted += fmt.Sprintf("%2d", t.Hour())
		case 0x6C: // %l
			formatted += fmt.Sprintf("%2d", t.Hour()/2)
		case 0x6D: // %m
			formatted += t.Format("01")
		case 0x4D: // %M
			formatted += t.Format("04")
		case 0x6E: // %n
			formatted += "\n"
		case 0x70: // %p
			formatted += Strftime_ampm(t)
		case 0x50: // %P
			formatted += strings.ToLower(Strftime_ampm(t))
		case 0x72: // %r
			formatted += t.Format("15:04:05") + " " + Strftime_ampm(t)
		case 0x52: // %R
			formatted += t.Format("15:04")
		case 0x73: // %s
			formatted += fmt.Sprint(t.Unix())
		case 0x53: // %S
			formatted += t.Format("05")
		case 0x74: // %t
			formatted += "\t"
		case 0x54: // %T
			formatted += t.Format("15:04:05")
		case 0x75: // %u
			formatted += fmt.Sprint(Strftime_weeks[t.Format("Mon")])
		case 0x55: // %U
			// FIXME: consider the first Sunday as the first day of week 01
			_, week := t.ISOWeek()
			formatted += fmt.Sprint(week)
		case 0x56: // %V
			_, week := t.ISOWeek()
			formatted += fmt.Sprint(week)
		case 0x77: // %w
			formatted += fmt.Sprint(7 - Strftime_weeks[t.Format("Mon")])
		case 0x57: // %W
			// FIXME: consider the first Monday as the first day of week 01
			_, week := t.ISOWeek()
			formatted += fmt.Sprint(week)
		case 0x78: // %x
			formatted += t.Format("01/02/06")
		case 0x58: // %X
			formatted += t.Format("15:04:05")
		case 0x79: // %y
			formatted += t.Format("06")
		case 0x59: // %Y
			formatted += t.Format("2006")
		case 0x7A: // %z
			formatted += t.Format("-0700")
		case 0x5A: // %Z
			formatted += t.Format("MST")
		case 0x25: // %%
			formatted += "%"

		// Modifiers
		case 0x45: // %E
			if i+1 > len(format)-1 {
				return "", errors.FunctionError(
					Strftime_Name, "Invalid format string: %s", format,
				)
			}
			i++
			vvv := format[i]
			switch vvv {
			case 0x63: // %Ec
				formatted += t.Format("Mon Jan _2 15:04:05 2001")
			case 0x43: // %EC
				formatted += fmt.Sprint(t.Year() / 100)
			case 0x6E: // %En
				formatted += "\n"
			case 0x70: // %Ep
				formatted += Strftime_ampm(t)
			case 0x50: // %EP
				formatted += Strftime_ampm(t)
			case 0x72: // %Er
				formatted += t.Format("15:04:05") + " " + Strftime_ampm(t)
			case 0x52: // %ER
				formatted += t.Format("15:04")
			case 0x73: // %Es
				formatted += fmt.Sprint(t.Unix())
			case 0x74: // %Et
				formatted += "\t"
			case 0x54: // %ET
				formatted += t.Format("15:04:05")
			case 0x75: // %Eu
				formatted += fmt.Sprint(Strftime_weeks[t.Format("Mon")])
			case 0x78: // %Ex
				formatted += t.Format("01/02/06")
			case 0x58: // %EX
				formatted += t.Format("15:04:05")
			case 0x79: // %Ey
				formatted += t.Format("06")
			case 0x59: // %EY
				formatted += t.Format("2006")
			case 0x7A: // %Ez
				formatted += t.Format("-0700")
			case 0x5A: // %EZ
				formatted += t.Format("MST")
			default:
				return "", errors.FunctionError(
					Strftime_Name, "Unexpected format token: %s at position %d", []byte{vvv}, i,
				)
			}

		case 0x4F: // %O
			if i+1 > len(format)-1 {
				return "", errors.FunctionError(
					Strftime_Name, "Invalid format string: %s", format,
				)
			}
			i++
			vvv := format[i]
			switch vvv {
			case 0x43: // %OC
				formatted += fmt.Sprint(t.Year() / 100)
			case 0x64: // %Od
				formatted += t.Format("02")
			case 0x65: // %Oe
				formatted += t.Format("_2")
			case 0x67: // %Og
				formatted += t.Format("06")
			case 0x47: // %OG
				formatted += t.Format("2006")
			case 0x48: // %OH
				formatted += t.Format("15")
			case 0x49: // %OI
				formatted += t.Format("03")
			case 0x6A: // %Oj
				formatted += t.Format("002")
			case 0x6B: // %Ok
				formatted += fmt.Sprintf("%2d", t.Hour())
			case 0x6C: // %Ol
				formatted += fmt.Sprintf("%2d", t.Hour()/2)
			case 0x6D: // %Om
				formatted += t.Format("01")
			case 0x4D: // %OM
				formatted += t.Format("04")
			case 0x6E: // %On
				formatted += "\n"
			case 0x70: // %Op
				formatted += Strftime_ampm(t)
			case 0x50: // %OP
				formatted += strings.ToLower(Strftime_ampm(t))
			case 0x72: // %Or
				formatted += t.Format("15:04:05") + " " + Strftime_ampm(t)
			case 0x52: // %OR
				formatted += t.Format("15:04")
			case 0x73: // %Os
				formatted += fmt.Sprint(t.Unix())
			case 0x53: // %OS
				formatted += t.Format("05")
			case 0x74: // %Ot
				formatted += "\t"
			case 0x54: // %OT
				formatted += t.Format("15:04:05")
			case 0x75: // %Ou
				formatted += fmt.Sprint(Strftime_weeks[t.Format("Mon")])
			case 0x55: // %OU
				// FIXME: consider the first Sunday as the first day of week 01
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x56: // %OV
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x77: // %Ow
				formatted += fmt.Sprint(7 - Strftime_weeks[t.Format("Mon")])
			case 0x57: // %OW
				// FIXME: consider the first Monday as the first day of week 01
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x79: // %Oy
				formatted += t.Format("06")
			case 0x7A: // %Oz
				formatted += t.Format("-0700")
			case 0x5A: // %OZ
				formatted += t.Format("MST")
			default:
				return "", errors.FunctionError(
					Strftime_Name, "Unexpected format token: %s at position %d", []byte{vvv}, i,
				)
			}

		// Extensions
		case 0x2D: // %-
			if i+1 > len(format)-1 {
				return "", errors.FunctionError(
					Strftime_Name, "Invalid format string: %s", format,
				)
			}
			i++
			vvv := format[i]
			switch vvv {
			case 0x61: // %-a
				formatted += t.Format("Mon")
			case 0x41: // %-A
				formatted += t.Format("Monday")
			case 0x62: // %-b
				formatted += t.Format("Jan")
			case 0x42: // %-B
				formatted += t.Format("January")
			case 0x63: // %-c
				formatted += t.Format("Mon Jan _2 15:04:05 2001")
			case 0x43: // %-C
				formatted += fmt.Sprint(t.Year() / 100)
			case 0x64: // %-d
				formatted += t.Format("02")
			case 0x44: // %-D
				formatted += t.Format("01/02/06")
			case 0x65: // %-e
				formatted += t.Format("_2")
			case 0x46: // %-F
				formatted += t.Format("2006-01-02")
			case 0x67: // %-g
				formatted += t.Format("06")
			case 0x47: // %-G
				formatted += t.Format("2006")
			case 0x68: // %-h
				formatted += t.Format("Jan")
			case 0x48: // %-H
				formatted += t.Format("15")
			case 0x49: // %-I, without any padding
				formatted += t.Format("3")
			case 0x6A: // %-j
				formatted += t.Format("002")
			case 0x6B: // %-k
				formatted += fmt.Sprintf("%2d", t.Hour())
			case 0x6C: // %-l, without any padding
				formatted += fmt.Sprintf("%d", t.Hour()/2)
			case 0x6D: // %-m, without any padding
				formatted += t.Format("1")
			case 0x4D: // %-M
				formatted += t.Format("04")
			case 0x6E: // %-n
				formatted += "\n"
			case 0x70: // %-p
				formatted += Strftime_ampm(t)
			case 0x50: // %-P
				formatted += strings.ToLower(Strftime_ampm(t))
			case 0x72: // %-r
				formatted += t.Format("15:04:05") + " " + Strftime_ampm(t)
			case 0x52: // %-R
				formatted += t.Format("15:04")
			case 0x73: // %-s
				formatted += fmt.Sprint(t.Unix())
			case 0x53: // %-S
				formatted += t.Format("05")
			case 0x74: // %-t
				formatted += "\t"
			case 0x54: // %-T
				formatted += t.Format("15:04:05")
			case 0x75: // %-u
				formatted += fmt.Sprint(Strftime_weeks[t.Format("Mon")])
			case 0x55: // %-U
				// FIXME: consider the first Sunday as the first day of week 01
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x56: // %-V
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x77: // %-w
				formatted += fmt.Sprint(7 - Strftime_weeks[t.Format("Mon")])
			case 0x57: // %-W
				// FIXME: consider the first Monday as the first day of week 01
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x78: // %-x
				formatted += t.Format("01/02/06")
			case 0x58: // %-X
				formatted += t.Format("15:04:05")
			case 0x79: // %-y
				formatted += t.Format("06")
			case 0x59: // %-Y
				formatted += t.Format("2006")
			case 0x7A: // %-z, without any padding
				offset := t.Format("-0700")
				sign := string(offset[0])
				trimmed := strings.TrimPrefix(offset[1:], "0")
				if trimmed == "" {
					trimmed = "0"
				}
				formatted += sign + trimmed
			case 0x5A: // %-Z
				formatted += t.Format("MST")
			default:
				return "", errors.FunctionError(
					Strftime_Name, "Unexpected format token: %s at position %d", []byte{vvv}, i,
				)
			}
		case 0x5F: // %_
			if i+1 > len(format)-1 {
				return "", errors.FunctionError(
					Strftime_Name, "Invalid format string: %s", format,
				)
			}
			i++
			vvv := format[i]
			switch vvv {
			case 0x61: // %_a
				formatted += t.Format("Mon")
			case 0x41: // %_A
				formatted += t.Format("Monday")
			case 0x62: // %_b
				formatted += t.Format("Jan")
			case 0x42: // %_B
				formatted += t.Format("January")
			case 0x63: // %_c
				formatted += t.Format("Mon Jan _2 15:04:05 2001")
			case 0x43: // %_C
				formatted += fmt.Sprint(t.Year() / 100)
			case 0x64: // %_d
				formatted += t.Format("02")
			case 0x44: // %_D
				formatted += t.Format("01/02/06")
			case 0x65: // %_e
				formatted += t.Format("_2")
			case 0x46: // %_F
				formatted += t.Format("2006-01-02")
			case 0x67: // %_g
				formatted += t.Format("06")
			case 0x47: // %_G
				formatted += t.Format("2006")
			case 0x68: // %_h
				formatted += t.Format("Jan")
			case 0x48: // %_H
				formatted += t.Format("15")
			case 0x49: // %_I, spaces are used for padding
				formatted += fmt.Sprintf("%2d", t.Hour())
			case 0x6A: // %_j
				formatted += t.Format("002")
			case 0x6B: // %_k
				formatted += fmt.Sprintf("%2d", t.Hour())
			case 0x6C: // %_l, spaces are used for padding
				formatted += fmt.Sprintf("%2d", t.Hour()/2)
			case 0x6D: // %_m, spaces are used for padding
				formatted += fmt.Sprintf("%2d", t.Month())
			case 0x4D: // %_M
				formatted += t.Format("04")
			case 0x6E: // %_n
				formatted += "\n"
			case 0x70: // %_p
				formatted += Strftime_ampm(t)
			case 0x50: // %_P
				formatted += strings.ToLower(Strftime_ampm(t))
			case 0x72: // %_r
				formatted += t.Format("15:04:05") + " " + Strftime_ampm(t)
			case 0x52: // %_R
				formatted += t.Format("15:04")
			case 0x73: // %_s
				formatted += fmt.Sprint(t.Unix())
			case 0x53: // %_S
				formatted += t.Format("05")
			case 0x74: // %_t
				formatted += "\t"
			case 0x54: // %_T
				formatted += t.Format("15:04:05")
			case 0x75: // %_u
				formatted += fmt.Sprint(Strftime_weeks[t.Format("Mon")])
			case 0x55: // %_U
				// FIXME: consider the first Sunday as the first day of week 01
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x56: // %_V
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x77: // %_w
				formatted += fmt.Sprint(7 - Strftime_weeks[t.Format("Mon")])
			case 0x57: // %_W
				// FIXME: consider the first Monday as the first day of week 01
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x78: // %_x
				formatted += t.Format("01/02/06")
			case 0x58: // %_X
				formatted += t.Format("15:04:05")
			case 0x79: // %_y
				formatted += t.Format("06")
			case 0x59: // %_Y
				formatted += t.Format("2006")
			case 0x7A: // %_z, spaces are used for padding
				offset := t.Format("-0700")
				sign := string(offset[0])
				padded := []byte(offset[1:])
				for i := range padded {
					if padded[i] != 0x30 { // not "0"
						break
					}
					padded[i] = 0x20 // replace to white space
				}
				if len(bytes.TrimSpace(padded)) == 0 {
					padded = []byte{0x20, 0x20, 0x20, 0x30} // "   0"
				}
				formatted += sign + string(padded)
			case 0x5A: // %_Z
				formatted += t.Format("MST")
			default:
				return "", errors.FunctionError(
					Strftime_Name, "Unexpected format token: %s at position %d", []byte{vvv}, i,
				)
			}

		case 0x30: // %0
			if i+1 > len(format)-1 {
				return "", errors.FunctionError(
					Strftime_Name, "Invalid format string: %s", format,
				)
			}
			i++
			vvv := format[i]
			switch vvv {
			case 0x61: // %0a
				formatted += t.Format("Mon")
			case 0x41: // %0A
				formatted += t.Format("Monday")
			case 0x62: // %0b
				formatted += t.Format("Jan")
			case 0x42: // %0B
				formatted += t.Format("January")
			case 0x63: // %0c
				formatted += t.Format("Mon Jan _2 15:04:05 2001")
			case 0x43: // %0C
				formatted += fmt.Sprint(t.Year() / 100)
			case 0x64: // %0d
				formatted += t.Format("02")
			case 0x44: // %0D
				formatted += t.Format("01/02/06")
			case 0x65: // %0e
				formatted += t.Format("_2")
			case 0x46: // %0F
				formatted += t.Format("2006-01-02")
			case 0x67: // %0g
				formatted += t.Format("06")
			case 0x47: // %0G
				formatted += t.Format("2006")
			case 0x68: // %0h
				formatted += t.Format("Jan")
			case 0x48: // %0H
				formatted += t.Format("15")
			case 0x49: // %0I, zeros are used for padding
				formatted += t.Format("03")
			case 0x6A: // %0j
				formatted += t.Format("002")
			case 0x6B: // %0k
				formatted += fmt.Sprintf("%2d", t.Hour())
			case 0x6C: // %0l, zeros are used for padding
				formatted += fmt.Sprintf("%02d", t.Hour()/2)
			case 0x6D: // %0m, zeros are used for padding
				formatted += fmt.Sprintf("%02d", t.Month())
			case 0x4D: // %0M
				formatted += t.Format("04")
			case 0x6E: // %0n
				formatted += "\n"
			case 0x70: // %0p
				formatted += Strftime_ampm(t)
			case 0x50: // %0P
				formatted += strings.ToLower(Strftime_ampm(t))
			case 0x72: // %0r
				formatted += t.Format("15:04:05") + " " + Strftime_ampm(t)
			case 0x52: // %0R
				formatted += t.Format("15:04")
			case 0x73: // %0s
				formatted += fmt.Sprint(t.Unix())
			case 0x53: // %0S
				formatted += t.Format("05")
			case 0x74: // %0t
				formatted += "\t"
			case 0x54: // %0T
				formatted += t.Format("15:04:05")
			case 0x75: // %0u
				formatted += fmt.Sprint(Strftime_weeks[t.Format("Mon")])
			case 0x55: // %0U
				// FIXME: consider the first Sunday as the first day of week 01
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x56: // %0V
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x77: // %0w
				formatted += fmt.Sprint(7 - Strftime_weeks[t.Format("Mon")])
			case 0x57: // %0W
				// FIXME: consider the first Monday as the first day of week 01
				_, week := t.ISOWeek()
				formatted += fmt.Sprint(week)
			case 0x78: // %0x
				formatted += t.Format("01/02/06")
			case 0x58: // %0X
				formatted += t.Format("15:04:05")
			case 0x79: // %0y
				formatted += t.Format("06")
			case 0x59: // %0Y
				formatted += t.Format("2006")
			case 0x7A: // %0z, zeros are used for padding
				formatted += t.Format("-0700")
			case 0x5A: // %0Z
				formatted += t.Format("MST")
			}
			return "", errors.FunctionError(
				Strftime_Name, "Unexpected format token: %s at position %d", []byte{vvv}, i,
			)
		default:
			return "", errors.FunctionError(
				Strftime_Name, "Unexpected format token: %s at position %d", []byte{vv}, i,
			)
		}
	}

	return formatted, nil
}
