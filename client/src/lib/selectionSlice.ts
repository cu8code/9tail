import { ExtendedNode } from './types';

interface SelectionSlice {
  selectedNodeId: string | null;
  selectNode: (id: string | null) => void;
  getSelectedNode: () => ExtendedNode | undefined;
}

export const createSelectionSlice = (set: any, get: any): SelectionSlice => ({
  selectedNodeId: null,

  selectNode: (id: string | null) => {
    set({ selectedNodeId: id });
  },

  getSelectedNode: () => {
    const { selectedNodeId, nodes } = get();
    return nodes.find((node: ExtendedNode) => node.id === selectedNodeId);
  },
});
