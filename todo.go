package yalzo

type TodoList struct {
    list []Todo
    labels []string
}

type Todo struct {
    label string
    title string
    isArchived bool
    no int
}

func NewTodoList(l []Todo, ls []string) TodoList{
    tl := &TodoList{
        list: l,
        labels: ls,
    }
    return (*tl)
}

func (tl *TodoList) ChangeLabelName(i int, l string) {
    (*tl).list[i].label = l
}

func (tl *TodoList) Delete(n int) {
    for i := n ; i < len((*tl).list) ; i++ {
        (*tl).list[i].setNumber(i)
    }
    tl.list = append(tl.list[:n], tl.list[n+1:]...)
}

func (tl *TodoList) AddTodo(l string, t string) {
    tl.list = append(tl.list, Todo{
        no:    len(tl.list),
        label: l,
        title: t,
    })
}

func (tl *TodoList) Exchange(i1 int, i2 int) {
    (*tl).list[i2].setNumber(i1+1)
    (*tl).list[i1].setNumber(i2+1)

    tmp := (*tl).list[i2]
    (*tl).list[i2] = (*tl).list[i1]
    (*tl).list[i1] = tmp
}

func (t *Todo) setNumber(n int) {
    (*t).no = n
}
