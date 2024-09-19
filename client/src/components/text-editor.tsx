import { useState } from 'react';
import MonacoEditor from '@monaco-editor/react';

export default function Editor() {
  const [code, setCode] = useState<string>('console.log("Hello, world!");');

  const handleEditorChange = (value: string | undefined) => {
    if (value !== undefined) {
      setCode(value);
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
