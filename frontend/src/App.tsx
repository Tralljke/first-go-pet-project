import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import StudentsPage from './views/students/index';

const App: React.FC = () => {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<StudentsPage />} />
            </Routes>
        </Router>
    );
};

export default App;