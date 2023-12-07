import { useState } from 'react';
import '../styles/globalui.css';
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
                <button type="submit">Add Task</button>
            </form>
            <div class="grid">
                <h5>List</h5>
                <h5>Completed</h5>
            </div>
            <div class="grid">
                <ul>
                    {todos.map((todo, index) => (
                        <li key={index}>
                            <input type="checkbox" onChange={() => markAsCompleted(index)} />
                            {todo}
                            <IconButton onClick={() => deleteTodo(index)}>
                                <DeleteIcon />
                            </IconButton>
                        </li>
                    ))}
                </ul>
                <ul>
                    {completed.map((todo, index) => (
                        <li key={index} style={{ textDecoration: 'line-through' }}>
                            {todo}
                        </li>
                    ))}
                </ul>
            </div>
        </main>
    );
};

export default TodoList;