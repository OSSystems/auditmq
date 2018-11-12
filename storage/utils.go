package storage

func DetermineMinStatus(statusList []Status) Status {
	lastValue := Sync

	if len(statusList) == 0 {
		return Unknown
	}

	for _, status := range statusList {
		if status < lastValue {
			lastValue = status
		}
	}

	return lastValue
}
