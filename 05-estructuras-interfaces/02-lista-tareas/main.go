package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}
func (tl *TaskList) printTaskList() {
	for i, task := range tl.tasks {
		fmt.Println(i, task.name)
	}
}

type Task struct {
	name        string
	description string
	completed   bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripcion: %s\nCompletado: %t\n", t.name, t.description, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}

func main() {
	t1 := Task{"Tarea 1", "Completar curso de go este mes", false}
	t2 := Task{"Tarea 2", "Segunda tarea", false}
	t3 := Task{"Tarea 3", "Ultima tarea", false}

	tl := TaskList{}
	tl.appendTask(&t1)
	tl.appendTask(&t2)
	tl.appendTask(&t3)
	tl.removeTask(1)
	tl.printTaskList()
}
