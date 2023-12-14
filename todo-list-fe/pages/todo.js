//todo.js

import { useState, useEffect, useContext } from 'react';
import axios from 'axios';
import '../styles/globalui.css';
import '../components/navbar.js';
import { AuthContext } from "../utils/AuthContext";
import DeleteIcon from '@mui/icons-material/Delete';
import IconButton from '@mui/material/IconButton';

const TodoList = () => {
    const [name, setName] = useState('');
    const [todos, setTodos] = useState([]);
    const [completed, setCompleted] = useState([]);
    const { token, setToken } = useContext(AuthContext);

    useEffect(() => {
        getTodos();
        getCompletedTodos();
    }, []);

    // const addTodo = (e) => {
    //     e.preventDefault();
    //     setTodos([...todos, name]);
    //     setName('');
    // };

    // const markAsCompleted = (index) => {
    //     const updatedCompleted = [...completed, todos[index]];
    //     const updatedTodos = todos.filter((_, i) => i !== index);
    //     setCompleted(updatedCompleted);
    //     setTodos(updatedTodos);
    // };

    // const deleteTodo = (index) => {
    //     const updatedTodos = todos.filter((_, i) => i !== index);
    //     setTodos(updatedTodos);
    // };
    // const deleteCompletedTodo = (index) => {
    //     const updatedCompleted = completed.filter((_, i) => i !== index);
    //     setCompleted(updatedCompleted);
    // };

    const getTodos = async () => {
        const response = await axios.get('http://localhost:5555/api/todos');
        // , {
        //     headers: {
        //         is_done: false,
        //     },
        //   });
        setTodos(response.data);
    };

    const getCompletedTodos = async () => {
        const response = await axios.get('http://localhost:5555/api/todos', {
            headers: {
                is_done: true,
            },
          });
        setCompleted(response.data);
    }; //saia bingung

    const addTodo = async (e) => {
        e.preventDefault();
        await axios.post('http://localhost:5555/api/todos', { text: name });
        setName('');
        getTodos();
    };

    const markAsCompleted = async (index) => {
        await axios.put(`http://localhost:5555/api/todos/${index}`);
        getTodos();
        getCompletedTodos();
    };

    const deleteTodo = async (index) => {
        await axios.delete(`http://localhost:5555/api/todos/${index}`);
        getTodos();
    };

    const deleteCompletedTodo = async (index) => {
        await axios.delete(`http://localhost:5555/api/todos/${index}`);
        getCompletedTodos();
    };

    return (
        <main class="container">
            <h1 style={{ textAlign: 'center', paddingTop: '50px' }}>Todo List</h1>
            <form onSubmit={addTodo}>
                <input
                    type="text"
                    id="name"
                    placeholder="Type your task here..."
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                />
                <div className="container" style={{ width: '200px' }}>
                    <button type="submit">Add Task</button>
                </div>
            </form>
            <div class="grid" style={{ paddingTop: '20px' }}>
                <h5 style={{ textAlign: 'center' }}>List</h5>
                <h5 style={{ textAlign: 'center' }}>Completed</h5>
            </div>
            <div class="grid">
                <ul>
                    {todos.map((todo, index) => (
                        <li key={index}>
                            <div class="grid">
                                <input type="checkbox" onChange={() => markAsCompleted(index)} />
                                {todo}
                                <IconButton onClick={() => deleteTodo(index)} style={{ width: '50px' }}>
                                    <DeleteIcon />
                                </IconButton>
                            </div>
                        </li>
                    ))}
                </ul>
                <ul>
                    {completed.map((todo, index) => (
                        <li key={index} style={{ textDecoration: 'line-through' }}>
                            <div class="grid">
                                {todo}
                                <IconButton onClick={() => deleteCompletedTodo(index)} style={{ width: '50px' }}>
                                    <DeleteIcon />
                                </IconButton>
                            </div>
                        </li>
                    ))}
                </ul>

            </div>
        </main>
    );
};

export default TodoList;