package model

import "math"

func (c *PageCondition) TotalPage(totalCount int) int {
	if c == nil {
		return 0
	}
	targetCount := 0

	if c.Backward == nil && c.Forward == nil {
		if c.Limit == nil {
			return 0
		} else {
			targetCount = *c.Limit
		}
	} else {
		if c.Backward != nil {
			targetCount = c.Backward.Last
		}
		if c.Forward != nil {
			targetCount = c.Forward.First
		}
	}

	return int(math.Ceil(float64(totalCount) / float64(targetCount)))
}

func (c *PageCondition) MoveToPageNo() int {
	if c == nil {
		return 1
	}
	if c.Backward == nil && c.Forward == nil {
		return c.PageNumber
	}
	if c.Backward != nil {
		if c.PageNumber <= 2 {
			return 1
		}
		return c.PageNumber - 1
	}
	if c.Forward != nil {
		return c.PageNumber + 1
	}
	return 1
}
