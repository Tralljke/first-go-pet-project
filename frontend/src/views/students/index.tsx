import React, { useEffect, useState } from 'react';
import { fetchStudents, Student } from '../../api/students';
import './styles.css';

const StudentsPage: React.FC = () => {
    const [students, setStudents] = useState<Student[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const loadStudents = async () => {
            try {
                const data = await fetchStudents();
                setStudents(data);
            } catch (err: any) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };

        loadStudents();
    }, []);

    if (loading) {
        return <div className="loading">Loading...</div>;
    }

    if (error) {
        return <div className="error">Error: {error}</div>;
    }

    return (
        <div className="students-page">
            <h1>Students List</h1>
            <table className="students-table">
                <thead>
                <tr>
                    <th>ID</th>
                    <th>First Name</th>
                    <th>Last Name</th>
                    <th>Group ID</th>
                    <th>Birthday</th>
                    <th>Comments</th>
                </tr>
                </thead>
                <tbody>
                {students.map((student) => (
                    <tr key={student.id}>
                        <td>{student.id}</td>
                        <td>{student.first_name}</td>
                        <td>{student.last_name}</td>
                        <td>{student.group_id}</td>
                        <td>{student.birthday}</td>
                        <td>{student.comments}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default StudentsPage;
