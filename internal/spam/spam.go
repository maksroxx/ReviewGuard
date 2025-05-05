package spam

import (
	"sync"
	"time"
)

type IPActiviry struct {
	Timestamps []time.Time
	mu         sync.Mutex
}

var activityMap = make(map[string]*IPActiviry)

func IsSuspicious(ip string) bool {
	now := time.Now()
	act, exists := activityMap[ip]
	if !exists {
		act = &IPActiviry{}
		activityMap[ip] = act
	}
	act.mu.Lock()
	defer act.mu.Unlock()

	cut := now.Add(-1 * time.Minute)
	var recent []time.Time
	for _, t := range act.Timestamps {
		if t.After(cut) {
			recent = append(recent, t)
		}
	}
	act.Timestamps = append(act.Timestamps, now)
	return len(act.Timestamps) > 5
}
