import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import Canvas from './components/canvas.tsx';
import TextEditor from './components/text-editor.tsx';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      // Define nested routes if needed
      { path: "editor", element: <Canvas /> },
      { path: "script", element: <TextEditor /> },
      // Add more routes here
    ],
  },
]);
createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
