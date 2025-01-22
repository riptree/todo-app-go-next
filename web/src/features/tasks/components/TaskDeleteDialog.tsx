"use client"

import { Trash2 } from "lucide-react"
import { useState } from "react"

import { Button } from "@/components/ui/button"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import { deleteTask, revalidateTasks } from "../actions"
import type { Task } from "../types"

export const TaskDeleteDialog = ({ task }: { task: Task }) => {
  const [open, setOpen] = useState(false)

  async function onSubmit() {
    try {
      const result = await deleteTask(task.id)

      if (!result.success) {
        throw new Error(result.error)
      }

      setOpen(false)
      await revalidateTasks()
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button className="text-red-600 hover:text-red-800" variant="ghost" size="icon">
          <Trash2 className="w-5 h-5" />
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle className="text-xl font-bold">タスクの削除</DialogTitle>
        </DialogHeader>
        <div className="py-6">
          <p className="text-center text-muted-foreground">このタスクを削除してもよろしいですか？</p>
          <p className="mt-2 text-center font-medium">{task.title}</p>
        </div>
        <div className="flex justify-end gap-4">
          <Button variant="outline" onClick={() => setOpen(false)}>
            キャンセル
          </Button>
          <Button variant="destructive" onClick={onSubmit}>
            削除する
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}
