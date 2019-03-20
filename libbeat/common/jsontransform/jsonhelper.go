package jsontransform

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/wangjia184/beats/libbeat/common"
	"github.com/wangjia184/beats/libbeat/logp"
)

var rex1 = regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})[^\d]*(?P<hour>\d{2})\:(?P<min>\d{2})\:(?P<sec>\d{2})`)
var rex2 = regexp.MustCompile(`Date\((?P<unixtimestamp>\d{10,})`)

func WriteJSONKeys(event common.MapStr, keys map[string]interface{}, overwriteKeys bool, errorKey string) {
	for k, v := range keys {
		if overwriteKeys {
			if k == "@timestamp" {
				vstr, ok := v.(string)
				if !ok {
					logp.Err("JSON: Won't overwrite @timestamp because value is not string")
					event[errorKey] = "@timestamp not overwritten (not string)"
					continue
				}

				// @timestamp must be of format RFC3339
				ts, err := safeParseTime(vstr)
				if err != nil {
					logp.Err("JSON: Won't overwrite @timestamp because of parsing error: %v", err)
					event[errorKey] = fmt.Sprintf("@timestamp not overwritten (parse error on %s)", vstr)
					continue
				}
				event[k] = common.Time(ts)
			} else if k == "type" {
				vstr, ok := v.(string)
				if !ok {
					logp.Err("JSON: Won't overwrite type because value is not string")
					event[errorKey] = "type not overwritten (not string)"
					continue
				}
				if len(vstr) == 0 || vstr[0] == '_' {
					logp.Err("JSON: Won't overwrite type because value is empty or starts with an underscore")
					event[errorKey] = fmt.Sprintf("type not overwritten (invalid value [%s])", vstr)
					continue
				}
				event[k] = vstr
			} else {
				event[k] = v
			}
		} else if _, exists := event[k]; !exists {
			event[k] = v
		}

	}
}

func safeParseTime(str string) (time.Time, error) {
	ts, err := time.Parse(time.RFC3339, str)
	if err == nil {
		return ts, nil
	}

	matches := rex1.FindStringSubmatch(str)
	if matches != nil && len(matches) > 6 {
		// [2019-03-20T08:04:38 2019 03 20 08 04 38]
		year, _ := strconv.Atoi(matches[1])
		month, _ := strconv.Atoi(matches[2])
		day, _ := strconv.Atoi(matches[3])
		hour, _ := strconv.Atoi(matches[4])
		min, _ := strconv.Atoi(matches[5])
		sec, _ := strconv.Atoi(matches[6])
		ts = time.Date(year, time.Month(month), day, hour, min, sec, 0, time.UTC)
		return ts, nil
	}

	matches = rex2.FindStringSubmatch(str)
	if matches != nil && len(matches) > 1 {
		//[Date(1553066016000 1553066016000]
		sec, _ := strconv.ParseInt(matches[1], 10, 64)
		if sec > 1000000000000 {
			sec = sec / 1000
		}

		ts = time.Unix(sec, 0)
		return ts, nil
	}

	return ts, err
}
