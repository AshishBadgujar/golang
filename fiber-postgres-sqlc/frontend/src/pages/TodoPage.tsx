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

    const completedCount = todos.filter((t) => t.completed).length;
    const remainingCount = todos.length - completedCount;

    return (
        <div className="app-shell">
            <div className="app-container">
                <header className="app-header">
                    <div>
                        <h1 className="app-title">Todos</h1>
                        <p className="app-subtitle">
                            {remainingCount} remaining • {completedCount} done
                        </p>
                    </div>
                </header>

                <section className="todo-card">
                    <form className="todo-form" onSubmit={handleSubmit}>
                        <input
                            className="todo-input"
                            value={title}
                            onChange={(e) => setTitle(e.target.value)}
                            placeholder="Add a new todo..."
                            aria-label="New todo title"
                        />
                        <button
                            className="todo-add-btn"
                            type="submit"
                            disabled={!title.trim()}
                        >
                            Add
                        </button>
                    </form>

                    <ul className="todo-list">
                        {todos.length === 0 ? (
                            <li className="todo-empty">
                                No todos yet. Add your first one above.
                            </li>
                        ) : (
                            todos.map((todo) => (
                                <TodoItem
                                    key={todo.id}
                                    todo={todo}
                                    onToggle={handleToggle}
                                    onDelete={handleDelete}
                                />
                            ))
                        )}
                    </ul>
                </section>
            </div>
        </div>
    );
}
