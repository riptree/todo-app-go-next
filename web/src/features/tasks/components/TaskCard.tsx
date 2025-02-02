import { Button } from "@/components/ui/button"
import { FilePenLine } from "lucide-react"
import Link from "next/link"
import type { Task } from "../types"
import { TaskDeleteDialog } from "./TaskDeleteDialog"

export function TaskCard({ task }: { task: Task }) {
  return (
    <div key={task.id} className="bg-white p-4 rounded-lg shadow border">
      <div className="flex justify-between items-start">
        <div>
          <h3 className="font-semibold">{task.title}</h3>
          <p className="text-gray-600 text-sm mt-1">{task.description}</p>
          <div className="mt-2 text-sm text-gray-500">
            期限: {task.due_date ? new Date(task.due_date).toLocaleDateString("ja-JP") : "なし"}
          </div>
        </div>
        <div className="flex gap-2">
          <Link href={`/tasks/${task.id}/edit`}>
            <Button variant="ghost" size="icon">
              <FilePenLine className="h-4 w-4" />
            </Button>
          </Link>
          <TaskDeleteDialog task={task} />
        </div>
      </div>
    </div>
  )
}
