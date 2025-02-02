export const TASKS_CACHE_KEY = "tasks"
export const getTaskCacheKey = (id: number) => `${TASKS_CACHE_KEY}-${id}` as const
