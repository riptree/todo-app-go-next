import { TaskForm } from "@/features/tasks/components/TaskForm"
import { TaskListContainer } from "@/features/tasks/components/TaskListContainer"
import { Suspense } from "react"

export default async function Home() {
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
          <Suspense
            fallback={
              <div className="flex items-center justify-center p-8">
                <div className="animate-spin h-8 w-8 border-4 border-blue-500 border-t-transparent rounded-full" />
                <span className="ml-3 text-lg text-gray-600">読み込み中...</span>
              </div>
            }
          >
            <TaskListContainer />
          </Suspense>
        </div>
      </div>
    </div>
  )
}
