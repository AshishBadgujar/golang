import type { Todo } from "../types/todo";

interface Props {
    todo: Todo;
    onToggle: (id: number, completed: boolean) => void;
    onDelete: (id: number) => void;
}

export function TodoItem({ todo, onToggle, onDelete }: Props) {
    return (
        <li>
            <input
                type="checkbox"
                checked={todo.completed}
                onChange={(e) => onToggle(todo.id, e.target.checked)}
            />
            <span style={{ textDecoration: todo.completed ? "line-through" : "" }}>
                {todo.title}
            </span>
            <button onClick={() => onDelete(todo.id)}>❌</button>
        </li>
    );
}
