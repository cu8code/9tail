import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import TextEditor from './components/text-editor.tsx';

const router = createBrowserRouter([
  {
    path: "*",
    element: <App />,
    children: [
      { path: "script", element: <TextEditor /> },
    ],
  },
]);
createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
