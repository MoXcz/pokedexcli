package pokecache

import "time"

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		<-ticker.C
		c.mu.Lock()

		now := time.Now()

		// for each tick
		for key, entry := range c.entries {
			age := now.Sub(entry.createdAt)
			if age > interval {
				delete(c.entries, key)
			}
		}

		c.mu.Unlock()
	}
}
