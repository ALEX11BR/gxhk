package main

import "sync"

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
