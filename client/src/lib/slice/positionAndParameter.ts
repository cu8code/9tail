import { applyNodeChanges } from '@xyflow/react';
import { Parameter } from '../types';

interface PositionAndParameterSlice {
  updateNodePosition: (id: string, x: number, y: number) => void;
  updateNodeParameters: (id: string, parameters: Parameter[]) => void;
  onNodesChange: (changes: any[]) => void;
}

export const createPositionAndParameterSlice = (set: any, get: any): PositionAndParameterSlice => ({
  updateNodePosition: (id, x, y) => {
    set((state: any) => ({
      nodes: state.nodes.map((node: any) =>
        node.id === id ? { ...node, position: { x, y } } : node
      ),
    }));
  },

  updateNodeParameters: (id, parameters) => {
    set((state: any) => ({
      nodes: state.nodes.map((node: any) =>
        node.id === id ? { ...node, parameters } : node
      ),
    }));
  },

  onNodesChange: (changes) => {
    set((state: any) => ({
      nodes: applyNodeChanges(changes, state.nodes) as any,
    }));
  },
});
