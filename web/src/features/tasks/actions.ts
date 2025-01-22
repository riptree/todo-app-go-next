"use server"

import { revalidateTag } from "next/cache"

import { TASKS_CACHE_KEY } from "./constants"

export async function createTask(title: string, description: string, due_date: Date) {
  const formatDate = (date: Date) => {
    return date.toISOString().split("T")[0]
  }

  try {
    const response = await fetch(`${process.env.API_BASE_URL}/tasks`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title,
        description,
        due_date: formatDate(due_date),
      }),
    })

    if (!response.ok) {
      throw new Error("タスクの作成に失敗しました")
    }

    const data = await response.json()
    return { success: true, data }
  } catch (error) {
    console.error("タスク作成エラー:", error)
    return { success: false, error: "タスクの作成中にエラーが発生しました" }
  }
}

export async function revalidateTasks() {
  revalidateTag(TASKS_CACHE_KEY)
}
