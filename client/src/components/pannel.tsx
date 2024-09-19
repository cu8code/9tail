import { useEffect, useState } from "react";
import useStore from "../lib/store";
import { createHTMLfrom9Tail } from "../lib/parser/main";

export default function SuperPannel() {
  const { code, getSelectedNode } : any = useStore();
  const [html, setHtml] = useState('');

  useEffect(() => {
    const generatedHtml = createHTMLfrom9Tail(code);
    console.log(generatedHtml);
    
    setHtml(generatedHtml);
  }, [getSelectedNode, code]);

  return (
    <div style={{ flex: 1, padding: '20px' }}>
      <h3>Node Parameters</h3>
      <div dangerouslySetInnerHTML={{ __html: html }} />
    </div>
  );
}