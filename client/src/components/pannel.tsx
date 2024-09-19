import { useEffect, useState } from "react";
import useStore from "../lib/store";
import { Parameter } from "../lib/types";

export default function SuperPannel() {
    const getSelectedNode = useStore((state: any) => state.getSelectedNode);  // from selectionSlice
    const updateNodeParameters = useStore((state: any) => state.updateNodeParameters);  // from positionAndParameterSlice

    const [parameters, setParameters] = useState<Parameter[]>([]);

    // Sync parameters when the selected node changes
    useEffect(() => {
        const selectedNode = getSelectedNode();
        if (selectedNode) {
            setParameters(selectedNode.parameters || []);
        }
        console.log("selectedNode", selectedNode);
        
    }, [getSelectedNode]);

    const handleParameterChange = (index: number, value: string) => {
        const newParameters = [...parameters];
        newParameters[index].name = value;  // Update the value of the parameter, not the name
        setParameters(newParameters);
    };

    const handleUpdateParameters = () => {
        const selectedNode = getSelectedNode();
        if (selectedNode) {
            updateNodeParameters(selectedNode.id, parameters);
        }
    };

    return (
        <div style={{ flex: 1, padding: '20px' }}>
            <h3>Node Parameters</h3>
            {parameters.length > 0 ? (
                parameters.map((param, index) => (
                    <div key={param.name}>
                        <label>{param.name}: </label>
                        <input
                            type={param.type}
                            value={param.name|| ''}  // Bind to the parameter's value
                            onChange={(e) => handleParameterChange(index, e.target.value)}
                        />
                    </div>
                ))
            ) : (
                <p>No parameters available for the selected node.</p>
            )}
            <button onClick={handleUpdateParameters}>Update Parameters</button>
        </div>
    );
}
