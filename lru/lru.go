package lru

import "fmt"

type list struct {
	value any
	key   string
	next  *list
	prev  *list
}

type Cache struct {
	Capacity int
	length   int
	head     *list
	tail     *list
	lookup   map[string]*list
}

func NewCache(capacity int) *Cache {
	return &Cache{
		Capacity: capacity,
		length:   0,
		head:     nil,
		tail:     nil,
		lookup:   make(map[string]*list),
	}
}

func (c *Cache) Set(key string, value any) {
	node, existOnCache := c.lookup[key]

	if existOnCache {
		node = c.lookup[key]
		node.value = value
	} else {
		node = createNode(key, value)
		c.lookup[key] = node
		c.length += 1
	}

	if c.length > c.Capacity {
		c.evictBlock()
	}

	c.moveToEnd(node)
}

func createNode(key string, value any) *list {
	return &list{
		key:   key,
		value: value,
	}
}

func (c *Cache) evictBlock() {
	node := c.head
	c.head = c.head.next
	node.next = nil
	if c.head.prev != nil {
		c.head.prev = nil
	}

	c.length -= 1
	delete(c.lookup, node.key)
}

func (c *Cache) moveToEnd(node *list) {
	if c.head == nil {
		c.head = node
		c.tail = node

		return
	}

	node.prev = c.tail
	c.tail.next = node
	c.tail = node
}

func (c *Cache) Get(key string) any {
	node := c.lookup[key]
	c.breakLink(node)
	c.moveToEnd(node)

	return node.value
}

func (c *Cache) breakLink(node *list) {
	if node.prev == nil {
		c.head = node.next
		node.next.prev = nil
		node.next = nil
	}

	if node.next != nil && node.prev != nil {
		node.prev.next = node.next
		if node.next != nil {
			node.next.prev = node.prev
		}

		node.next = nil
		node.prev = nil
	}
}

func (c *Cache) Print() {
	current := c.head
	arr := []any{}
	for i := 0; i < c.length; i++ {
		arr = append(arr, current.value)
		current = current.next
	}

	fmt.Println(arr)
}
