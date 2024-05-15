import React from 'react';
import './App.css';

export type TodoType = {
  task: string,
  description: string,
}

function Todo({ task, description }: TodoType) {
  return (
    <div className="todo">
      <div className="todo-details">
        <p className="todo-task">{task}</p>
        <p className="todo-description">{description}</p>
      </div>
    </div>
  );
}

export default Todo;
