// NodeSelectorPopup.tsx
import React, { useState } from 'react';
import useStore  from '../lib/store'; // Import Zustand store
import { FiActivity } from 'react-icons/fi';
import { FaCloudDownloadAlt, FaSpider } from 'react-icons/fa';

// Node types for selection
const nodeTypes : {
    type: 'hook' | 'fetch' | 'scrapper';
    icon: JSX.Element;
    label: string;
}[] = [
  { type: 'hook', icon: <FiActivity className="text-2xl text-gray-300" />, label: 'Hook Node' },
  { type: 'fetch', icon: <FaCloudDownloadAlt className="text-2xl text-gray-300" />, label: 'Fetch Node' },
  { type: 'scrapper', icon: <FaSpider className="text-2xl text-gray-300" />, label: 'Scrapper Node' },
];

const NodeSelectorPopup: React.FC<{ onClose: () => void }> = ({ onClose }) => {
  const { addNode } = useStore();
  const [selectedType, setSelectedType] = useState<"hook" | "fetch" | "scrapper">('hook');
  const [name, setName] = useState<string>('');
  const [author, setAuthor] = useState<string>('');
  const [position, setPosition] = useState<{ x: number, y: number }>({ x: 0, y: 0 });

  // Handle node type selection
  const handleSelectType = (type: "hook" | "fetch" | "scrapper") => {
    setSelectedType(type);
  };

  // Handle form submission
  const handleSubmit = () => {
    if (name && author) {
      addNode(selectedType, position.x, position.y, name, author);
      onClose(); // Close the popup after adding the node
    }
  };

  return (
    <div className="fixed inset-0 bg-gray-900 bg-opacity-75 flex items-center justify-center">
      <div className="bg-gray-800 p-6 rounded-lg shadow-lg w-80">
        <h2 className="text-lg font-semibold text-white mb-4">Select Node Type</h2>
        <div className="flex flex-col gap-2 mb-4">
          {nodeTypes.map((node) => (
            <div
              key={node.type}
              className={`flex items-center p-2 cursor-pointer rounded-lg ${selectedType === node.type ? 'bg-gray-700' : 'bg-gray-800'} hover:bg-gray-700`}
              onClick={() => handleSelectType(node.type)}
            >
              {node.icon}
              <span className="ml-2 text-sm font-medium text-gray-300">{node.label}</span>
            </div>
          ))}
        </div>
        <input
          type="text"
          placeholder="Node Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="w-full p-2 border border-gray-600 rounded mb-2 bg-gray-900 text-white placeholder-gray-500"
        />
        <input
          type="text"
          placeholder="Author"
          value={author}
          onChange={(e) => setAuthor(e.target.value)}
          className="w-full p-2 border border-gray-600 rounded mb-4 bg-gray-900 text-white placeholder-gray-500"
        />
        <button
          onClick={handleSubmit}
          className="w-full bg-blue-600 text-white py-2 rounded mb-2 hover:bg-blue-700"
        >
          Add Node
        </button>
        <button
          onClick={onClose}
          className="w-full bg-gray-700 text-white py-2 rounded hover:bg-gray-600"
        >
          Close
        </button>
      </div>
    </div>
  );
};

export default NodeSelectorPopup;
