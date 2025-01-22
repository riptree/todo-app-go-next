import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { TaskForm } from "./TaskForm";

export const TaskManager = () => {
  return (
    <div className="container mx-auto p-6 max-w-4xl">
      <div className="bg-white rounded-lg shadow-lg p-6">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold text-gray-800">タスク管理アプリ</h1>
          <Dialog>
            <DialogTrigger asChild>
              <Button className="pr-6 pl-4">
                <span>＋</span>新規追加
              </Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
              <DialogHeader>
                <DialogTitle className="text-xl font-bold">新規追加</DialogTitle>
              </DialogHeader>
              <TaskForm />
            </DialogContent>
          </Dialog>
        </div>
        
        <div className="space-y-6">
          <div className="border-b pb-4">
            <h2 className="text-xl font-semibold text-gray-700">タスク一覧</h2>
          </div>
          <div className="bg-gray-50 rounded-lg p-4 min-h-[300px]">
            {/* タスクリストがここに表示されます */}
          </div>
        </div>
      </div>
    </div>
  );
};
