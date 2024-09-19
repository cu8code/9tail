import MonacoEditor from '@monaco-editor/react';
import useStore from '../lib/store';

export default function Editor() {
  const { code, updateCode }: any = useStore();

  const handleEditorChange = (value: string | undefined) => {
    if (value !== undefined) {
      updateCode(value);
    }
  };

  return (
    <div className="h-full w-full">
      <MonacoEditor
        height="100%"
        width="100%"
        language="text"
        theme="vs-dark"
        value={code}
        onChange={handleEditorChange}
        options={{
          selectOnLineNumbers: true,
          minimap: { enabled: false },
          automaticLayout: true,
        }}
      />
    </div>
  );
}
