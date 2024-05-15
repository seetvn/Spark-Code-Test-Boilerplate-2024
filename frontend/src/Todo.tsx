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
  task: string,
  description: string,
  priority: string,

}

function Todo({ task, description, priority }: TodoType) {
  return (
    <div className="todo">
      <div className="todo-details">
        <p className="todo-task">{task}</p>
        <p className="todo-description">{description}</p>
        <p className="todo-priority">Priority : {priority}</p>
      </div>
    </div>
  );
}

export default Todo;
