// Типизация данных студента
export interface Student {
    id: number;
    first_name: string;
    last_name: string;
    group_id: number;
    birthday: string;
    comments: string;
}

// URL бэкенда
const API_BASE_URL = 'http://localhost:8080';

// Получение списка студентов
export const fetchStudents = async (): Promise<Student[]> => {
    const response = await fetch(`${API_BASE_URL}/students`);
    if (!response.ok) {
        throw new Error(`Failed to fetch students: ${response.statusText}`);
    }
    return await response.json();
};

// Пример функции для создания студента (если понадобится)
export const createStudent = async (student: Partial<Student>): Promise<Student> => {
    const response = await fetch(`${API_BASE_URL}/students`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(student),
    });
    if (!response.ok) {
        throw new Error(`Failed to create student: ${response.statusText}`);
    }
    return await response.json();
};