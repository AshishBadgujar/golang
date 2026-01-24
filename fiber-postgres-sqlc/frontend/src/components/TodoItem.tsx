import type { Todo } from "../types/todo";

interface Props {
    todo: Todo;
    onToggle: (id: number, completed: boolean) => void;
    onDelete: (id: number) => void;
}

export function TodoItem({ todo, onToggle, onDelete }: Props) {
    return (
        <li className="todo-item">
            <label className="todo-item__left">
                <input
                    className="todo-checkbox"
                    type="checkbox"
                    checked={todo.completed}
                    onChange={(e) => onToggle(todo.id, e.target.checked)}
                />
                <span className={todo.completed ? "todo-title todo-title--done" : "todo-title"}>
                    {todo.title}
                </span>
            </label>

            <div className="todo-actions">
                <button
                    className="todo-delete-btn"
                    type="button"
                    aria-label={`Delete todo: ${todo.title}`}
                    onClick={() => onDelete(todo.id)}
                >
                    ×
                </button>
            </div>
        </li>
    );
}
