package main

import "errors"

type TodoRepository interface {
	Create(todo *Todo)
	GetAll() []*Todo
	Get(id int) (t *Todo, err error)
	Update(*Todo) (err error)
	DeleteAll()
	Delete(id int) (err error)
}

type InMemoryTodoRepository struct {
	Todos  []*Todo
	nextId int
}

func NewInMemoryTodoRepository() TodoRepository {
	t := new(InMemoryTodoRepository)
	t.Todos = make([]*Todo, 0)
	t.nextId = 1
	return t
}

func (r *InMemoryTodoRepository) Create(todo *Todo) {
	todo.Id = r.nextId
	r.Todos = append(r.Todos, todo)
	r.nextId++
}

func (r *InMemoryTodoRepository) GetAll() []*Todo {
	return r.Todos
}

func (r *InMemoryTodoRepository) DeleteAll() {
	r.Todos = make([]*Todo, 0)
}

func (r *InMemoryTodoRepository) Get(id int) (t *Todo, err error) {
	for _, t = range r.Todos {
		if t.Id == id {
			return t, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (r *InMemoryTodoRepository) Delete(id int) (err error) {
	for i, t := range r.Todos {
		if t.Id == id {
			r.Todos = append(r.Todos[:i], r.Todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (r *InMemoryTodoRepository) Update(todo *Todo) (err error) {
	for i, t := range r.Todos {
		if t.Id == todo.Id {
			r.Todos[i] = todo
			return nil
		}
	}
	return errors.New("todo not found")
}
