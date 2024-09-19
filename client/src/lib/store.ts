import { create } from 'zustand';
import { createNodesSlice } from './slice/node';
import { createSelectionSlice } from './slice/selection';
import { createPositionAndParameterSlice } from './slice/positionAndParameter';
import { createCodeSlice  } from './slice/code';

const useStore = create((set, get) => ({
  ...createNodesSlice(set, get),
  ...createSelectionSlice(set, get),
  ...createPositionAndParameterSlice(set, get),
  ...createCodeSlice(set, get),
}));

export default useStore;
