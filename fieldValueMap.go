package structproto

var _ Iterator = new(FieldValueMap)

type FieldValueMap map[string]interface{}

func (values FieldValueMap) Iterate() <-chan FieldValueEntity {
	c := make(chan FieldValueEntity, 1)
	go func() {
		for k, v := range values {
			c <- FieldValueEntity{k, v}
		}
		close(c)
	}()
	return c
}
