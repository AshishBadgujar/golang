import { useEffect, useState } from "react";
import type { Todo } from "../types/todo";
import {
    fetchTodos,
    createTodo,
    toggleTodo,
    deleteTodo,
} from "../api/todos";
import { TodoItem } from "../components/TodoItem";

export function TodoPage() {
    const [todos, setTodos] = useState<Todo[]>([]);
    const [title, setTitle] = useState("");

    async function loadTodos() {
        const data = await fetchTodos();
        setTodos(data);
    }

    useEffect(() => {
        async function fetchData() {
            await loadTodos();
        }
        fetchData();
    }, []);

    async function handleSubmit(e: React.FormEvent) {
        e.preventDefault();
        if (!title.trim()) return;

        await createTodo(title);
        setTitle("");
        loadTodos();
    }

    async function handleToggle(id: number, completed: boolean) {
        await toggleTodo(id, completed);
        loadTodos();
    }

    async function handleDelete(id: number) {
        await deleteTodo(id);
        loadTodos();
    }

    return (
        <>
            <h1>Todo App</h1>

            <form onSubmit={handleSubmit}>
                <input
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    placeholder="New todo"
                />
                <button>Add</button>
            </form>

            <ul>
                {todos.map((todo) => (
                    <TodoItem
                        key={todo.id}
                        todo={todo}
                        onToggle={handleToggle}
                        onDelete={handleDelete}
                    />
                ))}
            </ul>
        </>
    );
}
