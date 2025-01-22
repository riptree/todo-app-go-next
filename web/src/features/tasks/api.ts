import { TASKS_CACHE_KEY } from "./constants"
import type { Task } from "./types"

export async function getTasks() {
  const apiUrl = process.env.API_BASE_URL || "http://localhost:8082"
  const res = await fetch(`${apiUrl}/tasks`, { next: { tags: [TASKS_CACHE_KEY] } })

  if (!res.ok) {
    throw new Error("Failed to fetch tasks")
  }

  const data = await res.json()
  return data.tasks as Task[]
}
