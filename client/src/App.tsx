import { useState, useEffect } from 'react';
import { Routes, Route, Link } from 'react-router-dom';
import { FaBeer, FaCog, FaQuestionCircle } from 'react-icons/fa';
import { PanelGroup, Panel, PanelResizeHandle } from "react-resizable-panels";
import Canvas from './components/canvas';
import TextEditor from './components/text-editor';
import Project from './components/projects';
import NodeSelectorPopup from './components/node-selector-popup'; // Import the NodeSelectorPopup component

function App() {
  const [dimensions, setDimensions] = useState({ width: 0, height: 0 });
  const [isNodeSelectorOpen, setIsNodeSelectorOpen] = useState(false); // State to manage node selector visibility

  // Update width and height when the window size changes
  useEffect(() => {
    const updateDimensions = () => {
      const headerHeight = 48; // Assuming header height is 3rem (48px)
      const availableHeight = window.innerHeight - headerHeight;
      const availableWidth = window.innerWidth * 0.8; // 80% width for canvas
      setDimensions({ width: availableWidth, height: availableHeight });
    };

    updateDimensions(); // Set initial dimensions
    window.addEventListener('resize', updateDimensions);

    // Cleanup the event listener
    return () => window.removeEventListener('resize', updateDimensions);
  }, []);

  return (
    <div className="bg-gray-950 text-white h-screen flex flex-col">
      <header className='flex flex-col w-full h-12 justify-between px-4 bg-gray-800 shadow-md'>
        <nav className='flex gap-4'>
          <Link to="/" className='flex items-center gap-2'>
            <FaBeer className='text-xl' />
            <span>Project</span>
          </Link>
          <div className='flex items-center gap-2 cursor-pointer'
            onClick={(e) => {
              e.preventDefault();
              setIsNodeSelectorOpen(true); // Open the node selector menu
            }}
          >
            <FaCog className='text-xl' />
            <span>Add Node</span>
          </div>
          <Link to="/editor" className='flex items-center gap-2'>
            <FaBeer className='text-xl' />
            <span>Editor</span>
          </Link>
          <Link to="/help" className='flex items-center gap-2'>
            <FaQuestionCircle className='text-xl' />
            <span>Help</span>
          </Link>
        </nav>
      </header>

      <main className='flex-grow flex'>
        <PanelGroup direction='horizontal' style={{ height: 'calc(100vh - 3rem)' }}>
          <Panel defaultSize={90} minSize={20} className='flex-grow'>
            <Routes>
              <Route path="/" element={<Project />} />
              <Route
                path="/editor"
                element={<Canvas width={`${dimensions.width}px`} height={`${dimensions.height}px`} />}
              />
              <Route path="/script" element={<TextEditor />} />
              {/* Add other routes as needed */}
            </Routes>
          </Panel>
          <PanelResizeHandle className='w-2 bg-gray-700' />
          <Panel defaultSize={10} minSize={20} className='flex-grow'>
            <div className='h-full w-full bg-gray-900 p-4'>Selection Menu</div>
          </Panel>
        </PanelGroup>
      </main>

      {isNodeSelectorOpen && <NodeSelectorPopup onClose={() => setIsNodeSelectorOpen(false)} />}
    </div>
  );
}

export default App;