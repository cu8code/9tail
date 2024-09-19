import { useCallback, useState } from 'react';
import {
  ReactFlow,
  MiniMap,
  Controls,
  Background,
  addEdge,
  BackgroundVariant,
  Connection,
  Edge,
  useReactFlow,
  NodeMouseHandler,
} from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import useStore from '../lib/store';  // Import the custom store with slices
import { Parameter } from '../lib/types';

interface Props {
  width: string;
  height: string;
}

export default function Canvas({ width, height }: Props) {
  // Access nodes and handlers from slices
  const nodes = useStore((state: any) => state.nodes);  // from nodesSlice
  const onNodesChange = useStore((state: any) => state.onNodesChange);  // from nodesSlice
  const getNodeById = useStore((state: any) => state.getNodeById);  // from selectionSlice
  const selectNode = useStore((state: any) => state.selectNode);  // from selectionSlice

  const { setEdges } = useReactFlow();
  const [_, setParameters] = useState<Parameter[]>([]);

  const onNodeClick: NodeMouseHandler = useCallback(
    (_, node) => {
      selectNode(node.id);  // Select node on click
      const selectedNode = getNodeById(node.id);  // Get selected node details
      if (selectedNode) {
        setParameters([...selectedNode.parameters]);  // Set parameters if node exists
      }
    },
    [selectNode, getNodeById]
  );

  const onConnect = useCallback(
    (params: Connection | Edge) => setEdges((eds) => addEdge(params, eds)),
    [setEdges]
  );

  return (
    <div style={{ width, height, display: 'flex' }}>
      <div style={{ flex: 3 }}>
        <ReactFlow
          nodes={nodes}
          onNodesChange={onNodesChange}
          onConnect={onConnect}
          onNodeClick={onNodeClick}
        >
          <Controls />
          <MiniMap />
          <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
        </ReactFlow>
      </div>
    </div>
  );
}
