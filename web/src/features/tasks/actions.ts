"use server"

import { revalidateTag } from "next/cache"

import { TASKS_CACHE_KEY, getTaskCacheKey } from "./constants"

export async function createTask(title: string, description: string, due_date: Date) {
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

export async function updateTask(id: number, title: string, description: string, due_date: Date) {
  try {
    const response = await fetch(`${process.env.API_BASE_URL}/tasks/${id}`, {
      method: "PUT",
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
      throw new Error("タスクの更新に失敗しました")
    }

    const data = await response.json()
    return { success: true, data }
  } catch (error) {
    console.error("タスク更新エラー:", error)
    return { success: false, error: "タスクの更新中にエラーが発生しました" }
  }
}

export async function deleteTask(id: number) {
  try {
    const response = await fetch(`${process.env.API_BASE_URL}/tasks/${id}`, {
      method: "DELETE",
    })

    if (!response.ok) {
      throw new Error("タスクの削除に失敗しました")
    }

    return { success: true }
  } catch (error) {
    console.error("タスク削除エラー:", error)
    return { success: false, error: "タスクの削除中にエラーが発生しました" }
  }
}

export async function revalidateTasks() {
  revalidateTag(TASKS_CACHE_KEY)
}

export async function revalidateTask(id: number) {
  revalidateTag(getTaskCacheKey(id))
}

const formatDate = (date: Date) => {
  return date.toISOString().split("T")[0]
}
