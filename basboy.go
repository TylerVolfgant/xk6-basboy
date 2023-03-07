package basboy

import (
	"fmt"
	"time"
    "regexp"
    "strings"
	"strconv"
	"math/rand"
	"sync/atomic"

	"go.k6.io/k6/js/modules"

)

var counter int64
var realCounter int64
var realCounterN int64
func init() {
	modules.Register("k6/x/basboy", new(BASBOY))
}
// BASBOY is the k6 extension
type BASBOY struct{}
//1 random Int between 2 int
func (*BASBOY) RandomIntBetween(min, max int) (int, error) {
	return rand.Intn(max-min+1) + min, nil
}
//2 random Item from array
func (*BASBOY) RandomItem(arrayOfItems []interface{}) (interface{}, error) {
	return arrayOfItems[rand.Intn(len(arrayOfItems))], nil
}
/*
    Counters
    only one element will be used by each VU
    the array doesn't need to be sharded between the VUs it will "dynamically" balance between them even if some elements take longer to process
*/
//3 global counter from 1
func (*BASBOY) CounterGlobal() (int64, error) {
	return atomic.AddInt64(&realCounter, 1), nil
}
//4 global counter from start with increment
func (*BASBOY) CounterPRGS(start int64, increment int64) func() int64 {
    if increment <= 0 {
        panic(fmt.Sprintf("Invalid increment: %s", increment))
    }
    atomic.StoreInt64(&realCounterN, start)
    return func() int64 {
        return atomic.AddInt64(&realCounterN, increment) - increment
    }
}
//5 global counter with format length
func (*BASBOY) CounterFormat(startingValue int64, increment int64, maxValue int64, format string) func() string {
    if increment <= 0 {
        panic(fmt.Sprintf("Invalid increment, increment can't be less then zero: %s", increment))
    }
    if startingValue > maxValue {
        panic(fmt.Sprintf("startingValue can't be more then maxValue: %s", maxValue))
    }
    atomic.StoreInt64(&counter, startingValue)
    return func() string {
        current := atomic.LoadInt64(&counter)
        result := strconv.FormatInt(current, 10)
        if format != "" {
            formatStr := fmt.Sprintf("%%0%dd", len(format))
            if num, err := strconv.ParseInt(result, 10, 64); err == nil {
                result = fmt.Sprintf(formatStr, num)
            }
        }
        next := current + increment
        if next > maxValue {
            next = startingValue
        }
        atomic.StoreInt64(&counter, next)
        return result
    }
}
//6 randString generates a random string of given length and characters.
func randStringGen(length int, chars string) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
//7 random uuid
func (*BASBOY) Uuidv4() (string, error) {
	const uuidFormat = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
	rand.Seed(time.Now().UnixNano())
	uuid := make([]byte, len(uuidFormat))
	for i, c := range uuidFormat {
		if c != 'x' && c != 'y' {
			uuid[i] = byte(c)
			continue
		}
		r := rand.Intn(16)
		if c == 'y' {
			r = r&0x03 | 0x08
		}
		uuid[i] = fmt.Sprintf("%x", r)[0]
	}
	return string(uuid), nil
}
//8 random Alphanumeric string with length, and bool uppercase
func (*BASBOY) Alphanumeric( length int, uppercase bool) (string, error) {
	result := randStringGen(length, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    if uppercase {
        result = strings.ToUpper(result)
    }
    return result, nil
}
//9 random Alphabetic string with length, and bool uppercase
func (*BASBOY) Alphabetic( length int, uppercase bool) (string, error) {
	result := randStringGen(length, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    if uppercase {
        result = strings.ToUpper(result)
    }
    return result, nil
}
//10 random Numeric string with length
func (*BASBOY) Numeric( length int) (string, error) {
	result := randStringGen(length, "0123456789")
    return result, nil
}
//11 random AlphanumericAndSymbolic string with length
func (*BASBOY) AlphanumericAndSymbolic( length int) (string, error) {
	result := randStringGen(length, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
    return result, nil
}
//12 random Hexadecimal string with length
func (*BASBOY) Hexadecimal( length int) (string, error) {
	result := randStringGen(length, "ABCDEF0123456789")
    return result, nil
}
//13 random Hexadecimal string with length, and bool uppercase
func (*BASBOY) RandomString( length int, uppercase bool) (string, error) {
	result := randStringGen(length, "abcdefghijklmnopqrstuvwxyz")
    if uppercase {
        result = strings.ToUpper(result)
    }
    return result, nil
}
//need for Rnow
func getOffset(offset string) (time.Time, error) {
    // Define regular expression pattern to match the offset
    pattern := `^([+\-])(\d+)([a-z]+)$`
    r := regexp.MustCompile(pattern)
    // Check if the offset is in the correct format
    matches := r.FindStringSubmatch(offset)
    if len(matches) == 0 {
        return time.Time{}, fmt.Errorf("Invalid offset format: %s", offset)
    }
    // Extract the value and unit components
    sign := matches[1]
    value, err := strconv.Atoi(matches[2])
    if err != nil {
        return time.Time{}, fmt.Errorf("Invalid offset value: %s", matches[2])
    }
    unit := matches[3]
    // Get the current time
    now := time.Now()
    // Add or subtract the offset
    switch unit {
    case "secs":
        if sign == "+" {
            now = now.Add(time.Second * time.Duration(value))
        } else {
            now = now.Add(time.Second * time.Duration(-value))
        }
    case "mins":
        if sign == "+" {
            now = now.Add(time.Minute * time.Duration(value))
        } else {
            now = now.Add(time.Minute * time.Duration(-value))
        }
    case "hours":
        if sign == "+" {
            now = now.Add(time.Hour * time.Duration(value))
        } else {
            now = now.Add(time.Hour * time.Duration(-value))
        }
    case "days":
        if sign == "+" {
            now = now.AddDate(0, 0, value)
        } else {
            now = now.AddDate(0, 0, -value)
        }
    case "weeks":
        if sign == "+" {
            now = now.AddDate(0, 0, 7*value)
        } else {
            now = now.AddDate(0, 0, -7*value)
        }
    case "months":
        if sign == "+" {
            now = now.AddDate(0, value, 0)
        } else {
            now = now.AddDate(0, -value, 0)
        }
    case "years":
        if sign == "+" {
            now = now.AddDate(value, 0, 0)
        } else {
            now = now.AddDate(-value, 0, 0)
        }
    default:
        return time.Time{}, fmt.Errorf("Invalid offset unit: %s", unit)
    }
    return now, nil
}
//14 time with format, timezone, offset
func (*BASBOY) Rnow(options ...map[string]interface{}) interface{} {
	var format string
	var timezone string
	var offset string

	if len(options) > 0 {
		if v, ok := options[0]["format"].(string); ok {
			format = v
		}
		if v, ok := options[0]["timezone"].(string); ok {
			timezone = v
		}
		if v, ok := options[0]["offset"].(string); ok {
			offset = v
		}
	}

    var t time.Time
    if offset != "" {
        // Call the getDateTime() function to get the time with the offset
        offsetTime, err := getOffset(offset)
        if err != nil {
            panic(err)
        }
        t = offsetTime
    } else {
        t = time.Now()
    }

	if timezone != "" {
		// Load the timezone based on the provided timezone string
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			panic(fmt.Sprintf("Invalid timezone: %s", timezone))
		}

		// Adjust the time to the specified timezone
		t = t.In(loc)
	}

	switch format {
	case "":
		return t.Format(time.RFC3339)
    case "unix":
        return t.Unix()
    case "epoch":
        return t.UnixNano()/int64(time.Millisecond)
	default:
		// Format the time based on the provided format string
		format = strings.ReplaceAll(format, "YYYY", "2006")
		format = strings.ReplaceAll(format, "yyyy", "2006")
		format = strings.ReplaceAll(format, "YY", "06")
		format = strings.ReplaceAll(format, "MMMM", "January")
		format = strings.ReplaceAll(format, "MMM", "Jan")
		format = strings.ReplaceAll(format, "MM", "01")
		format = strings.ReplaceAll(format, "M", "1")
		format = strings.ReplaceAll(format, "DDDD", "Monday")
		format = strings.ReplaceAll(format, "DDD", "Mon")
		format = strings.ReplaceAll(format, "DD", "02")
		format = strings.ReplaceAll(format, "_2", " 2")
		format = strings.ReplaceAll(format, "D", "2")
		format = strings.ReplaceAll(format, "ddd", "002")
		format = strings.ReplaceAll(format, "dd", " 02")
		format = strings.ReplaceAll(format, "__d", " 02")
		format = strings.ReplaceAll(format, "hh", "03")
		format = strings.ReplaceAll(format, "h", "3")
		format = strings.ReplaceAll(format, "am/pm", "PM")
		format = strings.ReplaceAll(format, "mm", "04")
		format = strings.ReplaceAll(format, "m", "4")
		format = strings.ReplaceAll(format, "ss", "05")
		format = strings.ReplaceAll(format, "s", "5")
		format = strings.ReplaceAll(format, "S", "999")
		format = strings.ReplaceAll(format, ".s", ".000000000")
		format = strings.ReplaceAll(format, ".9", ".9")
		format = strings.ReplaceAll(format, ".99", ".99")
		format = strings.ReplaceAll(format, ".999", ".999")
		format = strings.ReplaceAll(format, ".9999", ".9999")
		format = strings.ReplaceAll(format, ".99999", ".99999")
		format = strings.ReplaceAll(format, ".999999", ".999999")
		format = strings.ReplaceAll(format, ".9999999", ".9999999")
		format = strings.ReplaceAll(format, ".99999999", ".99999999")
		format = strings.ReplaceAll(format, ".999999999", ".999999999")

		return t.Format(format)
	}
}
