import { getTasks } from "../api"
import { TaskCard } from "./TaskCard"
import { TaskForm } from "./TaskForm"

export async function TaskManager() {
  const tasks = await getTasks()

  return (
    <div className="container mx-auto p-6 max-w-4xl">
      <div className="bg-white rounded-lg shadow-lg p-6">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold text-gray-800">タスク管理アプリ</h1>
          <TaskForm />
        </div>

        <div className="space-y-6">
          <div className="border-b pb-4">
            <h2 className="text-xl font-semibold text-gray-700">タスク一覧</h2>
          </div>
          <div className="bg-gray-50 rounded-lg p-4 min-h-[300px]">
            {tasks.length > 0 ? (
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
        </div>
      </div>
    </div>
  )
}
