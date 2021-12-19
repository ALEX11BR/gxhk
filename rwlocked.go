package main

import (
	"strings"
	"sync"
)

type RWLockedMap struct {
	Map    map[Hotkey]string
	RWLock sync.RWMutex
}

func (r *RWLockedMap) Get(hotkey Hotkey) string {
	r.RWLock.RLock()
	defer r.RWLock.RUnlock()

	return r.Map[hotkey]
}

func (r *RWLockedMap) Set(hotkey Hotkey, command string) {
	r.RWLock.Lock()
	defer r.RWLock.Unlock()

	r.Map[hotkey] = command
}

func (r *RWLockedMap) Delete(hotkey Hotkey) {
	r.RWLock.Lock()
	defer r.RWLock.Unlock()

	delete(r.Map, hotkey)
}

func (r *RWLockedMap) IsEmpty(hotkey Hotkey) bool {
	r.RWLock.RLock()
	defer r.RWLock.RUnlock()

	_, isNotEmpty := r.Map[hotkey]
	return !isNotEmpty
}

func NewRWLockedMap() RWLockedMap {
	return RWLockedMap{
		Map: make(map[Hotkey]string),
	}
}

type DescriptionsMap struct {
	Map    map[string]string
	RWLock sync.RWMutex
}

func (m *DescriptionsMap) Get(hotkey string) string {
	m.RWLock.RLock()
	defer m.RWLock.RUnlock()

	loweredHotkey := strings.ToLower(hotkey)
	return m.Map[loweredHotkey]
}

func (m *DescriptionsMap) Set(hotkey string, description string) {
	m.RWLock.Lock()
	defer m.RWLock.Unlock()

	loweredHotkey := strings.ToLower(hotkey)
	m.Map[loweredHotkey] = description
}

func (m *DescriptionsMap) Delete(hotkey string) {
	m.RWLock.Lock()
	defer m.RWLock.Unlock()

	loweredHotkey := strings.ToLower(hotkey)
	delete(m.Map, loweredHotkey)
}

func (m *DescriptionsMap) Iter(fn func(string, string)) {
	m.RWLock.RLock()
	defer m.RWLock.RUnlock()

	for hotkey, description := range m.Map {
		fn(hotkey, description)
	}
}

func NewDescriptionsMap() DescriptionsMap {
	return DescriptionsMap{
		Map: make(map[string]string),
	}
}
