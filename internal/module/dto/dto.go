package dto

import "github.com/samber/lo"

type States []State

func (s States) Valid() bool {
	for _, state := range s {
		if !state.Valid() {
			return false
		}
	}
	return true
}

type State int

func (s State) Int() int {
	return int(s)
}

func (s State) Valid() bool {
	return lo.Contains([]State{NEW, WORKING, IN_PROGRESS, DECLINED_AFTER_TRIAL, BEFORE_TRIAL, DECLINED_AFTER_LESSONS}, s)
}

const (
	NEW State = 1 << iota
	WORKING
	IN_PROGRESS
	DECLINED_AFTER_TRIAL
	BEFORE_TRIAL
	DECLINED_AFTER_LESSONS
)
