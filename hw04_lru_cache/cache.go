package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	keys     map[*ListItem]Key
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		keys:     make(map[*ListItem]Key, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok {
		item.Value = value
		c.queue.MoveToFront(item)
		return true
	}
	listItem := c.queue.PushFront(value)
	c.items[key] = listItem
	c.keys[listItem] = key

	if c.queue.Len() > c.capacity {
		back := c.queue.Back()
		c.queue.Remove(back)
		delete(c.items, c.keys[back])
		delete(c.keys, back)
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.keys = make(map[*ListItem]Key, c.capacity)
}
