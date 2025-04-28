package pokecache

import (
	"time"
	"sync"

	"github.com/EmmanuelBaranyk7/Pokecache-Go/internal/pokecache"
)

type Cache struct {
	entries		map[string]cacheEntry
	mu			*sync.Mutex 
	time	 	time.Duration
}

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) Cache{
	cache := &Cache{
		entries:	 make(map[string]cacheEntry),
		mu:			 sync.RWMutex{},
		time:		 interval,
	}

	go cache.reapLoop()

	return cache
}

func (cache Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: value,
	}
	cache.mu.Unlock()
}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	if entry, exists := cache.entries[key]; exists {
		return entry.val, true
	}
	return nil, false
}

func (cache Cache) reapLoop() {
	clock := time.NewTicker(cache.time)
	for {
		<-clock.C
		cache.mu.Lock()
		for name, entry := range cache.entries {
			if time.Since(entry.createdAt) > cache.time{
				delete(cache.entries, name)
			}
		}
		cache.mu.Unlock()
	}
}

