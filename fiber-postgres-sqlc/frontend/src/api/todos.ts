import type { Todo } from "../types/todo";

const BASE_URL = "http://localhost:8080/api";

export async function fetchTodos(): Promise<Todo[]> {
  const res = await fetch(`${BASE_URL}/todos`);
  return res.json();
}

export async function createTodo(title: string): Promise<void> {
  await fetch(`${BASE_URL}/todos`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ title }),
  });
}

export async function toggleTodo(id: number, completed: boolean) {
  await fetch(`${BASE_URL}/todos/${id}`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ completed }),
  });
}

export async function deleteTodo(id: number) {
  await fetch(`${BASE_URL}/todos/${id}`, {
    method: "DELETE",
  });
}
