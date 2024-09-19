interface CodeSlice {
  code: string;
  addCode: (newCode: string) => void;
  updateCode: (newCode: string) => void;
  resetCode: () => void;
  getCode: () => string;
}

const createCodeSlice = (set: any, get: any): CodeSlice => ({
  code: 'console.log("Hello, world!");',

  addCode: (newCode) => {
    set((state: any) => ({ code: state.code + '\n' + newCode }));
  },

  updateCode: (newCode) => {
    set({ code: newCode });
  },

  resetCode: () => {
    set({ code: 'console.log("Hello, world!");' });
  },

  getCode: () => {
    return get().code;
  },
});

export {createCodeSlice};