package structures

import (
	"time"
)

type List struct {
	title        string
	description  string
	createTime   time.Time
	completed    bool
	completeTime time.Time
}

func MakeList(
	title string,
	description string,
	createTime time.Time,
	completed bool,
	completeTime time.Time,
) *List {
	return &List{
		title:        title,
		description:  description,
		createTime:   createTime,
		completed:    completed,
		completeTime: completeTime,
	}
}

func (a *List) GetTitle() string {
	return a.title
}

func (a *List) GetDescription() string {
	return a.description
}

func (a *List) GetCreateTime() time.Time {
	return a.createTime
}

func (a *List) GetCompleted() bool {
	return a.completed
}

func (a *List) GetCompleteTime() time.Time {
	return a.completeTime
}

func (a *List) CompleteTask() {
	a.completed = true
	a.completeTime = time.Now()
}
