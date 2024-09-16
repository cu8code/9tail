# This documentaion contains our billion dollar plan

Hi I am Ankan the founder and probably still the core maintainer of the project. Below is GPT generated project backend structure, probably wrong but as of 16/Sep/2024 this is the future.

Good luck to myself!


### **Final Design: Merging Block and Workflow Logic**

In this final system design, we merge the concepts of **blocks** and **workflows** into a cohesive, process-like execution model where each workflow consists of a sequence of blocks. Each workflow can be treated as a standalone unit of execution, with individual blocks acting as stages within the workflow. The system behaves like a high-level **Operating System (OS)**, where workflows are treated as processes, and blocks are managed within the workflow as subprocesses.

---

### **1. Core Concepts: Unified Workflow-Block System**

#### **Unified Workflow-Block Entity**
- The **workflow** and its constituent **blocks** are now considered as a unified entity where blocks represent the sequential stages of a process, and workflows manage these blocks collectively.
- **Workflows** are dynamic processes, and each workflow consists of **one or more blocks**.
- Each workflow has its own **control structure**, containing necessary metadata and state tracking for both the overall workflow and the blocks within it.

### **2. Execution Engine Overview**

The execution engine manages both workflows and blocks. It schedules them for execution, coordinates communication between blocks, and ensures correct flow control.

#### **Execution Engine Architecture**
1. **Workflow Manager**: Responsible for starting, pausing, stopping, and managing the overall lifecycle of workflows.
2. **Block Processor**: Executes individual blocks inside the workflow.
3. **Scheduler**: Responsible for scheduling workflows, preempting them if necessary, and handling dependencies between blocks.
4. **IPC (Inter-Process Communication)**: Manages communication between blocks of different workflows.

---

### **3. Workflow Control Block (WCB)**

The **Workflow Control Block (WCB)** serves as the central control unit for each workflow. It contains metadata, state, and execution details for both the workflow and its blocks.

#### **Structure of WCB**

```typescript
interface WorkflowControlBlock {
  id: string;               // Unique workflow ID
  type: 'fetch' | 'html-scraper' | 'hook' | 'response';  // Workflow type
  state: 'READY' | 'RUNNING' | 'BLOCKED' | 'COMPLETED';  // Workflow state
  blocks: BlockControlObject[];  // List of blocks to execute in sequence
  input: any;                // Input data for the first block in the workflow
  output: any;               // Final output after all blocks are executed
  currentBlockIndex: number;  // Index of the currently executing block
  dependencies: string[];     // Dependencies on other workflows or blocks
  priority: number;           // Priority level of the workflow
  createdAt: Date;            // Creation time
  updatedAt: Date;            // Last updated timestamp
}
```

The **WCB** keeps track of the workflow’s execution state, manages transitions between blocks, and coordinates inputs/outputs between them.

---

### **4. Block Control Object (BCO)**

Each **block** within a workflow is represented by a **Block Control Object (BCO)**. This object manages individual blocks (like processes within a workflow) and handles block-specific execution logic.

#### **Structure of BCO**

```typescript
interface BlockControlObject {
  type: 'fetch' | 'html-scraper' | 'hook' | 'response';  // Block type
  input: any;              // Input to the block (can come from previous block or user-defined)
  output: any;             // Output from the block to be passed to the next block or workflow
  status: 'PENDING' | 'RUNNING' | 'COMPLETED' | 'FAILED';  // Execution status of the block
  retries: number;         // Number of retries in case of failure
  dependencies: string[];  // Dependencies on other blocks
}
```

- **Execution Order**: Blocks execute sequentially within a workflow, but can also execute in parallel if needed, depending on block type and dependencies.

---

### **5. Unified Execution Flow**

#### **Execution Model**

1. **Workflow Initialization**: 
   - The user provides a JSON configuration defining the workflow and its blocks.
   - The **Workflow Manager** parses this JSON, creating a **Workflow Control Block (WCB)** that contains all blocks as **Block Control Objects (BCO)**.

2. **Block Execution**:
   - The **Block Processor** picks blocks from the workflow (starting with the first) and executes them based on their type.
   - **Input Data**: Each block can either:
     - Receive input from a predefined value (user-defined in the JSON), or
     - Take the output from a previous block.
   - **Block Types**: Depending on the block type, different actions are performed (e.g., `fetch`, `html-scraper`, `response`).

3. **Parallel Execution**:
   - If a workflow has blocks that can run in parallel, the **Scheduler** assigns resources to these blocks accordingly.
   - The **Scheduler** handles CPU and memory allocation, making sure blocks do not compete for the same resources (to avoid race conditions).

4. **Block Chaining**:
   - Once a block completes, its **output** is passed to the next block.
   - If the block depends on another workflow or block (using IPC), the **Scheduler** ensures that the dependency is met before proceeding.

5. **Workflow Completion**:
   - The **Scheduler** monitors all blocks in the workflow. Once all blocks are executed, the workflow enters the **COMPLETED** state.

---

### **6. Scheduler (Inspired by OS Scheduling)**

The **Scheduler** ensures smooth and efficient execution of workflows and their blocks, inspired by OS-like scheduling policies.

#### **Features**:
- **Round-Robin Scheduling**: To ensure fair resource allocation, workflows are scheduled using round-robin, especially useful when multiple workflows compete for resources.
- **Priority Scheduling**: Workflows with higher priorities (e.g., time-sensitive tasks) get scheduled before lower-priority workflows.
- **Preemption**: If a block takes too long, it can be preempted by the Scheduler, freeing resources for other workflows.

---

### **7. Resource Management**

#### **Memory Management**
- **Isolated Memory**: Each workflow has isolated memory for its execution to avoid memory collisions. Blocks share memory only with their workflow.
- **Shared Memory**: If two workflows need to exchange data, the Scheduler manages the shared memory space.
  
#### **CPU Scheduling**
- **Multithreading/Multiprocessing**: Workflows are executed in parallel using either multithreading (for lightweight workflows) or multiprocessing (for heavy tasks that need process isolation).

---

### **8. Error Handling and Recovery**

#### **Error Management**
- If a block fails (e.g., network failure in a `fetch` block), the system can:
  - **Retry**: Attempt the block again after a delay.
  - **Skip**: Move to the next block if skipping is acceptable.
  - **Fail Workflow**: Mark the entire workflow as failed if critical blocks fail.

#### **Fault Tolerance**
- **Workflow Isolation**: Errors in one workflow do not affect other running workflows. Each workflow is isolated in its own execution context.

---

### **9. Example Workflow Execution**

Given the JSON configuration:

```json
[
  {
    "type": "hook",
    "path": "test"
  },
  {
    "type": "fetch",
    "parameter": {
      "url": "https://localhost:4000"
    }
  },
  {
    "type": "html-scraper",
    "parameter": [
      {
        "tags": "h1",
        "limit": 5
      },
      {
        "tags": "h2",
        "limit": 5
      }
    ]
  },
  {
    "type": "response",
    "header": [
      {
        "content-type": "json"
      }
    ],
    "body": "{}"
  }
]
```

1. **Step 1**: The **Workflow Manager** parses the JSON file and generates a **WCB** with four **BCOs** for each block (`hook`, `fetch`, `html-scraper`, `response`).

2. **Step 2**: The Scheduler starts executing the **`hook`** block, which triggers an internal function.

3. **Step 3**: The **`fetch`** block is executed next, making an HTTP request. While waiting for the response, the workflow enters the **BLOCKED** state.

4. **Step 4**: Once the `fetch` block completes, the **`html-scraper`** block processes the fetched HTML to extract data.

5. **Step 5**: Finally, the **`response`** block formats and returns the result.

6. **Step 6**: The workflow finishes execution and transitions to **COMPLETED**.

---

### **10. Tech Stack**

- **Backend**:
  - **Node.js** (for handling async tasks, I/O, and concurrency).
  - **TypeScript** (for type safety and scalability).
  - **Redis** (for shared memory/IPC between workflows).
  
- **Parallel Execution**:
  - **Worker Threads** or **Cluster API** in Node.js for multithreading/multiprocessing.

---

### **11. High-Level System Flow**

1. **User Input**: The user provides a JSON configuration file defining workflows and blocks.
2. **Workflow Manager**: The system parses the input, creates a **WCB** and associated **BCOs**, and hands it off to the Scheduler.
3. **Scheduler**: The Scheduler orchestrates the parallel execution of workflows, ensuring efficient resource management and IPC.
4. **Block Execution**: Blocks