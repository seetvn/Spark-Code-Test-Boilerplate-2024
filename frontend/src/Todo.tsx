import React from 'react';
import './App.css';

// type Todo struct {
//   ID   int    `json:"id"`
//   Task string `json:"task"`
// Description string `json:"description"`
// Priority string `json:"priority"`
// Finished bool `json:"finished"`
// }
export type TodoType = {
  id: number
  task: string,
  description: string,
  priority: string,
  onDelete: () => void // Callback function for delete action
}

function Todo({ id, task, description, priority, onDelete }: TodoType) {
  return (
    <div className="todo">
      <div className="todo-details">
          <p className="todo-task">{task}</p>
          <p className="todo-description">{description}</p>
          <p className="todo-priority">Priority: {priority}</p>
      </div>
      <button onClick={onDelete}>Delete</button>
    </div>
  );
}

export default Todo;