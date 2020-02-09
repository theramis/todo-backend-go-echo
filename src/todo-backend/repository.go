package main

type TodoRepository interface {
	Create(todo *Todo)
	GetAll() []*Todo
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
