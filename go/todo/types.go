package api

import "sync"

type Todo struct {
	Id      int    `json:"id,omitempty"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}

type Store struct {
	sync.Mutex
	count int
	list  map[int]*Todo
}

func (s *Store) Get(id int) *Todo {
	s.Lock()
	defer s.Unlock()
	return s.list[id]
}

func (s *Store) GetAll() []*Todo {
	s.Lock()
	defer s.Unlock()

	var todos []*Todo
	for _, v := range s.list {
		todos = append(todos, v)
	}
	return todos
}

func (s *Store) Put(t *Todo) {
	s.Lock()
	if t.Id == 0 {
		s.count = s.count + 1
		t.Id = s.count
	}
	s.list[t.Id] = t
	s.Unlock()
}

func (s *Store) Del(id int) {
	s.Lock()
	delete(s.list, id)
	s.Unlock()
}

func (s *Store) Exists(id int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.list[id]
	return ok
}
