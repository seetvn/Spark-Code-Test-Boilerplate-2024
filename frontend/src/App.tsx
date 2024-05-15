import React, { useEffect, useState, FormEvent, useRef } from 'react';
import './App.css';
import Todo, { TodoType } from './Todo';
import { stringify } from 'querystring';

function App() {
  const [todos, setTodos] = useState<TodoType[]>([]);
  const [reloadTodos, setReloadTodos] = useState(false); // Add reloadTodos state
  const formRef = useRef<HTMLFormElement>(null); // Create a ref for the form

  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const response = await fetch('http://localhost:8080/todos/');
        if (!response.ok) {
          throw new Error('Error fetching data');
        }
        const fetchedTodos = await response.json();
        setTodos(fetchedTodos);
      } catch (error) {
        console.error('Could not connect to server. Ensure it is running. ' + error);
      }
    };

    fetchTodos();
  }, [reloadTodos]); // Add reloadTodos to the dependency array

  // Function to reload todos
  const handleReloadTodos = () => {
    setReloadTodos(prevState => !prevState); // Toggle reloadTodos
  };

  // function to make POST request
  const submitToDo = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget);

    // Convert FormData to JSON
    const formDataJson: { [key: string]: string } = {};
    formData.forEach((value, key) => {
      const val = value as string;

      // Only add non-empty entries
      if (val.length !== 0 && key.length !== 0) {
        formDataJson[key] = val;
      }
    });
    // console.log(formDataJson);

    // Check if form data has correct length
    const isObjectLengthCorrect = (obj: object) => {
      const keys = Object.keys(obj);
      return keys.length === 3; // since 'text' and 'description' and 'priority'
    };

    if (isObjectLengthCorrect(formDataJson)) {
      try {
        const response = await fetch('http://localhost:8080/todos/', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(formDataJson)
        });

        if (response.ok) {
          setReloadTodos(prevState => !prevState); // Toggle the flag
          // Clear the form fields
          formRef.current?.reset();
        } else {
          console.log('Error adding todo');
        }
      } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
      }
    } else {
      console.log("Format is incorrect blud");
    }
  }

  // Function to delete a todo
  const deleteTodo = async (id: number) => {
    try {
      const response = await fetch(`http://localhost:8080/todos/${encodeURIComponent(id)}`, {
        method: 'DELETE'
      });
      if (response.ok) {
        setReloadTodos(prevState => !prevState); // Toggle the flag to reload todos
      } else {
        console.error('Error deleting todo');
      }
    } catch (error) {
      console.error('There was a problem with the fetch operation:', error);
    }
  };

  return (
    <div className="app">
      <header className="app-header">
        <h1>TODO</h1>
      </header>

      <div className="todo-list">
        {todos.map((todo) =>
          <Todo
            id = {todo.id}
            key={todo.task + todo.description}
            task={todo.task}
            description={todo.description}
            priority={todo.priority}
            onDelete={() => deleteTodo(todo.id)} // Pass deleteTodo function
          />
        )}
      </div>

      <h2>Add a Todo</h2>
      <form onSubmit={submitToDo} ref={formRef}>
        <select name="priority">
          <option value="High">High</option>
          <option value="Medium">Medium</option>
          <option value="Low">Low</option>
        </select>
        <input placeholder="Task" name="task" autoFocus={true} />
        <input placeholder="Description" name="description" />
        <button type="submit">Add Todo</button>
      </form>
    </div>
  );
}

export default App;
