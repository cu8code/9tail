import { create } from 'zustand';
import { v4 as uuidv4 } from 'uuid';
import { applyNodeChanges, Node } from '@xyflow/react';
import { ExtendedNode, Parameter } from './types';  // Assuming types are in a separate file

interface NodeStore {
  nodes: ExtendedNode[];
  addNode: (
    type: 'hook' | 'fetch' | 'scrapper',
    x: number,
    y: number,
    name: string,
    author: string
  ) => void;
  removeNode: (id: string) => void;
  updateNodePosition: (id: string, x: number, y: number) => void;
  updateNodeParameters: (id: string, parameters: Parameter[]) => void;
  getNodeById: (id: string) => ExtendedNode | undefined;
  getAllNodes: () => ExtendedNode[];
  onNodesChange: (changes: any[]) => void;  // Add this for XYFlow compatibility
}

const useStore = create<NodeStore>((set, get) => ({
  nodes: [],

  addNode: (type, x, y, name, author) => {
    const newNode: ExtendedNode = {
      id: `${type}.${uuidv4()}`,
      type: type,  // Add this for XYFlow compatibility
      position: { x, y },
      data: { label: type.charAt(0).toUpperCase() + type.slice(1) },  // Capitalized type as label
      name,
      author,
      parameters: getDefaultParameters(type),
    };

    set((state) => ({ nodes: [...state.nodes, newNode] }));
  },

  removeNode: (id) => {
    set((state) => ({
      nodes: state.nodes.filter((node) => node.id !== id),
    }));
  },

  updateNodePosition: (id, x, y) => {
    set((state) => ({
      nodes: state.nodes.map((node) =>
        node.id === id ? { ...node, position: { x, y } } : node
      ),
    }));
  },

  updateNodeParameters: (id, parameters) => {
    set((state) => ({
      nodes: state.nodes.map((node) =>
        node.id === id ? { ...node, parameters } : node
      ),
    }));
  },

  getNodeById: (id) => {
    return get().nodes.find((node) => node.id === id);
  },

  getAllNodes: () => {
    return get().nodes;
  },

  onNodesChange: (changes) => {
    set((state) => ({
      nodes: applyNodeChanges(changes, state.nodes as Node[]) as ExtendedNode[],
    }));
  },
}));

// Helper function to get default parameters based on node type
function getDefaultParameters(type: 'hook' | 'fetch' | 'scrapper'): Parameter[] {
  switch (type) {
    case 'hook':
      return [
        { name: 'name', type: 'string' },
        { name: 'age', type: 'number' },
      ];
    case 'fetch':
      return [
        { name: 'url', type: 'string' },
        { name: 'timeout', type: 'number' },
      ];
    case 'scrapper':
      return [
        { name: 'targetUrl', type: 'string' },
        { name: 'crawlDepth', type: 'number' },
      ];
  }
}

export default useStore;