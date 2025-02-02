import { TaskEditContainer } from "@/features/tasks/components/TaskEditContainer"
import { Suspense } from "react"

export default async function EditTaskPage({ params }: { params: { id: number } }) {
  const { id } = await params

  return (
    <div className="container mx-auto p-6 max-w-6xl">
      <div className="w-full max-w-2xl mx-auto bg-white rounded-lg shadow-lg p-6">
        <h1 className="text-xl font-semibold text-gray-700 mb-6">タスクの編集</h1>
        <Suspense fallback={<div>Loading...</div>}>
          <TaskEditContainer taskId={id} />
        </Suspense>
      </div>
    </div>
  )
}
