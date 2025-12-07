package dto

type State int

const (
	NEW State = 1 << iota
	WORKING
	IN_PROGRESS
	DECLINED_AFTER_TRIAL
	BEFORE_TRIAL
	DECLINED_AFTER_LESSONS
)
