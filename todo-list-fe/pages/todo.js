import { useState } from 'react';
import '../styles/globalui.css';
import '../components/navbar.js';
import DeleteIcon from '@mui/icons-material/Delete';
import IconButton from '@mui/material/IconButton';

const TodoList = () => {
    const [input, setInput] = useState('');
    const [todos, setTodos] = useState([]);
    const [completed, setCompleted] = useState([]);

    const addTodo = (e) => {
        e.preventDefault();
        setTodos([...todos, input]);
        setInput('');
    };

    const markAsCompleted = (index) => {
        const updatedCompleted = [...completed, todos[index]];
        const updatedTodos = todos.filter((_, i) => i !== index);
        setCompleted(updatedCompleted);
        setTodos(updatedTodos);
    };

    const deleteTodo = (index) => {
        const updatedTodos = todos.filter((_, i) => i !== index);
        setTodos(updatedTodos);
    };
    const deleteCompletedTodo = (index) => {
        const updatedCompleted = completed.filter((_, i) => i !== index);
        setCompleted(updatedCompleted);
    };
    return (
        <main class="container">
            <h1 style={{ textAlign: 'center', paddingTop: '50px' }}>Todo List</h1>
            <form onSubmit={addTodo}>
                <input
                    type="text"
                    placeholder="Type your task here..."
                    value={input}
                    onChange={(e) => setInput(e.target.value)}
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
                            <input type="checkbox" onChange={() => markAsCompleted(index)}/>
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