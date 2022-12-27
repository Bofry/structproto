package tagresolver

import "github.com/Bofry/structproto/common"

var _ common.TagResolver = NoneTagResolver

func NoneTagResolver(fieldname, token string) (*common.Tag, error) {
	var tag = &common.Tag{
		Name: fieldname,
	}
	return tag, nil
}
