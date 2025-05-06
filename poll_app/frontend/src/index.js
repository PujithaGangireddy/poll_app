import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import ErrorPage from './components/ErrorPage';
import Home from './components/Home';
import Polls from './components/Polls';
import CreatePoll from './components/CreatePoll';
import Login from './components/Login';
import PollDetail from './components/PollDetail'; // ✅ Import this

const root = ReactDOM.createRoot(document.getElementById('root'));

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <Home /> },
      { path: "/polls", element: <Polls /> },
      { path: "/polls/create", element: <CreatePoll /> },
      { path: "/polls/:id", element: <PollDetail /> }, 
      { path: "/login", element: <Login /> },
    ],
  },
]);

root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
