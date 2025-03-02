package pokecache

import "time"

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)

	for {
		<-ticker.C
		c.mu.Lock()

		now := time.Now()

		// for each tick
		for key, entry := range c.entries {
			age := now.Sub(entry.createdAt)
			if age > c.interval {
				delete(c.entries, key)
			}
		}

		c.mu.Unlock()
	}
}
