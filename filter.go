package main

import "github.com/kyokomi/twi-list-slack/twitter"

type StreamingFilter interface {
	filter(s twitter.Streaming) bool
}

type DefaultFilter struct {
}

func (f DefaultFilter) filter(_ twitter.Streaming) bool { return true }

type ListIDFilter struct {
	members *twitter.ListsMembers
}

func (f ListIDFilter) filter(s twitter.Streaming) bool {
	if f.members == nil {
		return false
	}

	for _, user := range f.members.Users {
		if s.User.ID == user.ID {
			return true
		}
	}
	return false
}

func NewListIDFilter(tc *twitter.Client, filterListID string) (StreamingFilter, error) {
	if filterListID == "" {
		return &DefaultFilter{}, nil
	}

	var f ListIDFilter
	members, err := tc.Lists.GetMembers(filterListID)
	if err != nil {
		return nil, err
	}

	f.members = members

	return &f, nil
}
