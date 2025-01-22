import type { Task } from "../types"
import { TaskEditForm } from "./TaskEditForm"

export function TaskCard({ task }: { task: Task }) {
  return (
    <div key={task.id} className="bg-white p-4 rounded-lg shadow border">
      <div className="flex justify-between items-start">
        <div>
          <h3 className="font-semibold">{task.title}</h3>
          <p className="text-gray-600 text-sm mt-1">{task.description}</p>
          <div className="mt-2 text-sm text-gray-500">期限: {new Date(task.due_date).toLocaleDateString("ja-JP")}</div>
        </div>
        <TaskEditForm task={task} />
      </div>
    </div>
  )
}
