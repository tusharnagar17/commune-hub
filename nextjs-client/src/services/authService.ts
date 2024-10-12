import axios from "axios";

interface LoginCredentials {
    email: string;
    password: string;
}

interface RegisterCredentials {
    username: string;
    email: string;
    password: string;
}

const api = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080",
    headers: {
        'Content-Type': 'application/json', 
    },
    withCredentials: true
})

export const loginService = async (credentials: LoginCredentials) => {
    try {
        const response = await api.post('/login', credentials);
        return response.data;
    } catch (error: any) {
        if(error.response && error.response.data) {
            throw new Error(error.response.data.error || 'Login failed');
        } else {
            throw new Error("Network error")
        }
    }
}

export const registerService = async (credentials: RegisterCredentials) => {
    try {
        const response = await api.post('/register', credentials);
        return response.data;
    } catch (error: any) {
        if(error.response && error.response.data) {
            throw new Error(error.response.data.error || 'Login failed');
        } else {
            throw new Error("Network error")
        }
    }
}