package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type mapEntry struct {
	Key   Key
	Value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	item, ok := lc.items[key]
	if ok {
		item.Value = mapEntry{key, value}
		lc.queue.MoveToFront(item)
	} else {
		item = lc.queue.PushFront(mapEntry{key, value})
		if lc.queue.Len() > lc.capacity {
			delete(lc.items, lc.queue.Back().Value.(mapEntry).Key)
			lc.queue.Remove(lc.queue.Back())
		}
	}
	lc.items[key] = item
	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := lc.items[key]
	if ok {
		lc.queue.MoveToFront(item)
		return item.Value.(mapEntry).Value, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc.queue = NewList()
	lc.items = make(map[Key]*ListItem, lc.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
