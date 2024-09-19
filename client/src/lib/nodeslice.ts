import { ExtendedNode, Parameter } from './types';
import { createFetchNode, createHookNode, createHttpScrapperNode } from './node';

interface NodesSlice {
  nodes: ExtendedNode[];
  addNode: (type: 'hook' | 'fetch' | 'scrapper', x: number, y: number, name: string, author: string) => void;
  removeNode: (id: string) => void;
  getNodeById: (id: string) => ExtendedNode | undefined;
  getAllNodes: () => ExtendedNode[];
}

export const createNodesSlice = (set: any, get: any): NodesSlice => ({
  nodes: [],

  addNode: (type, x, y, name, author) => {
    let newNode: ExtendedNode | null = null;
    switch (type) {
      case 'fetch': {
        newNode = createFetchNode(x, y, name, author);
        break;
      }
      case 'hook': {
        newNode = createHookNode(x, y, name, author);
        break;
      }
      case 'scrapper': {
        newNode = createHttpScrapperNode(x, y, name, author);
        break;
      }
    }
    set((state: any) => ({ nodes: [...state.nodes, newNode] }));
  },

  removeNode: (id) => {
    set((state: any) => ({
      nodes: state.nodes.filter((node: ExtendedNode) => node.id !== id),
    }));
  },

  getNodeById: (id) => {
    return get().nodes.find((node: ExtendedNode) => node.id === id);
  },

  getAllNodes: () => {
    return get().nodes;
  },
});
