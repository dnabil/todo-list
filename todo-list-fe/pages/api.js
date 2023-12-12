import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:5555/api',
});

export const registerUser = async (userData) => {
    try {
        const response = await api.post("/users/register", userData);
        return response.data;
    } catch (error) {
        console.error("Error registering user:", error);
        throw error;
    }
};

export const login = async (userData) => {
    try {
        const response = await api.post("/users/login", userData);
        return response.data;
    } catch (error) {
        console.error("Error logging in:", error);
        throw error;
    }
};