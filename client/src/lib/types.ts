import { Node } from "@xyflow/react";

// Define the parameter type
export interface Parameter {
  name: string;
  type: 'string' | 'number';
}

// Define the extended Node structure
export interface ExtendedNode extends Node {
  name: string;  // Custom name for the node
  author: string;  // Author of the node
  parameters: Parameter[];  // Array of parameters
}