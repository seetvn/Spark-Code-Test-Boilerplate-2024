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
  }, [reloadTodos]); 



  // Function to reload todos
  const handleReloadTodos = () => {
    setReloadTodos(prevState => !prevState); // Toggle reloadTodos
  };

  /* 
  function to make POST request + clears entry
  */
  const submitToDo = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const formData = new FormData(event.currentTarget); //formData 

    // Convert FormData to JSON
    const formDataJson: { [key: string]: string } = {};
    formData.forEach((value, key) => {
      const val = value as string;

      // Only add non-empty entries
      if (val.length !== 0 && key.length !== 0) {
        formDataJson[key] = val;
      }
    });

    // Check if form data has correct length
    const isObjectLengthCorrect = (obj: object) => {
      const keys = Object.keys(obj);
      return keys.length === 2; // since 'text' and 'description'
    };

    if (isObjectLengthCorrect(formDataJson)) { // only make POST request when format is correct
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
      console.log('Format is incorrect ');
    }
  }


  return (
    <div className="app">
      <header className="app-header">
        <h1>TODO</h1>
      </header>

      <div className="todo-list">
        {todos.map((todo) =>
          <Todo
            key={todo.title + todo.description}
            title={todo.title}
            description={todo.description}
          />
        )}
      </div>

      <h2>Add a Todo</h2>
      <form onSubmit={submitToDo} ref={formRef}>
        <input placeholder="Title" name="title" autoFocus={true} />
        <input placeholder="Description" name="description" />
        <button type="submit">Add Todo</button>
      </form>
    </div>
  );
}

export default App;
