import { getTasks } from "../api"
import { TaskList } from "./TaskList"

export async function TaskListContainer() {
  const tasks = await getTasks()

  return <TaskList tasks={tasks} />
}
