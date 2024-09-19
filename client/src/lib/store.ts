import { create } from 'zustand';
import { createNodesSlice } from './nodeslice';
import { createSelectionSlice } from './selectionSlice';
import { createPositionAndParameterSlice } from './positionAndParameterSlice';

const useStore = create((set, get) => ({
  ...createNodesSlice(set, get),
  ...createSelectionSlice(set, get),
  ...createPositionAndParameterSlice(set, get),
}));

export default useStore;
