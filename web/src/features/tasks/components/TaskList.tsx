"use client"

import type { Task } from "../types"
import { TaskCard } from "./TaskCard"

export function TaskList({ tasks }: { tasks: Task[] }) {
  return (
    <div className="bg-gray-50 rounded-lg p-4 min-h-[300px]">
      {tasks && tasks.length > 0 ? (
        <div className="space-y-4">
          {tasks.map((task) => (
            <TaskCard key={task.id} task={task} />
          ))}
        </div>
      ) : (
        <div className="flex justify-center items-center h-[300px] text-gray-500">
          <p>タスクがありません</p>
        </div>
      )}
    </div>
  )
}
