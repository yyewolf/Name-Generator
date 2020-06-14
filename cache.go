package namegenerator

import (
	"errors"
	"sync"
)

type cacheType struct {
	n []string
	l sync.Mutex
}

//Session is there to keep the caches.
//Nothing is useful for the user besides the functions.
type Session struct {
	lastNames    cacheType
	cityNames    cacheType
	kingdomNames cacheType
	townNames    cacheType
	countryNames cacheType
	adjectiveOne cacheType
	adjectiveTwo cacheType
}

func (c *cacheType) getNameFromCache() (r string, err error) {
	if len(c.n) == 0 {
		err = errors.New("cache empty")
		return
	}
	r = c.n[len(c.n)-1]
	c.n = c.n[:len(c.n)-1]
	return
}

func (c *cacheType) addNamesToCache(names []string) {
	for i := range names {
		if names[i] == "" {
			continue
		}
		c.n = append(c.n, names[i])
	}
	return
}
