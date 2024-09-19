import { useEffect, useState } from "react";
import useStore from "../lib/store";
import { createHTMLfrom9Tail } from "../lib/parser/main";

export default function SuperPannel() {
  const { code, getSelectedNode } : any = useStore();
  const [html, setHtml] = useState('');

  useEffect(() => {
    const parts = code.split(/^---[\s\S]*/m);
    const generatedHtml = createHTMLfrom9Tail(parts[0] ?? "");
    setHtml(generatedHtml);
  }, [getSelectedNode, code]);

  return (
    <div style={{ flex: 1, padding: '20px' }}>
      <h3>Node Parameters</h3>
      <div dangerouslySetInnerHTML={{ __html: html }} />
    </div>
  );
}