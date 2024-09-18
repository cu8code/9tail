import React, { useCallback } from 'react';
import {
  ReactFlow,
  Controls,
  Background,
  useNodesState,
  useEdgesState,
  addEdge,
  Edge,
  Node,
  Connection,
  MiniMap,
  BackgroundVariant
} from '@xyflow/react';
import '@xyflow/react/dist/style.css';

const nodeStyles = {
  backgroundColor: '#444', // Background color of the node
  color: '#fff',           // Text color inside the node
  padding: '10px',         // Padding inside the node
  borderRadius: '5px',     // Rounded corners
};

// Define the props interface
interface CanvasProps {
  width?: string;  // Width of the canvas (e.g., '80vw')
  height?: string; // Height of the canvas (e.g., '95vh')
}

const initialNodes: Node[] = [
  { id: '1', position: { x: 0, y: 0 }, data: { label: '1' }, style: nodeStyles },
  { id: '2', position: { x: 0, y: 100 }, data: { label: '2' }, style: nodeStyles },
];
const initialEdges: Edge[] = [{ id: 'e1-2', source: '1', target: '2' }];

const Canvas: React.FC<CanvasProps> = ({ width = '80vw', height = '95vh' }) => {
  const [nodes, setNodes, onNodesChange] = useNodesState<Node>(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge>(initialEdges);

  const onConnect = useCallback(
    (params: Connection) => setEdges((eds: Edge[]) => addEdge(params, eds)),
    [setEdges],
  );

  return (
    <div style={{ width, height }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onConnect={onConnect}
      >
        <Controls />
        <MiniMap />
        <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
      </ReactFlow>
    </div>
  );
}

export default Canvas;
