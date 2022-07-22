package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache    // Remove me after realization.
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	// fmt.Printf("Inside Set. key %v value %v\n", key, value)
	currentItem, ok := c.items[key]
	// newElement := cacheItem{key: key, value: value}
	if ok {
		// fmt.Printf("True. currentItem before. %v\n", currentItem)

		// новые данные пришли - надо обновить currentItem
		// fmt.Printf("set key is ok. %v %v %T\n", currentItem, currentItem.Value, currentItem.Value)

		currentItem.Value = cacheItem{key: key, value: value}
		// внутри листа создастся новый элемент с новыми ссылками. может пойти по другому ?
		c.queue.MoveToFront(currentItem)
		// node.Value.(*list.Element).Value = KeyPair{key: key, value: value}
		// fmt.Printf("True. currentItem after. %v\n", currentItem)
		return true
	}

	// todo - check for capacity
	// delete if capacity is full

	cacheItemValue := cacheItem{key: key, value: value}
	newCurrentItem := c.queue.PushFront(cacheItemValue)
	c.items[key] = newCurrentItem
	return false
}

// value := node.Value.(*list.Element).Value.(KeyPair).value
func (c *lruCache) Get(key Key) (interface{}, bool) {
	// fmt.Printf("Inside Get. key %v, c.items: %v \n", key, c.items)
	if currentItem, ok := c.items[key]; ok {
		// fmt.Printf("True. currentItem  %v\n", currentItem)
		c.queue.MoveToFront(currentItem)
		return currentItem.Value.(cacheItem).value, ok
	}
	// fmt.Printf("False. \n")
	return nil, false
}

func (c *lruCache) Clear() {
	// delete(c.queue, )
}
