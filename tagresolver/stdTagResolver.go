package tagresolver

import (
	"fmt"
	"strings"

	"github.com/Bofry/structproto/common"
)

var _ common.TagResolver = StdTagResolver

func StdTagResolver(fieldname, token string) (*common.Tag, error) {
	if len(token) > 0 {
		parts := strings.SplitN(token, ";", 2)
		var desc string
		if len(parts) == 2 {
			parts, desc = strings.Split(parts[0], ","), parts[1]
		} else {
			parts = strings.Split(token, ",")
		}
		name, flags := parts[0], parts[1:]

		if len(flags) == 0 {
			for ii := 0; ii < len(name); ii++ {
				ch := name[ii]
	
				if ch == '_' || ch == '-' ||
					(ch >= 'a' && ch <= 'z') ||
					(ch >= 'A' && ch <= 'Z') ||
					(ch >= '0' && ch <= '9') {
					name = name[ii:]
					break
				}
	
				switch ch {
				case '*':
					flags = append(flags, common.RequiredFlag)
				default:
					return nil, fmt.Errorf("unknow attribute symbol '%c'", ch)
				}
			}
		}

		var tag *common.Tag
		if len(name) > 0 && name != "-" {
			tag = &common.Tag{
				Name:  name,
				Flags: flags,
				Desc:  desc,
			}
		}
		return tag, nil
	}
	return nil, nil
}
