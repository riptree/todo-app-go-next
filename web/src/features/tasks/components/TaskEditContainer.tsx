import { notFound } from "next/navigation"
import { getTask } from "../api"
import { TaskEditForm } from "./TaskEditForm"

export async function TaskEditContainer({ taskId }: { taskId: number }) {
  const task = await getTask(taskId)

  if (!task) {
    notFound()
  }

  return <TaskEditForm task={task} />
}
