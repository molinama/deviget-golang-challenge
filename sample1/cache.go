package sample1

/*
Packege to get prices using a transparent cache.
*/

import (
	"fmt"
	"time"
)

// PriceService is a service that we can use to get prices for the items
// Calls to this service are expensive (they take time)
type PriceService interface {
	GetPriceFor(itemCode string) (float64, error)
}

// TransparentCache is a cache that wraps the actual service
// The cache will remember prices we ask for, so that we don't have to wait on every call
// Cache should only return a price if it is not older than "maxAge", so that we don't get stale prices
type TransparentCache struct {
	actualPriceService PriceService
	maxAge             time.Duration
	prices             map[string]Price
}

type Price struct {
	price     float64
	timestamp time.Time
}

type PriceFromCh struct {
	price float64
	err   error
}

func NewTransparentCache(actualPriceService PriceService, maxAge time.Duration) *TransparentCache {
	return &TransparentCache{
		actualPriceService: actualPriceService,
		maxAge:             maxAge,
		prices:             map[string]Price{},
	}
}

// GetPriceFor gets the price for the item, either from the cache or the actual service if it was not cached or too old
func (c *TransparentCache) GetPriceFor(itemCode string) (float64, error) {
	if cachedPrice, ok := c.prices[itemCode]; ok {
		// DONE: check that the price was retrieved less than "maxAge" ago!
		if time.Since(cachedPrice.timestamp) <= c.maxAge {
			return cachedPrice.price, nil
		}
	}
	price, err := c.actualPriceService.GetPriceFor(itemCode)
	if err != nil {
		return 0, fmt.Errorf("getting price from service : %v", err.Error())
	}
	c.prices[itemCode] = Price{price, time.Now()}
	return c.prices[itemCode].price, nil
}

func (c *TransparentCache) AsyncGetPriceFor(itemCode string, pricesCh chan PriceFromCh) {
	price, err := c.GetPriceFor(itemCode)
	pricesCh <- PriceFromCh{price, err}
}

// GetPricesFor gets the prices for several items at once, some might be found in the cache, others might not
// If any of the operations returns an error, it should return an error as well
func (c *TransparentCache) GetPricesFor(itemCodes ...string) ([]float64, error) {
	results := []float64{}
	pricesCh := make(chan PriceFromCh, len(itemCodes))
	for _, itemCode := range itemCodes {
		// DONE: parallelize this, it can be optimized to not make the calls to the external service sequentially
		go c.AsyncGetPriceFor(itemCode, pricesCh)
	}
	i := 1
	for p := range pricesCh {
		if p.err != nil {
			return []float64{}, p.err
		}
		results = append(results, p.price)
		if i >= len(itemCodes) {
			close(pricesCh)
		}
		i++
	}
	return results, nil
}
