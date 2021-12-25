package daemon

import "sync"

type HotkeyMap struct {
	Map    map[Hotkey]string
	RWLock sync.RWMutex
}

func (m *HotkeyMap) Get(hotkey Hotkey) string {
	m.RWLock.RLock()
	defer m.RWLock.RUnlock()

	return m.Map[hotkey]
}

func (m *HotkeyMap) Set(hotkey Hotkey, command string) {
	m.RWLock.Lock()
	defer m.RWLock.Unlock()

	m.Map[hotkey] = command
}

func (m *HotkeyMap) Delete(hotkey Hotkey) {
	m.RWLock.Lock()
	defer m.RWLock.Unlock()

	delete(m.Map, hotkey)
}

func (m *HotkeyMap) IsEmpty(hotkey Hotkey) bool {
	m.RWLock.RLock()
	defer m.RWLock.RUnlock()

	_, isNotEmpty := m.Map[hotkey]
	return !isNotEmpty
}

func (m *HotkeyMap) ForEach(fn func(Hotkey, string)) {
	m.RWLock.RLock()
	defer m.RWLock.RUnlock()

	for hotkey, description := range m.Map {
		fn(hotkey, description)
	}
}

func NewHotkeyMap() HotkeyMap {
	return HotkeyMap{
		Map: make(map[Hotkey]string),
	}
}
