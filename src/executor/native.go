package executor

import "time"

type clock struct{}

func (c *clock) call(executor *Executor, args []any) any {
	return time.Now()
}

func (c *clock) arity() int {
	return 0
}

func Clock() *clock {
	return &clock{}
}
