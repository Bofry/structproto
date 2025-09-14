package structproto

import (
	"reflect"
	"sort"
)

var (
	emptyFieldFlagSet = FieldFlagSet(nil)
)

type FieldFlagSet []string

func (s *FieldFlagSet) append(values ...string) {
	set := *s
	if len(values) > 0 {
		for _, v := range values {
			i := sort.SearchStrings(set, v)
			var existed = false
			if i < len(set) {
				existed = set[i] == v
			}
			if !existed {
				container := set
				if i < len(container) {
					container = append(container, "")
					copy(container[i+1:], container[i:])
					container[i] = v
				} else {
					container = append(container, v)
				}
				set = container
				*s = set
			}
		}
	}
}

func (s *FieldFlagSet) clone() *FieldFlagSet {
	set := *s
	if !reflect.ValueOf(set).IsZero() {
		var container = make([]string, len(set))
		copy(container, set)
		cloned := FieldFlagSet(container)
		return &cloned
	}
	return &emptyFieldFlagSet
}

func (s *FieldFlagSet) get(index int) (string, bool) {
	if !s.isEmpty() {
		set := *s
		if index >= 0 && index < len(set) {
			return set[index], true
		}
	}
	return "", false
}

func (s *FieldFlagSet) find(predicate func(v string) bool) bool {
	if s.isEmpty() {
		return false
	}

	set := *s
	for _, v := range set {
		found := predicate(v)
		if found {
			return true
		}
	}
	return false
}

func (s *FieldFlagSet) has(v string) bool {
	return s.indexOf(v) != -1
}

func (s *FieldFlagSet) indexOf(v string) int {
	if s.isEmpty() {
		return -1
	}

	set := *s
	i := sort.SearchStrings(set, v)
	if i < len(set) && set[i] == v {
		return i
	}
	return -1
}

func (s *FieldFlagSet) isEmpty() bool {
	return len(*s) == 0
}

func (s *FieldFlagSet) toArray() []string {
	return *s
}

func (s *FieldFlagSet) len() int {
	if s.isEmpty() {
		return 0
	}

	return len(*s)
}

func (s FieldFlagSet) iterate() <-chan string {
	if s.isEmpty() {
		return nil
	}

	c := make(chan string, len(s))
	for _, v := range s {
		c <- v
	}
	close(c)
	return c
}

func (s *FieldFlagSet) remove(v string) bool {
	if !s.isEmpty() {
		index := s.indexOf(v)
		deleted, _ := s.removeIndex(index)
		return deleted
	}
	return false
}

func (s *FieldFlagSet) removeIndex(index int) (bool, string) {
	if !s.isEmpty() {
		set := *s
		if index >= 0 && index < len(set) {
			value := set[index]
			copy(set[index:], set[index+1:])
			set = set[:len(set)-1]
			*s = set

			return true, value
		}
	}
	return false, ""
}
