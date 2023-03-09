package maps

type OrderedMap[K comparable, V any] struct {
	keys  []K
	inner map[K]*V
}

func (m *OrderedMap[K, V]) lazyInit() {
	if m.inner == nil {
		m.inner = make(map[K]*V)
	}
}

func (m *OrderedMap[K, V]) Get(key K) (*V, bool) {
	if m.inner == nil {
		return nil, false
	}
	value, ok := m.inner[key]
	return value, ok
}

func (m *OrderedMap[K, V]) Set(key K, value *V) {
	m.lazyInit()
	if _, ok := m.inner[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.inner[key] = value
}

func (m *OrderedMap[K, V]) Remove(key K) {
	if m.inner == nil {
		return
	}
	delete(m.inner, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

func (m *OrderedMap[K, V]) Clear() {
	if m.inner == nil {
		return
	}
	m.inner = nil
	m.keys = []K{}
}

func (m *OrderedMap[K, V]) ContainsKey(key K) bool {
	m.lazyInit()
	_, ok := m.inner[key]
	return ok
}

func (m *OrderedMap[K, V]) Size() int {
	return len(m.keys)
}

func (m *OrderedMap[K, V]) IsEmpty() bool {
	return len(m.keys) == 0
}

func (m *OrderedMap[K, V]) Values() []*V {
	if m.inner == nil {
		return make([]*V, 0)
	}
	values := make([]*V, len(m.keys))
	for i, key := range m.keys {
		values[i] = m.inner[key]
	}
	return values
}

func (m *OrderedMap[K, V]) Keys() []K {
	return m.keys
}

func (m *OrderedMap[K, V]) Each(f func(key K, value *V)) {
	if m.inner == nil {
		return
	}
	for _, key := range m.keys {
		f(key, m.inner[key])
	}
}
