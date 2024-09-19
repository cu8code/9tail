import { FiActivity } from 'react-icons/fi';
import { FaCloudDownloadAlt, FaSpider } from 'react-icons/fa';
import { v4 as uuidv4 } from 'uuid';
import { ExtendedNode } from './types';

// Common node style for all custom nodes
const commonNodeStyles = {
  backgroundColor: '#444',  // Background color of the node
  color: '#fff',            // Text color inside the node
  width: '80px',            // Square dimensions for all nodes
  height: '80px',
  display: 'flex',          // Flexbox to center the icon
  alignItems: 'center',
  justifyContent: 'center',
  fontSize: '30px',         // Size of the icon
  borderRadius: '10px',     // Slightly rounded corners
};


// Function to create a Hook node with additional information
export const createHookNode = (x: number, y: number, name: string, author: string): ExtendedNode => ({
  id: `hook.${uuidv4()}`,     // Use UUID for unique id
  position: { x, y },
  data: { label: <FiActivity /> },
  style: commonNodeStyles,
  name,  // Assign custom name
  author,  // Assign author name
  parameters: [
    { name: 'name', type: 'string' },
    { name: 'age', type: 'number' }
  ],  // Default parameters for Hook node
});

// Function to create a Fetch node with additional information
export const createFetchNode = (x: number, y: number, name: string, author: string): ExtendedNode => ({
  id: `fetch.${uuidv4()}`,    // Use UUID for unique id
  position: { x, y },
  data: { label: <FaCloudDownloadAlt /> },
  style: commonNodeStyles,
  name,  // Assign custom name
  author,  // Assign author name
  parameters: [
    { name: 'url', type: 'string' },
    { name: 'timeout', type: 'number' }
  ],  // Default parameters for Fetch node
});

// Function to create an HTTP Scrapper node with additional information
export const createHttpScrapperNode = (x: number, y: number, name: string, author: string): ExtendedNode => ({
  id: `scrapper.${uuidv4()}`, // Use UUID for unique id
  position: { x, y },
  data: { label: <FaSpider /> },
  style: commonNodeStyles,
  name,  // Assign custom name
  author,  // Assign author name
  parameters: [
    { name: 'targetUrl', type: 'string' },
    { name: 'crawlDepth', type: 'number' }
  ],  // Default parameters for Scrapper node
});
